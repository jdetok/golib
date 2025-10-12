package pgresd

import (
	"database/sql"
	"fmt"
	"sync"
	"time"
)

/* ORIGINAL DRAFT BUILT AND IN USE IN bball-etl-cli
- modified here for package use
*/

type InsertStmnt struct {
	Schema  string
	Tbl     string
	PrimKey string // define like "key" or "key1, key2"
	Cols    []string
	Vals    []any
	Rows    [][]any
	Chunks  [][][]any
}

type Table struct {
	Name    string
	PrimKey string
	PlTm    string
}

func MakeInsert(schema, tbl, primKey string, cols []string, rows [][]any) InsertStmnt {
	var ins = InsertStmnt{
		Schema:  schema,
		Tbl:     tbl,
		PrimKey: primKey,
		Cols:    cols,
		Rows:    rows,
	}
	ins.FlattenVals()
	ins.ChunkVals()
	return ins
}

func (ins *InsertStmnt) SchemaDotTable() string {
	return fmt.Sprintf("%s.%s", ins.Schema, ins.Tbl)
}

// flatten [][]any to []any with all values
func (ins *InsertStmnt) FlattenVals() {
	for _, r := range ins.Rows {
		ins.Vals = append(ins.Vals, r...)
	}
}

// flatten & return the values of a chunk of rows
func ValsFromSet(set [][]any) []any {
	var valsFlat []any
	for _, r := range set {
		valsFlat = append(valsFlat, r...)
	}
	return valsFlat
}

/*
populates ins.Chunks [][][]any with chunks
postgres sql.Exec() only allows 65,535 individual values to be inserted at once
ChunkVals populates ins.Chunks ([][][]any) with as many chunks ([][]any) with
as many full []any as necessary to keep the total number of values under 65,535.
* have found that setting the max vals in a chunk at 20,000 makes the individual
**execs much quicker
*/
func (ins *InsertStmnt) ChunkVals() {
	const PG_MAX int = 2000 // MUST BE < 65,535
	var totRows int = len(ins.Rows)
	var valsPer int = len(ins.Rows[0])
	var maxRows int = PG_MAX / valsPer
	var totVals int = len(ins.Vals)

	// number of chunks needed
	//subtracting by 1 enables ceiling integer division
	var numChunks int = (totVals + PG_MAX - 1) / PG_MAX

	// make slice of slice with 2 ends for start/end position in rows
	chunkPos := make([][2]int, 0, numChunks)
	for i := range numChunks {
		start := i * maxRows
		end := min((start + maxRows), totRows) // last row if < (start + tot)
		chunkPos = append(chunkPos, [2]int{start, end})
	}

	// append [][]any w/ start & end pos data from ins.Rows
	for _, c := range chunkPos {
		var valChunk [][]any = ins.Rows[c[0]:c[1]]
		ins.Chunks = append(ins.Chunks, valChunk)
	}
}

// loop through the chunks & attempt to insert all rows from each one
func (ins *InsertStmnt) InsertFast(db *sql.DB, global_row_count *int64) error {
	var wg sync.WaitGroup
	var mu sync.Mutex
	errCh := make(chan error, len(ins.Chunks))

	for i, c := range ins.Chunks {
		wg.Add(1)
		go func(i int, c [][]any) {
			defer wg.Done()
			st := time.Now()
			fmt.Printf("starting chunk %d/%d - %v\n", i+1, len(ins.Chunks), st)
			res, err := db.Exec(ins.BuildStmnt(c), ValsFromSet(c)...)
			if err != nil {
				chErr := fmt.Errorf("error inserting chunk %d/%d\n** %w", i+1, len(ins.Chunks), err)
				errCh <- chErr
			}
			ra, _ := res.RowsAffected()
			mu.Lock()
			*global_row_count += ra // add rows affected to total
			fmt.Println(
				fmt.Sprint(
					fmt.Sprintf("chunk %d/%d complete | rowsets: %d | vals: %d\n",
						i+1, len(ins.Chunks), len(c), len(ValsFromSet(c))),
					fmt.Sprintln("- ", time.Now()),
					fmt.Sprintln("- ", time.Since(st)),
					fmt.Sprintf("-- %d new rows inserted into %s\n", ra, ins.Tbl),
					fmt.Sprintln("-- total rows affected: ", *global_row_count),
				),
			)
			mu.Unlock()
			time.Sleep(1 * time.Second)
		}(i, c)
	}

	wg.Wait()
	close(errCh)
	if len(errCh) > 0 {
		err := <-errCh
		return fmt.Errorf("one or more chunks failed to insert| %e", err)
	}

	return nil
}

// construct the SQL statement to execute
func (ins *InsertStmnt) BuildStmnt(chunk [][]any) string {
	stmnt := fmt.Sprintf("insert into %s (", ins.Tbl)
	ins.addCols(&stmnt)
	ins.addChunkParams(&stmnt, chunk)
	return fmt.Sprintf("%s on conflict (%s) do nothing", stmnt, ins.PrimKey)
}

// use ins.Cols to add list of columns to sql statement
func (ins *InsertStmnt) addCols(stmnt *string) {
	for i, c := range ins.Cols {
		*stmnt += c
		if i < (len(ins.Cols) - 1) {
			*stmnt += ", "
		}
	}
	*stmnt += ")"
}

/*
creates list of placeholder params like ($1, $2, $3...)
postgres sql.Exec() function only accepts 65,535 total values per call
the chunk funcs break the vals into as many chunks of less than 65,535 as needed
*/
func (ins *InsertStmnt) addChunkParams(stmnt *string, chunk [][]any) {
	*stmnt += " values "
	for i, r := range chunk {
		*stmnt += "("
		for j := range r {
			// postgres uses 1-type placeholders, i*rows + idx of current val
			*stmnt += fmt.Sprintf("$%d", i*len(r)+(j+1))
			if j < (len(r) - 1) {
				*stmnt += ", "
			}
		}
		*stmnt += ")"
		if i < (len(chunk) - 1) {
			*stmnt += ", "
		}
	}
}

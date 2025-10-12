package pgresd

import "database/sql"

func ColumnsInTable(db *sql.DB, tbl string) ([]string, error) {
	q := `
select column_name
from information_schema.columns
where table_name = $1
order by ordinal_position
`
	rows, err := db.Query(q, tbl)
	if err != nil {
		return nil, err
	}
	var cols []string
	for rows.Next() {
		var col string
		if err := rows.Scan(&col); err != nil {
			return nil, err
		}
		cols = append(cols, col)
	}
	return cols, nil
}

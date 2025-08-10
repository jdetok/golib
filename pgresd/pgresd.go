package pgresd

import (
	"database/sql"
	"fmt"

	"github.com/jdetok/golib/envd"
	_ "github.com/lib/pq"
)

type PostGres struct {
	Host     string
	Port     int
	User     string
	Password string
	Database string
	ConnStr  string
}

func GetEnvFilePG(f string) PostGres {
	var pg PostGres
	envd.LoadDotEnvFile(f)
	pg.Host = envd.EnvStr("PG_HOST")
	pg.Port = envd.EnvInt("PG_PORT")
	pg.User = envd.EnvStr("PG_USER")
	pg.Password = envd.EnvStr("PG_PASS")
	pg.Database = envd.EnvStr("PG_DB")
	return pg
}

func GetEnvPG() PostGres {
	var pg PostGres
	envd.LoadDotEnv()
	pg.Host = envd.EnvStr("PG_HOST")
	pg.Port = envd.EnvInt("PG_PORT")
	pg.User = envd.EnvStr("PG_USER")
	pg.Password = envd.EnvStr("PG_PASS")
	pg.Database = envd.EnvStr("PG_DB")
	return pg
}

func (pg *PostGres) MakeConnStr() {
	pg.ConnStr = fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		pg.Host, pg.Port, pg.User, pg.Password, pg.Database)
}
func (pg *PostGres) Conn() (*sql.DB, error) {
	db, err := sql.Open("postgres", pg.ConnStr)
	if err != nil {
		return nil, err
	}
	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf(
			"error pining postgres after successful conn: %e", err)
	}
	return db, err
}

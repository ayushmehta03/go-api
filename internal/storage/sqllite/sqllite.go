package sqllite

import (
	"database/sql"

	"github.com/ayushmehta03/go-api/internal/config"
	_ "github.com/mattn/go-sqlite3"
)

type SqlLITE struct{
	Db *sql.DB
}


func New(cfg *config.Config)(*SqlLITE,error){
	db,err:=sql.Open("sqlite3",cfg.StoragePath)
	if err!=nil{
		return nil,err
	}
	
	_,err=db.Exec(`CREATE TABLE IF NOT EXISTS students(
	id INTEGER PRIMARY KEY AUTOINCREMENT,
	name TEXT,
	email TEXT,
     age INT
	
	
	
	)`)
	if err!=nil{
		return nil,err
	}

	return &SqlLITE{
		Db:db,

	},nil


}
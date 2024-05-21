package config

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

type ConnectorStruct struct {
	mysql *sql.DB
}

func (c ConnectorStruct) MySQL() *sql.DB {
	return c.mysql
}

type ConnectorIface interface {
	MySQL() *sql.DB
}

func InitDB() ConnectorIface {
	inst := new(ConnectorStruct)

	strConn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true",
		os.Getenv("MINDER_DB_USER"),
		os.Getenv("MINDER_DB_PASS"),
		os.Getenv("MINDER_DB_HOST"),
		os.Getenv("MINDER_DB_PORT"),
		os.Getenv("MINDER_DB_NAME"),
	)
	log.Println("strConn", strConn)
	db, err := sql.Open("mysql", strConn)
	if err != nil {
		log.Fatal(err)
	}
	// See "Important settings" section.
	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)
	if err = db.Ping(); err != nil {
		log.Fatal(err)

	}
	inst.mysql = db
	log.Print("InitDB() OK")
	return inst

}

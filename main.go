package main

import (
	"fmt"
	"log/slog"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"xorm.io/xorm"
)

type MysqlSettings struct {
	Username, Password string
	Ip, Port           string
	Database           string
}

var global MysqlSettings

func main() {
	err := read()
	if err != nil {
		slog.Error(err.Error())
		return
	}

	auth := global.Username + ":" + global.Password
	addr := global.Ip + ":" + global.Port
	database := global.Database
	dsn := fmt.Sprintf("%s@tcp(%s)/%s?charset=utf8", auth, addr, database)

	e, err := xorm.NewEngine("mysql", dsn)
	if err != nil {
		slog.Error(err.Error())
		return
	}
	f, err := os.Create("backup.sql")
	if err != nil {
		slog.Error(err.Error())
		return
	}
	defer f.Close()

	err = e.DumpAll(f)
	if err != nil {
		slog.Error(err.Error())
		return
	}
}

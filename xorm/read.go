package main

import (
	"bufio"
	"errors"
	"log/slog"
	"os"
	"strings"

	_ "github.com/go-sql-driver/mysql"
)

func read() error {
	f, err := os.Open("mysql.ini")
	if err != nil {
		return err
	}
	defer func() {
		cerr := f.Close()
		if cerr != nil {
			slog.Warn("", slog.Any("error", cerr))
		}
	}()
	s := bufio.NewScanner(f)
	for s.Scan() {
		ss := strings.Split(s.Text(), "=")
		if len(ss) != 2 {
			err = errors.New(`the format should be key="value"`)
			break
		}
		switch ss[0] {
		case "username":
			global.Username = strings.Trim(ss[1], `"`)
		case "password":
			global.Password = strings.Trim(ss[1], `"`)
		case "ip":
			global.Ip = strings.Trim(ss[1], `"`)
		case "port":
			global.Port = strings.Trim(ss[1], `"`)
		case "database":
			global.Database = strings.Trim(ss[1], `"`)
		}
	}
	return err
}

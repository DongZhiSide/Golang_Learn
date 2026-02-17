package main

import (
	"log/slog"
	"reflect"
	"testing"
	"time"

	"github.com/thanhpk/randstr"
	"github.com/xuri/excelize/v2"
)

type User struct {
	Username   string    `excel:"用户名"`
	Password   string    `excel:"-"`
	Nickname   string    `excel:"昵称"`
	Role       string    `excel:"角色"`
	Phone      string    `excel:"电话"`
	Email      string    `excel:"邮箱"`
	CreateTime time.Time `excel:"创建时间"`
}

func WriteUsersToXLSX(us []User, f *excelize.File) {
	n := reflect.TypeOf(User{}).NumField()
	u := User{}
	deleltEmptyRow := 0
	for ni := range n {
		key := reflect.TypeOf(u).Field(ni).Tag.Get("excel")
		if key == "-" {
			deleltEmptyRow--
			continue
		}
		cell, err := excelize.CoordinatesToCellName(ni+1+deleltEmptyRow, 1)
		if err != nil {
			slog.Error("", slog.Any("error", err))
			return
		}
		f.SetCellValue("Sheet1", cell, key)
	}
	for i := range us {
		deleltEmptyRow = 0
		for ni := range n {
			key := reflect.TypeOf(us[i]).Field(ni).Tag.Get("excel")
			if key == "-" {
				deleltEmptyRow--
				continue
			}
			cell, err := excelize.CoordinatesToCellName(ni+1+deleltEmptyRow, i+2)
			if err != nil {
				slog.Error("", slog.Any("error", err))
				return
			}
			value := reflect.ValueOf(us[i]).FieldByIndex([]int{ni}).Interface()
			slog.Info("", slog.Any(key, value))
			f.SetCellValue("Sheet1", cell, value)
		}
	}
}

func TestWriteUsersToXLSX(t *testing.T) {
	slog.SetLogLoggerLevel(slog.LevelError)
	n := 100
	us := []User{}
	for range n {
		u := User{
			Username:   "username: " + randstr.String(10),
			Password:   "password: " + randstr.String(10),
			Nickname:   "nickname: " + randstr.String(10),
			Role:       "admin",
			Phone:      randstr.String(10, "1234567890"),
			Email:      "email: " + randstr.String(10),
			CreateTime: time.Now(),
		}
		us = append(us, u)
	}
	f := excelize.NewFile()
	defer close(f.Close)
	WriteUsersToXLSX(us, f)
	err := f.SaveAs("Book2.xlsx")
	if err != nil {
		slog.Error(err.Error())
	}
}

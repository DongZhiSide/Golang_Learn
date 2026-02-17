package main

import (
	"fmt"
	"log/slog"
	"os"

	"github.com/xuri/excelize/v2"
)

func initLogger() {
	opts := &slog.HandlerOptions{
		AddSource: true,
		ReplaceAttr: func(groups []string, a slog.Attr) slog.Attr {
			if a.Key == slog.TimeKey {
				a.Value = slog.StringValue(a.Value.Time().Format("2006.01.02 15:04:05"))
			}
			return a
		}}
	slog.SetDefault(slog.New(slog.NewTextHandler(os.Stdout, opts)))
}

func close(f func() error) {
	err := f()
	if err != nil {
		slog.Error(err.Error())
	}
}

func main() {
	initLogger()
	f := excelize.NewFile()
	defer close(f.Close)
	// Create a new sheet.
	index, err := f.NewSheet("Sheet2")
	if err != nil {
		slog.Error(err.Error())
		return
	}
	// Set value of a cell.
	f.SetCellValue("Sheet2", "A2", "Hello world.")
	f.SetCellValue("Sheet1", "B2", 100)
	// Set active sheet of the workbook.
	f.SetActiveSheet(index)
	// Save spreadsheet by the given path.
	err = f.SaveAs("Book1.xlsx")
	if err != nil {
		slog.Error(err.Error())
	}
	cell, err := f.GetCellValue("Sheet1", "B2")
	if err != nil {
		slog.Error(err.Error())
	}
	fmt.Println(cell)
}

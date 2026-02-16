package main

import (
	"fmt"
	"log/slog"
	"os"

	"github.com/gocarina/gocsv"
)

func init() {
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

type NotUsed struct {
	Name string
}

// Our example struct, you can use "-" to ignore a field
type User struct {
	Id            string  `csv:"user_id"`
	Name          string  `csv:"user_name"`
	Age           string  `csv:"user_age"`
	NotUsedString string  `csv:"-"`
	NotUsedStruct NotUsed `csv:"-"`
}

func main() {
	readFile, err := os.OpenFile("read-user.csv", os.O_RDONLY, os.ModePerm)
	if err != nil {
		slog.Error(err.Error())
		return
	}
	defer readFile.Close()

	users := []*User{}
	err = gocsv.UnmarshalFile(readFile, &users) // Load user from file
	if err != nil {
		slog.Error(err.Error())
		return
	}

	for i := range users {
		fmt.Println("Hello", users[i].Name)
	}

	_, err = readFile.Seek(0, 0) // Go to the start of the file
	if err != nil {
		slog.Error(err.Error())
		return
	}

	users = append(users, &User{Id: "12", Name: "John", Age: "21"}) // Add user
	users = append(users, &User{Id: "13", Name: "Fred"})
	users = append(users, &User{Id: "14", Name: "James", Age: "32"})
	users = append(users, &User{Id: "15", Name: "Danny"})

	// Get all user as CSV string
	csvContent, err := gocsv.MarshalString(&users)
	if err != nil {
		slog.Error(err.Error())
		return
	}
	fmt.Println(csvContent) // Display all user as CSV string

	writeFile, err := os.OpenFile("write-user.csv", os.O_CREATE|os.O_TRUNC|os.O_WRONLY, os.ModePerm)
	if err != nil {
		slog.Error(err.Error())
		return
	}
	defer writeFile.Close()
	err = gocsv.MarshalFile(&users, writeFile) // Use this to save the CSV back to the file
	if err != nil {
		slog.Error(err.Error())
		return
	}
}

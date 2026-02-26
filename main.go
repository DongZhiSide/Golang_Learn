package main

import (
	"encoding/json"
	"log/slog"
	"net/http"

	_ "golanglearn/docs"

	httpSwagger "github.com/swaggo/http-swagger"
)

// @title 			Swagger Example API
// @version			2.0
// @description		Swagger Example API
// @termsOfService	http://swagger.io/terms/
// @contact.name	API Support
// @contact.url		http://www.swagger.io/support
// @contact.email	support@swagger.io
// @license.name	Apache 2.0
// @license.url		http://www.apache.org/licenses/LICENSE-2.0.html
// @schemes 		http https
// @host			localhost:8080
// @BasePath		/api
func main() {
	http.HandleFunc("/api/", apiHelloWorld)
	http.HandleFunc("/api/users", apiGetUser)

	// 渲染 swagger 文档
	http.Handle("/swagger/", httpSwagger.Handler(
		httpSwagger.URL("/swagger/doc.json"),
	))

	http.ListenAndServe(":8080", nil)
}

// @Summary			Hello World
// @Description		Hello World
//
// 第二个 string 是对应到 go 的类型
// 例如如果是对象, 则: Success	200	{object}	User	"Hello, World!"
//
// @Success			200	{string}	string	"Hello, World!"
// @Router			/ [get]
func apiHelloWorld(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("not found"))
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Hello, World!"))
}

type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

var users = []User{
	{ID: 1, Name: "John Doe"},
	{ID: 2, Name: "Jane Doe"},
}

// @Summary			List all users
// @Description		List all users
// @Produce      	json
// @Success			200	{array}	User	"List all users"
// @Router			/users [get]
func apiGetUser(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("not found"))
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	err := json.NewEncoder(w).Encode(users)
	if err != nil {
		slog.Error("failed to encode users", "error", err)
	}
}

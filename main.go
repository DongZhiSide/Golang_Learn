package main

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/fatih/color"
	"github.com/golang-jwt/jwt/v5"
)

var key []byte = []byte("secret key")

func Ver(token string) {
	p := jwt.NewParser()
	t, err := p.Parse(token, func(t *jwt.Token) (any, error) {
		return key, nil
	})
	if err != nil {
		color.Red("error: %s", err.Error())
	}

	if t.Valid {
		color.Green("token is valid")
	} else {
		color.Red("token is invalid")
	}

	color.Blue("header:")
	fmt.Println(t.Header)
	fmt.Println()

	// get Registered Claims
	//
	// type Claims interface {
	// 	GetExpirationTime() (*NumericDate, error)
	// 	GetIssuedAt() (*NumericDate, error)
	// 	GetNotBefore() (*NumericDate, error)
	// 	GetIssuer() (string, error)
	// 	GetSubject() (string, error)
	// 	GetAudience() (ClaimStrings, error)
	// }
	color.Blue("Registered Claims:")
	n, _ := t.Claims.GetAudience()
	fmt.Println("Audience claims:", n)
	e, _ := t.Claims.GetExpirationTime()
	fmt.Println("Expiration Time claims:", e)
	nb, _ := t.Claims.GetNotBefore()
	fmt.Println("Not Before claims:", nb)
	fmt.Println()

	color.Blue("custom claims:")
	// t.Claims is type MapClaims map[string]any
	fmt.Println("claims:", t.Claims)
}

type CustomClaims struct {
	jwt.RegisteredClaims
	User string `json:"user"`
}

func Gen(args []string) {
	// or use CustomClaims
	mc := jwt.MapClaims{}
	for _, s := range args {
		kv := strings.Split(s, ":")
		if len(kv) != 2 {
			color.Red("<claims> ::= <claim>:<value>")
			color.Red("example: user:jack")
			os.Exit(1)
		}
		mc[kv[0]] = kv[1]
	}
	mc["exp"] = jwt.NewNumericDate(time.Now().Add(time.Second * 5))

	t := jwt.NewWithClaims(jwt.SigningMethodHS256, mc)
	s, err := t.SignedString(key)
	if err != nil {
		color.Red("error: %s", err.Error())
	}
	fmt.Println("token is:", s)
}

func main() {

	if len(os.Args) <= 2 {
		fmt.Println("go run main.go gen <claims>")
		fmt.Println("\t<claims> ::= <claim>:<value>")
		fmt.Println("\texample: user:jack")
		fmt.Println("go run main.go ver <jwt>")
		os.Exit(1)
	}

	switch os.Args[1] {
	case "ver":
		Ver(os.Args[2])
	case "gen":
		Gen(os.Args[2:])
	default:
		fmt.Println("go run main.go gen <claims> <payload>")
		fmt.Println("\t<claims> ::= <claim>:<value>")
		fmt.Println("\texample: user:jack")
		fmt.Println("go run main.go ver <jwt>")
	}

	// go run main.go gen user:admin
}

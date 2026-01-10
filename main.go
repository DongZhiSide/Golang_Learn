package main

import (
	"fmt"
	"os"
	"time"

	"github.com/pquerna/otp/totp"
)

func main() {
	if len(os.Args) == 1 {
		key, err := totp.Generate(totp.GenerateOpts{
			Issuer:      "Example.com",
			AccountName: "alice@example.com",
		})
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		fmt.Println("Secret:", key.Secret())
		opts := totp.ValidateOpts{}
		passcode, err := totp.GenerateCodeCustom(key.Secret(), time.Now(), opts)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		fmt.Println("passcode:", passcode)
		// <secret> <passcode>
	} else if len(os.Args) == 3 {
		valid := totp.Validate(os.Args[2], os.Args[1])
		if valid {
			fmt.Println("Valid passcode!")
		} else {
			fmt.Println("Invalid passcode!")
		}
	}

}

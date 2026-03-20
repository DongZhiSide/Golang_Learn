package main

import (
	"encoding/json"
	"fmt"

	"github.com/samber/oops"
)

func main() {
	// Simple error with context
	err := oops.
		In("user-service").
		Tags("database", "postgres").
		Code("network_failure").
		User("user-123", "email", "foo@bar.com").
		With("path", "/hello/world").
		Errorf("failed to fetch user: %s", "connection timeout")

	// Error wrapping
	err = oops.
		Trace("req-123").
		With("product_id", "456").
		Wrapf(err, "user operation failed")

	// output
	fmt.Println(err.(oops.OopsError).Stacktrace())
	fmt.Println(err.(oops.OopsError).Sources())
	fmt.Printf("%+v", err)
	b, _ := json.MarshalIndent(err, "", "  ")
	fmt.Println(string(b))
}

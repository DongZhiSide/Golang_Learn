package main

import (
	"fmt"

	"github.com/casbin/casbin/v3"
)

func check(e *casbin.Enforcer, sub, obj, act string) {
	ok, _ := e.Enforce(sub, obj, act)
	if ok {
		fmt.Printf("%s CAN %s %s\n", sub, act, obj)
	} else {
		fmt.Printf("%s CANNOT %s %s\n", sub, act, obj)
	}
}
func main() {
	e, err := casbin.NewEnforcer("./model.pml", "./policy.csv")
	if err != nil {
		fmt.Printf("NewEnforecer failed:%v\n", err)
	}
	check(e, "dajun", "data1", "read")
	check(e, "lizi", "data2", "write")
	check(e, "dajun", "data1", "write")
	check(e, "dajun", "data2", "read")
	check(e, "root", "data2", "read")
}

package main

import (
	"fmt"
	"reflect"
)

type User struct {
	Name string `jun:"名前"`
	Age  int    `jun:"年齢"`
}

func main() {
	user := User{
		Name: "konojunya",
		Age:  21,
	}

	rt, rv := reflect.TypeOf(user), reflect.ValueOf(user)

	for i := 0; i < rt.NumField(); i++ {
		field := rt.Field(i)
		jun := field.Tag.Get("jun")

		fmt.Printf("jun: %v\n", jun)
		fmt.Printf("value: %v\n", rv.Field(i).Interface())
	}
}

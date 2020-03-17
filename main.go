package main

import (
	"fmt"
	"reflect"
	"strings"
)

func main() {
	type Person struct {
		Name string
		Age  int
	}

	john := Person{
		"John",
		12,
	}

	johnType := reflect.TypeOf(john)
	fmt.Println("johnType->", johnType) // main.Person

	johnTypeKind := reflect.TypeOf(john).Kind()
	fmt.Println("johnTypeKind->", johnTypeKind) // struct

	value := reflect.ValueOf(john)
	fmt.Println("ValueOf->", value) // {John 12}

	typeOf := value.Type()
	fmt.Println(typeOf) // main.Person

	valueField := value.Field(0)
	fmt.Println("ValueのField[0]->", valueField) // John

	typeField := typeOf.Field(0).Name
	fmt.Println("typeのField[0]->", typeField) // Name

	test := value.NumField()
	fmt.Println(test) // 2

	m := make(map[string]interface{})

	key := strings.ToLower(typeField)

	value2 := strings.ToLower(valueField.String())

	m[key] = value2

	fmt.Println(m)
}

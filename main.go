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

	person1 := Person{
		"John",
		12,
	}

	a := reflect.Indirect(reflect.ValueOf(person1).FieldByName(reflect.ValueOf(person1).Type().Field(0).Name)).Interface()

	fmt.Println("a->", a) // John

	johnType := reflect.TypeOf(person1)
	fmt.Println("johnType->", johnType) // main.Person

	johnTypeKind := reflect.TypeOf(person1).Kind()
	fmt.Println("johnTypeKind->", johnTypeKind) // struct

	value := reflect.ValueOf(person1)
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

package main

import (
	"fmt"

	"github.com/go-playground/validator/v10"
)

type Crowd struct {
	People []Person `validate:"unique=ID"`
}

type Person struct {
	ID   string
	Name string
	Age  int
}

func main() {
	crowd := Crowd{People: []Person{{"2", "Jean", 12}, {"2", "Anna", 12}, {"3", "Jean", 25}}}
	v := validator.New()
	fmt.Println(v.Struct(crowd)) // should not be nil as Jean and Anna are 12 years old
}

package main

import (
	"fmt"
)

type Student struct {
	Name     *StudentName `json:"name"`
	Age      int          `json:"age", sql, validation`
	Color    []string     `json:"color"`
	ColorMap map[string]int
}

type StudentName string

func (s StudentName) Validate() bool {
	return len(s) <= 32 && s != ""
}

var (
	colors   = []string{"red", "green", "blue"}
	colorMap = map[string]int{
		"red":   1,
		"green": 2,
	}
)

func (s Student) getStudent() Student {
	s.Color = append(s.Color, "yellow")
	s.Color[0] = "red"
	return s
}

func (s *Student) setAge(age int) {
	s.Age = age
}

func main() {
	name := StudentName("Thiep")
	student := Student{Name: &name, Age: 20}

	student.setAge(21)
	studentCopy := student.getStudent()
	fmt.Println(studentCopy)

	value := colorMap["red"]
	fmt.Println(value)

	colorMakeMap := make(map[string]int)

	colorMakeMap["red"] = 2

	fmt.Println(colorMakeMap)
}

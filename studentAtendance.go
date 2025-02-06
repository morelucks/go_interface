package main

import "fmt"

type Student struct{
	name string
	age uint8
}

func (s Student)addStudent(studentName string, studentAge int){
	s.name = studentName
	s.age = uint8(studentAge)
}

func main(){
	student1:= Student{
		name:"luckify",
		age: 98,
	}

	student1.addStudent(student1.name, int(student1.age))
	fmt.Printf("student %v \n", student1.name)

}
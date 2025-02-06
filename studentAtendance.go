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

func (s *Student)DeleteStudent(){
	s.name=""
	s.age=0
}


func (s Student)UdateStudent(NewName string, newAge int){
	s.name=NewName
	s.age=uint8(newAge)

}

func main(){
	student1:= Student{
		name:"luckify",
		age: 98,
	}


	student2:= Student{
		name:"luc",
		age: 98,
	}
	student2.addStudent(student2.name, int(student2.age))

	student1.addStudent(student1.name, int(student1.age))

	fmt.Println("before")
	fmt.Printf("student %v \n", student1.name)
	fmt.Printf("student %v \n", student2.name)

	student1.DeleteStudent()
	fmt.Println("after")
	// student1.UdateStudent("grace", 4)
	fmt.Printf("student %v \n", student1.name)
	fmt.Printf("student %v \n", student2.name)
}
package main

import "fmt"

type student struct{
  Name string
  Surname string
  Id string
  Grade string
}

type class struct{
  Name string
  Teacher string
  Students []student
}

type university struct{
  name string
}

func main() {
  var m int
  var KBTU university
  fmt.Print("Type the name of university: ")
  fmt.Scan(&KBTU.name)
  var temp class 
  fmt.Print("Type the name of discipline: ")
    fmt.Scan(&temp.Name)
    fmt.Print("Type the name of Teacher: ")
    fmt.Scan(&temp.Teacher)
    fmt.Print("Type the number of students: ")
    fmt.Scan(&m)
    temp.Students = make([]student, m)
    for j := 0; j < m; j++ {
      var stud student
      fmt.Print("Name of student: ")
      fmt.Scan(&stud.Name)
      fmt.Print("Surname of student: ")
      fmt.Scan(&stud.Surname)
      fmt.Print("Id of student: ")
      fmt.Scan(&stud.Id)
      fmt.Print("Grade: ")
      fmt.Scan(&stud.Grade)
      temp.Students[j] = stud
    }

  
  fmt.Print(temp)

}
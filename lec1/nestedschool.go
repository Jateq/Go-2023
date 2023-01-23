package main

import "fmt"

type student struct{
	name string
	surname string
	id int
}

// func (stud student) getName(s string) string{
// 	stud.name = s;
// 	return stud.name
// }
// func (stud student) getSurname() string{
// 	return stud.surname
// }
// func (stud student) getId() int{
// 	return stud.id
// }

type office_registrar struct{
	details student
}

func main(){
	var(
		name string
		surname string
		id int
	)
	fmt.Scanln(&name)
	fmt.Scanln(&surname)
	fmt.Scanln(&id)

	result := office_registrar{
		details : student{name, surname, id},
	}
	fmt.Println("\nDetails about student:")
	fmt.Println(result)
}

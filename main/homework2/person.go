package homework2

import "strconv"

type Person struct {
	Name string
	Age  int
}

type Employee struct {
	Person
	EmployeeID string
}

func (e *Employee) PrintInfo() string {
	return "Age: " + strconv.Itoa((*e).Age) + "Name: " + (*e).Name + ", EmployeeID: " + (*e).EmployeeID
}

// employee 使用 Employee 包
package main

import (
	"fmt"
	"go-learning/employee"
)

func main() {
	var dilbert employee.Employee
	dilbert = employee.Employee{ID: 1, Name: "Dilbert", Position: "Engineer", Salary: 27000}
	dilbert.Salary -= 5000
	fmt.Println(dilbert.Salary)
	position := &dilbert.Position
	*position = "Senior " + *position
	fmt.Println(dilbert.Position)
	fmt.Println(employee.EmployeeByID(dilbert.ID).Salary)
}
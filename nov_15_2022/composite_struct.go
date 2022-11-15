package main

import "fmt"

type Employee struct {
	ID   int
	Name string
	DOJ  string
	Department
}

type Department struct {
	DeptID   int
	DeptName string
}

func (e Employee) showEmpDetails() {
	fmt.Println(" ID : ", e.ID, "\n Name : ", e.Name, "\n Date of Joining : ", e.DOJ)
	e.Department.showDepartmentDetails()
}

func (d Department) showDepartmentDetails() {
	fmt.Println(" Dept ID : ", d.DeptID, "\n Dept Name : ", d.DeptName)
}

func main() {

	Employee1 := Employee{
		ID:   1001,
		Name: "tom",
		DOJ:  "26-03-22",
		Department: Department{
			DeptID:   123,
			DeptName: "Retail",
		},
	}

	Employee1.showEmpDetails()
}

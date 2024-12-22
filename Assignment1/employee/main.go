package main

import "fmt"

type Employee interface {
	GetDetails() string
}

type FullTimeEmployee struct {
	ID     string
	Name   string
	Salary uint32
}

type PartTimeEmployee struct {
	ID          string
	Name        string
	HourlyRate  uint64
	HoursWorked float32
}

func (f FullTimeEmployee) GetDetails() string {
	return fmt.Sprintf("Full-time Employee [ID: %s, Name: %s, Salary: %d KZT]", f.ID, f.Name, f.Salary)
}

func (p PartTimeEmployee) GetDetails() string {
	totalSalary := p.HourlyRate * uint64(p.HoursWorked)
	return fmt.Sprintf("Part-time Employee [ID: %s, Name: %s, Hourly Rate: %d KZT, Hours Worked: %.2f, Total Salary: %d KZT]",
		p.ID, p.Name, p.HourlyRate, p.HoursWorked, totalSalary)
}

type Company struct {
	Employees map[string]Employee
}

func (c *Company) AddEmployee(emp Employee) {
	empDetails := emp.GetDetails()
	c.Employees[empDetails] = emp
}

func (c *Company) ListEmployees() {
	fmt.Println("Employee List:")
	for _, emp := range c.Employees {
		fmt.Println(emp.GetDetails())
	}
}

func NewCompany() *Company {
	return &Company{Employees: make(map[string]Employee)}
}

func main() {
	c := NewCompany()

	c.AddEmployee(FullTimeEmployee{ID: "1", Name: "Alice", Salary: 70000})
	c.AddEmployee(PartTimeEmployee{ID: "2", Name: "Bob", HourlyRate: 416, HoursWorked: 140})
	c.AddEmployee(FullTimeEmployee{ID: "3", Name: "Jake", Salary: 82000})
	c.AddEmployee(PartTimeEmployee{ID: "4", Name: "Sofie", HourlyRate: 300, HoursWorked: 113})

	c.ListEmployees()
}

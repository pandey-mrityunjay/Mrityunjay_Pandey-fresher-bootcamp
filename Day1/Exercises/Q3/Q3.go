package main

import (
	"fmt"
)

// Employee interface
type Employee interface {
	CalculateSalary() int
}

// Full-time employee struct
type FullTime struct {
	SalaryPerMonth int
}

func (f FullTime) CalculateSalary() int {
	return f.SalaryPerMonth
}

// Contractor struct
type Contractor struct {
	SalaryPerMonth int
}

func (c Contractor) CalculateSalary() int {
	return c.SalaryPerMonth
}

// Freelancer struct
type Freelancer struct {
	PayPerHour  int
	HoursWorked int
}

func (f Freelancer) CalculateSalary() int {
	return f.PayPerHour * f.HoursWorked
}

// Main function
func main() {
	employees := []Employee{
		FullTime{SalaryPerMonth: 15000},
		Contractor{SalaryPerMonth: 3000},
		Freelancer{PayPerHour: 200, HoursWorked: 20},
	}
	// Calculate and print salaries
	for _, emp := range employees {
		fmt.Println("Salary:", emp.CalculateSalary())
	}
}

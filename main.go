package main

import (
	es "couchdb/db"
	"fmt"
	"os"
)

type DBInterface interface {
	Create(es.Employee) error
	Read(string) (es.Employee, error)
	Update(string, es.Employee) error
	Delete(string) error
}

var dbi DBInterface = &es.Employee{}

func main() {

	serviceCollection, err := es.GetCollection("localhost", "Employee", "loshik", "Nagaraja", "_default", "_default")
	if err != nil {
		fmt.Println("Connection Failed!!", err)
		os.Exit(1)
	} else {
		fmt.Println("Connected to cluster")
		es.SetDatasource(serviceCollection)
		crud()
	}
}

func crud() {

	var option int
	fmt.Println("Enter the operatin you want to do:")
	fmt.Println("1.Create \t 2.Read \t 3.Update \t 4.Delete \t 5.Exit")
	fmt.Scanln(&option)

	switch option {
	case 1:
		{
			var emp es.Employee
			fmt.Println("Enter the employee name:")
			fmt.Scanln(&emp.Name)
			fmt.Println("Enter the employee age:")
			fmt.Scanln(&emp.Age)
			fmt.Println("Enter the employee phone number:")
			fmt.Scanln(&emp.PhoneNumber)
			err := dbi.Create(emp)
			if err != nil {
				fmt.Println("Failed to add employee details!")
			} else {
				fmt.Println("Employee added successfully")
			}

			crud()
		}
	case 2:
		{
			var empId string
			fmt.Println("Enter the employee ID:")
			fmt.Scanln(&empId)

			emp, err := dbi.Read(empId)
			if err != nil {
				fmt.Println("Failed to get employee details!")
			} else {
				fmt.Println("Employee got successfully")
				fmt.Println(emp)
			}
			crud()
		}
	case 3:
		{
			var emp es.Employee
			var eId string
			fmt.Println("Enter the updated employee ID:")
			fmt.Scanln(&eId)
			fmt.Println("Enter the updated employee name:")
			fmt.Scanln(&emp.Name)
			fmt.Println("Enter the updated employee age:")
			fmt.Scanln(&emp.Age)
			fmt.Println("Enter the updated employee phone number:")
			fmt.Scanln(&emp.PhoneNumber)
			err := dbi.Update(eId, emp)
			if err != nil {
				fmt.Println("Failed to update employee details!")
			} else {
				fmt.Println("Employee updated successfully")
			}

			crud()
		}
	case 4:
		{
			var empID string
			fmt.Println("Enter the employee ID:")
			fmt.Scanln(&empID)

			err := dbi.Delete(empID)
			if err != nil {
				fmt.Println("Failed to delete employee details!")
			} else {
				fmt.Println("Employee deleted successfully")
			}

			crud()
		}
	case 5:
		{
			fmt.Println("Exited")
			os.Exit(0)
		}
	default:
		{
			fmt.Println("Wrong option!! Try Again")
			crud()
		}
	}
}

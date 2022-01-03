package db

import "github.com/couchbase/gocb/v2"

var serviceCollection *gocb.Collection

type Employee struct {
	Name        string `json:"name"`
	PhoneNumber string `json:"phoneNumber"`
	Age         int    `json:"age"`
}

func SetDatasource(couchbaseCollection *gocb.Collection) {
	serviceCollection = couchbaseCollection
}

func (e *Employee) Create(employee Employee) error {
	empId := employee.PhoneNumber + employee.Name
	_, err := serviceCollection.Insert(empId, employee, &gocb.InsertOptions{})
	if err != nil {
		return err
	}
	return nil
}

func (e *Employee) Read(empID string) (Employee, error) {
	getResult, err := serviceCollection.Get(empID, &gocb.GetOptions{})
	if err != nil {
		return Employee{}, err
	}

	var employee Employee
	err = getResult.Content(&employee)
	if err != nil {
		return Employee{}, err
	}
	return employee, nil
}

func (e *Employee) Update(empId string, employee Employee) error {
	_, err := serviceCollection.Replace(empId, employee, &gocb.ReplaceOptions{})
	if err != nil {
		return err
	}
	return nil
}

func (e *Employee) Delete(empID string) error {
	_, err := serviceCollection.Remove(empID, &gocb.RemoveOptions{})
	if err != nil {
		return err
	}
	return nil
}

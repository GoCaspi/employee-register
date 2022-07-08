package service

import (
	"example-project/model"
)

//go:generate go run github.com/maxbrunsfeld/counterfeiter/v6 . DatabaseInterface
type DatabaseInterface interface {
	UpdateMany(docs []interface{}) (interface{}, error)
	GetByID(id string) model.Employee
	DeleteByID(id string) (interface{}, error)
}

type EmployeeService struct {
	DbService DatabaseInterface
}

func NewEmployeeService(dbInterface DatabaseInterface) EmployeeService {
	return EmployeeService{
		DbService: dbInterface,
	}
}

func (s EmployeeService) CreateEmployees(employees []model.Employee) (interface{}, error) {

	var emp []interface{}
	for _, employee := range employees {
		emp = append(emp, employee)

	}

	response, err := s.DbService.UpdateMany(emp)

	if err != nil {
		return nil, err
	}

	return response, nil
}

func (s EmployeeService) GetEmployeeById(id string) model.Employee {
	return s.DbService.GetByID(id)
}
func (s EmployeeService) DeleteEmployee(id string) (interface{}, error) {
	return s.DbService.DeleteByID(id)

}

package service

import (
	"example-project/model"
	"go.mongodb.org/mongo-driver/mongo"
)

//go:generate go run github.com/maxbrunsfeld/counterfeiter/v6 . DatabaseInterface
type DatabaseInterface interface {
	UpdateMany(docs []interface{}) (interface{}, error)
	GetByID(id string) model.Employee
	DeleteByID(id string) (interface{}, error)
	GetPaginated(page int, limit int) (model.PaginatedPayload, error)
	UpdateEmp(update model.EmployeeReturn) (*mongo.UpdateResult, error)
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

func (s EmployeeService) GetPaginatedEmployees(page int, limit int) (model.PaginatedPayload, error) {
	result, err := s.DbService.GetPaginated(page, limit)
	return result, err
}

func (s EmployeeService) UpdateEmployee(update model.EmployeeReturn) (*mongo.UpdateResult, error) {
	result, err := s.DbService.UpdateEmp(update)
	return result, err
}

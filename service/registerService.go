package service

import (
	"errors"
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
	GetEmployeesByDepartment(department string) ([]model.EmployeeReturn, error)
	UpdateEmpShift(update model.Shift, id string) (model.Employee, error)
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

func (s EmployeeService) GetEmployeesDepartmentFilter(department string) ([]model.EmployeeReturn, error) {
	result, err := s.DbService.GetEmployeesByDepartment(department)
	if err != nil {
		return []model.EmployeeReturn{}, err
	}
	if len(result) == 0 && err == nil {
		noResultsErr := errors.New("No results could be found to your query")
		return result, noResultsErr
	}
	return result, err
}

func (s EmployeeService) AddShift(emp model.Employee, shift model.Shift) ([]model.Shift, error) {
	var shiftAlreadySet bool = false
	for _, s := range emp.Shifts {

		if s.Week == shift.Week {
			shiftAlreadySet = true
		}
	}
	if !shiftAlreadySet {
		response, err := s.DbService.UpdateEmpShift(shift, emp.ID)
		return response.Shifts, err
	} else {
		shiftErr := errors.New("The shift is already set for that week")
		return emp.Shifts, shiftErr
	}
}

func (s EmployeeService) GetRoster(employees []model.EmployeeReturn, week int) (map[string]map[string]model.Workload, error) {
	var roster map[string]map[string]model.Workload = map[string]map[string]model.Workload{}
	for _, e := range employees {
		emp := s.DbService.GetByID(e.ID)
		for _, s := range emp.Shifts {
			if s.Week == week {
				roster[e.FirstName+" "+e.LastName+" "+"id:"+" "+e.ID] = s.Duties
			}
		}
	}

	if len(roster) == 0 {
		emptyRosterErrMsg := "The recieved roster is empty."
		return roster, errors.New(emptyRosterErrMsg)
	}

	return roster, nil
}

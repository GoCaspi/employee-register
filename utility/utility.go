package utility

import (
	"crypto/sha256"
	"example-project/model"
)

func HashUsernameAndPassword(payLoad model.Auth) model.HashedAuth {
	hashedPassword := sha256.Sum256([]byte(payLoad.Password))
	hashedUsername := sha256.Sum256([]byte(payLoad.Username))

	hash := model.HashedAuth{Username: hashedUsername, Password: hashedPassword}
	return hash
}
func HashEmployees(emps []model.EmployeePayload) []model.Employee {
	var hashedEmps []model.Employee
	for _, emp := range emps {
		hashedEmp := model.Employee{
			ID:        emp.ID,
			FirstName: emp.FirstName,
			LastName:  emp.LastName,
			Email:     emp.Email,
			Auth: model.HashedAuth{
				Username: HashUsernameAndPassword(emp.Auth).Username,
				Password: HashUsernameAndPassword(emp.Auth).Password,
			},
		}
		hashedEmps = append(hashedEmps, hashedEmp)
	}
	return hashedEmps
}

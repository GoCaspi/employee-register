package model

import (
	"time"
)

type Employee struct {
	ID         string     `json:"id"`
	FirstName  string     `json:"first_name"`
	LastName   string     `json:"last_name"`
	Email      string     `json:"email"`
	Auth       HashedAuth `json:"auth" bson:"auth"`
	Department string     `json:"department" bson:"department"`
	Shifts     []Shift    `json:"shifts" bson:"shifts"`
}
type EmployeePayload struct {
	ID         string `json:"id"`
	FirstName  string `json:"first_name"`
	LastName   string `json:"last_name"`
	Email      string `json:"email"`
	Auth       Auth   `json:"auth" bson:"auth"`
	Department string `json:"department" bson:"department"`
}

type Auth struct {
	Username string `json:"username" bson:"username"`
	Password string `json:"password" bson:"password"`
}
type HashedAuth struct {
	Username [32]byte `json:"username" bson:"username"`
	Password [32]byte `json:"password" bson:"password"`
}

type Payload struct {
	Employees []EmployeePayload `json:"employees"`
}

type DbConfig struct {
	URL      string
	Database string
}

type EmployeeReturn struct {
	ID         string `json:"id"`
	FirstName  string `json:"first_name"`
	LastName   string `json:"last_name"`
	Email      string `json:"email"`
	Department string `json:"department" bson:"department"`
}

type PaginatedPayload struct {
	Employees []EmployeeReturn `json:"employees"`
	PageLimit int              `json:"pageLimit"`
}

type Workload struct {
	Duty  string    `json:"duty" bson:"duty"`
	Start time.Time `json:"start" bson:"start"`
	End   time.Time `json:"end" bson:"end"`
	Total int       `json:"total" bson:"total"`
}

type Shift struct {
	Week   int                 `json:"week" bson:"week"`
	Duties map[string]Workload `json:"duties" bson:"duties"`
}

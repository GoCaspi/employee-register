package model

type Employee struct {
	ID        string     `json:"id"`
	FirstName string     `json:"first_name"`
	LastName  string     `json:"last_name"`
	Email     string     `json:"email"`
	Auth      HashedAuth `json:"auth" bson:"auth"`
}
type EmployeePayload struct {
	ID        string `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	Auth      Auth   `json:"auth" bson:"auth"`
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
	ID        string `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
}

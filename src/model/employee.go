package model

type Employee struct {
	EmployeeID string  `json:"employeeId,omitempty" bson:"employeeId,omitempty"`
	FirstName  string  `json:"firstName,omitempty" bson:"firstName,omitempty"`
	LastName   string  `json:"lastName,omitempty" bson:"lastName,omitempty"`
	Age        int     `json:"age,omitempty" bson:"age,omitempty"`
	Position   string  `json:"position,omitempty" bson:"position,omitempty"`
	Salary     float64 `json:"salary,omitempty" bson:"salary,omitempty"`
}

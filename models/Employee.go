package models

type Employee struct {
	ID       uint
	Name     string
	Mobile   string
	Address  string
	Salary   int
	Age      int
	Username string
	Password string
}

func (b *Employee) TableName() string {
	return "Employee"
}

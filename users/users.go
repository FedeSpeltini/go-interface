package users

import "fmt"

type User struct {
	Id   int
	Name string
}

type Employee struct {
	User
	Active bool
}

type Cashier interface {
	CalcTotal(item ...float32) float32
	desactivate()
	reactivate()
}

type Admin interface {
	DesactivateEmployee(c Cashier)
}

func NewEmployee(id int, name string) *Employee {
	return &Employee{User{Id: id, Name: name}, true}
}

func (e *Employee) CalcTotal(item ...float32) float32 {
	if !e.Active {
		fmt.Println("Cashier is inactive")
		return 0
	}

	var total float32
	for _, v := range item {
		total += v
	}
	return total
}

func (e *Employee) DesactivateEmployee(c Cashier) {
	c.desactivate()
}

func (e *Employee) desactivate() {
	e.Active = false
}

func (e *Employee) reactivate() {
	e.Active = true
}

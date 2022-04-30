package main

import (
	"fmt"

	"github.com/my-first-interface/users"
)

func main() {

	var frank users.Cashier = users.NewEmployee(1, "Frank")
	var robert users.Admin = users.NewEmployee(2, "Robert")
	total := frank.CalcTotal(1.0, 2.0, 3.0)
	fmt.Println(total)

	robert.DesactivateEmployee(frank)

	fmt.Println(frank.CalcTotal(1.0, 2.0, 3.0))

}

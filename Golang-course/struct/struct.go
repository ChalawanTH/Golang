package main

import "fmt"

type employee struct {
	employeeID   string
	employeeName string
	phone        string
}

func main() {
	//Array
	//employeeList := [3]employee{}
	// employeeList[0] = employee{employeeID: "101", employeeName: "Pradoo", phone: "+66"}
	// employeeList[1] = employee{employeeID: "102", employeeName: "Prayat", phone: "+67"}
	// employeeList[2] = employee{employeeID: "103", employeeName: "Pranee", phone: "+68"}

	//Slice
	employeeList := []employee{}
	employeeList1 := employee{employeeID: "101", employeeName: "Pradoo", phone: "+66"}
	employeeList2 := employee{employeeID: "102", employeeName: "Prayat", phone: "+67"}
	employeeList3 := employee{employeeID: "103", employeeName: "Pranee", phone: "+68"}
	employeeList = append(employeeList, employeeList1)
	employeeList = append(employeeList, employeeList2)
	employeeList = append(employeeList, employeeList3)

	fmt.Println("employee =", employeeList)
}

package structure

import "fmt"

//create structure of type Employee
type Employee struct {
	name string
	post string
	employeeCode string
	isActive bool
	experience float32
}

//array contains the employee
var headCount []Employee;

func AddEmployee()  {
	Employee1 := Employee{"Rizvi","SDE","HD-353",true,1.8};
	Employee2 := Employee{"Afzal","SDE 2","HD-393",true,2.3};
	headCount = append(headCount,Employee1)
	headCount = append(headCount,Employee2);
	PrintEmployee();
}

func PrintEmployee()  {
	for _, empObj := range headCount{
		fmt.Printf("%s is %s with experience of %.1f years\n",empObj.name,empObj.post,empObj.experience);
	}
}


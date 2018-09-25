package conditionalLoops

import "fmt"

func SwitchExp(name string)  {

	switch name {

	case "afzal":
		fmt.Printf("This is the profile for Afzal\n")

	case "rizvi":
		fmt.Println("This is the profiel for Rizvi\n")

	default:
		fmt.Println("profile doesn't matches with anyone\n")

	}
}

func IfElseExp(flag bool) bool {
	if(flag){
		return true;
	}
	return false;
}

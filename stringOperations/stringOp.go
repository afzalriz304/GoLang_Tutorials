package stringOperations

import "strings"

func StringManipulation(str string)  {

	println("length of string ",len(str))
	println(strings.Contains(str,"a"))
	println(strings.Compare(str,"text"))
	HW := []string{"hello","world"};
	println(strings.Join(HW," to my "));
}

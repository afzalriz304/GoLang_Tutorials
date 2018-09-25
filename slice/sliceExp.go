package slice

import "fmt"

func CreateDynamiceArray(length, cap int) {

	dynamicArr := make([]int, length, cap);
	fmt.Println("length of array ", len(dynamicArr))
	var i int

	for i=0; i<len(dynamicArr);i++  {
		dynamicArr[i]= i;
	}

	for _, arr := range dynamicArr {
		fmt.Println(arr)
	}
	city := make(map[string]string)


}

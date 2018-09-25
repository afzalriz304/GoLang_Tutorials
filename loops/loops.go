package loops

import "fmt"

func ForLoop()  {
	var i int;
	numbers :=[6]int{1,2,3,4,5};

	//simple For loop
	for i=0; i<5 ;i++  {
		fmt.Println("simple ",i);
	}

	//range for loop
	for _, num := range numbers{
		fmt.Println("range ",num);
	}


}

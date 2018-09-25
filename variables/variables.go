package variables

import "fmt"

func Variables() {
	var a,b int
	var name string
	var flag bool
	const g  float32= 9.8


	flag=true
	name="Afzal"
	a=10
	b=20

	fmt.Printf("type of a is %T\n",a);
	fmt.Printf("type of b is %T\n",b);
	fmt.Printf("type of name is %T\n",name);
	fmt.Printf("type of flag is %T\n",flag);
	fmt.Printf("value of g is %f\n",g);
}

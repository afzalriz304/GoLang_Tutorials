package Interfaces

import "fmt"



//create an interface for increment the number
type Incrementor interface {
	Increment() int
}

type NeedIncrement int;

func (ic *NeedIncrement) Increment() int {
	*ic++;
	return int(*ic);
}

type NewAdd interface{
	New_Add() int
}

type Adder interface {
	Addition() int
}

type Values int

func (n Values) Addition() int {
	n++;
	return int(n);
}


func Local_main()  {


	/*polymorphism*/
	var rec rectangle;
	//rec.breath=10
	//rec.length=8
	fmt.Println("Area of rectangle",AreaCalulator.Area(rec))

	var c circle
	//c.radius="";
	fmt.Println("Area of circle",AreaCalulator.Area(c))

//-------------------polymorphism ends here----------------------/


	/*passing data to the interfaces method*/
	var i int
	n:=NeedIncrement(0)
	var num = &n;
	for i=0;i<10 ;i++ {
		fmt.Println(Incrementor.Increment(num))
	}

	no := Values(1)
	var a= &no;
	fmt.Println("its adder return",Adder.Addition(a))

	var response interface {} ="aaaa"

	switch response.(type) {

	case int:
		fmt.Println("int")
	case string:
		fmt.Println("string")
	case bool:
		fmt.Println("it's boolean")
	}
}



//create an interface for calculation of area (polymorphism)
type AreaCalulator interface {
	Area() interface{};
}

type circle struct {
	radius int
}

type rectangle struct {
	length, breath int
}

type square struct {
	side int
}


func (c circle) Area() interface{} {
	//area := 3.14 * float32(c.radius*c.radius);
	var a string ="Rizvi";
	return a;
}

func (r rectangle) Area() interface{} {
	area := float32(r.length) * float32(r.breath);
	return area;
}

func (s square) Area() interface{}  {
	area := float32(s.side*s.side);
	return area
}




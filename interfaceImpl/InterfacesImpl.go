package interfaceImpl

import (
	"../Interfaces"
	"fmt"
)

type addMe int;

type a int;

func (a addMe) New_Add() int  {
	a++
	return int(a);
}

func (a addMe) Addition() int {
	a++;
	return int(a);
}

func InterfaceImpl()  {
	
	n := addMe(5)
	var a = &n;
	//fmt.Println("added",Interfaces.Adder.Addition(a));

	fmt.Println("add me works ",Interfaces.NewAdd.New_Add(a))

	
}


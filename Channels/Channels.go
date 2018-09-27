package Channels

import (
	"fmt"
	"sync"
)

var wg = sync.WaitGroup{}

func ImplementingChannels(n int, ch chan<- int)  {

	fmt.Println("Implementing channels")

	result := n*2;

	ch<- result


	cha := make(chan int)
	wg.Add(2)

	go getter(cha)
	go setter(cha)

	wg.Wait()


}

func setter(ch chan<- int)  {
	ch<- 45
	ch<- 30
	wg.Done()
}

func getter(ch <-chan int)  {
	/*i := <-ch
	fmt.Println("i val",i)
	i = <-ch
	fmt.Println("i val",i)*/
	for i := range ch{
		fmt.Println("val",i)
	}
	wg.Done()
}

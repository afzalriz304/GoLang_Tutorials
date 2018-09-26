package GoRoutines

import (
	"fmt"
	"log"
	"sync"
)

var counter int =0;


func PrintHello(wg sync.WaitGroup, m sync.RWMutex)  sync.WaitGroup{
	fmt.Printf("Hello %d\n",counter)
	m.RUnlock()
	log.Printf("unlocking hello for counter %d",counter)
	wg.Done()
	log.Printf("remove hello from wait group")
	return wg
}

func Increment(wg sync.WaitGroup, m sync.RWMutex) sync.WaitGroup {
	log.Printf("inside increament")
	counter++
	m.Unlock()
	log.Printf("increament done with counter %d",counter)
	wg.Done()
	log.Printf("releasing increament from  wait group")

	return wg
}

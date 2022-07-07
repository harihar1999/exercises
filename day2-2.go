package main

import (
	"fmt"
	"math/rand"
	"sync"
	"sync/atomic"
	"time"
)


func student (i int, ) int32{
	//fmt.Printf("%d start\n", i)
	time.Sleep(time.Duration(rand.Int63n(1))*time.Second)
	//fmt.Printf("%d done\n", i)
	return rand.Int31n(10)
}

func main(){

	var wg sync.WaitGroup



    var sum int32 =0

	for i:=0; i<200; i++ {
		wg.Add(1)

		i:=i

		go func() {

			defer wg.Done()


			atomic.AddInt32(&sum, student(i))


		}()

	}


    wg.Wait()
	fmt.Println(float32(sum)/float32(200))
}

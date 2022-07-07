package main



import (
	"fmt"
	"sync"
)

var (
	mutex          sync.Mutex
	accountBalance int
)


func main() {
    accountBalance=500
	fmt.Printf("starting balance %d\n", accountBalance)


	var wg sync.WaitGroup
    for i:=0;i<3 ;i++ {
		wg.Add(2)

		go withdraw(300, &wg)
		go deposit(500, &wg)

	}

	wg.Wait()
	fmt.Printf("New balance %d\n", accountBalance)

}

func withdraw(value int, wg *sync.WaitGroup) {

	mutex.Lock()
	defer mutex.Unlock()
	if accountBalance>=value {
		accountBalance -= value
		fmt.Printf("Current balance: %d\n", accountBalance)
	} else {
		fmt.Println("balance is insufficient")
	}



	wg.Done()

}
func deposit(value int, wg *sync.WaitGroup) {

	mutex.Lock()
	defer mutex.Unlock()
	accountBalance += value
	fmt.Printf("Current balance: %d\n", accountBalance)


	wg.Done()

}


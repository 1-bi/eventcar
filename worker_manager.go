package eventcar

import (
	"fmt"
	"time"
)

// workerManager monitor all running services, if facade mode
type WorkerManager struct {
	controlCh <-chan int
}

func NewWorkerManager() *WorkerManager {
	wm := new(WorkerManager)
	wm.controlCh = make(chan int)
	return wm
}

func (myself *WorkerManager) SentSign(sign <-chan int) {
	myself.controlCh = sign
}

func (myself *WorkerManager) Run() {
	/*
		for x := range myself.controlCh{

			fmt.Println(" control value ")
			fmt.Println(  x )

		}
	*/
	ch1 := make(chan int)
	go func() {

		time.Sleep(2 * time.Second)

		ch1 <- 1

	}()
	// --- connect to message

	for x := range ch1 {

		fmt.Println(" control value ")
		fmt.Println(x)
		if x == 1 {
			go func() {
				ch1 <- 2

			}()
		}
		if x == 2 {

			go func() {
				ch1 <- 3

			}()

		}
		if x == 3 {
			break
		}

	}

	/*
		for {
			select {
			case ka, ok := <- ch1:

				fmt.Println( ok )

				fmt.Println(  ka )

			}

		}
	*/

}

package test

import (
	"fmt"
	"testing"
)

func Test_Out(t *testing.T) {
	c := make(chan int)
	defer close(c)

	go func() {
		c <- 3 + 4
		fmt.Println("------999999")
	}()
	i := <-c
	fmt.Println(i)
}

type SuccessCallbackImpl struct {
}

func (myself *SuccessCallbackImpl) Succeed(content []byte) {
	fmt.Println(" callback successfully .")
	fmt.Println(string(content))
}

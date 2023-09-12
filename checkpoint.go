package main

import (
	"fmt"
	"sync"
	"time"
)

func worker(id int, checkpoint, resume chan struct{}, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Printf("worker %d:starting\n", id)
	time.Sleep(time.Duration(id) * time.Second)
	fmt.Printf("worker %d:reached checkpoint\n", id)
	checkpoint <- struct{}{}
	<-resume
	fmt.Printf("worker %d:resuming\n", id)
}
func main() {
	var num int
	fmt.Print("enter num of workers:")
	fmt.Scan(&num)
	checkpoint := make(chan struct{})
	resume := make(chan struct{})
	var wg sync.WaitGroup
	for i := 1; i <= num; i++ {
		wg.Add(1)
		go worker(i, checkpoint, resume, &wg)

	}
	for i := 1; i <= num; i++ {
		<-checkpoint
	}
	fmt.Println("all workers reached checkpoint")
	close(resume)
	wg.Wait()
	fmt.Println("allworkers completd work")

}

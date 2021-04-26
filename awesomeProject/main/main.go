package main

import (
	"fmt"
	"time"
)

func main() {
	c := make(chan int) // Делает канал для связи
	for i := 0; i < 5; i++ {
		go sleepyGopher(i, c)
	}
	for i := 0; i < 5; i++ {
		gopherID := <-c // Получает значение от канала
		fmt.Println("gopher ", gopherID, " has finished sleeping")
	}
}

func sleepyGopher(id int, c chan int) { // Объявляет канал как аргумент
	time.Sleep(3 * time.Second)
	fmt.Println("... ", id, " snore ...")
	c <- id // Отправляет значение обратно к main
}
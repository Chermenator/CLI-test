package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"io"
	"os"
	"strings"
	"time"
)

func main() {
	timeout := flag.Int("timer", 30, "timeout for quiz")
	flag.Parse()
	file, err := os.Open("main/data.csv")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	reader.FieldsPerRecord = 2
	reader.Comment = '#'
	reader.Comma = ','

	k1 := 0
	k2 := 0

	c  := make(chan string, 1)

	timer := time.After(time.Duration(*timeout) * time.Second)

LOOP:
	for {
		record, e := reader.Read()
		if e != nil {
			if e == io.EOF {
				break
			}
		}

		fmt.Print(record[0] + "= ")
		ans := ""

		go fff(ans,record[1],c)
		ans = strings.TrimSpace(ans)




		select {
		case <-timer:
			fmt.Println("\nYour time is up")
			break LOOP
		case b:=<-c:
			if b == record[1] {
				k1++
			} else {
				k2++
			}
			continue LOOP
		}

	}

	fmt.Printf("\nПравильных ответов: %d\nНеправильных ответов: %d", k1, k2)
}

func fff(ans string, rec1 string, c1 chan string) {
	fmt.Fscan(os.Stdin, &ans)
	if ans == rec1 {
		fmt.Println("Правильно!")

	} else {
		fmt.Println("Неправильно!")
	}
	c1<-ans

}

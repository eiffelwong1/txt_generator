package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func fatalCheck(e error) {
	if e != nil {
		log.Fatal(e)
	}
}

func main() {
	file, err := os.Create("LongText.txt")
	fatalCheck(err)
	defer file.Close()

	w := bufio.NewWriter(file)
	defer w.Flush()

	fmt.Println("start")

	var i int64 = 0
	for {
		_, err = w.WriteString(fmt.Sprintf("%d %x\n", i, i))
		if err != nil {
			break
		}
		i++
	}

	fmt.Printf("completed, total of %d\n", i)
}

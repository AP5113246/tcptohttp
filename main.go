package main

import (
	"errors"
	"fmt"
	"io"
	"log"
	"os"
)

const inputFilePath = "messages.txt"

func main() {
	// open file and defer close

	f, err := os.Open(inputFilePath)
	if err != nil {
		log.Fatalf("Could not open file titled : %s . Error : %s", inputFilePath, err)

	}
	defer f.Close()

	fmt.Printf("Reading data from %s\n", inputFilePath)
	fmt.Println("=====================================")

	for {
		b := make([]byte,8,8)
		n, err := f.Read(b)
		if err != nil {
			if errors.Is(err, io.EOF){
				break
			}
			fmt.Printf("error: %s\n",err.Error())
			break
		}
		str := string(b[0:n])
		fmt.Printf("read: %v\n",str)
	}
}
package main

import (
	"errors"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

const inputFilePath = "messages.txt"

func main() {
	f, err := os.Open(inputFilePath)
	if err != nil {
		log.Fatalf("Could not open file titled : %s . Error : %s", inputFilePath, err)

	}
	defer f.Close()

	fmt.Printf("Reading data from %s\n", inputFilePath)
	fmt.Println("=====================================")
	newLineContent := ""

	for {
		buffer := make([]byte,8,8)
		n, err := f.Read(buffer)
		if err != nil {
			if newLineContent != "" {
				fmt.Printf("read: %s",newLineContent)
				newLineContent = ""
			}
			if errors.Is(err, io.EOF){
				break
			}
			fmt.Printf("error: %s\n",err.Error())
			break
		}

		str := string(buffer[0:n])
		parts := strings.Split(str,"\n")

		for i := 0; i < len(parts) - 1; i++ {
			fmt.Printf("read: %s%s\n", newLineContent,parts[i])
			newLineContent = ""
		}
		newLineContent += parts[len(parts) - 1]
	}
}
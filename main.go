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

func getLinesChannel(f io.ReadCloser) <-chan string {
	lines := make(chan string)

	go func(){
		newLineContent := ""
		for {
			buffer := make([]byte, 8,8)
			number_bytes_read, err := f.Read(buffer)
			if err != nil {
				if newLineContent != "" {
					lines <- newLineContent
					newLineContent = ""
				}
				if errors.Is(err, io.EOF) {
					break
				}
				fmt.Println("Error in opening file", err)
				break
			}

			str := string(buffer[0:number_bytes_read])
			parts := strings.Split(str, "\n")

			for i := 0; i < len(parts) - 1; i++ {
				lines <- (newLineContent + parts[i])
				newLineContent = ""
			}
			newLineContent += parts[len(parts) - 1]			
		}
		close(lines) // what does close () do?
	}()

	
	return lines 
}

func main() {
	f, err := os.Open(inputFilePath)
	if err != nil {
		log.Fatalf("Could not open file titled : %s . Error : %s", inputFilePath, err)

	}
	defer f.Close()

	fmt.Printf("Reading data from %s\n", inputFilePath)    
	fmt.Println("=====================================")
	
	linesFromChannel := getLinesChannel(f)

	for line := range linesFromChannel {
		fmt.Printf("read: %s\n",line)
	}
}
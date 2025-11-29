package main

import (
	"errors"
	"fmt"
	"io"
	"os"
	"strings"
)


func getLinesChannel(files io.ReadCloser) <- chan string {

	outputChannel := make(chan string)
	go func () {
		bytes_extracted := make([]byte, 8, 8)
		num_bytes_read, err := files.Read(bytes_extracted)
	
	}()

	return outputChannel
}

func main() {

	filepointer, err := os.Open("messages.txt")
	if err != nil {
		fmt.Printf("Unable to open file error: %v", err)
		return 
	}
	
	defer filepointer.Close()
	newLineContent := ""
	
	for {
		bytes_read := make([]byte,8,8)
		number_of_bytes_read, err := filepointer.Read(bytes_read)
		
		if err != nil {
			if newLineContent != "" {
				fmt.Printf("read: %s\n",newLineContent)
			}
			if errors.Is(err, io.EOF) {
				break
			}
			fmt.Printf("Unable to open file error: %w", err)
			return
		}
		
		outputString := string(bytes_read[0:number_of_bytes_read])
		parts := strings.Split(outputString,"\n")
		for i := 0; i < len(parts) - 1; i++ {
			newLineContent += parts[i]
			fmt.Printf("read: %s\n",newLineContent)
			newLineContent = ""
		}
		newLineContent += parts[len(parts) - 1]
	}
}
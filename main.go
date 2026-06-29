package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {

	filePointer, err :=  os.Open("messages.txt")
	if err != nil {
		fmt.Println("Error in reading text file: ", err)
		return 
	}

	defer filePointer.Close()

	linesSlice := make([]byte, 8)
	var currentLine string

	for {
	
		byteCount, err := filePointer.Read(linesSlice)
		if err == io.EOF {
			if currentLine != "" {
				fmt.Printf("read: %s\n", currentLine)
			}
			return 
		}
		
		parts := strings.Split(string(linesSlice[:byteCount]), "\n")
		for _, parts := range parts[:len(parts) - 1] {
			
			currentLine += parts
			fmt.Printf("read: %s\n",currentLine)
			currentLine = ""
		}
		currentLine += parts[len(parts) - 1]

	}
}
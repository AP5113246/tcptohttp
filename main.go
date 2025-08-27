package main

import (
	"io"
	"log"
	"os"
)

const inputFilePath = "messages.txt"

func main() {

	file, err := os.Open(inputFilePath)
	if err != nil {
		log.Fatal("oopsies in opening the file: ",err)
		return 
	}
	defer file.Close()

	buffer := make([]byte, 8, 8)
	var currentLine string

	for {
		number_of_bytes, err := file.Read(buffer)
		if err != nil {
			if err == io.EOF {
				break
			} else {
				log.Fatal("issue reading bytes from file!")
			}
			
		}
		currentRead := string(buffer[:number_of_bytes])
		currentLine += currentRead 
	}
	
}
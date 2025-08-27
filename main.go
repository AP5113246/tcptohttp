package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
)

const inputFilePath = "messages.txt"


func getLinesCahnnel(f io.ReadCloser) <- chan string {
	
}



func main() {

	file, err := os.Open(inputFilePath)
	if err != nil {
		log.Fatal("oopsies in opening the file: ", err)
		return
	}
	defer file.Close()

	str := ""

	for {
		file_data := make([]byte, 8, 8)

		_, err := file.Read(file_data)
		if err != nil {
			if err == io.EOF {
				break
			} else {
				log.Fatal("issue reading bytes from file!")
			}

		}

		// file_data = file_data[:number_of_bytes]
		if i := bytes.IndexByte(file_data, '\n'); i != -1 {
			str += string(file_data[:i])
			file_data = file_data[i+1:]
			fmt.Printf("read: %s\n", str)
			str = ""
		}
		str += string(file_data)
	}

	if len(str) != 0 {
		fmt.Printf("read: %s\n", str)
	}

}

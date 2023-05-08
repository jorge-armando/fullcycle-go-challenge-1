package main

import (
	"fmt"
	"jorge/client"
	"jorge/env"
	"os"
	"time"
)

func main() {
	data := client.GetLastQuote()

	_, err := os.Stat(env.LOG_FILE_NAME)

	if err != nil {
		os.Create(env.LOG_FILE_NAME)
	}

	file, err := os.OpenFile(env.LOG_FILE_NAME, os.O_WRONLY|os.O_APPEND, os.ModeAppend)

	if err != nil {
		panic(err)
	}

	currentTime := time.Now().UTC().String()

	textToWrite := fmt.Sprintf("%s -> Cotacao: %s\n", currentTime, data.CurrentQuote)
	_, err = file.Write([]byte(textToWrite))

	if err != nil {
		panic(err)
	}

	file.Close()

	fmt.Println(textToWrite)
}

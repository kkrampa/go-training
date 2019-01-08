package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

type WordCount struct {
	fileName string
	count int
}

func countWordsInFile(fileName string, c chan<- WordCount) {
	data, err := ioutil.ReadFile(fileName)
	if err != nil {
		panic(err)
	}
	dataStr := string(data)
	splitByNewLine := strings.Split(dataStr, "\n")
	counter := 0
	for _, line := range splitByNewLine {
		splitBySpace := strings.Split(line, " ")
		counter += len(splitBySpace)
	}
	c <- WordCount{fileName: fileName, count: counter}
}

func main() {
	files := os.Args[1:]
	filesNum := len(files)

	if filesNum == 0 {
		fmt.Println("Usage: wc [FILE]...")
		return
	}

	c := make(chan WordCount, filesNum)
	for _, v := range files {
		go countWordsInFile(v, c)
	}

	for i := 0; i < filesNum; i++ {
		wc := <-c
		fmt.Printf("%s %d\n", wc.fileName, wc.count)
	}
}

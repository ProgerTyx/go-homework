package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
)

func main() {
	const root = "./test/task/"

	if err := filepath.Walk(root, walkFunc); err != nil {
		fmt.Printf("Какая-то ошибка: %v\n", err)
	}
}

func walkFunc(path string, info os.FileInfo, err error) error {
	file, err := os.Open(info.Name())
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	r := csv.NewReader(file)

	//record, _ := r.ReadAll()
	//if len(record) == 10 {
	//	fmt.Println(record[4][2])
	//}

	for {
		// Read each record from csv
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("Question: %s Answer %s\n", record[0], record[1])
	}

	//fmt.Printf("Name: %s\tSize: %d byte\tPath: %s\n", info.Name(), info.Size(), path)
	return nil

}

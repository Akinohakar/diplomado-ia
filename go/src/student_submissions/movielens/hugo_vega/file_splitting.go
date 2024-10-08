package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"
)

func SplitBigFile(file_name string, number_of_chunks int, directory string) []string {
	var filesCreated []string
	data := readCSVFile(directory + file_name)
	rowsPerChunk := getRowsPerChunk(len(data), number_of_chunks)
	fmt.Println("Rows in file: ", len(data))
	fmt.Println("Rows per chunk: ", rowsPerChunk)
	i := 0
	for i < number_of_chunks-1 {
		chunkName := "output_file_" + strconv.Itoa(i) + ".csv"
		chunkPath := directory + chunkName
		chunkData := data[(rowsPerChunk * i):(rowsPerChunk * (i + 1))]
		fmt.Println("Will write file: ", chunkName)
		fmt.Println("File will contain: ", len(chunkData), " rows")
		writeCSVFile(chunkData, chunkPath)
		filesCreated = append(filesCreated, chunkName)
		i += 1
	}
	chunkName := "output_file_" + strconv.Itoa(i) + ".csv"
	chunkPath := directory + chunkName
	chunkData := data[rowsPerChunk*i:]
	fmt.Println("Will write file: ", chunkName)
	fmt.Println("File will contain: ", len(chunkData), " rows")
	writeCSVFile(chunkData, chunkPath)
	filesCreated = append(filesCreated, chunkName)
	return filesCreated
}

func getRowsPerChunk(numberOfRows int, numberOfChunks int) int {
	var rowsPerChunk int
	rowsPerChunk = numberOfRows / numberOfChunks
	return rowsPerChunk
}

func readCSVFile(fileName string) [][]string {
	file, err := os.Open(fileName)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	csvReader := csv.NewReader(file)
	data, err := csvReader.ReadAll()
	if err != nil {
		panic(err)
	}
	return data
}

func writeCSVFile(dataToWrite [][]string, fileName string) {
	csvFile, err := os.Create(fileName)
	if err != nil {
		log.Fatalf("failed creating file: %s", err)
	}
	defer csvFile.Close()
	writer := csv.NewWriter(csvFile)
	defer writer.Flush()
	err = writer.WriteAll(dataToWrite)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(len(dataToWrite), " records written to ", fileName)
}

func main() {
	SplitBigFile("ratings.csv", 3, "./test/")
}

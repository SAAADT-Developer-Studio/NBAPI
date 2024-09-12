package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"
)

func convertStrArrToIntArr(strArr []string) []int {
	intArr := []int{}
	for _, i := range strArr {
		j, err := strconv.Atoi(i)
		if err != nil {
			log.Fatalf("Failed to convert string to int: %s", err)
		}
		intArr = append(intArr, j)
	}
	return intArr
}

func main() {

	sqlDump, err := os.Create("dump.sql") 
	if err != nil {
		log.Fatal(err)
	}
	defer sqlDump.Close()

	args := os.Args
	if len(args) < 4 {
		log.Fatalf("Usage: go run idk.go <csvfile> <table_name> <column_indexes>")
	}
	csvFile := args[1]
	tableName := args[2]
	indexArgs := convertStrArrToIntArr(os.Args[3:])

	file, err := os.Open("csv/" + csvFile)
	if err != nil {
		log.Fatalf("Failed to open the file: %s", err)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	reader.Comma = ','
	reader.FieldsPerRecord = -1

	records, err := reader.ReadAll()
	if err != nil {
		log.Fatalf("Error reading CSV: %s", err)
	}

	for _, record := range records {
		currentRecords := []string{}
		query := fmt.Sprintf("INSERT INTO %s VALUES(", tableName)

		for _, fieldIndex := range indexArgs {
			currentRecords = append(currentRecords, fmt.Sprintf("\"%s\"", record[fieldIndex]))
		}

		query += fmt.Sprintf("%s);", joinStrings(currentRecords, ", "))
		_, err := fmt.Fprintln(sqlDump, query)
		if err != nil {
			log.Fatal(err)
		}
	}
}

func joinStrings(arr []string, sep string) string {
	result := ""
	for i, s := range arr {
		if i > 0 {
			result += sep
		}
		result += s
	}
	return result
}

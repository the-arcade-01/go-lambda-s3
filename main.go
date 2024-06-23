package main

import (
	"encoding/csv"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"github.com/the-arcade-01/go-lambda-s3/config"
	"github.com/the-arcade-01/go-lambda-s3/models"
	"github.com/the-arcade-01/go-lambda-s3/repository"
	"github.com/the-arcade-01/go-lambda-s3/utils"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}
}

func main() {
	pool := config.NewThreadPool()
	defer pool.Wait()

	file, err := os.Open(os.Getenv("FILE_NAME"))
	if err != nil {
		panic(err)
	}
	defer file.Close()

	csvReader := csv.NewReader(file)
	err = utils.CheckHeader(csvReader)
	if err != nil {
		panic(err)
	}

	ParseRecordsConcurrent(pool, csvReader)
	// ParseRecordsSequentially(csvReader)
}

// Concurrent Code
func ParseRecordsConcurrent(pool *config.ThreadPool, reader *csv.Reader) {
	start := time.Now()
	for {
		row, err := reader.Read()
		if err != nil {
			if err.Error() == "EOF" {
				break
			}
			log.Println(err)
			continue
		}
		pool.Add(processConcurrent(row))
	}
	log.Printf("Concurrent Time Taken: %v ms\n", time.Since(start).Milliseconds())
}

func processConcurrent(row []string) config.Process {
	return func() {
		csvObject, err := utils.ParseRow(row)
		if err != nil {
			log.Println(err)
			return
		}
		dbObject := models.NewBhavCopyDB(csvObject)
		err = repository.Insert(config.NewDBClient(), dbObject)
		if err != nil {
			log.Println(err)
		}
	}
}

// Sequential Code
func ParseRecordsSequentially(reader *csv.Reader) {
	start := time.Now()
	for {
		row, err := reader.Read()
		if err != nil {
			if err.Error() == "EOF" {
				break
			}
			log.Println(err)
			continue
		}
		processSequentially(row)
	}
	log.Printf("Sequential Time Taken: %v ms\n", time.Since(start).Milliseconds())
}

func processSequentially(row []string) {
	csvObject, err := utils.ParseRow(row)
	if err != nil {
		log.Println(err)
		return
	}
	dbObject := models.NewBhavCopyDB(csvObject)
	err = repository.Insert(config.NewDBClient(), dbObject)
	if err != nil {
		log.Println(err)
	}
}

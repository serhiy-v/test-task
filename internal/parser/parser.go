package parser

import (
	"github.com/gocarina/gocsv"
	"io"
	"log"
	"mime/multipart"
	"os"
	"test-task/internal/repository"
)

func GetData(file multipart.File, header *multipart.FileHeader) []*repository.Transaction {
	tmpfile, err := os.CreateTemp("./", header.Filename)
	defer os.Remove(tmpfile.Name())
	_, err = io.Copy(tmpfile, file)
	if err != nil {
		log.Fatal("Parsing Error")
	}

	return parse(tmpfile)

}

func parse(tmpfile *os.File) []*repository.Transaction {
	f, err := os.Open(tmpfile.Name())
	if err != nil {
		log.Fatal("Unable to read input file ")
	}
	defer f.Close()

	transactions := []*repository.Transaction{}
	if err := gocsv.UnmarshalFile(f, &transactions); err != nil {
		log.Fatal("Parsing Error")
	}

	return transactions
}

package main

import (
	"encoding/csv"
	"io"
	"os"
)

type DictReader struct {
	csvReader *csv.Reader
	header    []string
}

func NewDictReader(csvReader *csv.Reader) (*DictReader, error) {
	var err error
	reader := &DictReader{
		csvReader: csvReader,
	}
	reader.header, err = reader.csvReader.Read()
	if err != nil {
		return nil, err
	}
	return reader, nil
}

func NewDictReaderFromReader(reader io.Reader) (*DictReader, error) {
	csvReader := csv.NewReader(reader)
	return NewDictReader(csvReader)
}

func NewDictReaderFromPath(path string) (*DictReader, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	return NewDictReaderFromReader(file)
}

func (reader *DictReader) Read() (map[string]string, error) {
	row, err := reader.csvReader.Read()
	if err != nil {
		return nil, err
	}
	ret := make(map[string]string, len(reader.header))
	for pos, val := range row {
		ret[reader.header[pos]] = val
	}
	return ret, nil
}

package main

import (
	"bufio"
	"encoding/csv"
	"errors"
	"io"
	"os"
	"testing"
)

const (
	FilePath   = "utf_ken_all.csv"
	FieldCount = 15
)

func Benchmark_Csv_ReuseRecord(b *testing.B) {
	for b.Loop() {
		readCsv(true)
	}
}

func Benchmark_Csv_No_ReuseRecord(b *testing.B) {
	for b.Loop() {
		readCsv(false)
	}
}

func readCsv(reuse bool) error {
	var (
		file *os.File
		err  error
	)
	if file, err = os.Open(FilePath); err != nil {
		return err
	}
	defer file.Close()

	var (
		bufR = bufio.NewReader(file)
		csvR = csv.NewReader(bufR)
	)
	csvR.ReuseRecord = reuse
	csvR.FieldsPerRecord = FieldCount

	var (
		rec []string
	)
	for {
		if rec, err = csvR.Read(); errors.Is(err, io.EOF) {
			break
		}

		if err != nil {
			return err
		}

		for _, field := range rec {
			io.Discard.Write([]byte(field))
		}
	}

	return nil
}

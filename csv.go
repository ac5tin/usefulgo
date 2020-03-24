package utils

import (
	"encoding/csv"
	"io"
	"log"
)

// CSVParser parses csv and runs callback function
func CSVParser(r *csv.Reader, cb func(res []string, chunk int), rows int) error {
	ins := []string{}
	chunk := 0
	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Println(err.Error())
			return err
		}
		// no error, safe to read record
		ins = append(ins, record...)
		// if ins length is ROWSIZE, chunk size reached, make this a chunk (++chunk)
		if len(ins) == rows {
			// run chunk callback function
			go cb(ins, chunk)
			// empty slice
			ins = ins[:0]
			chunk++
		}
	}
	if len(ins) > 0 {
		go cb(ins, chunk)
	}

	log.Printf("finish parsing csv, total chunks : %d\n", chunk)
	return nil
}

// CSVParserH parses csv with headers and runs callback function
func CSVParserH(r *csv.Reader, cb func(res []map[string]string, chunk int), rows int) error {
	ins := []map[string]string{}
	headers := []string{}
	chunk := 0
	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Println(err.Error())
			return err
		}
		// no error, safe to read record
		if len(headers) == 0 {
			headers = record
			continue
		}
		// headers inserted
		tmp := map[string]string{}
		for i := range headers {
			tmp[headers[i]] = record[i]
		}
		ins = append(ins, tmp)
		// if ins length is ROWSIZE, chunk size reached, make this a chunk (++chunk)
		if len(ins) == rows {
			// run chunk callback function
			go cb(ins, chunk)
			// empty slice
			ins = ins[:0]
			chunk++
		}
	}

	if len(ins) > 0 {
		go cb(ins, chunk)
	}

	return nil
}

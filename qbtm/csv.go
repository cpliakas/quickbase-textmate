package qbtm

import (
	"bytes"
	"encoding/csv"
	"errors"
	"io"
)

// ParseCSV parses the CSV file.
func ParseCSV(opts *Opts, parser Parser) error {
	buf := bytes.NewBufferString(opts.File)
	reader := csv.NewReader(buf)

	line := -1
	for {
		line++

		row, err := reader.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			return err
		}

		if len(row) != 8 {
			return errors.New("expected eight columns")
		}

		if line == 0 {
			continue
		}

		if err = parser.Row(row); err != nil {
			return err
		}
	}

	return nil
}

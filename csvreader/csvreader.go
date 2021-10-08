package csvreader

import (
	"bytes"
	"encoding/csv"
	"errors"
	"io"
	"net/http"
	"time"

	"github.com/hashicorp/go-uuid"

	"github.com/foomo/busser/config"

	"github.com/foomo/busser/table"
)

// ReaderConfig is essentially a copy of encoding/csv.Reader and is used to configure it
type ReaderConfig struct {
	// Comma is the field delimiter.
	// It is set to comma (',') by NewReader.
	// Comma must be a valid rune and must not be \r, \n,
	// or the Unicode replacement character (0xFFFD).
	Comma rune

	// Comment, if not 0, is the comment character. Lines beginning with the
	// Comment character without preceding whitespace are ignored.
	// With leading whitespace the Comment character becomes part of the
	// field, even if TrimLeadingSpace is true.
	// Comment must be a valid rune and must not be \r, \n,
	// or the Unicode replacement character (0xFFFD).
	// It must also not be equal to Comma.
	Comment rune

	// FieldsPerRecord is the number of expected fields per record.
	// If FieldsPerRecord is positive, Read requires each record to
	// have the given number of fields. If FieldsPerRecord is 0, Read sets it to
	// the number of fields in the first record, so that future records must
	// have the same field count. If FieldsPerRecord is negative, no check is
	// made and records may have a variable number of fields.
	FieldsPerRecord int

	// If LazyQuotes is true, a quote may appear in an unquoted field and a
	// non-doubled quote may appear in a quoted field.
	LazyQuotes bool

	// If TrimLeadingSpace is true, leading white space in a field is ignored.
	// This is done even if the field delimiter, Comma, is white space.
	TrimLeadingSpace bool

	// ReuseRecord controls whether calls to Read may return a slice sharing
	// the backing array of the previous call's returned slice for performance.
	// By default, each call to Read returns newly allocated memory owned by the caller.
	ReuseRecord bool
}

func (conf *ReaderConfig) configure(csvReader *csv.Reader) {
	if conf == nil {
		return
	}
	csvReader.Comma = conf.Comma
	csvReader.Comment = conf.Comment
	csvReader.FieldsPerRecord = conf.FieldsPerRecord
	csvReader.ReuseRecord = conf.ReuseRecord
	csvReader.LazyQuotes = conf.LazyQuotes
	csvReader.TrimLeadingSpace = conf.TrimLeadingSpace
}

var DefaultConfig = &ReaderConfig{
	Comma: ',',
}

func configureReader(reader *csv.Reader, conf *ReaderConfig) {
	if conf != nil {
		conf.configure(reader)
		return
	}
	if DefaultConfig != nil {
		DefaultConfig.configure(reader)
	}
}

func GetURLTableLoader(id table.ID, csvURL string, conf *ReaderConfig) config.TableLoader {
	return func() (*table.Table, error) {
		t, err := ReadURL(csvURL, conf)
		if err != nil {
			return nil, err
		}
		err = initTable(t, id)
		if err != nil {
			return nil, err
		}
		return t, nil
	}
}

func GetByteTableLoader(id table.ID, csvBytes []byte, conf *ReaderConfig) config.TableLoader {
	return func() (*table.Table, error) {
		t, err := Read(bytes.NewBuffer(csvBytes), conf)
		if err != nil {
			return nil, err
		}
		err = initTable(t, id)
		if err != nil {
			return nil, err
		}
		return t, nil
	}
}

func initTable(t *table.Table, id table.ID) error {
	t.ID = id
	v, err := uuid.GenerateUUID()
	if err != nil {
		return err
	}
	t.Version = table.Version(v)
	return nil
}

func ReadURL(csvURL string, conf *ReaderConfig) (t *table.Table, err error) {
	resp, err := http.Get(csvURL)
	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		return nil, errors.New("unexpected http status code - non 200")
	}
	return Read(resp.Body, conf)
}

func Read(r io.Reader, conf *ReaderConfig) (t *table.Table, err error) {
	csvReader := csv.NewReader(r)
	configureReader(csvReader, conf)
	headerLine := map[int]table.ColumnName{}
	t = &table.Table{
		Timestamp: time.Now().UnixMicro(),
	}
	i := -1
	for {
		i++
		record, lineError := csvReader.Read()
		if lineError == io.EOF {
			if len(headerLine) == 0 {
				return nil, errors.New("file ended before header line")
			}
			return t, nil
		}
		if i == 0 {
			if lineError != nil {
				return t, errors.New("line error in header line: " + lineError.Error())
			}
			for i, colNameString := range record {
				headerLine[i] = table.ColumnName(colNameString)
			}
			continue
		}
		row := table.Row{}
		t.AppendRow(row, lineError)
		for i, colName := range headerLine {
			if len(record) > i {
				row[colName] = record[i]
			} else {
				row[colName] = ""
			}

		}
	}
}

package utils

import (
	"encoding/csv"
	"fmt"
	"strconv"
	"time"

	"github.com/the-arcade-01/go-lambda-s3/models"
)

var colsMapping = []string{
	"SC_CODE",
	"SC_NAME",
	"SC_GROUP",
	"SC_TYPE",
	"OPEN",
	"HIGH",
	"LOW",
	"CLOSE",
	"LAST",
	"PREVCLOSE",
	"NO_TRADES",
	"NO_OF_SHRS",
	"NET_TURNOV",
	"TDCLOINDI",
	"ISIN_CODE",
	"TRADING_DATE",
	"FILLER2",
	"FILLER3",
}

func CheckHeader(reader *csv.Reader) error {
	header, err := reader.Read()
	if err != nil {
		panic(err)
	}

	if len(header) != len(colsMapping) {
		err := fmt.Errorf("mismatch in columns length, CSV: %v, CODE: %v", len(header), len(colsMapping))
		return err
	}

	for i, col := range colsMapping {
		if header[i] != col {
			err := fmt.Errorf("mismatch in columns order, COL NO.: %v, CSV: %v, CODE: %v", i, header[i], col)
			return err
		}
	}
	return nil
}

func ParseRow(row []string) (models.BhavCopyCSV, error) {
	var record models.BhavCopyCSV
	var err error

	// Parse SC_CODE
	record.SC_CODE, err = strconv.Atoi(row[0])
	if err != nil {
		return record, fmt.Errorf("invalid SC_CODE: %v", err)
	}

	// Parse other fields
	record.SC_NAME = row[1]
	record.SC_GROUP = row[2]
	record.SC_TYPE = row[3]

	// Parse OPEN
	record.OPEN, err = strconv.ParseFloat(row[4], 64)
	if err != nil {
		return record, fmt.Errorf("invalid OPEN: %v", err)
	}

	// Parse HIGH
	record.HIGH, err = strconv.ParseFloat(row[5], 64)
	if err != nil {
		return record, fmt.Errorf("invalid HIGH: %v", err)
	}

	// Parse LOW
	record.LOW, err = strconv.ParseFloat(row[6], 64)
	if err != nil {
		return record, fmt.Errorf("invalid LOW: %v", err)
	}

	// Parse CLOSE
	record.CLOSE, err = strconv.ParseFloat(row[7], 64)
	if err != nil {
		return record, fmt.Errorf("invalid CLOSE: %v", err)
	}

	// Parse LAST
	record.LAST, err = strconv.ParseFloat(row[8], 64)
	if err != nil {
		return record, fmt.Errorf("invalid LAST: %v", err)
	}

	// Parse PREVCLOSE
	record.PREVCLOSE, err = strconv.ParseFloat(row[9], 64)
	if err != nil {
		return record, fmt.Errorf("invalid PREVCLOSE: %v", err)
	}

	// Parse NO_TRADES
	record.NO_TRADES, err = strconv.Atoi(row[10])
	if err != nil {
		return record, fmt.Errorf("invalid NO_TRADES: %v", err)
	}

	// Parse NO_OF_SHRS
	record.NO_OF_SHRS, err = strconv.Atoi(row[11])
	if err != nil {
		return record, fmt.Errorf("invalid NO_OF_SHRS: %v", err)
	}

	// Parse NET_TURNOV
	record.NET_TURNOV = row[12]

	// Parse TDCLOINDI
	record.TDCLOINDI = row[13]

	// Parse ISIN_CODE
	record.ISIN_CODE = row[14]

	// Parse TRADING_DATE
	record.TRADING_DATE, err = time.Parse("02-Jan-06", row[15])
	if err != nil {
		return record, fmt.Errorf("invalid TRADING_DATE: %v", err)
	}

	return record, nil
}

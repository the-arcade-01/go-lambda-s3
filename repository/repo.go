package repository

import (
	"database/sql"

	"github.com/the-arcade-01/go-lambda-s3/models"
)

func Insert(client *sql.DB, bc *models.BhavCopyDB) error {
	query := `INSERT INTO bse_bhavcopy (code, name, open, high, low, close, pclose, isin, date) 
VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?) 
ON DUPLICATE KEY UPDATE 
name = VALUES(name), 
open = VALUES(open), 
high = VALUES(high), 
low = VALUES(low), 
close = VALUES(close), 
pclose = VALUES(pclose), 
isin = VALUES(isin), 
date = VALUES(date)`

	_, err := client.Exec(query, bc.Code, bc.Name, bc.Open, bc.High, bc.Low, bc.Close, bc.PClose, bc.Isin, bc.Date)

	if err != nil {
		return err
	}
	return nil
}

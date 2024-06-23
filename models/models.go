package models

import "time"

type BhavCopyDB struct {
	Code   int     `json:"scode"`
	Name   string  `json:"name"`
	Open   float64 `json:"open"`
	High   float64 `json:"hight"`
	Low    float64 `json:"low"`
	Close  float64 `json:"close"`
	PClose float64 `json:"pclose"`
	Isin   string  `json:"isin"`
	Date   string  `json:"date"`
}

type BhavCopyCSV struct {
	SC_CODE      int       `json:"SC_CODE"`
	SC_NAME      string    `json:"SC_NAME"`
	SC_GROUP     string    `json:"SC_GROUP"`
	SC_TYPE      string    `json:"SC_TYPE"`
	OPEN         float64   `json:"OPEN"`
	HIGH         float64   `json:"HIGH"`
	LOW          float64   `json:"LOW"`
	CLOSE        float64   `json:"CLOSE"`
	LAST         float64   `json:"LAST"`
	PREVCLOSE    float64   `json:"PREVCLOSE"`
	NO_TRADES    int       `json:"NO_TRADES"`
	NO_OF_SHRS   int       `json:"NO_OF_SHRS"`
	NET_TURNOV   string    `json:"NET_TURNOV"`
	TDCLOINDI    string    `json:"TDCLOINDI"`
	ISIN_CODE    string    `json:"ISIN_CODE"`
	TRADING_DATE time.Time `json:"TRADING_DATE"`
}

func NewBhavCopyDB(csvObject BhavCopyCSV) *BhavCopyDB {
	bc := new(BhavCopyDB)
	bc.Code = csvObject.SC_CODE
	bc.Name = csvObject.SC_NAME
	bc.Open = csvObject.OPEN
	bc.High = csvObject.HIGH
	bc.Low = csvObject.LOW
	bc.Close = csvObject.CLOSE
	bc.PClose = csvObject.PREVCLOSE
	bc.Isin = csvObject.ISIN_CODE
	bc.Date = csvObject.TRADING_DATE.Format("2006-01-02")
	return bc
}

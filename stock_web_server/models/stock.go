package models

import (
	"time"
	"strings"
	"fmt"
)

type Response struct {
	Code int64
	Msg string
	Data map[string]map[string]interface{}
}

type SortAnalysDayKLins []*AnalysDayKLine

func (a SortAnalysDayKLins) Len() int           { return len(a) }
func (a SortAnalysDayKLins) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a SortAnalysDayKLins) Less(i, j int) bool {
	if a[i].UpCount > a[j].UpCount {
		return true
	}

	return false
}

type SortKLine []*KLine
func (a SortKLine) Len() int           { return len(a) }
func (a SortKLine) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a SortKLine) Less(i, j int) bool {
	if a[i].Date.Before( a[j].Date) {
		return true
	}
	return false
}

type AnalysDayKLine struct {
	Stock *Stock
	Days int64
	RedCount int64
	GreenCount int64
	UpCount int64
	DownCount int64
	UpRateCount float64
	DownRateCount float64
	UpDownRateTotal float64
}

type Stock struct {
	StockId int64       `orm:"pk;auto"`
	Code string
	SyncTime time.Time
	SyncOk bool
	
	Infos []*StockInfo	`orm:"-"`
	DeltaVal float64
}

func (s *Stock)CodeStr() string {
	if s == nil {
		return ""
	}
	if strings.HasPrefix(s.Code, "60") {
		return fmt.Sprintf("sh%s",s.Code)
	} else if strings.HasPrefix(s.Code, "00") || strings.HasPrefix(s.Code, "30") {
		return fmt.Sprintf("sz%s",s.Code)
	}

	return s.Code
}

type KLine struct {
	KLineId int64       `orm:"pk;auto"`
	StockId int64
	OpeningPrice float64
	ClosingPrice float64
	MaxPrice float64
	MinPrice float64
	Date time.Time
	Vol float64 	//万手
	Type int		//1 日K 2 周K 3月K 4年K
}

func (k *KLine)IsRed() bool {
	if k.ClosingPrice > k.OpeningPrice {
		return true
	}
	return false
}

func (k *KLine)GetAddRate(last *KLine) float64 {
	if last == nil {
		return 0
	}
	return (k.ClosingPrice - last.ClosingPrice)/last.ClosingPrice
}

type StockInfo struct {
	StockInfoId int64 `orm:"pk;auto"`
	StockId int64
	MainIn float64
	MainOut float64
	MainTotal float64
	RetailIn float64
	RetailOut float64
	RetailTotal float64
	Date time.Time
}

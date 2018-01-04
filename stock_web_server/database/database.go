package database

import (
	"github.com/cheneylew/goutil/utils/beego"
	_ "github.com/go-sql-driver/mysql"
	"github.com/cheneylew/goutil/stock_web_server/models"
	"math"
	"time"
)

var DB DataBase

func init() {
	//db := beego.InitRegistDB("cheneylew","12344321","47.91.151.207","3308","stock")
	//db := beego.InitRegistDB("root","cnldj1988","127.0.0.1","3306","stock")
	db := beego.InitRegistDB("root","cnldj1988","127.0.0.1","13306","stock")
	db.Orm.Using("default")
	DB = DataBase{
		BaseDataBase:*db,
	}

}

type DataBase struct {
	beego.BaseDataBase
}

func (db *DataBase)GetStockWithCode(code string) *models.Stock {
	var objects []*models.Stock

	qs := db.Orm.QueryTable("Stock")
	_, err := qs.Filter("Code", code).RelatedSel().All(&objects)

	if err == nil && len(objects) > 0 {
		return objects[0]
	}

	return nil
}

func (db *DataBase)GetStockWithCodePrefix(prefix string) []*models.Stock {
	var objects []*models.Stock

	qs := db.Orm.QueryTable("Stock")
	_, err := qs.Filter("Code__istartswith", prefix).Limit(math.MaxInt32, 0).RelatedSel().All(&objects)

	if err == nil && len(objects) > 0 {
		return objects
	}

	return nil
}

func (db *DataBase)GetNotSyncStocksWithCodePrefix(prefix string) []*models.Stock {
	var objects []*models.Stock

	qs := db.Orm.QueryTable("Stock")
	_, err := qs.Filter("Code__istartswith", prefix).Filter("SyncOk__isnull", true).Limit(math.MaxInt32, 0).RelatedSel().All(&objects)

	if err == nil && len(objects) > 0 {
		return objects
	}

	return nil
}

func (db *DataBase)GetSyncFailedStocks() []*models.Stock {
	var objects []*models.Stock

	qs := db.Orm.QueryTable("Stock")
	_, err := qs.Filter("SyncOk", 0).RelatedSel().All(&objects)

	if err == nil && len(objects) > 0 {
		return objects
	}

	return nil
}

func (db *DataBase)GetKLineAll() []*models.KLine {
	var objects []*models.KLine

	qs := db.Orm.QueryTable("k_line")
	_, err := qs.Limit(math.MaxInt32, 0).All(&objects)
	if err != nil {
		return nil
	}
	return objects
}

func (db *DataBase)GetKLineAllForStock(stock *models.Stock) []*models.KLine {
	var objects []*models.KLine

	qs := db.Orm.QueryTable("k_line")
	_, err := qs.Filter("StockId", stock.StockId).Limit(math.MaxInt32, 0).All(&objects)
	if err != nil {
		return nil
	}
	return objects
}

func (db *DataBase)GetKLineAllForStockCode(code string) []*models.KLine {
	stock := db.GetStockWithCode(code)
	if stock == nil {
		return nil
	}

	var objects []*models.KLine

	qs := db.Orm.QueryTable("k_line")
	_, err := qs.Filter("StockId", stock.StockId).Limit(math.MaxInt32, 0).All(&objects)
	if err != nil {
		return nil
	}
	return objects
}

func (db *DataBase)GetKLineAllForStockCodeAndDays(code string, days int) []*models.KLine {
	stock := db.GetStockWithCode(code)
	if stock == nil {
		return nil
	}

	var objects []*models.KLine

	date := time.Now().Add(-time.Hour * 24*time.Duration(days))
	qs := db.Orm.QueryTable("k_line")
	_, err := qs.Filter("StockId", stock.StockId).Filter("Date__gte", date).Limit(math.MaxInt32, 0).All(&objects)
	if err != nil {
		return nil
	}
	return objects
}

func (db *DataBase)GetUser() *models.User {
	return nil
}

func (db *DataBase)GetStockInfoAll() []*models.StockInfo {
	var objects []*models.StockInfo

	qs := db.Orm.QueryTable("stock_info")
	_, err := qs.Limit(math.MaxInt32, 0).All(&objects)
	if err != nil {
		return nil
	}
	return objects
}


func (db *DataBase)GetStockInfoAllForStock(stock *models.Stock) []*models.StockInfo {
	var objects []*models.StockInfo

	qs := db.Orm.QueryTable("stock_info")
	_, err := qs.Filter("StockId", stock.StockId).Limit(math.MaxInt32, 0).All(&objects)
	if err != nil {
		return nil
	}
	return objects
}
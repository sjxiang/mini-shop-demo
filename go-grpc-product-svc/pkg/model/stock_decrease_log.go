package model

type StockDecreaseLog struct {
	Id           int64  `json:"id" gorm:"primaryKey"`
	OrderId      int64  `json:"order_id"`
	ProductRefer int64  `json:"product_id"`
}

func (s *StockDecreaseLog) TableName() string {
	return "stock_decrease_log"
}
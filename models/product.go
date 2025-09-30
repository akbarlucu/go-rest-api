package models

type Product struct {
	Id    int64   `gorm:"primaryKey" json:"id"`
	Nama  string  `gorm:"type:varchar(300)" json:"nama"`
	Stock int32   `gorm:"type:int(5)" json:"stock"`
	Harga float64 `gorm:"type:decimal(14,2)" json:"harga"`
}

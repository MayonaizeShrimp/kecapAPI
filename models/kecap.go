package models

type Kecap struct {
	ID    uint   `json:"id" gorm:"primary_key"`
	Nama  string `json:"nama"`
	Harga int    `json:"harga"`
}

type KecapInsertData struct {
	Nama  string `json:"nama" binding:"required"`
	Harga int    `json:"harga" binding:"required"`
}

type KecapUpdateData struct {
	Nama  string `json:"nama"`
	Harga int    `json:"harga"`
}

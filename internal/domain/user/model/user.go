package model

type User struct {
	Nama       string `json:"nama"`
	Alamat     string `json:"alamat"`
	Umur       string `json:"umur"`
	Pendidikan string `json:"pendidikan"`
}

type ListUser struct {
	Nama        string `gorm:"nama"`
	Alamat      string `gorm:"alamat"`
	Pendidikan  string `gorm:"pendidikan"`
	FilterCount int    `gorm:"count"`
}

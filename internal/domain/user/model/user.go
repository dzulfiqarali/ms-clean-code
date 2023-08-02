package model

type User struct {
	Nama       string `json:"nama"`
	Alamat     string `json:"alamat"`
	Umur       string `json:"umur"`
	Pendidikan string `json:"pendidikan"`
}

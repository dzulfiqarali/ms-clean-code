package fakeapi

type RequestFakeAPI struct {
	Tittle      string  `json:"title"`
	Price       float32 `json:"price"`
	Description string  `json:"description"`
	Image       string  `json:"image"`
	Category    string  `json:"category"`
}

type ResponseFakeAPI struct {
	Id          int     `json:"id"`
	Tittle      string  `json:"title"`
	Price       float32 `json:"price"`
	Description string  `json:"description"`
	Image       string  `json:"image"`
	Category    string  `json:"category"`
}

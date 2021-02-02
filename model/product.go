package model
type Product struct {
	Id int `json:"id"`
	Name string `json:"name"`
	BrandDetail Brand `json:"brand"`
}

type RetErr struct{
	StCode int `json:"st_code"`
	Errmessage string `json:"errmessage"`
}

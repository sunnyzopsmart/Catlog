package service

import "Catlog/model"

type Product interface {
	GetById(id int)  (model.NewProduct,error)
	InsertProductBrand(pName string,bName string) (model.NewProduct,error)
}
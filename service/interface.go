package service

import "Catlog/model"

type Product interface {
	GetById(id int)  (model.Product,error)
	InsertProductBrand(pName string,bName string) (model.Product,error)
}
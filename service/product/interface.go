package product

import "Catlog/model"

type Product interface {
	GetById(id int)  (model.NewProduct,error)
	InsertProductBrand(pName string,bName string) (model.Product,error)
}
package store

import "Catlog/model"

type Store interface {
	GetProductById(id int)  (model.Product,error)
	InsertProduct(p model.Product) (int,error)
}

type Brand interface {
	GetBrandById(id int)  (model.Brand,error)
	InsertBrand(p model.Brand) (int,error)
	GetBrandByName(name string)  (model.Brand,error)
}
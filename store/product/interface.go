package product

import "Catlog/model"

type Store interface {
	GetById(id int)  (model.Product,error)
	InsertProduct(p model.Product) (int,error)
}

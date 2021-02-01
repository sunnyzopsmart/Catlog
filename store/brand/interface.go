package brand

import "Catlog/model"

type Store interface {
	GetById(id int)  (model.Brand,error)
	InsertBrand(p model.Brand) (int,error)
	GetByName(name string)  (model.Brand,error)
}

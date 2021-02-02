package product

import (
	"Catlog/model"
	"Catlog/service"
	"Catlog/store"
	"fmt"
)

type ProductBrand struct {
	p store.Store
	b store.Brand
}

func New(p store.Store,b store.Brand) service.Product {
	return &ProductBrand{p,b}
}
func (pb *ProductBrand) GetById(id int) (model.Product,error){
	var newProduct model.Product
	res,err := pb.p.GetProductById(id)
	if err!=nil {
		fmt.Println(err)
		return newProduct,err
	}

	fmt.Println(res)
	newProduct.Id,newProduct.Name = res.Id,res.Name
	id = res.BrandDetail.Id
	newProduct.BrandDetail.Id = id
	bres,err := pb.b.GetBrandById(id)

	fmt.Println(bres)
	newProduct.BrandDetail.Name = bres.Name
	return newProduct, nil
}

func (pb *ProductBrand) InsertProductBrand(pName string,bName string) (model.Product,error){
	brr,err := pb.b.GetBrandByName(bName)
	br := model.Brand{brr.Id,bName}
	if err!=nil{
		br.Id,err = pb.b.InsertBrand(br)
		if err!=nil{
			return model.Product{},err
		}
	}
	pr := model.Product{0,pName,model.Brand{br.Id,bName}}
	var npr model.Product
	pr.Id,err = pb.p.InsertProduct(pr)
	if err!=nil{
		return npr,err
	}
	npr = model.Product{pr.Id,pName,model.Brand{br.Id,bName}}
	fmt.Println(npr)
	return npr,nil
}

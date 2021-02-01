package product

import (
	"Catlog/model"
	"Catlog/store/brand"
	"Catlog/store/product"
	"fmt"
)

type ProductBrand struct {
	p product.Store
	b brand.Store
}

func New(p product.Store,b brand.Store) Product{
	return &ProductBrand{p,b}
}
func (pb *ProductBrand) GetById(id int) (model.NewProduct,error){
	var newProduct model.NewProduct
	res,err := pb.p.GetById(id)
	if err!=nil {
		fmt.Println(err)
		return newProduct,err
	}

	fmt.Println(res)
	newProduct.Id,newProduct.Name = res.Id,res.Name
	id = res.BId
	bres,err := pb.b.GetById(id)

	fmt.Println(bres)
	newProduct.BidName = bres.Name
	return newProduct, nil
}

func (pb *ProductBrand) InsertProductBrand(pName string,bName string) (model.Product,error){
	brr,err := pb.b.GetByName(bName)
	br := model.Brand{brr.Id,bName}
	if err!=nil{
		br.Id,err = pb.b.InsertBrand(br)
		if err!=nil{
			return model.Product{},err
		}
	}
	pr := model.Product{0,pName,br.Id}
	pr.Id,err = pb.p.InsertProduct(pr)
	if err!=nil{
		return model.Product{},err
	}
	fmt.Println(pr)
	return pr,nil
}

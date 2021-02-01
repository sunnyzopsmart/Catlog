package product

import (
	"Catlog/model"
	"Catlog/store/brand"
	"Catlog/store/product"
	"errors"
	_ "github.com/go-sql-driver/mysql"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"reflect"
	"testing"
)


var pd = []model.Product{
	{1,"Mouse",1},
	{2,"Laptop",2},
}

var bd = []model.Brand{
	{1,"Dell"},
	{2,"Lenevo"},
}

var npd = []model.NewProduct{
	{1,"Mouse","Dell"},
	{2,"Laptop","Lenevo"},
}
func TestGetById(t *testing.T)  {

	testCases := []struct{
		id int
		exp model.NewProduct
		expectedProdStore model.Product
		expectedBrandStore model.Brand
		err error
	}{
		{1,npd[0],pd[0], bd[0],nil},
		{2,npd[1],pd[1],bd[1],nil},
		{3,model.NewProduct{},model.Product{}, model.Brand{},model.Err{3}},
	}
	//fmt.Println(testCases)
	ctrl := gomock.NewController(t)
	ps := product.NewMockStore(ctrl)
	bs := brand.NewMockStore(ctrl)
	pserv := New(ps,bs)

	for  _,k := range testCases{
		ps.EXPECT().GetById(k.id).Return(k.expectedProdStore,k.err)
		if k.err==nil {
			bs.EXPECT().GetById(k.id).Return(k.expectedBrandStore,k.err)
		}
		np,err := pserv.GetById(k.id)
		//fmt.Println(err)
		if err!=nil{
			assert.Error(t, err,k.err)
			if !reflect.DeepEqual(err, k.err) {
				t.Errorf("Error, got: %v, want: %v.", err, k.err)
			}
		} else {
			//assert.Equal(t, np, k.exp)
			if !reflect.DeepEqual(np, k.exp) {
				t.Errorf("Error, got: %v, want: %v.", np, k.exp)
			}
		}
	}

}


func TestInsert(t *testing.T){
	testCases := []struct{
		pname string
		bname string
		prod model.Product
		bran model.Brand
		branerr bool
		expecProductId int
		expecBrandId int
		experr error
	}{
		{pd[0].Name,bd[0].Name,model.Product{0,pd[0].Name,pd[0].BId},bd[0],true,1,1,nil},
		{pd[1].Name,bd[1].Name,model.Product{0,pd[1].Name,pd[1].BId},bd[1],true,2,2,nil},
		{"Laptop","Hp",model.Product{0,"Laptop",3},model.Brand{0,"Hp"},false,3,3,errors.New("Brand not Found!")},
	}

	ctrl := gomock.NewController(t)
	ps := product.NewMockStore(ctrl)
	bs := brand.NewMockStore(ctrl)
	pserv := New(ps,bs)
	for  _,k := range testCases{
		if !k.branerr {
			bs.EXPECT().GetByName(k.bname).Return(model.Brand{},k.experr)
			bs.EXPECT().InsertBrand(k.bran).Return(k.expecBrandId, nil)
		}else {
			bs.EXPECT().GetByName(k.bname).Return(k.bran,k.experr)
		}
		ps.EXPECT().InsertProduct(k.prod).Return(k.expecProductId,nil)
		mp,err := pserv.InsertProductBrand(k.pname,k.bname)
		if err!=nil{
			//fmt.Println(err)
			if !reflect.DeepEqual(err, k.experr) {
				t.Errorf("Error, got: %v, want: %v.", err, k.experr)
			}
		} else {
			if !reflect.DeepEqual(mp.Id, k.expecProductId){
				t.Errorf("Error, got: %v, want: %v.", mp.Id, k.expecProductId)
			}
		}
	}
}
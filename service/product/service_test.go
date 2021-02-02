package product

import (
	"Catlog/model"
	"Catlog/store"
	"errors"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"reflect"
	"testing"
)


var pd = []model.Product{
	{1,"Mouse",model.Brand{1,"Dell"}},
	{2,"Laptop",model.Brand{2,"Lenevo"}},
}

var bd = []model.Brand{
	{1,"Dell"},
	{2,"Lenevo"},
}


func TestGetById(t *testing.T)  {

	testCases := []struct{
		id int
		expectedProdStore model.Product
		expectedBrandStore model.Brand
		err error
	}{
		{1,pd[0], bd[0],nil},
		{2,pd[1],bd[1],nil},
		{3,model.Product{}, model.Brand{},model.Err{3}},
	}
	//fmt.Println(testCases)
	ctrl := gomock.NewController(t)
	ps := store.NewMockStore(ctrl)
	bs := store.NewMockBrand(ctrl)
	pserv := New(ps,bs)

	for  _,k := range testCases{
		ps.EXPECT().GetProductById(k.id).Return(k.expectedProdStore,k.err)
		if k.err==nil {
			bs.EXPECT().GetBrandById(k.id).Return(k.expectedBrandStore,k.err)
		}
		np,err := pserv.GetById(k.id)
		if err!=nil{
			assert.Error(t, err,k.err)
			if !reflect.DeepEqual(err, k.err) {
				t.Errorf("Error, got: %v, want: %v.", err, k.err)
			}
		} else {
			if !reflect.DeepEqual(np, k.expectedProdStore) {
				t.Errorf("Error, got: %v, want: %v.", np, k.expectedProdStore)
			}
		}
	}

}


func TestInsert(t *testing.T){
	testCases := []struct{
		pname string
		bname string
		prod model.Product
		branerr bool
		expecProductId int
		expecBrandId int
		experr error
	}{
		{pd[0].Name,bd[0].Name,model.Product{0,pd[0].Name,pd[0].BrandDetail},true,1,1,nil},
		{pd[1].Name,bd[1].Name,model.Product{0,pd[1].Name,pd[1].BrandDetail},true,2,2,nil},
		{"Laptop","Hp",model.Product{0,"Laptop",model.Brand{3,"Hp"}},false,3,3,errors.New(fmt.Sprintf(model.BrandNotFound,"Hp"))},
	}

	ctrl := gomock.NewController(t)
	ps := store.NewMockStore(ctrl)
	bs := store.NewMockBrand(ctrl)
	pserv := New(ps,bs)
	for  _,k := range testCases{
		if !k.branerr {
			bs.EXPECT().GetBrandByName(k.bname).Return(model.Brand{},k.experr)
			bs.EXPECT().InsertBrand(model.Brand{0,k.prod.BrandDetail.Name}).Return(k.expecBrandId, nil)
		}else {
			bs.EXPECT().GetBrandByName(k.bname).Return(k.prod.BrandDetail,k.experr)
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
package product

import (
	"Catlog/model"
	"Catlog/service"
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/golang/mock/gomock"
	"github.com/gorilla/mux"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"
)

var pd = []model.Product{
	{1,"Mouse",model.Brand{1,"Dell"}},
	{2,"Laptop",model.Brand{2,"Lenevo"}},
}
var pbrand = []ProductBrand{
	{"Mouse","Dell"},
}
func TestHandler_ProductWithId(t *testing.T) {
	testCases := []struct{
		id string
		expectednpd model.Product
		stCode int
		err error
	}{
		{"1",pd[0],http.StatusOK,nil},
		{"2",pd[1],http.StatusOK,nil},
		{"3",model.Product{},http.StatusBadRequest,errors.New(fmt.Sprintf(model.ProductNotAvailable,3))},
		{"abcd",model.Product{},http.StatusInternalServerError,errors.New(model.InvalidId)},
	}
	ctrl := gomock.NewController(t)
	serv := service.NewMockProduct(ctrl)
	h := Handler{serv}
	for _,testCase := range testCases{

		link := "/product/%s"
		r := httptest.NewRequest("GET",fmt.Sprintf(link,testCase.id),nil)
		w  := httptest.NewRecorder()
		//router := mux.NewRouter()
		r = mux.SetURLVars(r, map[string]string{
			"id": testCase.id,
		})
		id,err := strconv.Atoi(testCase.id)
		if err == nil {
			serv.EXPECT().GetById(id).Return(testCase.expectednpd,testCase.err)
		}
		h.ProductWithId(w,r)
		if w.Code != testCase.stCode {
			t.Fatalf("ProductWithId() = %v , want %v", w.Code,testCase.stCode)
		}
	}
}

func TestHandler_InsertProduct(t *testing.T) {
	testCases := []struct{
		pb ProductBrand
		expectedpd model.Product
		stcode int
		err error
	}{
		{pbrand[0],pd[0],http.StatusCreated,nil},
	}

	ctrl := gomock.NewController(t)
	serv := service.NewMockProduct(ctrl)
	h := Handler{serv}
	for _,testCase := range testCases{
		l, _ := json.Marshal(testCase.pb)
		m := bytes.NewBuffer(l)
		link := "/product"
		r := httptest.NewRequest("POST",link,m)
		w  := httptest.NewRecorder()
		serv.EXPECT().InsertProductBrand(testCase.pb.Pname,testCase.pb.Bname).Return(testCase.expectedpd,testCase.err)
		h.InsertProduct(w,r)
		if w.Code != testCase.stcode {
			t.Fatalf("InsertProduct() = %v , want %v", w.Code,testCase.stcode)
		}
	}

}


func TestHandler_InsertBrandErr(t *testing.T) {
	testCases := []struct{
		pb string
		expectedpd model.Product
		stcode int
		err error
	}{
		{`{"name": "}`,model.Product{},http.StatusInternalServerError,nil},
	}

	ctrl := gomock.NewController(t)
	serv := service.NewMockProduct(ctrl)
	h := Handler{serv}
	for _,testCase := range testCases{
		l, _ := json.Marshal(testCase.pb)
		m := bytes.NewBuffer(l)
		link := "/product"
		r := httptest.NewRequest("POST",link,m)
		w  := httptest.NewRecorder()
		h.InsertProduct(w,r)
		if w.Code != testCase.stcode {
			t.Fatalf("InsertProduct() = %v , want %v", w.Code,testCase.stcode)
		}
	}
}

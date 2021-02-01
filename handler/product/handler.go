package product

import (
	"Catlog/model"
	"Catlog/service"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

type Handler struct{
	Sev service.Product
}

func (serv Handler) ProductWithId(w http.ResponseWriter,r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	v := params["id"]

	id,er:= strconv.Atoi(v)

	//fmt.Println(id)
	if er != nil {
		herr := model.RetErr{http.StatusInternalServerError,model.InvalidId}
		w.WriteHeader(http.StatusInternalServerError)
		res, er := json.Marshal(herr)
		if er == nil {
			w.Write(res)
		}
		//http.Error(w, "Invalid Id!", http.StatusInternalServerError)
		return
	}
	productData,err := serv.Sev.GetById(id)
	if err!=nil {
		herr := model.RetErr{http.StatusBadRequest,fmt.Sprintf(model.ProductNotAvailable,id)}
		res, er := json.Marshal(herr)
		w.WriteHeader(http.StatusBadRequest)
		if er == nil {
			w.Write(res)
		}
		return
	}
	res, er := json.Marshal(productData)
	if er == nil {
		w.Write(res)
	}
}

type ProductBrand struct {
	Pname string `json:"pname"`
	Bname string `json:"bname"`
}
var pb ProductBrand
func (serv *Handler) InsertProduct(w http.ResponseWriter,r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	body := r.Body
	err := json.NewDecoder(body).Decode(&pb)
	if err != nil {
		herr := model.RetErr{http.StatusInternalServerError,model.InvalidJson}
		res, er := json.Marshal(herr)
		w.WriteHeader(http.StatusInternalServerError)
		if er == nil {
			w.Write(res)
		}
		return
	}
	prod,err := serv.Sev.InsertProductBrand(pb.Pname,pb.Bname)
	res,er :=json.Marshal(prod)
	if er == nil {
		w.WriteHeader(http.StatusCreated)
		w.Write(res)
	}
}

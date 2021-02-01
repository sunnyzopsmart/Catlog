package product

import (
	"Catlog/service/product"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"strconv"
)

type Handler struct{
	Sev product.Product
}

func (serv Handler) ProductWithId(w http.ResponseWriter,r *http.Request) {

	if r.Method == "GET" || r.Method == "PUT" {
		w.Header().Set("Content-Type", "application/json")
		params := mux.Vars(r)
		v := params["id"]

		id,er:= strconv.Atoi(v)

		//fmt.Println(id)
		if er != nil {
			http.Error(w, "Invalid Id!", http.StatusInternalServerError)
			return
		}
		productData,err := serv.Sev.GetById(id)
		if err!=nil {
			http.Error(w,fmt.Sprintf("Product with id %v is not available!",id),http.StatusBadRequest)
			return
		}
		res, er := json.Marshal(productData)
		if er == nil {
			w.Write(res)
		}

	}
}

type ProductBrand struct {
	Pname string `json:"pname"`
	Bname string `json:"bname"`
}
var pb ProductBrand
func (serv *Handler) InsertProduct(w http.ResponseWriter,r *http.Request){
	if r.Method=="POST"{
		body := r.Body
		err := json.NewDecoder(body).Decode(&pb)
		if err != nil {
			log.Print(err)
		}
		x,y := pb.Pname,pb.Bname
		fmt.Println(x,y)
		prod,err := serv.Sev.InsertProductBrand(pb.Pname,pb.Bname)
		res,er :=json.Marshal(prod)
		if er == nil {
			w.WriteHeader(http.StatusCreated)
			w.Write(res)
		}
	}
}

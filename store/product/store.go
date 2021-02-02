package product

import (
	"Catlog/model"
	"database/sql"
	"errors"
	"fmt"
	"Catlog/store"
)

type dbStore struct {
	db *sql.DB
}
func New(db *sql.DB) store.Store{
	return &dbStore{db}
}
func (s *dbStore) GetProductById(id int) (model.Product,error){
	var pd model.Product
	if id < 1{
		return pd,errors.New(model.NegativeId)
	}
	dis,err := s.db.Query("SELECT id, name, bid FROM product WHERE id = ?",id)
	//fmt.Println(err)
	if err != nil || !dis.Next(){
		return pd,model.Err{id}
	}
	defer dis.Close()
	err = dis.Scan(&pd.Id,&pd.Name,&pd.BrandDetail.Id)
	return pd,nil
}

func (s *dbStore) InsertProduct(p model.Product) (int,error){
	sq := "INSERT INTO product(name, bid) VALUES (?,?)"
	res, err := s.db.Exec(sq,p.Name,p.BrandDetail.Id)
	if err != nil {
		return -1,errors.New(model.SQLProblem)
	}
	lastId, err := res.LastInsertId()
	if err != nil {
		return -1,errors.New(model.LastInsertId)
	}
	fmt.Printf("The last inserted row id: %d\n", lastId)
	return int(lastId),nil
}
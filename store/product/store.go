package product

import (
	"Catlog/model"
	"database/sql"
	"errors"
	"fmt"
)

type dbStore struct {
	db *sql.DB
}
func New(db *sql.DB) Store{
	return &dbStore{db}
}
func (s *dbStore) GetById(id int) (model.Product,error){
	dis,err := s.db.Query("SELECT id, name, bid FROM product WHERE id = ?",id)
	var pd model.Product
	if err != nil || !dis.Next(){
		return pd,model.Err{id}
	}
	defer dis.Close()
	err = dis.Scan(&pd.Id,&pd.Name,&pd.BId)
	return pd,nil
}

func (s *dbStore) InsertProduct(p model.Product) (int,error){
	sq := "INSERT INTO product(name, bid) VALUES (?,?)"
	res, err := s.db.Exec(sq,p.Name,p.BId)
	if err != nil {
		return -1,errors.New("Problem in Executing Query!")
	}
	lastId, err := res.LastInsertId()
	if err != nil {
		return -1,err
	}
	fmt.Printf("The last inserted row id: %d\n", lastId)
	return int(lastId),nil
}
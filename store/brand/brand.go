package brand

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


func New(db *sql.DB) store.Brand{
	return &dbStore{db}
}


func (s *dbStore) GetBrandById(id int) (model.Brand,error){
	var bd model.Brand
	if id < 1{
		return bd,errors.New(model.NegativeId)
	}
	db := s.db
	dis,err := db.Query("SELECT id, name FROM brand WHERE id = ?",id)

	if err != nil || !dis.Next(){
		return bd,model.Err{id}
	}
	defer dis.Close()
	dis.Scan(&bd.Id,&bd.Name)
	return bd,nil
}



func (s *dbStore) InsertBrand(p model.Brand) (int,error){
	bname := p.Name
	sq := "INSERT INTO brand(name) VALUES (?)"
	res, err := s.db.Exec(sq,bname)
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

func (s *dbStore) GetBrandByName(name string) (model.Brand,error){
	db := s.db
	dis,err := db.Query("SELECT id, name FROM brand WHERE name = ?",name)

	var bd model.Brand
	if err != nil || !dis.Next(){
		return bd,errors.New(fmt.Sprintf(model.BrandNotFound,name))
	}
	defer dis.Close()
	dis.Scan(&bd.Id,&bd.Name)
	return bd,nil
}

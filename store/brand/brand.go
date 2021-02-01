package brand

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


func (s *dbStore) GetById(id int) (model.Brand,error){
	db := s.db
	dis,err := db.Query("SELECT id, name FROM brand WHERE id = ?",id)
	var bd model.Brand
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
		return -1,errors.New("Problem in Excecuting command!")
	}
	lastId, err := res.LastInsertId()
	if err != nil {
		return -1,err
	}
	fmt.Printf("The last inserted row id: %d\n", lastId)
	return int(lastId),nil
}

func (s *dbStore) GetByName(name string) (model.Brand,error){
	db := s.db
	dis,err := db.Query("SELECT id, name FROM brand WHERE name = ?",name)

	var bd model.Brand
	if err != nil || !dis.Next(){
		return bd,errors.New("Brand not Found!")
	}
	defer dis.Close()
	dis.Scan(&bd.Id,&bd.Name)
	return bd,nil
}

package brand

import (
	"Catlog/model"
	"errors"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"reflect"
	"testing"
	"fmt"
)

var u = []model.Brand{{
	Id:     1,
	Name:   "Dell",
},
	{
		Id:     2,
		Name:   "Hp",
	},
}

func TestFindByID(t *testing.T) {
	db, mock,err := sqlmock.New()
	if err!=nil {
		fmt.Sprintf(model.SQLMockError)
		return
	}
	dbHandler := New(db)
	defer db.Close()

	testCases := []struct{
		id int
		expbrand model.Brand
		err error
	}{
		{1,u[0],nil},
		{2,u[1],nil},
	}

	query := "SELECT id, name FROM brand WHERE id = ?"

	for _,testCase := range testCases{
		rows := sqlmock.NewRows([]string{"id", "name"}).
			AddRow(testCase.expbrand.Id, testCase.expbrand.Name)
		mock.ExpectQuery(query).WithArgs(testCase.id).WillReturnRows(rows)
		res,_ := dbHandler.GetBrandById(testCase.id)

		if !reflect.DeepEqual(res, testCase.expbrand){
			t.Errorf("Product result was incorrect, got: %v, want: %v.", res, testCase.expbrand)
		}
	}
}

func TestFindByIDError(t *testing.T){
	db, mock,err := sqlmock.New()
	if err!=nil {
		fmt.Sprintf(model.SQLMockError)
		return
	}
	dbHandler := New(db)
	defer db.Close()

	testCases := []struct{
		id int
		expbrand model.Brand
		err error
	}{
		{3,model.Brand{},model.Err{3}},
		{-3,model.Brand{},errors.New(model.NegativeId)},
	}
	query := "SELECT id, name FROM brand WHERE id = ?"

	for _,testCase := range testCases {
		mock.ExpectQuery(query).WithArgs(testCase.id).WillReturnError(testCase.err)
		_, err := dbHandler.GetBrandById(testCase.id)
		fmt.Println(err)
		if !reflect.DeepEqual(err, testCase.err) {
			t.Errorf("Error, got: %v, want: %v.", err, testCase.err)
		}
	}
}

func TestInsert(t *testing.T){
	db,mock,err := sqlmock.New()
	dbHandler := New(db)
	query := "INSERT INTO brand"

	mock.ExpectExec(query).WithArgs(u[0].Name).WillReturnResult(sqlmock.NewResult(1,1))

	lastID,err := dbHandler.InsertBrand(u[0])
	assert.Equal(t, u[0].Id,lastID)
	assert.NoError(t, err)
}

func TestFindByName(t *testing.T) {
	db, mock,err := sqlmock.New()
	dbHandler := New(db)
	defer db.Close()

	query := "SELECT id, name FROM brand WHERE name = ?"

	rows := sqlmock.NewRows([]string{"id", "name"}).
		AddRow(u[0].Id, u[0].Name)
	//fmt.Println(rows)
	mock.ExpectQuery(query).WithArgs(u[0].Name).WillReturnRows(rows)
	res,err := dbHandler.GetBrandByName(u[0].Name)

	if !reflect.DeepEqual(res, u[0]){
		t.Errorf("Brand result was incorrect, got: %v, want: %v.", res, u[0])
	}

	bname := "Acer"
	berr := errors.New(fmt.Sprintf(model.BrandNotFound,bname))
	mock.ExpectQuery(query).WithArgs(bname).WillReturnError(berr)
	res,err = dbHandler.GetBrandByName(bname)
	//fmt.Println(err)
	if !reflect.DeepEqual(err, berr) {
		t.Errorf("Error, got: %v, want: %v.", err, berr)
	}
}
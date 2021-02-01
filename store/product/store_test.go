package product

import (
	"Catlog/model"
	"errors"
	"fmt"
	"github.com/DATA-DOG/go-sqlmock"
	"reflect"
	"testing"
)

var u = []model.Product{
	{1,"Headphone", 1,},
    {2,"Mouse",1},
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
    	expproduct model.Product
    	err error
	}{
    	{1,u[0],nil},
    	{2,u[1],nil},
    	//{3,model.Product{},model.Err{3}},
	}
	query := "SELECT id, name, bid FROM product WHERE id = ?"

	for _,testCase := range testCases{
		rows := sqlmock.NewRows([]string{"id", "name", "bid"}).
			AddRow(testCase.expproduct.Id, testCase.expproduct.Name, testCase.expproduct.BId)
		mock.ExpectQuery(query).WithArgs(testCase.id).WillReturnRows(rows)
		res,_ := dbHandler.GetProductById(testCase.id)

		if !reflect.DeepEqual(res, testCase.expproduct){
			t.Errorf("Product resulr was incorrect, got: %v, want: %v.", res, testCase.expproduct)
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
		expproduct model.Product
		err error
	}{
		{3,model.Product{},model.Err{3}},
		{-3,model.Product{},errors.New(model.NegativeId)},
	}
	query := "SELECT id, name, bid FROM product WHERE id = ?"

	for _,testCase := range testCases {
		mock.ExpectQuery(query).WithArgs(testCase.id).WillReturnError(testCase.err)
		_, err := dbHandler.GetProductById(testCase.id)
		if !reflect.DeepEqual(err, testCase.err) {
			t.Errorf("Error, got: %v, want: %v.", err, testCase.err)
		}
	}
}

func TestInsert(t *testing.T){
	db,mock,err := sqlmock.New()
	if err!=nil {
		fmt.Sprintf(model.SQLMockError)
		return
	}
	dbHandler := New(db)
	query := "INSERT INTO product"

	mock.ExpectExec(query).WithArgs(u[0].Name, u[0].BId).WillReturnResult(sqlmock.NewResult(1,1))

	lastID,err := dbHandler.InsertProduct(u[0])
	if !reflect.DeepEqual(lastID, u[0].Id){
		t.Errorf("LastId, got: %v, want: %v.", lastID, u[0].Id)
	}
}

func TestDbStore_InsertErr(t *testing.T) {
	db,mock,err := sqlmock.New()
	if err!=nil {
		fmt.Sprintf(model.SQLMockError)
		return
	}
	dbHandler := New(db)
	query := "INSERT INTO product"

	mock.ExpectExec(query).WithArgs(u[0].Name, 0).WillReturnResult(sqlmock.NewResult(0,0))

	lastID,err := dbHandler.InsertProduct(u[0])
	if !reflect.DeepEqual(lastID, -1){
		t.Errorf("LastId, got: %v, want: %v.", lastID, -1)
	}
}



package model

import "fmt"

const (
	ErrPro = "Not Found any data for id %v"
)

type Err struct {
	Id int
}

func (er Err) Error() string{
	return fmt.Sprintf(ErrPro,er.Id)
}


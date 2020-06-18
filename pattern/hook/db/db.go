package db

import "fmt"

type Db struct {
	DbName string
}

func (d *Db) Save(i interface{}) error {
	if v, ok := i.(interface{ BeforeSave() }); ok {
		v.BeforeSave()
	}

	fmt.Println("Saving db")

	if v, ok := i.(interface{ AfterSave() }); ok {
		v.AfterSave()
	}

	return nil
}

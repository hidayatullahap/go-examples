package main

import (
	"fmt"
	"reflect"
)

type DbConfig struct {
	Name string `env:"db.name,required"`
	Host string `env:"db.host"`
	Port string
}

func main() {
	var foo DbConfig

	value := reflect.ValueOf(&foo)
	if value.Kind() == reflect.Ptr {
		value = value.Elem()

		v := reflect.TypeOf(&foo)
		if v.Kind() == reflect.Ptr {
			v = v.Elem()
		}

		var fieldNames []reflect.StructField

		for i := 0; i < v.NumField(); i++ {
			fieldNames = append(fieldNames, v.Field(i))
		}

		for _, field := range fieldNames {
			f := value.FieldByName(field.Name)
			f.SetString("test")

			tag := field.Tag.Get("env")
			fmt.Println(tag)
		}
	}

	fmt.Println(foo)
}

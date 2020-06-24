package main

import (
	"fmt"
	"os"
	"reflect"
	"strings"
)

type DbConfig struct {
	Name string `env:"db.name,required"`
	Host string `env:"db.host"`
	Port string
}

func main() {
	_ = os.Setenv("db.name", "prodcts")
	_ = os.Setenv("db.host", "http://127.0.0.1")

	var foo DbConfig

	valueOf := reflect.ValueOf(&foo)
	if valueOf.Kind() == reflect.Ptr {
		valueOf = valueOf.Elem()

		typeOf := reflect.TypeOf(&foo)
		if typeOf.Kind() == reflect.Ptr {
			typeOf = typeOf.Elem()
		}

		var fieldNames []reflect.StructField

		for i := 0; i < typeOf.NumField(); i++ {
			fieldNames = append(fieldNames, typeOf.Field(i))
		}

		for _, field := range fieldNames {
			f := valueOf.FieldByName(field.Name)

			tag := field.Tag.Get("env")
			key := getEnv(tag)
			f.SetString(key)
		}
	}

	fmt.Println(foo)
}

func getEnv(key string) string {
	value := ""
	s := strings.Split(key, ",")
	sLen := len(s)

	switch {
	case sLen > 1:
		if s[1] == "required" {
			value = os.Getenv(s[0])
			if value == "" {
				panic(fmt.Sprintf("env config key %s not set", s[0]))
			}
		}
	case sLen == 1:
		value = os.Getenv(s[0])
	}

	return value
}

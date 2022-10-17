package utils

import "reflect"

// Function to check if a struct has a field with provided name
func HasStructField(v interface{}, name string) bool {
	rv := reflect.ValueOf(v)
	if rv.Kind() == reflect.Ptr {
	  rv = rv.Elem()
	}
	if rv.Kind() != reflect.Struct {
	  return false
	}
	return rv.FieldByName(name).IsValid()
}

// Function to retrieve the value of a field in a struct provided the field name
func GetStructFieldValue(v interface{}, name string) interface{} {
	rv := reflect.ValueOf(v)
	if rv.Kind() == reflect.Ptr {
	  rv = rv.Elem()
	}
	if rv.Kind() != reflect.Struct {
	  return nil
	}
	return rv.FieldByName(name).Interface()
}
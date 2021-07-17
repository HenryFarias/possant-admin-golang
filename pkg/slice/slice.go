package slice

import (
	"fmt"
	dynamicstruct "github.com/Ompluscator/dynamic-struct"
	"gopkg.in/jeevatkm/go-model.v1"
	"reflect"
)

func CopyList(source interface{}, dest interface{}) []error {
	sourceSlice := toSliceInterface(source)
	destSlice := toSliceInterface(dest)
	for _, value := range sourceSlice {
		result := newStruct(dest)
		errs := model.Copy(&result, value)
		if errs != nil {
			fmt.Println(errs)
			return errs
		}
		dest = append(destSlice, result)
	}
	return nil
}

func newStruct(reply interface{}) interface{} {
	fields := getAllFields(reply)
	instance := dynamicstruct.NewStruct()
	for f, typeOfStruct := range fields {
		field := fmt.Sprintf("%v", f)
		instance.AddField(field, typeOfStruct, "")
	}
	return instance.Build().New()
}

func getAllFields(reply interface{}) map[interface{}]interface{} {
	m := make(map[interface{}]interface{})
	t := reflect.TypeOf(reply)
	for i := 0; i < t.NumField(); i++ {
		m[t.Field(i)] = "teste"
	}
	return m
}

func toSliceInterface(slice interface{}) []interface{} {
	s := reflect.ValueOf(slice)
	if s.Kind() != reflect.Slice {
		panic("InterfaceSlice() given a non-slice type")
	}
	if s.IsNil() {
		return nil
	}
	ret := make([]interface{}, s.Len())
	for i:=0; i<s.Len(); i++ {
		ret[i] = s.Index(i).Interface()
	}
	return ret
}

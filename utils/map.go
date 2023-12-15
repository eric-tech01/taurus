package utils

import "reflect"

func MapToStruct(m map[string]interface{}, s interface{}) {
	rv := reflect.ValueOf(s).Elem()
	for key, value := range m {
		field := rv.FieldByName(key)
		if field.IsValid() {
			if field.Type().AssignableTo(reflect.TypeOf(value)) {
				field.Set(reflect.ValueOf(value))
			}
		}
	}
}

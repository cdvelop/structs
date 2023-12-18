package structs

import "reflect"

// obtiene el nombre de las interfaces nulas
func NilInterfaces(struct_in any) []string {
	var nullFields []string
	// Obtener el valor reflect.Value de MainHandler
	handlersValue := reflect.ValueOf(struct_in)
	// Iterar sobre los campos de MainHandler
	for i := 0; i < handlersValue.NumField(); i++ {
		fieldValue := handlersValue.Field(i)
		// Verificar si el campo es una interfaz es nula
		if fieldValue.Kind() == reflect.Interface && fieldValue.IsNil() {
			fieldName := handlersValue.Type().Field(i).Name
			nullFields = append(nullFields, fieldName)
		}
	}

	return nullFields
}

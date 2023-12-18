package structs

import "reflect"

// obtiene el nombre de los capos de una estructura
func GetFieldNames(struct_in any) (field_names []string) {
	// Obtener el tipo reflect.Type de la estructura
	structType := reflect.TypeOf(struct_in)
	// Iterar sobre los campos de la estructura
	for i := 0; i < structType.NumField(); i++ {
		field := structType.Field(i)
		field_names = append(field_names, field.Name)
	}

	return field_names
}

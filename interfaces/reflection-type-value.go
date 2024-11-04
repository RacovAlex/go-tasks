package interfaces

import (
	"fmt"
	"reflect"
)

// PrintDetails принимает аргумент типа interface{} и выводит информацию об объекте: его тип и значение.
func PrintDetails(i interface{}) {
	v := reflect.ValueOf(i)
	fmt.Printf("type: %T, value: %v\n", i, v)
}

// IsNil принимает значение любого типа и проверяет, является ли оно nil
func IsNil(i interface{}) bool {
	v := reflect.ValueOf(i)
	if !v.IsValid() {
		return true
	}
	switch v.Kind() {
	case reflect.Chan, reflect.Func, reflect.Interface, reflect.Map, reflect.Ptr, reflect.Slice:
		return v.IsNil()
	default:
		return false
	}
}

// SetField принимает указатель на структуру, название поля в виде строки и значение.
// Функция устанавливает значение указанного поля в структуре через рефлексию.
func SetField(v any, name string, value any) {
	// Получаем значение переданного объекта
	refVal := reflect.ValueOf(v)

	// Проверяем, что передан указатель на структуру
	if refVal.Kind() != reflect.Ptr || refVal.Elem().Kind() != reflect.Struct {
		fmt.Println("Expected a pointer to a struct")
		return
	}

	// Разыменовываем указатель, чтобы работать с самой структурой
	refVal = refVal.Elem()

	// Ищем поле по имени
	field := refVal.FieldByName(name)

	// Проверяем, существует ли поле и можно ли его изменить
	if !field.IsValid() {
		fmt.Println("No such field:", name)
		return
	}

	if !field.CanSet() {
		fmt.Println("Cannot set field:", name)
		return
	}

	// Устанавливаем новое значение
	val := reflect.ValueOf(value)
	if field.Type() != val.Type() {
		fmt.Println("Provided value type doesn't match field type")
		return
	}

	field.Set(val)
}

func ImplementsInterface(v any, i interface{}) bool {
	structType := reflect.ValueOf(v).Type()
	if structType.Implements(reflect.TypeOf(i).Elem()) {
		return true
	}
	return false
}

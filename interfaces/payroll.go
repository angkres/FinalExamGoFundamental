package interfaces

//Interface
type PayrollInterface interface {
	Add(map[string]interface{})
	GetLastId() int
	GetFieldsAll()map[string]interface{}
}

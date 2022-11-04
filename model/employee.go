package model

//Struct Employee
type Employee struct {
	id      int
	name    string
	gender  string
	grade   int
	married bool
}

//Slice Employee
var Employees []Employee

// Method Add milik Employee.
// Mengimplementasi method pada interface PayrollInterface.
// Untuk memasukkan dalam struct employee lalu di append ke dalam slice employee.
func (emp *Employee) Add(datas map[string]interface{}) {
	emp.id = datas["id"].(int)
	emp.name = datas["name"].(string)
	emp.gender = datas["gender"].(string)
	emp.grade = datas["grade"].(int)
	emp.married = datas["married"].(bool)
	Employees = append(Employees, *emp)
}

// Method GetFields milik Employee.
// Mengimplementasi method pada interface PayrollInterface.
// Untuk mengembalikan nilai dari employee karena properti dalam employee bersifat private.
func (emp *Employee) GetFieldsAll() map[string]any {
	datas := map[string]any{
		"id":      emp.id,
		"name":    emp.name,
		"gender":  emp.gender,
		"grade":   emp.grade,
		"married": emp.married,
	}
	return datas
}

// Method GetLastId milik Employee.
// Mengimplementasi method pada interface PayrollInterface.
// untuk mengambil id terkhir dari id slice Employee.
func (emp *Employee) GetLastId() int {
	if Employees == nil {
		return 1
	} else {
		var tempId int
		for _, v := range Employees {
			if tempId < v.id {
				tempId = v.id
			}
		}
		return tempId + 1
	}
}

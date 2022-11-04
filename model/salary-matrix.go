package model

//Struct Salary Matrix
type SalaryMatrix struct {
	id           int
	grade        int
	basic_salary uint
	pay_cut      uint
	allowance    uint
	hof          uint
}

//Slice Salary Matrix
var SalaryMatrixs []SalaryMatrix

//Method Add milik Salary Matrix
//Mengimplementasi method pada interface PayrollInterface
//Untuk memasukkan dalam struct salary matrix lalu di append ke dalam slice Salary Matrix
func (sal *SalaryMatrix) Add(datas map[string]interface{}) {
	sal.id = datas["id"].(int)
	sal.grade = datas["grade"].(int)
	sal.basic_salary = datas["basic_salary"].(uint)
	sal.pay_cut = datas["pay_cut"].(uint)
	sal.allowance = datas["allowance"].(uint)
	sal.hof = datas["hof"].(uint)
	SalaryMatrixs = append(SalaryMatrixs, *sal)
}

//Method GetFields milik Salary Matrix
//Mengimplementasi method pada interface PayrollInterface
//Untuk mengembalikan nilai dari salary matrix karena properti dalam salary matrix bersifat private
func (sal *SalaryMatrix) GetFieldsAll() map[string]any {
	datas := map[string]any{
		"id":           sal.id,
		"grade":        sal.grade,
		"basic_salary": sal.basic_salary,
		"pay_cut":      sal.pay_cut,
		"allowance":    sal.allowance,
		"hof":          sal.hof,
	}
	return datas
}

//Method GetLastId milik Salary Matrix
//Mengimplementasi method pada interface PayrollInterface
//untuk mengambil id terkhir dari id slice Salary Matrix
func (sal *SalaryMatrix) GetLastId() int {
	if SalaryMatrixs == nil {
		return 1
	} else {
		var tempId int
		for _, v := range SalaryMatrixs {
			if tempId < v.id {
				tempId = v.id
			}
		}
		return tempId + 1
	}
}

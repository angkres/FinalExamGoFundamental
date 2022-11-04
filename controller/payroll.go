package controller

import "payroll-program/interfaces"

//Controller untuk input pegawai
func InsertEmployeeHandler(pay interfaces.PayrollInterface, name string, gender string, grade int, married bool) {
	datas := map[string]interface{}{
		"id":      pay.GetLastId(),
		"name":    name,
		"gender":  gender,
		"grade":   grade,
		"married": married,
	}
	pay.Add(datas)
}

//Controller untuk input salary matrix
func InsertSalaryMatrixHandler(pay interfaces.PayrollInterface, grade int, basic_salary uint, pay_cut uint, allowance uint, hof uint) {
	datas := map[string]interface{}{
		"id":           pay.GetLastId(),
		"grade":        grade,
		"basic_salary": basic_salary,
		"pay_cut":      pay_cut,
		"allowance":    allowance,
		"hof":          hof,
	}
	pay.Add(datas)
}

//Controller untuk input payroll
func InsertPayrollHandler(pay interfaces.PayrollInterface, employee string, present uint, absent uint) {
	datas := map[string]interface{}{
		"id":       pay.GetLastId(),
		"employee": employee,
		"present":  present,
		"absent":   absent,
	}
	pay.Add(datas)
}

//Controller untuk mengembalikan semua nilai
func GetFieldsHandler(pay interfaces.PayrollInterface) map[string]any {
	return pay.GetFieldsAll()
}

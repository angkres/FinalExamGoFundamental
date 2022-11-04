package model

import (
	"errors"
	"strconv"
)

// Struct Payroll
type Payroll struct {
	id                int
	basic_salary      uint
	pay_cut           uint
	additional_salary uint
	Employee
}

// Slice Payroll
var Payrolls []Payroll

// Method GetIndexName milik Payroll.
// Untuk mengambil index ketika nama sama.
func (pay *Payroll) GetIndexName(name *string) (Employee, error) {
	var emp Employee
	for _, v := range Employees {
		if v.name == *name {
			return v, nil
		}
	}
	return emp, errors.New("Id " + *name + " tidak ditemukan")
}

// Method GetIndexGrade milik Payrol.
// Untuk mengambil index ketika grade sama.
func (pay *Payroll) GetIndexGrade(grade *int) (SalaryMatrix, error) {
	var sal SalaryMatrix
	for _, v := range SalaryMatrixs {
		if v.grade == *grade {
			return v, nil
		}
	}
	return sal, errors.New("Id " + strconv.Itoa(*grade) + " tidak ditemukan")
}

// Method Add milik Payroll
// Mengimplementasi method pada interface PayrollInterface
// Untuk memasukkan dalam struct payroll lalu di append ke dalam slice payroll
func (pay *Payroll) Add(datas map[string]interface{}) {
	name := datas["employee"].(string)
	emp, err := pay.GetIndexName(&name)
	if err == nil {
		pay.name = name
		pay.Employee.id = emp.id
		pay.grade = emp.grade
		pay.gender = emp.gender
		pay.married = emp.married
		sal, errS := pay.GetIndexGrade(&pay.grade)
		if errS == nil {
			pay.id = datas["id"].(int)
			pay.basic_salary = sal.basic_salary
			pay.pay_cut = datas["absent"].(uint) * sal.pay_cut
			if emp.gender == "L" && emp.married {
				pay.additional_salary = (datas["present"].(uint) * sal.allowance) + sal.hof
			} else {
				pay.additional_salary = datas["present"].(uint) * sal.allowance
			}
		}
	}
	Payrolls = append(Payrolls, *pay)
}

// Method GetFields milik Payroll
// Mengimplementasi method pada interface PayrollInterface
// Untuk mengembalikan nilai dari payroll karena properti dalam payroll bersifat private
func (pay *Payroll) GetFieldsAll() map[string]any {
	datas := map[string]any{
		"id":                pay.id,
		"basic_salary":      pay.basic_salary,
		"pay_cut":           pay.pay_cut,
		"additional_salary": pay.additional_salary,
		"employeeId":        pay.Employee.id,
		"name":              pay.name,
		"gender":            pay.gender,
		"grade":             pay.grade,
		"married":           pay.married,
		"total":             pay.basic_salary - pay.pay_cut + pay.additional_salary,
	}
	return datas
}

// Method GetLastId milik Payroll
// Mengimplementasi method pada interface PayrollInterface
// untuk mengambil id terkhir dari id slice Payroll
func (pay *Payroll) GetLastId() int {
	if Payrolls == nil {
		return 1
	} else {
		var tempId int
		for _, v := range Payrolls {
			if tempId < v.id {
				tempId = v.id
			}
		}
		return tempId + 1
	}
}

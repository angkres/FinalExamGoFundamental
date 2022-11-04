package template

import (
	"fmt"
	"payroll-program/controller"
	"payroll-program/helper"
	"payroll-program/model"
)

// Fungsi Salary Matrix
func MatrixSallaryEmployee() {
	var salaryMatrix model.SalaryMatrix
	controller.InsertSalaryMatrixHandler(&salaryMatrix, 1, 5000000, 100000, 150000, 1000000)
	controller.InsertSalaryMatrixHandler(&salaryMatrix, 2, 9000000, 200000, 300000, 2000000)
	controller.InsertSalaryMatrixHandler(&salaryMatrix, 3, 15000000, 400000, 600000, 3000000)
}

// Fungsi tampilan untuk Salary Matrix
func ListMatrixSalary() {
	helper.ClearScreen()
	fmt.Println("===================================================================================")
	fmt.Printf("| %-5v| %-8v| %-15v| %-13v| %-14v| %-15v|\n", "ID", "Grade", "Basic Salary", "PayCut/day", "Allowance/day", "Head of family")
	fmt.Println("===================================================================================")
	for _, v := range model.SalaryMatrixs {
		datas := controller.GetFieldsHandler(&v)
		id := datas["id"].(int)
		grade := datas["grade"].(int)
		basic_salary := datas["basic_salary"].(uint)
		pay_cut := datas["pay_cut"].(uint)
		allowance := datas["allowance"].(uint)
		hof := datas["hof"].(uint)
		fmt.Printf("| %-5v| %-8v| Rp. %-11v| Rp. %-9v| Rp. %-9v | Rp. %-11v|\n", id, grade, basic_salary, pay_cut, allowance, hof)
	}
	fmt.Println("===================================================================================")

	//untuk kembali ke menu
	helper.BackHandler()
	Menu()
}

package template

import (
	"fmt"
	"payroll-program/controller"
	"payroll-program/helper"
	"payroll-program/model"
)

// Fungsi menginputkan Employee
func FormEmployee() {
	helper.ClearScreen()
	fmt.Println("Form Employee")
	fmt.Println("==================")
	name := InputName()
	gender := InputGender()
	grade := InputGrade()
	married := InputMarried()

	//proses employee
	var employee model.Employee
	controller.InsertEmployeeHandler(&employee, name, gender, grade, married)
	fmt.Println("Data berhasil di input")

	//untuk kembali ke menu
	helper.BackHandler()
	Menu()
}

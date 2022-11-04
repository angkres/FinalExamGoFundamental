package template

import (
	"fmt"
	"payroll-program/controller"
	"payroll-program/helper"
	"payroll-program/model"
)

// Fungsi proses dari payroll
func ProsesPayroll() {
	helper.ClearScreen()
	fmt.Println("Proses Gaji Pegawai")
	fmt.Println("=======================")
	name := InputNamePayroll()
	present := InputPresent()
	absent := InputAbsent()
	//untuk memproses payroll
	var payroll model.Payroll
	controller.InsertPayrollHandler(&payroll, name, present, absent)
	//get nilai payroll
	datas := controller.GetFieldsHandler(&payroll)
	idPayroll := datas["id"].(int)
	basic_salary := datas["basic_salary"].(uint)
	pay_cut := datas["pay_cut"].(uint)
	additional_salary := datas["additional_salary"].(uint)
	idEmployee := datas["employeeId"].(int)
	nameEmployee := datas["name"].(string)
	gender := datas["gender"].(string)
	grade := datas["grade"].(int)
	married := datas["married"].(bool)
	total := datas["total"].(uint)

	genderInd := Gender(gender)
	marriedInd := Married(married)
	//tampilan slip gaji
	helper.ClearScreen()
	fmt.Println("===================================================================")
	fmt.Println("|                            PT. Phincon                          |")
	fmt.Println("===================================================================")
	fmt.Println("|                   Salary Slip for November 2022                 |")
	fmt.Println("===================================================================")
	fmt.Printf("| ID Pegawai : %-17v   Departemen : Golang Trainee    |\n", idEmployee)
	fmt.Printf("| Nama : %-23v   Grade : %-22v |\n", nameEmployee, grade)
	fmt.Printf("| Gender : %-21v   Married : %-20v |\n", genderInd, marriedInd)
	fmt.Println("===================================================================")
	fmt.Println("|            Earnings            |           Deductions           |")
	fmt.Println("===================================================================")
	fmt.Printf("| Basic Salary : %15v | Pay Cut : %20v |\n", basic_salary, pay_cut)
	fmt.Printf("| Additional Salary : %10v | %30v |\n", additional_salary, "")
	fmt.Println("===================================================================")
	fmt.Printf("|             Net Pay            | Rp. %26v |\n", total)
	fmt.Println("===================================================================")
	fmt.Printf("%66v\n", idPayroll)

	//untuk kembali ke menu
	helper.BackHandler()
	Menu()
}


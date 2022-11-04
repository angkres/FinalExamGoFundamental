package template

import (
	"fmt"
	"payroll-program/controller"
	"payroll-program/helper"
	"payroll-program/model"
)

// Fungsi tampilan Employee
func ListEmployee() {
	helper.ClearScreen()
	fmt.Println("============================================================================")
	fmt.Printf("| %-5v| %-20v| %-14v| %-8v| %-17v |\n", "ID", "Nama", "Gender", "Grade", "Married")
	fmt.Println("============================================================================")
	//jika data kosong
	if len(model.Employees) == 0 {
		fmt.Printf("| %-56v |\n", "Data Kosong")
		//jika data ada
	} else {
		for _, v := range model.Employees {
			datas := controller.GetFieldsHandler(&v)
			id := datas["id"].(int)
			name := datas["name"].(string)
			gender := datas["gender"].(string)
			grade := datas["grade"].(int)
			married := datas["married"].(bool)
			
			genderInd := Gender(gender)
			marriedInd := Married(married)
			fmt.Printf("| %-5v| %-20v| %-14v| %-8v| %-17v |\n", id, name, genderInd, grade, marriedInd)
		}
	}
	fmt.Println("============================================================================")
	//untuk kembali ke menu
	helper.BackHandler()
	Menu()
}

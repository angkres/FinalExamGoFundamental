package template

import (
	"fmt"
	"os"
	"payroll-program/helper"
)

//Fungsi untuk menjalankan Matrix Salary terlebih dahulu, karena dia statis
func init() {
	MatrixSallaryEmployee()
}

//Fungsi menu
func Menu() {
	helper.ClearScreen()
	fmt.Println("==========================")
	fmt.Println("|          Menu          |")
	fmt.Println("==========================")
	fmt.Println("| 1. Daftar Pegawai      |")
	fmt.Println("| 2. Tambah Pegawai      |")
	fmt.Println("| 3. Proses Gaji Pegawai |")
	fmt.Println("| 4. Matrix Gaji Pegawai |")
	fmt.Println("| 5. Exit                |")
	fmt.Println("==========================")

	var menu int
	fmt.Println("Pilih Menu")
	fmt.Scanln(&menu)

	switch menu {
	case 1:
		ListEmployee()
	case 2:
		FormEmployee()
	case 3:
		ProsesPayroll()
	case 4:
		ListMatrixSalary()
	case 5:
		os.Exit(0)
	default:
		Menu()
	}
}

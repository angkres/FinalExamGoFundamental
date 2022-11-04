package template

import (
	"bufio"
	"fmt"
	"os"
	"payroll-program/controller"
	"payroll-program/model"
	"regexp"
	"strings"
)

// Fungsi validasi input nama
func InputName() string {
	fmt.Print("Name: ")
	//agar inputan bisa menggunakan spasi
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	name := scanner.Text()

	//hanya bisa mengandung a-z, A-Z dan spasi
	a := regexp.MustCompile(`^[a-zA-Z ]+$`).MatchString

	//panjang nama lebih dari 20 huruf
	if len(name) > 20 {
		fmt.Println("Name tidak boleh lebih dari 20 huruf")
		return InputName()
		//nama sesuai regex maka di return nama
	} else if a(name) {
		return name
		//nama tidak sesuai regex
	} else {
		fmt.Println("Name tidak boleh kosong dan hanya mengandung alfabet")
		return InputName()
	}
}

// Fungsi validasi input gender
func InputGender() string {
	var gender string
	fmt.Print("Gender: ")
	fmt.Scanln(&gender)

	//jika inputan L/l/p/P maka return gender
	if strings.ToUpper(gender) == "L" || strings.ToUpper(gender) == "P" {
		return gender
		//jika salah, kembali input gender
	} else {
		fmt.Println("Gender harus diinput 'L' untuk laki-laki dan 'P' untuk perempuan")
		return InputGender()
	}
}

// Fungsi validasi input grade
func InputGrade() int {
	var grade int
	fmt.Print("Grade: ")
	fmt.Scanln(&grade)
	//untuk mencari grade pada Salary Matrix
	for _, v := range model.SalaryMatrixs {
		datas := controller.GetFieldsHandler(&v)
		gradeSalary := datas["grade"].(int)
		//_, gradeSalary, _, _, _, _ := v.GetFields()
		//jika grade pada salary dan yang kita input sama maka return grade
		if gradeSalary == grade {
			return grade
		}
	}
	//jika tidak ada, maka input ulang grade
	fmt.Println("Grade tidak ditemukan dalam Matrix Sallary")
	return InputGrade()
}

// Fungsi validasi input married
func InputMarried() bool {
	var married string
	fmt.Print("Married: ")
	fmt.Scanln(&married)

	//jika ya/Ya/Yes/yes maka nilai married true
	if strings.ToLower(married) == "ya" || strings.ToLower(married) == "yes" {
		return true
		//jika tidak/Tidak/No/no maka nilai married false
	} else if strings.ToLower(married) == "tidak" || strings.ToLower(married) == "no" {
		return false
		//jika tidak semuanya, input married lagi
	} else {
		fmt.Println("Married  harus diinput 'Ya/Yes' untuk kondisi true dan 'Tidak/No' untuk kondisi false")
		return InputMarried()
	}
}

// Fungsi validasi input nama pada payroll
func InputNamePayroll() string {
	fmt.Print("Name: ")
	//agar inputan bisa menggunakan spasi
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	name := scanner.Text()
	//untuk mencari name pada Employee
	//var employee model.Employee
	for _, v := range model.Employees {
		datas := controller.GetFieldsHandler(&v)
		nameEmployee := datas["name"].(string)
		// nameEmployee = datas["name"].(string)
		//_, nameEmployee, _, _, _ := v.GetFields()
		//jika nama pada Employee dan yang kita input sama maka return nama
		if nameEmployee == name {
			return name
		}
	}
	//jika tidak ada maka input ulang nama pada payroll
	fmt.Println("Nama tidak ditemukan dalam daftar pegawai")
	return InputNamePayroll()
}

func Gender(gender string) string {
	if gender == "L" {
		return "Laki-laki"
	} else {
		return "Perempuan"
	}
}

func Married(married bool) string {
	if !married {
		return "Belum Menikah"
	} else {
		return "Sudah Menikah"
	}
}

// Fungsi validasi input
func InputPresent() uint {
	var present int
	fmt.Print("Jumlah hari masuk: ")
	fmt.Scanln(&present)
	//Jika present sama dengan 0, input ulang
	if present <= 0 {
		fmt.Println("Anda kemana saja tidak masuk?")
		return InputPresent()
		//jika present lebih dari 31 hari, input ulang
	} else if present > 31 {
		fmt.Println("Tidak ada hari dalam satu bulan yang lebih dari 31 hari")
		return InputPresent()
	} else {
		return uint(present)
	}
}

// // Fungsi validasi input absent
func InputAbsent() uint {
	var absent int
	fmt.Print("Jumlah hari tidak masuk: ")
	fmt.Scanln(&absent)
	//absent tidak boleh kurang dari 0 dan lebih dari 31 hari
	if absent > 31 || absent < 0 {
		fmt.Println("Anda seharusnya udah dipecat")
		return InputAbsent()
	} else {
		return uint(absent)
	}
}

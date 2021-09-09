package homework

/*
Bài 4 Một nhân viên trong công ty bao gồm các thuộc tính sau : Tên, Hệ số lương, Tiền trợ cấp
Tạo 1 mảng nhân viên (số lượng tuỳ ý) và thực hiện các chức năng sau:

Sắp xếp tên nhân viên tăng dần theo bảng chữ cái
Sắp xếp nhân viên theo mức lương giảm dần (lương = Hệ số lương * 1.500.000 + Tiền trợ cấp)
Lấy ra danh sách nhân viên có mức lương lớn thứ 2 trong mảng nhân viên
*/

type Staff struct {
	Name               string
	CoefficientsSalary int
	Allowance          int
}

type SortName []Staff

func (s SortName) Len() int {
	return len(s)
}

func (s SortName) Less(i, j int) bool {
	return s[i].Name < s[j].Name
}

func (s SortName) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

type SortSalary []Staff

func (s SortSalary) Len() int {
	return len(s)
}

func (s SortSalary) Less(i, j int) bool {
	return s[i].CoefficientsSalary*1500000+s[i].Allowance > s[j].CoefficientsSalary*1500000+s[j].Allowance
}

func (s SortSalary) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

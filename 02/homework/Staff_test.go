package homework

import (
	"fmt"
	"sort"
	"testing"
)

func TestStaff(t *testing.T) {
	staff := []Staff{
		{"Denis", 12, 1000000},
		{"Ascolin", 50, 1000000},
		{"IIsus", 99, 1000000},
		{"Mario", 2, 10000},
		{"Gogaie", 37, 10000000},
		{"Dentistul", 76, 100000},
	}
	sort.Sort(SortName(staff))
	fmt.Println("Sort by Name:", staff)
	fmt.Println("-------------------------------------------")

	sort.Sort(SortSalary(staff))
	fmt.Println("Sort by Salary:", staff)
	fmt.Println("-------------------------------------------")

	fmt.Println(staff[1])
}

package mapreduce

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	employees = []Employee{
		{"Hao", 44, 0, 8000},
		{"Bob", 34, 10, 5000},
		{"Alice", 23, 5, 9000},
		{"Jack", 26, 0, 4000},
		{"Tom", 48, 9, 7500},
		{"Marry", 29, 0, 6000},
		{"Mike", 32, 8, 4000},
	}
)

func Test_GeneralExample(t *testing.T) {
	// 统计有多少员工大于40岁
	oldCnt := EmployeeCountIf(employees, func(e *Employee) bool {
		return e.Age > 40
	})
	assert.Equal(t, 2, oldCnt)

	// 统计有多少员工薪水大于6000
	highPayCnt := EmployeeCountIf(employees, func(e *Employee) bool {
		return e.Salary >= 6000
	})
	assert.Equal(t, 4, highPayCnt)

	// 统计列出没有休假的员工
	hasNoVacationsCnt := EmployeeFilterIn(employees, func(e *Employee) bool {
		return e.Vacation == 0
	})
	assert.Equal(t, 3, len(hasNoVacationsCnt))

	// 统计所有员工的薪资总和
	totalPay := EmployeeSumIf(employees, func(e *Employee) int {
		return e.Salary
	})
	assert.Equal(t, 43500, totalPay)

	// 统计30岁以下员工的薪资总和
	totalYoungPay := EmployeeSumIf(employees, func(e *Employee) int {
		if e.Age < 30 {
			return e.Salary
		}
		return 0
	})
	assert.Equal(t, 19000, totalYoungPay)
}

func Test_Transform(t *testing.T) {
	// 用于字符串数组
	names := []string{"Tom", "Terry", "Marry"}
	Transform(names, func(s string) string {
		return strings.ToLower(s)
	})
	assert.Equal(t, []string{"Tom", "Terry", "Marry"}, names)

	// 用于整形数据
	numbers := []int{1, 2, 3, 4, 5}
	TransformInPlace(numbers, func(n int) int {
		return n * n
	})
	assert.Equal(t, []int{1, 4, 9, 16, 25}, numbers)

	// 用于结构体
	employees = []Employee{
		{"Hao", 44, 0, 8000},
		{"Bob", 34, 10, 5000},
	}
	TransformInPlace(employees, func(e Employee) Employee {
		e.Salary += 1000
		e.Age += 1
		return e
	})
	assert.Equal(t, []Employee{
		{"Hao", 45, 0, 9000},
		{"Bob", 35, 10, 6000},
	}, employees)
}

func Test_ReflectReduce(t *testing.T) {
	names := []string{"Tom", "Terry", "Marry"}
	v := ReflectReduce(names, func(s1, s2 string) string {
		return s1 + " " + s2
	}, 0)
	assert.Equal(t, "Tom Terry Marry", v)

	numbers := []int{1, 2, 3, 4, 5}
	v = ReflectReduce(numbers, func(x, y int) int {
		return x * y
	}, 1)
	assert.Equal(t, 120, v)
}

func Test_ReflectFilter(t *testing.T) {
	var numbers = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	odds := ReflectFilter(numbers, func(n int) bool {
		return n%2 == 1
	})
	assert.Equal(t, []int{1, 3, 5, 7, 9}, odds)

	ReflectFilterInPlace(&numbers, func(n int) bool {
		return n > 5
	})
	assert.Equal(t, []int{6, 7, 8, 9, 10}, numbers)
}

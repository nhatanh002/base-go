package utils

import (
	"fmt"
	"testing"
	"time"
)

type TestCase[A any, B any] struct {
	Case     string
	Given    A
	Expected B
}

type TestSuite[A any, B any] []TestCase[A, B]

func TestMapReduce(t *testing.T) {
	test1 := TestCase[[]int, []string]{
		"Test 1",
		[]int{1, 2, 3, 4, 5, 6},
		[]string{"1", "2", "3", "4", "5", "6"},
	}
	test2 := TestCase[[]int, string]{
		"Test 2",
		[]int{1, 2, 3, 4, 5, 6},
		"!1!2!3!4!5!6",
	}

	fmt.Println(Map(func(a int) string { return fmt.Sprintf("%d", a) }, test1.Given), test1.Expected)
	fmt.Println(Reduce(func(b string, a int) string { return fmt.Sprintf("%s!%d", b, a) }, "", test2.Given) == test2.Expected)

}

func TestDropWhile(t *testing.T) {
	test := []string{"a", "b", "c", "d", "e", "e", "f", "e"}
	fmt.Println("Result: ", DropWhile(func(a string) bool { return a != "e" }, test))
}

func TestTakeWhile(t *testing.T) {
	test := []string{"a", "b", "c", "d", "e", "e", "f", "e"}
	fmt.Println("Result: ", TakeWhile(func(a string) bool { return a != "e" }, test))
}

func TestMember(t *testing.T) {
	test := []string{"a", "b", "c", "d", "e", "e", "f", "e"}
	fmt.Println("Result: ", Member("e", test))
	fmt.Println("Result: ", Member("m", test))
}

func TestCompose(t *testing.T) {
	test := []string{"a", "b", "c", "d", "e", "e", "f", "e"}
	length := func(ip []string) int { return len(ip) }
	double := func(i int) int { return i * 2 }
	fmt.Println("Result: ", Compose(double, length)(test))
	fmt.Println("Result: ", Pipe(length, double)(test))
}

func TestGroupBy(t *testing.T) {
	test := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12}
	partition := func(i int) int { return i % 3 }
	parted := GroupBy(partition, test)
	for k, arr := range parted {
		fmt.Printf("Modulo %d: %+v\n", k, arr)
	}
}

func TestGetDays(t *testing.T) {
	now := time.Now()
	lastFirst := GetLastFirstDayOfMonth(now)
	fmt.Printf("Last first day of the month: %+v\n\n", lastFirst)
	lastMon := GetLastMonday(now)
	fmt.Printf("Last first day of the week: %+v\n\n", lastMon)
}

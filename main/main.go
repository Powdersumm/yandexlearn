package main
import (
	"fmt"
	"go/version"
)

func main() { 
	fmt.Println(git --version)
}

// package main

// import (
// 	"errors"
// 	"unicode/utf8"
// 	"testing"
// )

// var ErrInvalidUTF8 = errors.New("invalid utf8")

// func GetUTFLength(input []byte) (int, error) {
// 	if !utf8.Valid(input) {
// 		return 0, ErrInvalidUTF8
// 	}

// 	return utf8.RuneCount(input), nil
// }

// func TestGetUTFLength(t *testing.T ){
// 	tests := []struct{
// 		input []byte
// 		expected int
// 		err error
// 	}{
// 		{[]byte("hello"), 5, nil},
// 	}
// 	for _, test := range tests{
//         result, err := GetUTFLength(test.input)
// 		if result != test.expected {
// 			t.Errorf("GetUTFLength(%q) = %d, %v; want %d, %v", test.input, result, err, test.expected, test.err)
// 		}

// 	}
// }

// {[]byte{0xff}, 0, ErrInvalidUTF8}

// package main
// import(
// 	"testing"
// )

// func GetUTFLength(t *testing.T ){
// 	tests := []struct{
// 		input []byte
// 		expected int
// 		err error
// 	}{
// 		{[]byte("hello"), 5, nil},
// 	}
// 	result, err := GetUTFLength(test.input)
// 	for _, test := range tests{
// 		if result != test.expected{
// 			t.Errorf("GetUTFLength(%q)=%d, %v want %d,,%v ", test.input,result,err, test.expected, test.err )
// 		}

// 	}
// }

// package main
// import(
// 	"testing"
// )

// func DeleteVowels(s string) string{
// 	tests := []struct{
// 		input int
// 		expected int
// 	}
// 	result := DeleteVowels(test.input)
// 	for _, test := range tests{
// 		if result!= test.expected{
// 			t.Fatalf("DeleteVowels(%s)=%s, want %s ", test.input,result, test.expected )
// 		}

// 	}
// }

// package main
// import(
// 	"testing"
// )

// func Multiply(a, b int) int{
// 	tests := []struct{
// 		a,b int
// 		_ int
// 	}
// 	result := Multiply(test.a,test.b)
// 	for _, test := range tests{
// 		if result!= test.expected{
// 			t.Fatalf("Multiply()")
// 		}

// 	}
// }

// package main
// import(
// 	"testing"
// )

// func Testingsum(t *testing.T){
// 	tests:= []struct {
// 		a,b int
// 		expected int
// 	}
// 	for _, test := range tests{
// 		result := Sum(test.a, test.b)
// 		if result!= test.expected{
// 			t.Errorf("Sum(%d,%d) = %d, want %d", test.a, test.b, result, test.expected)
// 		}

// 	}
// }
// package main

// import (
// 	"time"
// 	"fmt"
// )

// func main() {
// 	duration := 2 * time.Hour
// 	hours := duration.Minutes()
// 	fmt.Println(hours)
// }

// package main
// import(
// 	"fmt"
// 	"time"
// )
// func  NextWorkday(start time.Time) time.Time{

// }

// package main
// import(
// 	"fmt"
// 	"time"
// )

// func  NextWorkday(start time.Time) time.Time{
// 	nexdday:= strt.Add(0,0,1)
// 	switch nextday{
// 	case time.Sunday:
// 		nextday = nextday.Add(0,0,2)
// 	case time.Saturday:
// 		nextday = nextday.Add(0,0,1)
// 	}
// 	return time.nextday
// }

// package main

// import (
// 	"testing"
// 	"time"
// )

// func TestTimeAgo(t *testing.T) {
// 	pastTime := time.Now().Add(-2 * time.Hour)
// 	timeAgo := TimeAgo(pastTime)
// 	expected := "2 hours ago"
// 	if timeAgo != expected {
// 		t.Errorf("Expected time ago string to be %s, but got %s", expected, timeAgo)
// 	}
// }

// func TimeAgo(pastTime time.Time) string{
//     result := time.Since(pastTime)
// 	years:= int(result.Hours() /24 /365)
// 	months := int(result.Hours() /24/12)
// 	days := int(result.Hours() /24)
// 	hours := int(result.Hours())
// 	minutes:= int(result.Minutes())
// 	seconds:= int(result.Seconds())
// 	if years >0 {
// 		return fmt.Sprintf("%d years ago",years)
// 	}else if months >0 {
// 		return fmt.Sprintf("%d months ago", months)
// 	}else if days >0 {
// 		return fmt.Sprintf("%d months ago", days)
// 	}else if hours >0 {
// 		return fmt.Sprintf("%d months ago", hours)
// 	}else if minutes >0 {
// 		return fmt.Sprintf("%d months ago", minutes)
// 	}else if seconds >0 {
// 		return fmt.Sprintf("%d months ago", seconds)
//     }
// }

// package main

// import (
// 	"fmt"
// 	"time"
// )

// func  ParseStringToTime(dateString, format string) (time.Time, error){
// 	result, err:= dateString.Format(format)
// 	if err != nil {
// 		return 0, fmt.Errorf("Оишибка")
// 	} else{
// 		return dateString.Format(format)
// 	}
// }

// package main
// import(
// 	"time"
// )

// func  FormatTimeToString(timestamp time.Time, format string) string{
// 	return timestamp.Format(format)

// }

// package main
// import(
// 	"time"
// )
// func TimeDifference(start, end time.Time) {
// 	diff := end.Sub(start)
// 	return diff
// }

// package main
// import("fmt")
// func main(){
// 	val:= string(0)
// 	fmt.Println(val)
// }

// package main

// import (
// 	"fmt"
// 	"strconv"
// 	"error"
// )

// type Num int
// func (n Num) String(){
// 	return fmt.Printf("%s", n)
// }

// func IntToBinary(num int) (string, error){
// 	if num<0{
// 		return  fmt.Errorf("negative numbers are not allowed")
// 	} else {
// 		num:= int64(num)
// 		int10 := strconv.FormatInt(num, 2)
// 		var result Num = Num(int10)
// 		return result,nil
// 	}
// }

// func main() {
// 	fmt.Println(IntToBinary(20))
// }

// package main

// import (
// 	"fmt"
// )

// func main() {
// 	var a string = "Зенит чемпион!"
// 	fmt.Println(len(a))
// }

// func GetCharacterAtPosition(str string, position int) (rune, error) {
// 	var result string = str
// 	if len(result) >= position && position >=0{
// 		return []rune(str)[position], nil
// 	} else {
// 		return 0, fmt.Errorf("position out of range")
// 	}

// func main() {
// 	fmt.Println(GetCharacterAtPosition("Зенит чемпион!", 20))
// }

// package main

// import (
// 	"fmt"
// )

// func Factorial(n int) (int, error){
// 	var result int = 1
// 	if n<0 {
// 		return 0, fmt.Errorf("factorials is not defined for negative numbers")
// 	} else{
// 		for i:= 1; i<=n; i++{
// 			result = result * i
// 		}
// 		return result, nil
// 	}
// }

// package main

// import (
// 	"fmt"
// )

// func GetCharacterAtPosition(str string, position int) (rune, error) {
// 	if len(str) >= (position - 1) {
// 		return []rune(str)[position], nil
// 	} else {
// 		return 0, fmt.Errorf("position out of range")
// 	}
// }
// func main() {
// 	fmt.Println(GetCharacterAtPosition("Оле ола вперёд Спартак Москва", 2))
// }

// package main
// import(
// 	"fmt"
// 	"errors"
// )

// type Num int
// func (n Num) Float64() float64{
// 	return fmt.Printf("%f",n)
// }
// func DivideIntegers(a,b int) (float64, error){
// 	if b == 0{
// 		return 0, fmt.Errorf(("ошибка"))
// 	}
// 	var resulta Num
// 	var resultb Num
// 	resulta = Num(a)
// 	resultb = Num(b)
// 	return resulta.Float64()/resultb.Float64(), nil
// }

// package main

// import "fmt"

// type Num int

// func (n Num) String() string {
// 	return fmt.Sprintf("%d", n)
// }

// func ConcatStringsAndInt(str1, str2 string, num int) string {
// 	var result Num
// 	result = Num(num)
// 	return str1 + "" + str2 + "" + result.String()
// }

// func main() {
// 	fmt.Println(ConcatStringsAndInt("al", "himik", 239))
// }

// package main
// import(
// 	"fmt"
// )
// type Logger interface{
// 	Log(message string)
// }
// type LogLevel string
// type Log struct{
// 	level LogLevel
// }
// const(
// 	Error LogLevel = "ERROR"
// 	Unfo LogLevel = "INFO"
// )
// func (l Log) Log(message string) string {
// 	fmt.Printf("%s: %s\n", l.level, message)

// }

// type Animal interface{
// 	MakeSound()
// }
// type Dog struct{}
// type Cat struct{}

// func (c Cat) MakeSound(){
// 	fmt.Println("Мяу!")

// }
// func (d Dog) MakeSound(){
// 	fmt.Println("Гав!")
// }

// package main

// import (
// 	"fmt"
// 	"math"
// )

// type Circle struct {
// 	radius float64
// }

// type Rectangle struct {
// 	height float64
// 	width  float64
// }

// func (r Rectangle) Area() float64 {
// 	return r.height * r.width
// }
// func (c Circle) Area() float64 {
// 	return math.Pi * math.Pow(c.radius, 2)
// }

// type Shape interface {
// 	Area() float64
// }

// func main() {
// 	figure_1 := Circle{radius: 1.0}
// 	fmt.Println(figure_1.Area())
// 	figure_2 := Rectangle{width: 57.2, height: 10.2}
// 	fmt.Println(figure_2.Area())
// }

// package main

// import (
// 	"time"
// )

// type Task struct {
// 	summary     string
// 	description string
// 	deadline    time.Time
// 	priority    int
// }
// type Note struct {
// 	title string
// 	text  string
// }

// type ToDoList struct {
// 	name  string
// 	tasks []Task
// 	notes []Note
// }

// func (t Task) IsOverdue() bool {
// 	if time.Now().After(t.deadline) {
// 		return true
// 	} else {
// 		return false
// 	}
// }
// func (t Task) IsTopPriority() bool {
// 	if t.priority == 4 || t.priority == 5 {
// 		return true
// 	} else {
// 		return false
// 	}
// }
// func (t ToDoList) TasksCount() int{
// 	return len(t.tasks)
// }
// func (t ToDoList) NotesCount() int{
// 	return len(t.notes)
// }
// func (t ToDoList)CountOverdueTasks() int{
// 	count:=0
// 	for _, task := range t.tasks{
// 		if task.IsOverdue()== true {
// 			count++
// 		}
// 	}
// 	return count
// }
// func (t ToDoList)  CountTopPrioritiesTasks()int{
// 	count := 0
// 	for _, task := range t.tasks{
// 		if task.IsTopPriority()==true{
// 			count ++
// 		}
// 	}
// 	return count
// }

// type Task struct {
// 	summary    string
// 	description string
// 	deadline   time.Time
// 	priority   int
// }

// func (t Task) IsOverdue() bool {
// 	if time.Now().After(t.deadline) {
// 		return true
// 	} else {
// 		return false
// 	}
// }
// func (t Task) IsTopPriority() bool {
// 	if t.priority == 4 || t.priority == 5 {
// 		return true
// 	} else {
// 		return false
// 	}
// }

// package main

// import "fmt"

// func main() {
// 	student := Student{name: "Gosha", solvedProblems: 30, scoreForOneTask: 10.0, passingScore: 290.0}
// 	fmt.Print(student.IsExcellentStudent())
// }

// type Student struct {
// 	name            string
// 	solvedProblems  int
// 	scoreForOneTask float64
// 	passingScore    float64
// }

// func (s Student) IsExcellentStudent() bool {
// 	if float64(s.solvedProblems)*s.scoreForOneTask >= s.passingScore {
// 		return true
// 	} else {
// 		return false
// 	}
// }

// package main
// import(
// 	"fmt"
// )
// type Person struct{
// 	name string
// 	age int
// 	address string
// }

// func (p Person) Print() {
// 	fmt.Printf("Name: %s \nAge: %d \nAddress: %s\n",p.name, p.age, p.address )

// }

// func main() {
// 	student1 := students.Student{}
// 	// user := students.User{}
// 	fmt.Println(student1)
// 	// fmt.Println(user)
// }

// package main

// func DeleteLongKeys(m map[string]int) map[string]int {
// 	result := make(map[string]int)
// 	for key, val := range m {
// 		if len(key) >= 6 {
// 			result[key] = val
// 		}
// 	}
// 	return result
// }

// package main

// func CountingSort(contacts []string) map[string]int{
//     result := make(map[string]int)
//     for _ , val := range contacts {
//         result[val]++
//     }
//     return result
// }

// package main
// func  SwapKeysAndValues(m map[string]string) map[string]string{
//     result:= make(map[string]string)
//     for key, val := range m{
//         result[val] = key
//     }
//     return result

// }

// import (
//     "fmt"
// )
// func  SumOfValuesInMap(m map[int]int) int{
//     var result int
//     for _, num := range m{
//         result += num

//     }
//     return result
// }

// package main

// import (
//     "fmt"
// )
// func FindMaxKey(m map[int]int) int{
//     var maxnum int
//     tr := true
//     for num :== range m{
//         if tr{
//             maxnum = numnum
//             tr = false
//         } else if maxnum < num{
//             maxnum = num
//         }
//     }return maxnum
// }

// package main

// import (
// 	"sort"
// 	"strings"
// )

// func IsPalindrome(input string) bool {
//     var result []string
//     input = strings.ReplaceAll(input, " ", "" )
//     for chis := range strings.ToLower(input){
//         result = append(result, (string(input[(len(input)-chis)-1])))
//     }
// 	var stroka string
//     stroka = strings.Join(result, "")
//     if stroka == input {
//         return true
//     } else{
//         return false
//     }
// }

// package main

// import "fmt"

// func isLatin(input string) bool {
// 	var result string = ""
// 	for str := range input {
// 		if input[str] <= 128 {
// 			resturn true
// 		} else if
// 			return false
// 	}
// }

// func main() {
// 	word := "yandexlyceumgolang"
// 	fmt.Println(isLatin(word))
// }

// package main

// import (
// 	"fmt"
// )

// func ConcatenateStrings(str1, str2 string) string{
// 	result := strings.Join(str1, "", str2)
// 	return result
// }

// package main

// import (
// 	"fmt"
// )

// func StringLength(input string) int {
// 	return(len(input)
// }

// package main

// import (
// 	"fmt"
// )

// func Clear(nums []int, x int) []int {
// 	var b []int
// 	for c := range nums {
// 		if x == nums[c] {
// 			continue
// 		}
// 		b = append(b, nums[c])
// 	}
// 	return b
// }
// func main() {
// 	d := []int{}
// 	fmt.Println(Clear(d, 2))

// }

// package main

// import (
// 	"errors"
// 	"testing"
// )

// func UnderLimit(nums []int, limit int, n int) ([]int, error) {
// 	if n < 0 {
// 		return nil, errors.New("n должно быть больше или равно 0")
// 	}
// 	if len(nums) == 0 {
// 		return nil, errors.New("Слайс пустой")
// 	}
// 	c := make([]int, 0)
// 	for i := 0; i < len(nums); i++ {
// 		if nums[i] < limit {
// 			c = append(c, nums[i])
// 		}
// 		if len(c) >= n {
// 			break
// 		}
// 	}
// 	if len(c) == 0 {
// 		return nil, errors.New("нет элементов меньше limit")
// 	}
// 	return c, nil
// }

// func TestUnderLimit(t *testing.T) {
// 	type test struct {
// 		nums      []int
// 		n         int
// 		limit     int
// 		expected  []int
// 		wantError bool
// 	}

// 	tests := []test{
// 		{
// 			nums:      []int{4, 7, 89, 3, 21, 2, 5, 7, 32, 4, 6, 8, 0, 3, 4, 6, 2, 115, 12},
// 			n:         5,
// 			limit:     3,
// 			expected:  []int{2, 0, 2},
// 			wantError: false,
// 		},
// 		{
// 			nums:      nil,
// 			wantError: true,
// 		},
// 		{
// 			nums:      []int{},
// 			n:         5,
// 			limit:     3,
// 			expected:  []int{},
// 			wantError: false,
// 		},
// 		{
// 			nums:      []int{3, 5, 6},
// 			n:         5,
// 			limit:     10,
// 			expected:  []int{3, 5, 6},
// 			wantError: false,
// 		},
// 		{
// 			nums:      []int{-13, 0, 6},
// 			n:         1,
// 			limit:     -5,
// 			expected:  []int{-13},
// 			wantError: false,
// 		},
// 		{
// 			nums:      []int{},
// 			n:         -1,
// 			limit:     5,
// 			wantError: true,
// 		},
// 	}

// 	for _, tc := range tests {
// 		err := UnderLimit(tc.nums, tc.limit, tc.n)
// 		if tc.wantError {
// 			if err == nil {
// 				t.Fatalf("expec... File is too long to be displayed fully")
// 			}
// 		}
// 	}
// }

// package main

// import (
// 	"fmt"
// 	"sort"
// )

// func main() {
// 	slice := []int{1, 2, 3, 4, 6, 5, 7, 8, 9}
// 	fmt.Println(sort.IntsAreSorted(slice))
// 	sort.Ints(slice)
// 	fmt.Println(slice)
// }

// package main

// import (
// 	"fmt"
// )

// func main() {
// 	a := make([]int, 2, 3)
// 	a[0], a[1] = 0, 1
// 	b := append(a, 1)
// 	b[0] = 21

// 	fmt.Println(a)
// 	fmt.Println(b)
// }

// package main

// import (
// 	"fmt"
// )
// var a int
// a := make([]int, 2)
// a[0]=0
// a[1]=1
// fmt.Println(a)

// package main

// import (
// 	"fmt"
// )

// func PrettyArrayOutput(array [9]string) {
// 	for i := 0; i < (len(array) - 2); i++ {
// 		fmt.Printf("%d я уже сделал: %s\n", i+1, array[i])
// 	}
// 	for i := 7; i < len(array); i++ {
// 		fmt.Printf("%d не успел сделать: %s\n", i+1, array[i])
// 	}
// }

// func main() {
// 	input := [9]string{
// 		"проснуться",
// 		"позавтракать",
// 		"сходить в школу",
// 		"пообедать",
// 		"погулять с друзьями",
// 		"сделать домашнюю работу",
// 		"попрограммировать на Go",
// 		"поужинать",
// 		"лечь спать",
// 	}
// 	PrettyArrayOutput(input)
// }

// package main

// import (
// 	"fmt"
// )

// func FindMinMaxInArray(array [10]int) (int, int) {
// 	max, min := array[0], array[0]
// 	for i := 1; i < len(array); i++ {
// 		if min > array[i] {
// 			min = array[i]
// 		}
// 		if max < array[i] {
// 			max = array[i]
// 		}
// 	}
// 	return max, min
// }

// func main() {
// 	input := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
// 	max, min := FindMinMaxInArray(input)

// 	fmt.Print(max, min)
// }

// package main

// import (
// 	"fmt"
// 	"math"
// )

// func SqRoots() {
// 	var a, b, c float64
// 	fmt.Scanf("%f %f %f", &a, &b, &c)
// 	d := math.Pow(b, 2) - 4*a*c
// 	if d > 0 {
// 		r1 := (-b + math.Sqrt(d)) / (2 * a)
// 		r2 := (-b - math.Sqrt(d)) / (2 * a)
// 		if r1 < r2 {
// 			fmt.Printf("%.6f %.6f\n", r1, r2)
// 		} else {
// 			fmt.Printf("%.6f %.6f\n", r2, r1)
// 		}
// 	} else if d == 0 {
// 		r := (-b) / (2 * a)
// 		fmt.Printf("%.6f\n", r)
// 	} else {
// 		fmt.Println("0 0")
// 	}
// }
// func main() {
// 	SqRoots()
// }

// package main

// import (
// 	"fmt"
// )

// func FiveSteps(array [5]int) [5]int {
// 	var c [5]int
// 	for i := len(array) - 1; i >= 0; i-- {
// 		c[4-i] = array[i]
// 	}
// 	return c
// }
// func main() {
// 	input := [5]int{1, 2, 3, 4, 5}
// 	output := FiveSteps(input)
// 	for i := 0; i < len(output); i++ {
// 		fmt.Printf("%d ", output[i])
// 	}
// }

// package main

// import (
// 	"fmt"
// )

// func IsPowerOfTwoRecursive(N int) {
// 	if N==1 {
// 		fmt.Println("YES")
// 	} else if N%2==0{
// 		IsPowerOfTwoRecursive(N/2)
// 	} else if N<1{
// 		fmt.Println("NO")
// 	} else {
// 		fmt.Println("NO")
// 	}
// }

// package main

// import (
// 	"fmt"
// )

// func main() {
// 	fmt.Println(10 / 2)
// }

// package main

// import (
// 	"fmt"
// )

// func SumDigitsRecursive(n int) int {
// 	if n == 0 {
// 		return 0
// 	}
// 	return n%10 + SumDigitsRecursive(n/10)
// }
// func main() {
// 	fmt.Println(SumDigitsRecursive(123))
// }

// package main

// import (
// 	"fmt"
// )

// func Add(a, b float64) float64 {
// 	return a + b
// }

// func Multiply(a, b float64) float64 {
// 	return a * b
// }
// func PrintNumbersAssending(n int) {
// 	for i := 1; i <= n; i++ {
// 		fmt.Printf("%d", i)
// 	}
// }
// func main() {
// 	PrintNumbersAssending(25)
// }

// package main

// import (
// 	"fmt"
// 	"math"
// )

// func SqRoots() {
// 	var a, b, c float64
// 	fmt.Scanf("%f %f %f", &a, &b, &c)
// 	d := math.Pow(b, 2) - 4*a*c
// 	if d > 0 {
// 		r1 := (-b + math.Sqrt(d)) / (2 * a)
// 		r2 := (-b - math.Sqrt(d)) / (2 * a)
// 		if r1 < r2 {
// 			fmt.Printf("%.6f %.6f\n", r1, r2)
// 		} else {
// 			fmt.Printf("%.6f %.6f\n", r2, r1)
// 		}
// 	} else if d == 0 {
// 		r := (-b) / (2 * a)
// 		fmt.Printf("%.6f\n", r)
// 	} else {
// 		fmt.Println("0 0")
// 	}
// }

// func main() {
// 	fmt.Println(SqRoots)
// }

// package main

// import (
// 	"fmt"
// )

// func fib(n int) int {
// 	if n < 2 {
// 		return n
// 	}
// 	return fib(n-1) + fib(n-2)
// }
// func main() {
// 	fmt.Println(fib(4))
// }

// package main

// import (
// 	"fmt"
// )

// func main() {
// 	var a int
// 	var c int
// 	fmt.Scanln(&a)
// 	for i := range 100 {
// 		c += i
// 	}
// 	for i := range 100 {
// 		if a == i+c {
// 			fmt.Println(a, i, c)
// 		}
// 	}
// }

// package main

// import (
// 	"fmt"
// )

// func main() {
// 	var a int
// 	var c int
// 	fmt.Scanln(&a)
// 	for i := range a {
// 		if (i+1)%3 != 0 && (i+1)%5 != 0 {
// 			c += (i + 1)
// 		}
// 	}
// 	fmt.Println(c)
// }

// package main

// import (
// 	"fmt"
// )

// func main() {
// 	var a int
// 	fmt.Scanln(&a)
// 	for i := range a {
// 		if (i+1)%3 == 0 {
// 			fmt.Println((i + 1))
// 		}
// 	}
// }

// package main

// import (
// 	"fmt"
// )

// func main() {
// 	var a int
// 	var c int
// 	fmt.Scanln(&a)
// 	if a < 0 {
// 		fmt.Println("Некорректный ввод")
// 	} else {
// 		for i := 0; i < a; i++ {
// 			if (i+1)%2 != 0 {
// 				c += (i + 1)
// 			}

// 		}
// 		fmt.Println(c)
// 	}
// }

// package main

// import (
// 	"fmt"
// )

// func main() {
// 	var a int
// 	var c int
// 	fmt.Scanln(&a)
// 	if a < 0 {
// 		fmt.Println("Некорректный ввод")
// 	}
// 	for i := 0; i < a; i++ {
// 		if (i+1)%2 != 0 {
// 			c += (i + 1)
// 		}

// 	}
// 	fmt.Println(c)
// }

// package main

// import (
// 	"fmt"
// )

// func main() {
// 	var a int
// 	var c int
// 	c += 1
// 	fmt.Scanln(&a)
// 	for i := 0; i < a; i++ {
// 		c += c * (i)
// 	}
// 	fmt.Println(c)
// }

// package main

// import (
// 	"fmt"
// )

// func main() {
// 	var a int
// 	fmt.Scanln(&a)
// 	for i := range 10 {
// 		fmt.Printf("%d * %d = %d\n", a, a, a*(i+16))
// 	}
// }

//  package main

//  import "fmt"
// 	func main() {
// 		var age int
// 		fmt.Println("Введите ваш возраст")
// 		fmt.Scanln(&age)
// 		fmt.Println(age)
// 	}
// package main

// import "fmt"

// func mint() {
// 	a := "abcdefg"
// 	b.parselnt
// }

// package main

// import (
// 	"fmt"
// )

//  func main() {
// 	var a int
// 	fmt.Scanln(&a)
// 	if a >= 50 {
// 		fmt.Println("Отличная работа")
// 	} else {
// 		fmt.Println("Нужно усерднее потрудиться")
// 	}
// }

// package main

// import (
// 	"fmt"
// )

// func main() {
// 	for i := range "Hello" {
// 		fmt.Println(string(i))
// 	}
// }

// package main

// import (
// 	"fmt"
// )

// func main() {
// 	for i := range 10 {
// 		fmt.Println(i + 1)
// 	}
// }

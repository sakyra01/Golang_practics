package main

import "fmt"

// Функции с переменным количеством аргументов

func sums(nums ...int){
	total := 0
	for _, v := range nums {
		total += v
	}
	fmt.Println(total) // Можно вызвать как sums(2, 4, 6) - output 12 или sums(42, 8) - output 50
}
	s := []int{42,33,22,11}
	sums(s...)  // передача массива s в качетве аргумента функции sums

// There is no clasesses in go but go have structures
type Contact struct {
	name string
	age  int
}

//Example
lx := Contact{name: "Andre", age: 42}
lx.age = 9
fmt.Println(lx.age)  // return 9
fmt.Println(lx.name) // return Andre
	
// Указатели на структуры
x := Contact {name: "James", age: 42}
p := &x
fmt.Println(p.age)

p := &Contcats("Jo", 14) // p - это указатель на только что созданный экземпляр структуры Contact.


// Методы 
func (x Cont) welcome(){   // структура метода func (x Car) test() {}
	fmt.Println(x.name)
	fmt.Println(x.age)
}

// Method using example
package main

import "fmt"

type Contact struct {
  name string									
  age int
}

func (x Contact) welcome() {
  fmt.Println(x.name)
  fmt.Println(x.age)
}

func main() {
  x := Contact{"Джеймс", 42}
  x.welcome()
}



// Function

func welcome() {
	fmt.Println("Welcome!")
}

// Arguments in Function

func welcome2(name string) {
	fmt.Println("Welcome!, " + name)
}

// Func with few arguments
func convert(kli, zli int) {
	fmt.Println(kli + zli)
}

// Func with return value for main code
func concat(a, b string) string {
	return a + b
}

// Func returns few values
func swap(r, u int) (int, int) {
	return u, r
}

// Defer, function will work after main function, Defer usually using to clear some used files, links
/*
package main

import "fmt"

func welcome() {
  fmt.Println("Добро пожаловать")
}

func main() {
  defer welcome() 2)
  fmt.Println("Эй!") 1)
}
*/

// Указатели
// Функция принимающая аргумент - указатель на целое число и устанавливает его значение в 0
func zero(cd *int) {
	*cd = 0
}

// Вывод адреса переменной в памяти
func adres() {
	var x string
	fmt.Scanln(&x)
	fmt.Println(&x)
}

func main() {
	adres()
	welcome()
	welcome2("Amer")
	welcome2("Jhon")
	convert(7, 9)
	concat("concat", "enation")
	swap(10, 20)

	fmt.Println(swap(10, 20))


	// Карты(maps) - аналог словарей в python
	m := make(map[string]int)
	m["James"] = 42
	m["Amy"] = 24
	fmt.Println(m["James"]) // выведет 42

	m := map[string]int{
		"James": 42,
		"Amy": 24
	}
	fmt.Println(m["Amy"]) // выведет 24 

	delete (m, "James") // delete element with key "James" from m map 

	// Mассивы
	var a [5]int
	aa := [5]int{0,2,4,6,8}
	a[0] = 8
	a[1] = 42
	fmt.Println(a[1])

	//Срезы
	aaa := [5]int{0,2,4,6,8}
	var s []int = aaa[1:3]
	fmt.Println(s)  // [2 4]

	a := make([]int,5)
	a = append(a, 8)
	fmt.Println(a)  // [0 0 0 0 0 8]

	//Range


	// Input and output string value
	var input string
	fmt.Scanln(&input)
	fmt.Println("Hey man")

	// Input and output integer value*2
	var input2 int
	fmt.Scanln(&input2)
	fmt.Println(input2 * 2)

	// Reand view values on input
	var h, p, l int
	fmt.Scanln(&h, &p, &l)
	fmt.Println(h, p, l)

	// Simple work
	const pi = 3.14
	var a, b int = 9, 10
	i, j, k := 3, 6, 8
	a = a + b
	num := "Veselo"
	var x bool

	fmt.Println("Hello world!")
	fmt.Println(a, i, j, k)
	fmt.Println(num)
	fmt.Println(x)
	fmt.Println(pi)

	//Operator if, else

	z := 14
	if z > 18 {
		fmt.Println("Accepted")
	} else {
		fmt.Println("Forbiden")
	}

	//Using value only in if,else structure
	if t := 42; t > 18 {
		fmt.Println("Accepted")
	} else {
		fmt.Println("Forbidden")
	}

	//Switch like elif
	good := 3
	switch good {
	// compare good value with each case value
	case 1:
		fmt.Println("One")
	case 2:
		fmt.Println("Two")
	case 3:
		fmt.Println("Three")
	default:
		fmt.Println(good)
	}

	// For
	for i := 0; i < 5; i++ {
		fmt.Println(i) // 0 1 2 3 4
	}

	sums := 0
	for i := 1; i <= 3; i++ {
		sums += i
	}
	fmt.Println(sums) // 6

	// Short For
	sum := 1
	res := 0
	for sum <= 1000 {
		res += sum
		sum++
	}
	fmt.Println(res)

	rez := 0
	for i := 0; i < 10; i += 3 {
		rez += i
	}
	fmt.Println(rez)


	/* Task  
	Вы создаете программу для анализа результатов спортивных матчей и подсчета очков заданной команды.
	Результаты матчей хранятся в массиве results.
	Каждый матч имеет один из следующих результатов:
	"w" - выиграл
	"l" - проиграл
	"d" - ничья
	Победа добавляет три очка, ничья - одно очко, а проигранный матч не добавляет очков.
	Ваша программа должна принять на вход результат последнего матча и добавить его в массив результатов. 
	После этого необходимо вычислить и вывести общее количество очков, набранных командой по результатам матчей из массива results.

	Sample Input:
	w
	Sample Output:
	22
	----------------------------------------------------------------------------------------------------------------------------------
*/
	package main

import "fmt"

func main() {

    results := []string{"w", "l", "w", "d", "w", "l", "l", "l", "d", "d", "w", "l", "w", "d"}
    var x string
    fmt.Scanln(&x)
    results = append(results, x)
    get_results(results...)
}

func get_results(values ...string){
    total := 0 
    for _, v := range values {
        switch v{
            case "w":
                total += 3
            case "d":
                total += 1
            default:
                continue
        }
    }
    fmt.Println(total)
}



// Gorutin

package main
import (
    "fmt"
    "time"
)

func out(from, to int) {
  for i:=from; i<=to; i++ {
    time.Sleep(50 * time.Millisecond)
    fmt.Println(i)
  }
}
func main() {
  out(0, 5)
  out(6, 10)
}

// Каналы и горутины

package main

import (
    "fmt"
    "time"
    )

func out(from, to int, ch chan bool) {
    for i:=from; i<=to; i++ {
        time.Sleep(50 * time.Millisecond)
        fmt.Println(i)
    }
    ch <- true
}

func main() {
    ch := make(chan bool)

    go out(0, 5, ch)
    go out(6, 10, ch)

    <-ch
    <-ch
}


// Select - Оператор select используется для ожидания нескольких операций канала.
evenCh := make(chan int)
sqCh := make(chan int)

go evenSum(0, 100, evenCh)
go squareSum(0, 100, sqCh)

select {
  case x := <- evenCh:
     fmt.Println(x)
  case y := <- sqCh:
     fmt.Println(y)		
}

/* 
Оператор select ожидает, пока канал получит данные, и выполняет его case.
Это означает, что будет выполнен только один из case - тот, который соответствует каналу, получившему данные первым.
Если оба канала получают данные одновременно, один из case выбирается случайным образом.
*/



// Loader
/*
Вы создаете программу для загрузки файлов.
Чтобы ускорить загрузку, вы решили использовать многопоточность. Каждая загрузка файла будет выполняться как отдельная горутина.

Для имитации загрузки файла функция download() должна принимать в качестве аргумента целое число 
(которое будет играть роль размера файла) и подсчитывать сумму всех целых чисел от 0 до этого числа (включительно).

Данная программа принимает три размера файлов в качестве входных данных и вызывает функцию download() как горутину для каждого файла.
Завершите программу, объявив функции download() и отправив результат их работы в main() с помощью каналов. 
Результаты работы трех функций должны быть сложены и выведены на экран.
*/
package main
import "fmt"

//определите функцию download()

func download(size int, ch chan int) {
    total := 0
    for i:=0; i<=size; i++{
        total += i
    }
    ch <- total

}

func main() {
    ch1 := make(chan int)
    ch2 := make(chan int)
    ch3 := make(chan int)

    var s1, s2, s3 int
    fmt.Scanln(&s1)
    fmt.Scanln(&s2)
    fmt.Scanln(&s3)

    go download(s1, ch1)
    go download(s2, ch2)
    go download(s3, ch3)
    
    chennel := <-ch1 + <-ch2 + <-ch3
    fmt.Println(chennel)
}

}

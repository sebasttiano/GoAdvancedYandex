// Избегайте коммуникаций через общую, разделённую память, делитесь данными через сообщения
//(Don't communicate by sharing memory, share memory by communicating).

package main

var (
	salary = 100000
	bonus  = 20000
	rent   = 5000
)

// переменная, которую нужно менять из нескольких потоков
var Balance int

// канал коммуникации между горутинами
var ch chan int

// эту функцию можно безопасно вызывать из разных горутин конкурентно,
// она не изменяет состояние переменной в памяти, а только пишет в канал
func Deposit(i int) {
	ch <- i
}

// эта функция тоже
func Withdraw(i int) {
	ch <- -i
}

// а это единственная функция с доступом к Balance,
// она принимает данные из канала от многих источников
// и непосредственно изменяет состояние памяти
func Account() {
	for {
		Balance += <-ch
	}
}

func main() {
	// ...
	go Account()
	// ...
	go Deposit(salary)
	// ...
	go Deposit(bonus)
	// ...
	go Withdraw(rent)
	// ...
}

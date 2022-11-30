// Напишите код, в котором несколько горутин увеличивают значение целочисленного счётчика
// и синхронизируют свою работу через канал.
// Нужно предусмотреть возможность настройки количества используемых горутин
// и конечного значения счётчика, до которого его следует увеличивать.
//
// Попробуйте реализовать счётчик с элементами ООП (в виде структуры и методов структуры).
// Попробуйте реализовать динамическую проверку достижения счётчиком нужного значения
package main

import (
	"fmt"
	"sync"
)

type Counter struct {
	ch  chan int
	sum int
	wg  sync.WaitGroup
}

const gorutineCount = 50
const maxCounter = 1000

// Transaction Метод пишет в канал ch заранее определенно раз
func (c *Counter) Transaction() {
	defer c.wg.Done()
	for {
		if c.sum >= maxCounter {
			return
		}
		c.ch <- 1
		c.sum--
	}
}

// NewTransaction Метод синхронизирует канал ch со счетчиком
func NewTransaction() (c *Counter) {
	c = &Counter{make(chan int), 0, sync.WaitGroup{}}

	go func() {
		defer c.wg.Done()
		for increment := range c.ch {
			c.sum = c.sum + increment
		}
	}()
	return
}

func main() {
	c := NewTransaction()

	c.wg.Add(gorutineCount)
	for i := 1; i <= gorutineCount; i++ {
		go c.Transaction()
	}
	c.wg.Wait()

	fmt.Printf("Счетчик равен макксимуму: = %d", c.sum)
}

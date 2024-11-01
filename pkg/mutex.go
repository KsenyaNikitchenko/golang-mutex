package mutex

import "fmt"

type Mutex struct {
	Count int
	done  chan bool
}

func (m *Mutex) Unlock() {
	m.done <- true
}

func (m *Mutex) Wait() {
	for i := 0; i < m.Count; i++ {
		<-m.done
	}
}

func WorkMutex() {
	m := Mutex{Count: 5, done: make(chan bool, 5)}

	for i := 0; i < 5; i++ {
		go func() {
			defer m.Unlock()
			fmt.Println("Hello, 世界")
		}()
	}

	m.Wait()
	fmt.Println("Работа всех потоков завершена")
}

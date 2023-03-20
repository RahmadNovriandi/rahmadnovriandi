package main

import (
	"fmt"
)

func main() {

	for i := 1; i <= 4; i++ {
		w.Add(1)
		go func(id int) {
			m.Lock()
			defer m.Unlock()
			if id%2 == 1 {
				fmt.Printf("%v %d\n", coba, id)
			} else {
				fmt.Printf("%v %d\n", bisa, id)
			}
			w.Done()
		}(i)
	}
}

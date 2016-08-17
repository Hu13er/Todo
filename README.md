`Todo` is a s simple Concurrency package

## Sample:
```go
package main

import (
    "fmt"
	"log"
	"time"

	"github.com/Hub13er/Todo"
)

func main(){

	todo := Todo.NewTodo(func(i interface{}){
			log.Println("my", i , "th beat =)")
			time.Sleep(1 * time.Second)
		})

	for i := 1 ; i <= 5 ; i++ {
		todo.Push(i)
	}

	p := ""

	todo.Run()

	time.Sleep(2 * time.Second)

	todo.Stop()

	fmt.Scanln(&p)

	// todo.Run()

	todo.Run()

	for i := 100 ; i <= 105 ; i++ {
		todo.Push(i)
	}

	fmt.Scanln(&p)
}

```
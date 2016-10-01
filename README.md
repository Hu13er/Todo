`Todo` is a simple Concurrency package
It's like a buffered channel with unlimited size.

## Install:
```bash
go get -u -v github.com/Hu13er/Todo
```

## Sample:
```go
package main

import (
        "fmt"
        "time"

	"github.com/Hub13er/Todo"
)

func main(){
	todo := Todo.NewTodo(func(i interface{}){
			fmt.Println("my", i , "th beat =)")
			time.Sleep(1 * time.Second)
		})

	for i := 1 ; i <= 5 ; i++ {
		todo.Push(i)
	}
	todo.Run()
	fmt.Scanln()
	todo.Stop()
	fmt.Scanln()
	todo.Run()
	for i := 100 ; i <= 105 ; i++ {
		todo.Push(i)
	}
        fmt.Println("pushed.")
	fmt.Scanln()
}

```


package main

import (
	"encoding/json"
	"fmt"
	jpc "github.com/quantum73/go_package_example/json_placeholder_client"
	"log"
	"sync"
)

var todoBaseUrl = "https://jsonplaceholder.typicode.com/todos"
var todosCount = 100

func main() {
	jpClient := jpc.NewJSONPlaceholderClient(todosCount)
	var wg sync.WaitGroup

	wg.Add(todosCount)
	for i := 0; i < todosCount; i++ {
		go func(idx int) {
			defer wg.Done()

			detailTodoBaseUrl := fmt.Sprintf("%s/%d", todoBaseUrl, idx+1)
			err := jpClient.AddFromUrl(idx, detailTodoBaseUrl)
			if err != nil {
				log.Printf("Error during request to `%s`: %s", detailTodoBaseUrl, err)
			}
		}(i)
	}
	wg.Wait()
	fmt.Println("All todos has been received!")

	resultExample, err := jpClient.GetResultById(0)
	if err != nil {
		log.Fatalf("Error during get result by id: %s", err)
	}

	todo := jpc.TodoResponse{}
	err = json.Unmarshal([]byte(resultExample), &todo)
	if err != nil {
		log.Fatalf("Error during JSON unmarshalling todo object: %s", err)
	}

	fmt.Printf("Todo example:\n%s\n", todo)
}

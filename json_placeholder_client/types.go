package json_placeholder_client

import (
	"encoding/json"
	"fmt"
	"log"
)

type TodoResponse struct {
	UserId    int    `json:"userId"`
	Id        int    `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

func (t TodoResponse) String() string {
	todoAsJson, err := json.MarshalIndent(t, "", "    ")
	if err != nil {
		log.Printf("todo json marshal error: %v\n", err)
		return fmt.Sprintf(
			"TodoResponse[userId=%d, Id=%d, Title=%s, Completed=%t]",
			t.UserId, t.Id, t.Title, t.Completed,
		)
	}
	return string(todoAsJson)
}

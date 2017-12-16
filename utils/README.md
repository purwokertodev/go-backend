
## Http Request

### Usage

```go
package main

import (
	"bytes"
	"encoding/json"
	"fmt"

	"github.com/purwokertodev/go-backend/utils"
)

type Post struct {
	UserID int    `json:"userId"`
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Body   string `json:"body"`
}

type User struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Address  struct {
		Street  string `json:"street"`
		Suite   string `json:"suite"`
		City    string `json:"city"`
		Zipcode string `json:"zipcode"`
		Geo     struct {
			Lat string `json:"lat"`
			Lng string `json:"lng"`
		} `json:"geo"`
	} `json:"address"`
	Phone   string `json:"phone"`
	Website string `json:"website"`
	Company struct {
		Name        string `json:"name"`
		CatchPhrase string `json:"catchPhrase"`
		Bs          string `json:"bs"`
	} `json:"company"`
}

type Users []User

func main() {
	var post Post
	post.ID = 101
	post.UserID = 1
	post.Title = "Golang"
	post.Body = "Golang is awesome"

	payload, _ := json.Marshal(post)

	var resp Post
	req := utils.NewRequest(10)

	headers := make(map[string]string)
	headers["Content-Type"] = "application/json"
	headers["Accept"] = "application/json"

	err := req.Req("POST", "https://jsonplaceholder.typicode.com/users", bytes.NewBuffer(payload), &resp, headers)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(resp)
}
```

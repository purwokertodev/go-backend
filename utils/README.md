
## Http Request

### Usage

  - `POST`
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

  - `GET`

    ```go
    package main

    import (
    	"fmt"

    	"github.com/purwokertodev/go-backend/utils"
    )

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
    	var users Users

    	req := utils.NewRequest(10)

    	headers := make(map[string]string)
    	headers["Content-Type"] = "application/json"
    	headers["Accept"] = "application/json"

    	err := req.Req("GET", "https://jsonplaceholder.typicode.com/users", nil, &users, headers)
    	if err != nil {
    		fmt.Println(err)
    	}

    	fmt.Println(users)
    }
    ```

## Email

### Usage

  - Html Email Template

    ```html
    <!DOCTYPE html PUBLIC "-//W3C//DTD XHTML 1.0 Transitional//EN"
            "http://www.w3.org/TR/xhtml1/DTD/xhtml1-transitional.dtd">
    <html>

    </head>

    <body>
    <p>
        Hi {{.Username}}
        <h4>Your Account Registration Success</h4>
        <p>Please follow link bellow, to complete your Registration</p>
        <a href="{{.URL}}">Confirm email address</a>
    </p>

    </body>

    </html>
    ```

  - Golang Code

    ```go
    import (
    	"fmt"

    	"github.com/purwokertodev/go-backend/utils"
    )

    func main() {
    	authEmail := "wuriyanto007@gmail.com"
    	authPassword := "sudahlah"
    	authHost := "smtp.gmail.com"
    	address := "smtp.gmail.com:587"
    	to := []string{"wuriyanto48@yahoo.co.id"}
    	from := "wuriyanto007@gmail.com"
    	subject := "Golang email"
    	body := "Golang email sent..."
    	email := utils.NewEmail(to, address, from, subject, body, authEmail, authPassword, authHost)

    	emailData := struct {
    		Username string
    		URL      string
    	}{
    		Username: "Wuriyanto",
    		URL:      "wuriyanto.com",
    	}

    	err := execute(email, "email_template.html", emailData)
    	if err != nil {
    		fmt.Println(err)
    	}
    	fmt.Println("email sent")

    }

    func execute(u utils.EmailSender, fileName string, data interface{}) error {
    	err := u.SetTemplate(fileName, data)
    	if err != nil {
    		return err
    	}

    	err = u.Send()
    	if err != nil {
    		return err
    	}

    	return nil
    }
    ```

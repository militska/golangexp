package main

import (
	"fmt"
	"github.com/go-resty/resty/v2"
	"net/http"
)

type RequestError struct {
	StatusCode   int
	ErrorMessage string
}

func (req *RequestError) Error() string {
	return req.ErrorMessage
}

type ValidateError struct {
	Field        string
	ErrorMessage string
}

func (req *ValidateError) Error() string {
	return req.ErrorMessage
}

func test(name string) (bool, error) {
	if len(name) < 5 {
		return false, &ValidateError{
			"name",
			"Имя должно содержать больше 5 символов",
		}
	}

	resp, err := resty.New().R().Post(fmt.Sprintf("test.ru?name=%s", name))

	if err != nil {
		return false, &RequestError{
			StatusCode:   http.StatusNotFound,
			ErrorMessage: "Ошибка получения данных",
		}
	}

	switch resp.StatusCode() {
	case 200:
		return true, nil
	default:
		return false, err

	}
}
func main() {
	res, err := test("4dddddd4")

	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(res)
	fmt.Println("Hello, playground")
}

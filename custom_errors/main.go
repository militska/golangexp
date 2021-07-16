package main

import (
	"fmt"
	"github.com/go-resty/resty/v2"
	"net/http"
	"strconv"
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
	return fmt.Sprintf("%s: %s", req.Field, req.ErrorMessage)
}

func test(name string) (bool, error) {

	// вот так делать не нужно
	// сюда должны попадать уже отвалидированные данные
	// но это для того что бы показать
	if len(name) < 5 {
		return false, &ValidateError{
			"name",
			"Имя должно содержать больше 5 символов",
		}
	}

	resp, err := resty.New().R().Post(fmt.Sprintf("test.ru?name=%s", name))

	if err != nil {
		return false, &RequestError{
			StatusCode:   http.StatusBadRequest,
			ErrorMessage: "Ошибка запроса: " + err.Error(),
		}
	}

	switch resp.StatusCode() {
	case http.StatusOK:
		return true, nil
	default:
		return false, &RequestError{
			StatusCode:   resp.StatusCode(),
			ErrorMessage: "Ошибка получения данных",
		}

	}
}

func getHttpCodeFromError(err error) int {
	var httpCode int
	switch v := err.(type) {
	case *RequestError:
		errResponse := err.(*RequestError)
		httpCode = errResponse.StatusCode
	case *ValidateError:
		httpCode = http.StatusNotFound
	default:
		fmt.Println("Unknown type error %T", v)
		httpCode = http.StatusUnprocessableEntity
	}

	return httpCode
}

func main() {
	res, err := test("4dddddd4")

	if err != nil {
		fmt.Println(err)
		fmt.Println("Status code: " + strconv.Itoa(getHttpCodeFromError(err)))
	}
	fmt.Println(res)
	fmt.Println("Hello, playground")
}

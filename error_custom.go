package main

import "fmt"

type validationError struct {
	Message string
}

type notFoundError struct {
	Message string
}

func (v validationError) Error() string {
	return v.Message
}

func (n notFoundError) Error() string {
	return n.Message
}

func SaveData(id string, data any) error {
	if id == "" {
		return &validationError{"Validation Error"}
	}

	if id != "LOL" {
		return &validationError{"Data not Found"}
	}

	return nil
}

func main() {
	err := SaveData("1", nil)
	if err != nil {
		// if validationErr, ok := err.(*validationError); ok {
		// 	fmt.Println("validation Error:", validationErr.Error())
		// } else if notFoundError, ok := err.(*notFoundError); ok {
		// 	fmt.Println("Not Found Error:", notFoundError.Message)
		// } else {
		// 	fmt.Println("Unknown Error:", err.Error())
		// }

		switch finalError := err.(type) {
		case *validationError:
			fmt.Println("validation Error:", finalError.Error())
		case *notFoundError:
			fmt.Println("Data Not Found:", finalError.Error())
		default:
			fmt.Println("Unknown Error:", finalError.Error())
		}
	} else {
		fmt.Println("Sukses")
	}
}

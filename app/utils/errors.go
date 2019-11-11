package utils

import (
	"fmt"
)

func FormatErrors(tag string, field string, param string) string {
	switch tag {
	case "required":
		return fmt.Sprintf("%v is required.", field)
	case "email":
		return fmt.Sprintf("%v is not a valid email", field)
	case "gte":
		return fmt.Sprintf("%v should be greater than %v characters", field, param)
	}

	return ""
}

func ExtractErrorMessages(errors []error) []string {
	messages := []string{}
	for _, err := range errors {
		messages = append(messages, err.Error())
	}

	return messages
}

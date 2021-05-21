package utility

import (
	"errors"
	"regexp"
	"strconv"
)

// ValidateRequireAndLengthAndRegex is used to validate any input data but in string type
func ValidateRequireAndLengthAndRegex(value string, isRequired bool, minLength, maxLength int, regex, fieldName string) error {
	length := len(value)
	Re := regexp.MustCompile(regex)

	if isRequired == true && length < 1 {
		return errors.New(fieldName + "is required")
	}

	// Min length check
	// If params min length value is zero that indicates, there will be not min length check
	if minLength != 0 && length > 1 && length < minLength {
		return errors.New(fieldName + " must be min " + strconv.Itoa((minLength)))
	}

	// Max length check
	// If params max length value is zero that indicates, there will be no max length check
	if maxLength != 0 && length > 1 && length > maxLength {
		return errors.New(fieldName + "mus be max " + strconv.Itoa(maxLength))
	}

	// Regex check
	if !Re.MatchString(value) {
		return errors.New("Invalid " + fieldName)
	}

	return nil
}

// NewHTTPError creates error model that will send as http response
// if any error occors
func NewHTTPError(errorCode string, statusCode int) map[string]interface{} {

	m := make(map[string]interface{})
	m["error"] = errorCode
	m["error_description"], _ = errorMessage[errorCode]
	m["code"] = statusCode

	return m
}

// NewHTTPCustomError creates error model that will send as http response
// If any error occurs
func NewHTTPCustomError(errorCode, errorMsg string, statusCode int) map[string]interface{} {

	m := make(map[string]interface{})
	m["error"] = errorCode
	m["error_description"] = errorMsg
	m["code"] = statusCode

	return m
}

// error code
const (
	InvalidUserId       = "InvalidUserId" // in case userId is not exists
	InternalError       = "InternalError" // in case, any internal server error occurs
	UserNotFound        = "userNotFound"  // if user not found
	InvalidBindingModel = "invalidBindingModel"
	EntityCreationError = "entityCreationError"
	Unauthorized        = "unauthorized" // in case, try to access restricted resource
	BadRequest          = "badRequest"
	UserAlreadyExists   = "userAlreadyExists"
)

// Error code with
var errorMessage = map[string]string{
	"invalidUserID":       "invalid user id",
	"internalError":       "an internal error occurs",
	"userNotFound":        "user could not be found",
	"invalidBindingModel": "model could not be bound",
	"entityCreationError": "could not create entity",
	"unauthorized":        "an unauthorized access",
	"userAlreadyExists":   "user already exists",
}

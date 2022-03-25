package twilio

import (
	twilioClient "github.com/twilio/twilio-go/client"
)

// HTTPCode represents the HTTP status code
type HTTPCode int

const (
	// HTTPCodeOK represents the HTTP status code 200
	HTTPCodeOK HTTPCode = 200
	// HTTPCodeBadRequest represents the HTTP status code 400
	HTTPCodeBadRequest HTTPCode = 400
	// HTTPCodeUnauthorized represents the HTTP status code 401
	HTTPCodeUnauthorized HTTPCode = 401
	// HTTPCodeNotFound represents the HTTP status code 404
	HTTPCodeNotFound HTTPCode = 404
	// HTTPCodeInternalServerError represents the HTTP status code 500
	HTTPCodeInternalServerError HTTPCode = 500
)

func httpCodeFromError(err error) HTTPCode {
	switch err.(type) {
	case *twilioClient.TwilioRestError:
		return HTTPCode(err.(*twilioClient.TwilioRestError).Status)
	default:
		return HTTPCodeInternalServerError
	}
}

func newError(code HTTPCode, msg string) error {
	return &twilioClient.TwilioRestError{
		Status:  int(code),
		Message: msg,
	}
}

// Ok returns true if the error is nil or the error is a TwilioRestError with a status code of 200
func Ok(err error) bool {
	return httpCodeFromError(err) == HTTPCodeOK
}

// NotFound returns true if the error is nil or the error is a TwilioRestError with a status code of 404
func NotFound(err error) bool {
	return httpCodeFromError(err) == HTTPCodeNotFound
}

// BadRequest returns true if the error is nil or the error is a TwilioRestError with a status code of 400
func BadRequest(err error) bool {
	return httpCodeFromError(err) == HTTPCodeBadRequest
}

// Unauthorized returns true if the error is nil or the error is a TwilioRestError with a status code of 401
func Unauthorized(err error) bool {
	return httpCodeFromError(err) == HTTPCodeUnauthorized
}

// Is returns true if the error is nil or the error is a TwilioRestError with a status code equal to the given status code
func Is(err error, code ...HTTPCode) bool {
	c := httpCodeFromError(err)
	for _, v := range code {
		if c == v {
			return true
		}
	}
	return false
}

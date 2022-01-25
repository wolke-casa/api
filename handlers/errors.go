package handlers

import "fmt"

func ErrInvalidAuthorization() (string, int) { return "Invalid authorization token provided.", 403 }
func ErrMissingData(s string) (string, int)  { return fmt.Sprintf(`Missing "%s" in request.`, s), 400 }
func ErrInvalidDataInRequest(s string) (string, int) {
	return fmt.Sprintf("Invalid %s in request.", s), 400
}
func ErrUnknownErrorOccurred() (string, int) {
	return "Unknown error occurred, please try again in a second.", 500
}
func ErrFileTooBig(s string) (string, int) {
	return fmt.Sprintf("Uploaded file is too big, file size limit is %s.", s), 413
}
func ErrInvalidData(s string, v string) (string, int) {
	return fmt.Sprintf("Invalid %s, valid are %s.", s, v), 400
}
func ErrUserAlreadyHasApiKey() (string, int) { return "That user already has an API key.", 400 }
func ErrResourceNotFound() (string, int)     { return "That resource wasnt found.", 404 }

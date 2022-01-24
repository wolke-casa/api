package handlers

type statusError struct {
	Error  string
	Status int
}

var (
	ErrInvalidAuthorization = statusError{"Invalid authorization token provided.", 401}
	ErrMissingData          = statusError{`Missing "{}" in request.`, 400}
	ErrInvalidInRequest     = statusError{"Invalid {} in request.", 400}
	ErrUnknownErrorOccurred = statusError{"Unknown error occurred, please try again in a second.", 500}
	ErrFileTooBig           = statusError{"Uploaded file is too big, file size limit is {}.", 413}
	ErrInvalid              = statusError{"Invalid {}, valid are {}.", 400}
	ErrUserAlreadyHasApiKey = statusError{"That user already has an API key.", 400}
	ErrResourceNotFound     = statusError{"That resource wasnt found.", 404}
)

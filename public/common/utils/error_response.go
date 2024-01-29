package utils

import (
	"github.com/go-chi/render"
	"net/http"
)

type ErrResponse struct {
	Err            error `json:"-"` // low-level runtime error
	HTTPStatusCode int   `json:"-"` // http response status code

	StatusText []string `json:"message"`         // user-level status message
	ErrorText  string   `json:"error,omitempty"` // application-level error message, for debugging
}

const (
	INVALID_REQUEST = "Invalid Request"
	NOT_AUTHORIZED  = "Not authorized"
	NOT_FOUND       = "Resource not found."
	INTERNAL_ERROR  = "Internal error."
	TEST            = "Skip test."
)

func (e *ErrResponse) Render(w http.ResponseWriter, r *http.Request) error {
	render.Status(r, e.HTTPStatusCode)
	return nil
}

func ErrInvalidRequest(err error, errMessage ...string) render.Renderer {
	return &ErrResponse{
		Err:            err,
		HTTPStatusCode: 400,
		StatusText:     errMessage,
		ErrorText:      err.Error(),
	}
}

func ErrNotAuthorized(err error, errMessage ...string) render.Renderer {
	return &ErrResponse{
		Err:            err,
		HTTPStatusCode: 401,
		StatusText:     errMessage,
		ErrorText:      err.Error(),
	}
}

func ErrNotFound(err error, errMessage ...string) render.Renderer {
	return &ErrResponse{
		Err:            err,
		HTTPStatusCode: 404,
		StatusText:     errMessage,
		ErrorText:      err.Error(),
	}
}

func ErrInternal(err error, errMessage ...string) render.Renderer {
	return &ErrResponse{
		Err:            err,
		HTTPStatusCode: 500,
		StatusText:     errMessage,
		ErrorText:      err.Error(),
	}
}

func NewErrInvalidRequest(err error, errMessage ...string) render.Renderer {
	if err != nil {
		errMessage = append(errMessage, err.Error())
	}
	return &ErrResponse{
		Err:            err,
		HTTPStatusCode: 400,
		StatusText:     errMessage,
		//ErrorText:      err.Error(),
	}
}

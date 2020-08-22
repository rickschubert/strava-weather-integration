package response

import (
	"fmt"
	"net/http"
)

type Response struct {
	StatusCode int
	Body       string
}

func SuccessResponse(message string) Response {
	return Response{
		StatusCode: 200,
		Body:       message,
	}
}

func WriteResponse(writer http.ResponseWriter, response Response) {
	writer.WriteHeader(response.StatusCode)
	fmt.Fprintf(writer, response.Body)
}

func InternalServerError(message string) Response {
	return Response{
		StatusCode: 500,
		Body:       message,
	}
}

func ForwardError(err error) Response {
	return InternalServerError(err.Error())
}

func CustomResponse(message string, statusCode int) Response {
	return Response{
		StatusCode: statusCode,
		Body:       message,
	}
}

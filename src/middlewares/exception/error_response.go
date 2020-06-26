package exception

import "net/http"

// NotFound -> response empty data
func NotFound(message string, errors []map[string]interface{}) {
	response := map[string]interface{}{
		"status":  http.StatusNotFound,
		"message": message,
		"data":    nil,
		"errors":  errors,
	}
	panic(response)
}

// BadRequest -> response for bad request
func BadRequest(message string, errors []map[string]interface{}) {
	response := map[string]interface{}{
		"status":  http.StatusBadRequest,
		"message": message,
		"data":    nil,
		"errors":  errors,
	}
	panic(response)
}

func Conflict(message string, errors []map[string]interface{}) {
	response := map[string]interface{}{
		"status":  http.StatusConflict,
		"message": message,
		"data":    nil,
		"errors":  errors,
	}
	panic(response)
}

// InternalServerError -> response for internal server error
func InternalServerError(message string, errors []map[string]interface{}) {
	response := map[string]interface{}{
		"message": message,
		"data":    nil,
		"status":  http.StatusInternalServerError,
		"errors":  errors,
	}
	panic(response)
}

func Unauthorized(message string, errors []map[string]interface{}) {
	response := map[string]interface{}{
		"status":  http.StatusUnauthorized,
		"message": message,
		"data":    nil,
		"errors":  errors,
	}
	panic(response)
}

func Forbidden(message string, errors []map[string]interface{}) {
	response := map[string]interface{}{
		"status":  http.StatusForbidden,
		"message": message,
		"data":    nil,
		"errors":  errors,
	}
	panic(response)
}

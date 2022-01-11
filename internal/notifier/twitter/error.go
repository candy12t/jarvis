package twitter

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type ErrorResponse struct {
	Method     string
	URL        string
	StatusCode int
	Message    string
}

func checkResponse(resp *http.Response) error {
	errorResponse := &ErrorResponse{
		Method:     resp.Request.Method,
		URL:        resp.Request.URL.String(),
		StatusCode: resp.StatusCode,
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err == nil && body != nil {
		var parseBody struct {
			Message string `json:"message"`
		}
		json.Unmarshal(body, &parseBody)
		errorResponse.Message = parseBody.Message
	}
	defer resp.Body.Close()

	return errorResponse
}

func (err ErrorResponse) Error() string {
	return fmt.Sprintf("%v %v: %d %v", err.Method, err.URL, err.StatusCode, err.Message)
}

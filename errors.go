package stream

import (
	"encoding/json"
	"fmt"
	"time"
)

var (
	errMissingCredentials = fmt.Errorf("missing API key or secret")
)

// APIError is an error returned by Stream API when the request cannot be
// performed or errored server side.
type APIError struct {
	Code       int    `json:"code"`
	Detail     string `json:"detail"`
	Duration   time.Duration
	Exception  string `json:"exception"`
	StatusCode int    `json:"status_code"`
}

func (e APIError) Error() string {
	return e.Detail
}

// UnmarshalJSON decodes the provided JSON byte payload to the APIError.
func (e *APIError) UnmarshalJSON(b []byte) error {
	type alias APIError
	aux := &struct {
		Duration string `json:"duration"`
		*alias
	}{alias: (*alias)(e)}
	err := json.Unmarshal(b, &aux)
	if err != nil {
		return err
	}
	e.Duration, err = time.ParseDuration(aux.Duration)
	if err != nil {
		return err
	}
	return nil
}

// ToAPIError tries to cast the provided error to APIError type, returning the
// obtained APIError and whether the operation was successful.
func ToAPIError(err error) (APIError, bool) {
	se, ok := err.(APIError)
	return se, ok
}
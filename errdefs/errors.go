package errdefs

import (
	"net/http"
)

// GetHTTPErrorStatusCode retrieves status code from error message.
func GetHTTPErrorStatusCode(err error) int {
	if err == nil {
		return http.StatusInternalServerError
	}

	var statusCode int

	// Stop right there
	// Are you sure you should be adding a new error class here? Do one of the existing ones work?

	// Note that the below functions are already checking the error causal chain for matches.
	switch {
	case IsNotFound(err):
		statusCode = http.StatusNotFound
	case IsInvalidParameter(err):
		statusCode = http.StatusBadRequest
	case IsConflict(err):
		statusCode = http.StatusConflict
	case IsUnauthorized(err):
		statusCode = http.StatusUnauthorized
	case IsUnavailable(err):
		statusCode = http.StatusServiceUnavailable
	case IsForbidden(err):
		statusCode = http.StatusForbidden
	case IsNotModified(err):
		statusCode = http.StatusNotModified
	case IsNotImplemented(err):
		statusCode = http.StatusNotImplemented
	case IsSystem(err) || IsUnknown(err) || IsDataLoss(err) || IsDeadline(err) || IsCancelled(err):
		statusCode = http.StatusInternalServerError
	default:
		if e, ok := err.(causer); ok {
			return GetHTTPErrorStatusCode(e.Cause())
		}
	}

	if statusCode == 0 {
		statusCode = http.StatusInternalServerError
	}

	return statusCode
}

// FromStatusCode creates an errdef error, based on the provided HTTP status-code
func FromStatusCode(err error, statusCode int) error {
	if err == nil {
		return err
	}
	switch statusCode {
	case http.StatusNotFound:
		err = NotFound(err)
	case http.StatusBadRequest:
		err = InvalidParameter(err)
	case http.StatusConflict:
		err = Conflict(err)
	case http.StatusUnauthorized:
		err = Unauthorized(err)
	case http.StatusServiceUnavailable:
		err = Unavailable(err)
	case http.StatusForbidden:
		err = Forbidden(err)
	case http.StatusNotModified:
		err = NotModified(err)
	case http.StatusNotImplemented:
		err = NotImplemented(err)
	case http.StatusInternalServerError:
		if !IsSystem(err) && !IsUnknown(err) && !IsDataLoss(err) && !IsDeadline(err) && !IsCancelled(err) {
			err = System(err)
		}
	}
	return err
}

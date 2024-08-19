package web

import (
	"context"
	"encoding/json"

	"net/http"

	"github.com/pkg/errors"
)

func Response(ctx context.Context, w http.ResponseWriter, data interface{}, statusCode int) error {
	if statusCode == http.StatusNoContent {
		w.WriteHeader(statusCode)
	}

	jsonData, err := json.Marshal(data)
	if err != nil {
		return err
	}

	w.Header().Set("Content-Type", "application/json")
	if _, err := w.Write(jsonData); err != nil {
		return err
	}
	return nil

}

func ResponseError(ctx context.Context, w http.ResponseWriter, err error) error {
	if webErr, ok := errors.Cause(err).(*Error); ok {
		er := ErrorResponse{
			Error:  webErr.Err.Error(),
			Fields: webErr.Fields,
		}
		if err := Response(ctx, w, er, webErr.Status); err != nil {
			return err
		}
		return nil
	}
	er := ErrorResponse{
		Error: http.StatusText(http.StatusInternalServerError),
	}
	if err := Response(ctx, w, er, http.StatusInternalServerError); err != nil {
		return err
	}
	return nil
}

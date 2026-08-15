package middleware

import (
	"encoding/json"
	"fmt"
	"net/http"

	"go.uber.org/zap"
)

func StandardRespond(
	zapstruct *zap.Logger,
	response http.ResponseWriter,
	r *http.Request,
	status int,
	message any,
) {
	response.Header().Set("Content-Type", "application/json; charset=utf-8")

	beforeWriteHeader(zapstruct, response, r, status, message)
	if message != nil && message != "" {
		err := json.NewEncoder(response).Encode(message)
		if err != nil {
			zapstruct.Error("[HTTP] Failed to send response", zap.Error(err))
		}
	}
}

func beforeWriteHeader(
	zapstruct *zap.Logger, w http.ResponseWriter, r *http.Request, status int, message any,
) {
	if zapstruct == nil {
		return
	}
	var messageErr any
	if message == nil {
		messageErr = "unrecognizable error"
	} else {
		messageErr = message
	}
	path := r.URL.Path

	if r.URL.RawQuery != "" {
		path = fmt.Sprintf("%v?%v", path, r.URL.RawQuery)
	}

	if status >= 400 {
		zapstruct.Error("respond error", zap.String("method", r.Method), zap.String("path", path),
			zap.Int("code", status), zap.Reflect("message", messageErr))
	} else {
		zapstruct.Info("respond OK", zap.String("method", r.Method), zap.String("path", path),
			zap.Int("code", status))
	}
	w.WriteHeader(status)
}

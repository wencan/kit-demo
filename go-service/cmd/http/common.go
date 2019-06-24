package http

/*
 * 通用公共http处理辅助
 *
 * wencan
 * 2019-06-24
 */

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"strings"

	"github.com/go-playground/form"
)

// decodeRequest 从http请求解出请求参数结构体对象，支持查询参数、表单
func decodeRequest(_ context.Context, r *http.Request, request interface{}) error {
	switch r.Method {
	case http.MethodGet:
		err := form.NewDecoder().Decode(request, r.URL.Query())
		if err != nil {
			return err
		}
	case http.MethodPost, http.MethodPut, http.MethodPatch:
		contentType := r.Header.Get("Content-Type")
		if contentType != "" {
			contentType = strings.Split(contentType, ";")[0]
		}

		switch contentType {
		case "application/x-www-form-urlencoded", "multipart/form-data":
			err := r.ParseForm()
			if err != nil {
				return nil
			}
			err = form.NewDecoder().Decode(request, r.Form)
			if err != nil {
				return err
			}
		case "application/json":
			if r.Body == nil {
				err := errors.New("missing form body")
				return err
			}
			defer r.Body.Close()
			err := json.NewDecoder(r.Body).Decode(request)
			if err != nil {
				return err
			}
		default:
		}
	default:
	}
	return nil
}

// encodeResponse 将响应编码响应
func encodeResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	return json.NewEncoder(w).Encode(response)
}

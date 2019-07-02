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
	"github.com/wencan/errmsg"
	"gopkg.in/go-playground/validator.v9"
)

var validate *validator.Validate

func init() {
	validate = validator.New()
}

// makeRequestDecoder 构建一个请求对象的解码器（含参数检查逻辑）。需要提供一个请求对象构建函数
func makeRequestDecoder(New func() interface{}) func(context.Context, *http.Request) (interface{}, error) {
	return func(ctx context.Context, r *http.Request) (interface{}, error) {
		request := New() // 这里复用对象，可能造成污染
		err := decodeRequest(ctx, r, request)
		return request, err
	}
}

// decodeRequest 从http请求解出请求参数结构体对象，并检查参数。支持查询参数、表单、json实体
func decodeRequest(_ context.Context, r *http.Request, request interface{}) error {
	switch r.Method {
	case http.MethodGet:
		err := form.NewDecoder().Decode(request, r.URL.Query())
		if err != nil {
			err = errmsg.WrapError(errmsg.ErrInvalidArgument, err)
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
				return err
			}
			err = form.NewDecoder().Decode(request, r.Form)
			if err != nil {
				err = errmsg.WrapError(errmsg.ErrInvalidArgument, err)
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
				err = errmsg.WrapError(errmsg.ErrInvalidArgument, err)
				return err
			}
		default:
		}
	default:
	}

	// 参数检查
	err := validate.Struct(request)
	if err != nil {
		err = errmsg.WrapError(errmsg.ErrInvalidArgument, err)
		return err
	}
	return nil
}

// encodeResponse 将响应编码作http响应实体
func encodeResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	return json.NewEncoder(w).Encode(response)
}

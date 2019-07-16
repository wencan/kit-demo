package fasthttp

/*
 * 通用公共http处理辅助
 *
 * wencan
 * 2019-06-24
 */

import (
	"context"
	"net/http"
	"net/url"
	"strings"

	"github.com/go-playground/form"
	"github.com/valyala/fasthttp"
	"github.com/wencan/errmsg"
	fasthttp_transport "github.com/wencan/kit-plugins/transport/fasthttp"
	"gopkg.in/go-playground/validator.v9"
)

var validate *validator.Validate

func init() {
	validate = validator.New()
}

// 将fasthttp.Args转为url.Values
func valuesFromArgs(args *fasthttp.Args) url.Values {
	values := make(url.Values, args.Len())
	args.VisitAll(func(k, v []byte) {
		values.Add(string(k), string(v))
	})
	return values
}

// decodeRequest 从http请求解出请求参数结构体对象，并检查参数。支持查询参数、表单、json实体
func decodeRequest(ctx context.Context, r *fasthttp.Request, request interface{}) error {
	switch string(r.Header.Method()) {
	case http.MethodGet:
		values := valuesFromArgs(r.URI().QueryArgs())
		err := form.NewDecoder().Decode(request, values)
		if err != nil {
			err = errmsg.WrapError(errmsg.ErrInvalidArgument, err)
			return err
		}
	case http.MethodPost, http.MethodPut, http.MethodPatch:
		contentType := string(r.Header.ContentType())
		if contentType != "" {
			contentType = strings.Split(contentType, ";")[0]
		}

		switch contentType {
		case "application/x-www-form-urlencoded":
			values := valuesFromArgs(r.PostArgs())
			err := form.NewDecoder().Decode(request, values)
			if err != nil {
				err = errmsg.WrapError(errmsg.ErrInvalidArgument, err)
				return err
			}
		case "application/json":
			err := fasthttp_transport.DecodeJSONRequest(ctx, r, request)
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

package http

/*
 * 通用公共http辅助处理
 * 包含编码请求和解码响应的支持
 *
 * wencan
 * 2019-07-08
 */

import (
	"context"
	"encoding/json"
	"errors"
	"io"
	"log"
	"net/http"
	"net/url"

	transport "github.com/go-kit/kit/transport/http"
	"github.com/go-playground/form"
	"github.com/wencan/errmsg"
)

var (
	Encoder = form.NewEncoder()
)

// AddrURL 用于支持urljoin
type AddrURL struct {
	headURL *url.URL
}

func newAddrURL(addr string) (*AddrURL, error) {
	headURL, err := url.Parse(addr)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return &AddrURL{
		headURL: headURL,
	}, nil
}

func (addrURL *AddrURL) Join(paths ...string) (*url.URL, error) {
	headURL := &(*addrURL.headURL) // copy一份，以免污染
	for _, path := range paths {
		pathURL, err := url.Parse(path)
		if err != nil {
			log.Println(err)
			return nil, err
		}

		headURL = headURL.ResolveReference(pathURL)
	}

	return headURL, nil
}

func encodeQueryRequest(ctx context.Context, r *http.Request, req interface{}) error {
	// 构建查询请求，只适用于GET请求
	if r.Method != http.MethodGet {
		err := errors.New("request method error")
		log.Println(err)
		return err
	}

	// 将请求结构体编码为map[string][][]string
	forms, err := Encoder.Encode(req)
	if err != nil {
		log.Println(err)
		return err
	}

	// 解析原路径中的查询参数
	query, err := url.ParseQuery(r.URL.RawQuery)
	if err != nil {
		log.Println(err)
		return err
	}
	// 将请求结构体中字段添加到查询
	for key, values := range forms {
		for _, value := range values {
			query.Add(key, value)
		}
	}

	r.URL.RawQuery = query.Encode()
	return nil
}

func decodeErrMsg(r io.Reader) (*errmsg.ErrMsg, error) {
	errMsg := &errmsg.ErrMsg{}

	err := json.NewDecoder(r).Decode(errMsg)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return errMsg, nil
}

// makeResponseDecoder 创建HTTP响应解码器，需要传递一个响应对象工厂函数
// 如果状态码非2XX，解码为errmsg.ErrMsg
// 目前全当作json解码
func makeResponseDecoder(New func() interface{}) transport.DecodeResponseFunc {
	return func(ctx context.Context, resp *http.Response) (interface{}, error) {
		if resp.Body == nil {
			return nil, nil
		}
		defer resp.Body.Close()

		if resp.StatusCode/200 == 2 {
			// 解析ErrMsg
			errMsg, err := decodeErrMsg(resp.Body)
			if err != nil {
				return nil, err
			}
			return nil, errMsg
		}

		response := New()
		err := json.NewDecoder(resp.Body).Decode(response)
		if err != nil {
			log.Println(err)
			return nil, err
		}

		return response, nil
	}
}

package grpc

/*
 * 通用公共grpc处理辅助
 * 提供grpc请求/响应和公共请求/响应之间的编解码器
 *
 *
 * wencan
 * 2019-07-01
 */

import (
	"context"

	"github.com/wencan/copier"
	"github.com/wencan/errmsg"
	"gopkg.in/go-playground/validator.v9"
)

var (
	Decoder = copier.NewCopier("protobuf.name", "form")

	Encoder = copier.NewCopier("resp", "protobuf.name")

	validate = validator.New()
)

// makeRequestDecoder 构建一个将grpc请求转换为公共请求的解码器，（含参数检查逻辑）
func makeRequestDecoder(New func() interface{}) func(context.Context, interface{}) (interface{}, error) {
	req := New()
	return func(_ context.Context, grpcReq interface{}) (interface{}, error) {
		err := Decoder.Copy(req, grpcReq)
		if err != nil {
			err = errmsg.WrapError(errmsg.ErrInvalidArgument, err)
			return nil, err
		}

		// 参数检查
		err = validate.Struct(req)
		if err != nil {
			err = errmsg.WrapError(errmsg.ErrInvalidArgument, err)
			return nil, err
		}
		return req, nil
	}
}

// makeResponseEncoder 构建一个将公共响应编码为grpc响应的编码器
func makeResponseEncoder(New func() interface{}) func(context.Context, interface{}) (interface{}, error) {
	grpcResp := New()
	return func(_ context.Context, resp interface{}) (interface{}, error) {
		err := Encoder.Copy(grpcResp, resp)
		if err != nil {
			// unknown error ...
			return nil, err
		}
		return grpcResp, err
	}
}

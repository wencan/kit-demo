package grpc

/*
 * 通用公共grpc处理辅助
 * 提供grpc请求/响应和公共请求/响应之间的编解码器
 *
 * 与服务端的部分相比，就是编码变解码，解码变编码
 *
 *
 * wencan
 * 2019-07-01
 */

import (
	"context"

	"github.com/wencan/copier"
	"github.com/wencan/errmsg"
)

var (
	Encoder = copier.NewCopier("protobuf.name", "form")

	Decoder = copier.NewCopier("resp", "protobuf.name")
)

// makeRequestEncoder 构建一个将grpc请求转换为公共请求的编码器
// 参数检查留给服务端
func makeRequestEncoder(New func() interface{}) func(context.Context, interface{}) (interface{}, error) {
	grpcReq := New()
	return func(_ context.Context, req interface{}) (interface{}, error) {
		err := Encoder.Copy(grpcReq, req)
		if err != nil {
			err = errmsg.WrapError(errmsg.ErrInvalidArgument, err)
			return nil, err
		}
		return grpcReq, nil
	}
}

// makeResponseDecoder 构建一个将公共响应编码为grpc响应的解码器
func makeResponseDecoder(New func() interface{}) func(context.Context, interface{}) (interface{}, error) {
	resp := New()
	return func(_ context.Context, grpcResp interface{}) (interface{}, error) {
		err := Decoder.Copy(resp, grpcResp)
		if err != nil {
			// unknown error ...
			return nil, err
		}
		return resp, err
	}
}

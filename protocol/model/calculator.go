package model

/*
 * 计算器请求/响应数据模型
 * 请求模型中form标签用于支持github.com/go-playground/form解码
 * 请求模型中validate标签用于支持github.com/go-playground/validator做请求参数检查
 * 请求模型中reply标签用于支持github.com/wencan/copier深拷贝
 *
 * wencan
 * 2019-06-24
 */

type CalculatorAddRequest struct {
	A int32 `form:"a" validate:"required"`
	B int32 `form:"b" validate:"required"`
}

type CalculatorSubRequest struct {
	C int32 `form:"c" validate:"required"`
	D int32 `form:"d" validate:"required"`
}

type CalculatorMulRequest struct {
	E int32 `form:"e" validate:"required"`
	F int32 `form:"f" validate:"required"`
}

type CalculatorDivRequest struct {
	M int32 `form:"m" validate:"required"`
	N int32 `form:"n" validate:"required,ne=0"`
}

type CalculatorInt32Response struct {
	Result int32 `reply:"result" json:"result"`
}

type CalculatorFloatResponse struct {
	Result float32 `reply:"result" json:"result"`
}

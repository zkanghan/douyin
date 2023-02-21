package errno

import (
	"douyin/kitex_gen/base"
	"errors"
)

// BuildBaseResp build baseResp from error
func BuildBaseResp(err error) *base.BaseResp {
	if err == nil {
		return baseResp(Success)
	}

	e := ErrNo{}
	if errors.As(err, &e) {
		return baseResp(e)
	}

	s := ServiceErr.WithMessage(err.Error())
	return baseResp(s)
}

func baseResp(err ErrNo) *base.BaseResp {
	return &base.BaseResp{StatusCode: int32(err.ErrCode), StatusMsg: &err.ErrMsg}
}

package server

import (
	"net/http"

	khttp "github.com/go-kratos/kratos/v2/transport/http"
	kstatus "github.com/go-kratos/kratos/v2/transport/http/status"
	"google.golang.org/grpc/status"
)

// HTTP Encoder
// 自定义HTTP响应编码器,按照自定义的响应格式

type httpResponse struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

// 自定义编码方法
func ResponseEncoder(w http.ResponseWriter, r *http.Request, v interface{}) error {
	if v == nil {
		return nil
	}
	// 判断是不是重定向
	if rd, ok := v.(khttp.Redirector); ok {
		url, code := rd.Redirect()
		http.Redirect(w, r, url, code)
		return nil
	}

	// 构造自定义的相应结构体
	resp := &httpResponse{
		Code: http.StatusOK,
		Msg:  "Success",
		Data: v,
	}

	// Codec
	codec, _ := khttp.CodecForRequest(r, "Accept")
	data, err := codec.Marshal(resp)
	if err != nil {
		return err
	}
	w.Header().Set("Content-Type", "application/"+(codec.Name()))
	_, err = w.Write(data)
	if err != nil {
		return err
	}
	return nil
}

// 自定义错误编码
func ErrorEncoder(w http.ResponseWriter, r *http.Request, err error) {
	// 0.if err == nil ,return
	if err == nil {
		return
	}
	// 1.if can resolve info from err
	resp := new(httpResponse)
	if gs, ok := status.FromError(err); ok {
		resp = &httpResponse{
			Code: kstatus.FromGRPCCode(gs.Code()),
			Msg:  gs.Message(),
			Data: nil,
		}
	} else {
		// 2.can not resolve info from err
		resp = &httpResponse{
			Code: http.StatusInternalServerError,
			Msg:  "内部错误",
			Data: err,
		}
	}

	codec, _ := khttp.CodecForRequest(r, "Accpet")
	body, err := codec.Marshal(resp)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/"+codec.Name())
	w.WriteHeader(resp.Code)
	_, _ = w.Write(body)
}

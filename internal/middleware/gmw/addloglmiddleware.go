package gmw

import (
	"bytes"
	"github.com/wordpress-plus/app-api/internal/svc"
	"github.com/zeromicro/go-zero/core/logx"
	"io"
	"net/http"
	"time"
)

type AddLogMiddleware struct {
	svcCtx *svc.ServiceContext
}

func NewAddLogMiddleware(svcCtx *svc.ServiceContext) *AddLogMiddleware {
	return &AddLogMiddleware{svcCtx: svcCtx}
}

// Deprecated: Handle move to operation mw
func (m *AddLogMiddleware) Handle(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		uri := r.RequestURI
		startTime := time.Now()

		// 读取请求主体
		body, err := io.ReadAll(r.Body)
		if err != nil {
			logx.WithContext(r.Context()).Errorf("Failed to read request body: %v", err)
		}

		// 创建一个新的请求主体用于后续读取
		r.Body = io.NopCloser(bytes.NewBuffer(body))

		// 打印请求参数日志
		logx.WithContext(r.Context()).Infof("Request: %s %s %s", r.Method, uri, body)

		// 创建一个自定义的 ResponseWriter，用于记录响应
		recorder := &responseRecorder{
			ResponseWriter: w,
			statusCode:     http.StatusOK,
			body:           make([]byte, 0),
		}

		// 调用下一个处理器，捕获响应
		next(recorder, r)

		// 打印响应日志
		logx.WithContext(r.Context()).Infof("Response: %s %s %s duration: %d", r.Method, r.RequestURI, string(recorder.body), time.Since(startTime))
	}
}

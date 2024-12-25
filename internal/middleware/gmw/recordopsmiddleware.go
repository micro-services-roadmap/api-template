package gmw

import (
	"bytes"
	"github.com/micro-services-roadmap/kit-common/ipx"
	"github.com/wordpress-plus/app-api/internal/svc"
	//"github.com/wordpress-plus/rpc-tracing/pbtms"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/rest/httpx"
	"io"
	"net/http"
	"strings"
	"time"
)

type RecordOpsMiddleware struct {
	svcCtx *svc.ServiceContext
}

func NewRecordOpsMiddleware(svcCtx *svc.ServiceContext) *RecordOpsMiddleware {
	return &RecordOpsMiddleware{svcCtx: svcCtx}
}

// Handle record url and uid and ip with rpc-trace
func (m *RecordOpsMiddleware) Handle(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		startTime := time.Now()

		uri := r.RequestURI
		reqBody, err := io.ReadAll(r.Body)
		if err != nil {
			logx.WithContext(r.Context()).Errorf("Failed to read request reqBody: %v", err)
		}
		r.Body = io.NopCloser(bytes.NewBuffer(reqBody))
		logx.WithContext(r.Context()).Infof("Request: %s %s %s", r.Method, uri, reqBody)

		// 创建一个自定义的 ResponseWriter，用于记录响应
		recorder := &responseRecorder{
			ResponseWriter: w,
			statusCode:     http.StatusOK,
			body:           make([]byte, 0),
		}
		next(recorder, r)
		// 打印响应日志
		logx.WithContext(r.Context()).Infof("Response: %s %s %s duration: %d", r.Method, r.RequestURI, string(recorder.body), time.Since(startTime))

		//go func(tq *pbtms.ViewAddReq, startTime time.Time) {
		//	tq.Cost = time.Since(startTime).Milliseconds()
		//	if _, err := m.svcCtx.TracingViewService.ViewAdd(context.Background(), tq); err != nil {
		//		logx.Errorf("Failed to TracingViewService-ViewAdd: %v", err)
		//	}
		//}(buildTraceAddReq(r, trace.TraceIDFromContext(r.Context()), trace.SpanIDFromContext(r.Context()), string(reqBody), string(recorder.body)), startTime)
	}
}

//func buildTraceAddReq(r *http.Request, traceID, spanID, req, resp string) *pbtms.ViewAddReq {
//	p := &pbtms.ViewAddReq{
//		Utype:     0,
//		IP:        GetRemoteAddr(r),
//		Url:       r.RequestURI,
//		UserAgent: r.Header.Get("User-Agent"),
//		TraceID:   traceID,
//		SpanID:    spanID,
//		Method:    r.Method,
//		Response:  resp,
//		Request:   req,
//	}
//	if uid, ok := r.Context().Value(modelo.ID).(int64); ok {
//		p.Uid = uid
//	}
//	if uname, uok := r.Context().Value(modelo.Name).(string); uok {
//		p.Uname = uname
//	}
//	return p
//}

const xRealIP = "X-Real-IP"

// GetRemoteAddr returns the peer address, supports X-Forward-For.
func GetRemoteAddr(r *http.Request) string {
	logx.Info("x-real-ip: ", r.Header.Get(xRealIP))
	ip := httpx.GetRemoteAddr(r)
	logx.Info("x-forward-for: ", ip)
	logx.Info("RemoteAddr: ", r.RemoteAddr)

	return extractFirstPart(ip)
}

func extractFirstPart(v string) string {
	idx := strings.Index(v, ",")
	if idx != -1 {
		v = v[0:idx]
	}

	if ipx.IsValidIP(v) {
		return v
	}

	portIndex := strings.LastIndex(v, ":") // ipv6(240e:46c:8910:219e:41a2:1185:be37:5f61)
	if portIndex == -1 {
		return v
	}
	return v[0:portIndex]
}

// 自定义的 ResponseWriter
type responseRecorder struct {
	http.ResponseWriter
	statusCode int
	body       []byte
}

// WriteHeader 重写 WriteHeader 方法，捕获状态码
func (r *responseRecorder) WriteHeader(statusCode int) {
	r.statusCode = statusCode
	r.ResponseWriter.WriteHeader(statusCode)
}

// 重写 Write 方法，捕获响应数据
func (r *responseRecorder) Write(body []byte) (int, error) {
	r.body = body
	return r.ResponseWriter.Write(body)
}

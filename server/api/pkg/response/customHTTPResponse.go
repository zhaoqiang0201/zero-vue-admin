package response

import (
	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zhaoqiang0201/zero-vue-admin/server/api/pkg/errorx"
	"google.golang.org/grpc/status"
)

func ErrorHandle(err error) (int, interface{}) {
	defer func() {
		logx.Errorf("[API]-->%+v", err)
	}()

	causeErr := errors.Cause(err)
	if e, ok := causeErr.(*errorx.Errorx); ok {
		return 550, *e
	} else {
		if gstatus, ok := status.FromError(causeErr); ok { // grpc err错误
			grpcCode := errorx.ErrorCode(gstatus.Code())
			if errorx.IsErrorxCode(grpcCode) {
				//errcode = grpcCode
				//errmsg = gstatus.Message()
				return 551, errorx.Errort{551, "t1"}
			}
		}
		return 552, errorx.Errort{552, "t2"}
	}
}

//func CustomHttpResponse(w http.ResponseWriter, r *http.Request, data interface{}, err error) {
//	if err != nil {
//		errcode := errorx.SERVER_COMMON_ERROR
//		errmsg := err.Error()
//		causeErr := errors.Cause(err)
//		if e, ok := causeErr.(*errorx.Errorx); ok {
//			errcode = errorx.ErrorCode(e.ErrorxCode())
//			errmsg = e.ErrorxMsg()
//		} else {
//			if gstatus, ok := status.FromError(causeErr); ok { // grpc err错误
//				grpcCode := errorx.ErrorCode(gstatus.Code())
//				if errorx.IsErrorxCode(grpcCode) {
//					errcode = grpcCode
//					errmsg = gstatus.Message()
//				}
//			}
//		}
//		logx.WithContext(r.Context()).Errorf("[API]-->%+v", err)
//		httpx.WriteJson(w, http.StatusOK, BeanFail(uint32(errcode), errmsg))
//	} else {
//		responseBean := BeanOKWithData(data)
//		httpx.WriteJson(w, http.StatusOK, responseBean)
//	}
//}
//
//func ParaErrorResult(w http.ResponseWriter, r *http.Request, err error) {
//	code := uint32(errorx.REUQEST_PARAM_ERROR)
//	httpx.WriteJson(w, http.StatusBadRequest, BeanFail(code, errorx.ErrorxMessage(errorx.REUQEST_PARAM_ERROR)))
//}

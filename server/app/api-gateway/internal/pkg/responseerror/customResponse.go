package responseerror

import (
	"database/sql"
	"fmt"
	"github.com/pkg/errors"
	"github.com/zhaoqiang0201/zero-vue-admin/server/app/api-gateway/internal/pkg/responseerror/errorx"
	"net/http"

	"google.golang.org/grpc/status"
)

type ErrorResponse struct {
	Code     int32       `json:"code"`
	Msg      string      `json:"msg"`
	Meta     interface{} `json:"meta"`
	CauseErr string      `json:"cause_err"`
}

/*
code: 1-17   ==>   rpc错误
code: 9999   ==>   其他错误
*/
func ErrorHandle(err error) (int, interface{}) {
	causeErr := errors.Cause(err)
	switch e := causeErr.(type) {
	case *errorx.Errorx: //自定义错误
		fmt.Printf("==> %T %#v\n", e, e)
		if e.Code == errorx.UNAUTHORIZATION {
			return http.StatusUnauthorized, ErrorResponse{Code: int32(e.Code), Msg: e.Msg, CauseErr: e.Cause.Error(), Meta: e.Meta}
		}
		return 450, ErrorResponse{Code: int32(e.Code), Msg: e.Msg, CauseErr: e.Cause.Error(), Meta: e.Meta}
	default:
		fmt.Printf("==> %T\n", e)
		//fmt.Printf("=====>  default  %T %v\n", err, err)
		//=====>  default  *status.Error      rpc error: code = Unavailable desc = last resolver error: produced zero addresses
		//s, o := status.FromError(e)
		//fmt.Println(o)           // true -> rpc错误， false 其他错误
		//fmt.Println(s.Err())     // rpc error: code = Unavailable desc = last resolver error: produced zero addresses
		//fmt.Println(s.Message()) //last resolver error: produced zero addresses
		//fmt.Println(s.Code())    //Unavailable
		//fmt.Println(s.Proto())   //code:14  message:"last resolver error: produced zero addresses"
		//fmt.Println(s.Details()) //[]
		//fmt.Println(s.String())  //rpc error: code = Unavailable desc = last resolver error: produced zero addresses
		e2, isRpcErr := status.FromError(e)
		if isRpcErr { // rpc错误
			if e2.Message() == sql.ErrNoRows.Error() { //rpc查询为空
				return 450, ErrorResponse{Code: e2.Proto().Code, Msg: fmt.Sprintf("grpc错误， %s", e2.Message()), CauseErr: e2.String()}
			}
			return 450, ErrorResponse{Code: e2.Proto().Code, Msg: fmt.Sprintf("grpc错误， %s", e2.Message()), CauseErr: e2.String()}
		}
		return 450, ErrorResponse{Code: 9999, Msg: fmt.Sprintf("XX: default: %s", e.Error()), CauseErr: e.Error()}
	}
}

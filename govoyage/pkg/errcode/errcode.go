package errcode

import (
	"log/slog"

	"github.com/kelein/cookbook/govoyage/pbgen"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var internalErr = make(map[int]string)

// Preset Error Codes
var (
	Success          = NewError(0, "成功")
	Fail             = NewError(10000000, "内部错误")
	InvalidParams    = NewError(10000001, "无效参数")
	Unauthorized     = NewError(10000002, "认证错误")
	NotFound         = NewError(10000003, "没有找到")
	Unknown          = NewError(10000004, "未知")
	DeadlineExceeded = NewError(10000005, "超出最后截止期限")
	AccessDenied     = NewError(10000006, "访问被拒绝")
	LimitExceed      = NewError(10000007, "访问限制")
	MethodNotAllowed = NewError(10000008, "不支持该方法")
)

// Error stands for RPC error
type Error struct {
	code int
	msg  string
}

// NewError creates a new PRC Error
func NewError(code int, msg string) *Error {
	if _, ok := internalErr[code]; ok {
		slog.Warn("error code already exits", "code", code)
	}
	return &Error{code: code, msg: msg}
}

// Code return the error code
func (e *Error) Code() int { return e.code }

// Msg return the error message
func (e *Error) Msg() string { return e.msg }

// RPCError converts common error to RPC error
func RPCError(err *Error) error {
	pbErr := &pbgen.CommonError{Code: int32(err.Code()), Message: err.Msg()}
	s, serr := status.New(RPCCode(err.Code()), err.Msg()).WithDetails(pbErr)
	if serr != nil {
		slog.Error("status.WithDetails failed", "error", serr)
		return serr
	}
	return s.Err()
}

// RPCCode converts common error to RPC code
func RPCCode(code int) codes.Code {
	var statusCode codes.Code
	switch code {
	case Success.Code():
		statusCode = codes.OK
	case Fail.Code():
		statusCode = codes.Internal
	case InvalidParams.Code():
		statusCode = codes.InvalidArgument
	case Unauthorized.Code():
		statusCode = codes.Unauthenticated
	case AccessDenied.Code():
		statusCode = codes.PermissionDenied
	case NotFound.Code():
		statusCode = codes.NotFound
	case DeadlineExceeded.Code():
		statusCode = codes.DeadlineExceeded
	case LimitExceed.Code():
		statusCode = codes.ResourceExhausted
	case MethodNotAllowed.Code():
		statusCode = codes.Unimplemented
	default:
		statusCode = codes.Unknown
	}
	return statusCode
}

// CommonStatus stands for GRPC status
type CommonStatus struct {
	*status.Status
}

// FromError convert common error to GRPC status
func FromError(err error) *CommonStatus {
	s, _ := status.FromError(err)
	return &CommonStatus{s}
}

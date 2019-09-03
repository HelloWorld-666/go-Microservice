package common

import (
	"encoding/json"
	"net/http"
)

// 规则：GC-{1位系统序号}{1位业务首字母大写}{3位错误序号} 3位错误序号首位数字代表系统下面不同的子业务

var (
	///////////////////////////////////////////////////////////////////////////////////////
	// 全局错误码。错误码前缀为GC-1G
	///////////////////////////////////////////////////////////////////////////////////////
	ErrUnkownError         = NewError("GC-1G999", "未知错误", http.StatusInternalServerError)
	ErrInternalServerError = NewError("GC-1G998", "系统内部错误", http.StatusInternalServerError)
	ErrParamterError       = NewError("GC-1G997", "参数错误", http.StatusBadRequest)
	ErrPageSizeTooLarge    = NewError("GC-1G996", "每页查询数据过大", http.StatusBadRequest)
	ErrRecordNotExist      = NewError("GC-1G101", "此条数据不存在", http.StatusBadRequest) // 此条错误码太过笼统，需要逐渐弃用

	///////////////////////////////////////////////////////////////////////////////////////
	// 基础微服务相关错误码。用于goodacdining-base-server，错误码前缀为GC-1B
	///////////////////////////////////////////////////////////////////////////////////////
	// 登录授权
	ErrUsernameOrPasswordError = NewError("GC-1B001", "用户名或密码错误", http.StatusBadRequest)
	ErrUserDisabled            = NewError("GC-1B002", "用户已被禁用", http.StatusForbidden)
	ErrUnauthorized            = NewError("GC-1B003", "用户未登录", http.StatusUnauthorized)
	ErrMchtUnselected          = NewError("GC-1B004", "未选择商户", http.StatusForbidden)
	ErrStoreUnselected         = NewError("GC-1B005", "未选择门店", http.StatusForbidden)
	ErrMchtNoEmpty             = NewError("GC-1B006", "商户号为空", http.StatusForbidden)
	ErrMchtVerifyFailed        = NewError("GC-1B007", "验证商户信息失败", http.StatusForbidden)
	ErrSnKeyWrong              = NewError("GC-1B010", "sn密钥错误", http.StatusBadRequest)
	// 点餐码
	ErrDeskSnNotExist = NewError("GC-1B015", "该点餐码不存在", http.StatusForbidden)
	ErrDeskSnNotBind  = NewError("GC-1B016", "该点餐码未绑定", http.StatusForbidden)
	ErrBindOtherStore = NewError("GC-1B017", "已绑定其它的门店", http.StatusForbidden)
	// 菜品
	ErrDishIdNotExist       = NewError("GC-1B021", "菜品Id不存在", http.StatusNotAcceptable)
	ErrDishSpecIdNotExist   = NewError("GC-1B022", "菜品规格Id不存在", http.StatusNotAcceptable)
	ErrDishStockNotEnough   = NewError("GC-1B023", "菜品数量不够", http.StatusNotAcceptable)
	ErrDishTasteIdNotExist  = NewError("GC-1B024", "菜品口味Id不存在", http.StatusNotAcceptable)
	ErrDishNameAlreadyExist = NewError("GC-1B024", "菜品名已存在", http.StatusNotAcceptable)
	// 菜品分类
	ErrDishCategoryNameAlreadyExist = NewError("GC-1B031", "菜品分类名已存在", http.StatusNotAcceptable)
	// 桌台
	ErrDeskNameAlreadyExist = NewError("GC-1B041", "桌台名已存在", http.StatusNotAcceptable)
	// 桌台区域
	ErrDeskAreaNameAlreadyExist = NewError("GC-1B051", "桌台区域名已存在", http.StatusNotAcceptable)
	// 桌牌
	ErrDeskCardNameAlreadyExist = NewError("GC-1B061", "桌牌名已存在", http.StatusNotAcceptable)
	// 配置
	ErrConfigNameAlreadyExist = NewError("GC-1B071", "配置名已存在", http.StatusNotAcceptable)
	// 支付方式
	ErrPayMethodCodeNotExist = NewError("GC-1B081", "支付方式编码不存在", http.StatusNotAcceptable)
	// 授权
	ErrAuthorizationAlreadyExist = NewError("GC-1B091", "授权已存在", http.StatusNotAcceptable)
	ErrAuthorizationNotExist     = NewError("GC-1B092", "授权不存在", http.StatusNotAcceptable)

	///////////////////////////////////////////////////////////////////////////////////////
	// 订单微服务相关错误码，用于goodacdining-order-server，错误码前缀为GC-1O
	///////////////////////////////////////////////////////////////////////////////////////
	ErrTransAmountNotMatch = NewError("GC-1O101", "交易金额不匹配", http.StatusNotAcceptable)
	ErrOrderNotExist       = NewError("GC-1O103", "未找到对应的订单", http.StatusNotAcceptable)
	ErrOrderCheckouted     = NewError("GC-1O105", "该订单已结账", http.StatusNotAcceptable)
	ErrOrderCanceled       = NewError("GC-1O106", "该订单已取消", http.StatusNotAcceptable)
	ErrOutTradeNoNotExist  = NewError("GC-1O111", "第三方支付单号不存在", http.StatusNotAcceptable)
)

// Error provide a way to return detailed information
// for an gRPC request error. The error is normally JSON encoded.
type Error struct {
	ErrorCode  string `json:"errorCode"`
	ErrorDesc  string `json:"errorDesc"`
	StatusCode int    `json:"statusCode"`
}

func (e *Error) Error() string {
	b, _ := json.Marshal(e)
	return string(b)
}

func NewError(errorCode, errorDesc string, statusCode int) *Error {
	return &Error{
		ErrorCode:  errorCode,
		ErrorDesc:  errorDesc,
		StatusCode: statusCode,
	}
}

func ParseError(errString string) *Error {
	e := new(Error)
	err := json.Unmarshal([]byte(errString), e)
	if err != nil {
		return ErrUnkownError
	}
	return e
}

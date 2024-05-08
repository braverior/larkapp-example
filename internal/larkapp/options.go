package larkapp

type LarkAppConfig struct {
	// 应用APP ID
	AppId string
	// 应用APP Secret
	AppSecret string
	// API请求超时时间，单位：秒
	ReqTimeout int
	// 是否开启Token缓存
	EnableTokenCache bool
	// 日志级别
	// LogLevelDebug LogLevel = 1 ;
	// LogLevelInfo  LogLevel = 2 ;
	// LogLevelWarn  LogLevel = 3 ;
	// LogLevelError LogLevel = 4 ;;
	LogLevel int
	// 用于加密事件或回调的请求内容，校验请求来源
	VerificationToken string
	// 用于加密事件或回调的请求内容，用于解密请求内容
	EncryptKey string
	// 消息卡片事件，如果为0，则不开启卡片回传
	CardEventPort int
	// 卡片事件回调路径
	CardEventPath string
}

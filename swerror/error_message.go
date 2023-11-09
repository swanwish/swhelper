package swerror

const (
	// Database operation error messages
	ErrorMessageGetConnectionFailed   = "Failed to get database connection"
	ErrorMessageNoConnectionProvider  = "Connection provider not specified"
	ErrorMessageNoTransactionFunction = "Transaction function not specified"

	ErrorMessageNotExist            = "Not exist"
	ErrorMessageAlreadyExist        = "Already exist"
	ErrorMessageInvalidParameter    = "Invalid Parameter"
	ErrorMessageInvalidToken        = "Invalid token"
	ErrorMessageNoPrivileges        = "No privileges"
	ErrorMessageRefreshTokenExpired = "Refresh token expired"
	ErrorMessageTokenExpired        = "Token expired"
	ErrorMessageNoDeviceInformation = "No deviceinformation"
	ErrorMessageLimitExceed         = "Limit exceed"

	ErrorMessageStatusNotPublished = "The status is not published"

	ErrorMessageNotSupport    = "Not support"
	ErrorMessageBadResponse   = "Bad response"
	ErrorMessageForbidden     = "Forbidden"
	ErrorMessageInternalError = "Internal error"
	ErrorMessageInvalidStatus = "Invalid status"

	ErrorMessageExecuteSqlFailed     = "Failed to execute sql %s, the error is %v"
	ErrorMessageNotImplemented       = "Not implemented"
	ErrorMessageUrlExpired           = "Url expired"
	ErrorMessageConfigurationMissing = "Configuration missing."

	ErrorMessageOssClientNotInited = "Oss client not inited"

	ErrorMessageDataExistFailed = "Data already exist"
)

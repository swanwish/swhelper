package wserror

import "errors"

var (
	// Database Related Error
	ErrNoConnectionProvider  = errors.New(ErrorMessageNoConnectionProvider)
	ErrNoTransactionFunction = errors.New(ErrorMessageNoTransactionFunction)

	ErrNotExist            = errors.New(ErrorMessageNotExist)
	ErrAlreadyExist        = errors.New(ErrorMessageAlreadyExist)
	ErrInvalidParameter    = errors.New(ErrorMessageInvalidParameter)
	ErrInvalidToken        = errors.New(ErrorMessageInvalidToken)
	ErrNoPrivileges        = errors.New(ErrorMessageNoPrivileges)
	ErrRefreshTokenExpired = errors.New(ErrorMessageRefreshTokenExpired)
	ErrTokenExpired        = errors.New(ErrorMessageTokenExpired)
	ErrNoDeviceInformation = errors.New(ErrorMessageNoDeviceInformation)
	ErrLimitExceed         = errors.New(ErrorMessageLimitExceed)

	ErrStatusNotPublished = errors.New(ErrorMessageStatusNotPublished)
	ErrNotSupport         = errors.New(ErrorMessageNotSupport)
	ErrBadResponse        = errors.New(ErrorMessageBadResponse)
	ErrInternalError      = errors.New(ErrorMessageInternalError)
	ErrInvalidStatus      = errors.New(ErrorMessageInvalidStatus)

	ErrOssClientNotInited = errors.New(ErrorMessageOssClientNotInited)

	ErrConfigurationMissing = errors.New("Configuration missing.")
)

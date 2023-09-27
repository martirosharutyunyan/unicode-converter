package static_errors

import "errors"

var (
	NOT_SUPPORTED_FORMAT_ERR = errors.New("Not supported file type")
	TRANSFER_SOURCE_ERR      = errors.New("Transfer source")
)

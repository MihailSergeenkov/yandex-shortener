package common

import "errors"

type ContextValueKey string

const KeyUserID ContextValueKey = "userID"

var ErrFetchUserIDFromContext = errors.New("failed to fetch user id from context")
var EncRespErrStr = "error encoding response"
var ContentTypeHeader = "Content-Type"
var JSONContentType = "application/json"

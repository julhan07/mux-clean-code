package entity

const DATA_CREATED = "DATA_CREATED"
const DATA_UPDATED = "DATA_UPDATED"
const DATA_DELETED = "DATA_DELETED"
const DATA_RECEIVED = "DATA_RECEIVED"
const ERROR_RECEIVED = "ERROR_RECEIVED"

type ResponseApi struct {
	Error   error       `json:"error,omitempty"`
	Data    interface{} `json:"data,omitempty"`
	Message string      `json:"msg,omitempty"`
}

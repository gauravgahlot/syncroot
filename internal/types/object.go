package types

// Object represents an object that can be enqueued for processing.
type Object interface {
	GetID() string
	GetType() string
}

// register all types to be synced
var (
	_ = Object(&Contact{})
)

type Operation string

const (
	OperationCreate Operation = "create"
	OperationUpdate Operation = "update"
	OperationDelete Operation = "delete"
)

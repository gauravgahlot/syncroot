package types

type WorkerType string

const (
	WorkerTypeForwarder      WorkerType = "forwarder"
	WorkerTypeSyncer         WorkerType = "syncer"
	WorkerTypeDLQHandler     WorkerType = "dlq-handler"
	WorkerTypeWebhookHandler WorkerType = "webhook-handler"
)

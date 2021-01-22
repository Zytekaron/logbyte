package types

type Level string

const (
	LevelTrace Level = "trace"
	LevelDebug       = "debug"
	LevelInfo        = "info"
	LevelWarn        = "warn"
	LevelError       = "error"
	LevelFatal       = "fatal"
)

type LogEntry struct {
	// The log entry id
	Id string `json:"id" bson:"_id"`

	// The service from which this log entry originated
	Service string `json:"service" bson:"service"`

	// The severity of the log entry
	Level Level `json:"level" bson:"level"`

	// The message of the log entry
	Message string `json:"message" bson:"message"`

	// The data associated with this log entry, if any
	Data map[string]interface{} `json:"data" bson:"data"`

	// The time this log entry was created
	CreatedAt int64 `json:"created_at" bson:"created_at"`
}

// Create a new LogEntry
func NewLogEntry(id string, service string, level Level, message string, data map[string]interface{}, createdAt int64) *LogEntry {
	return &LogEntry{
		Id:        id,
		Service:   service,
		Level:     level,
		Message:   message,
		Data:      data,
		CreatedAt: createdAt,
	}
}

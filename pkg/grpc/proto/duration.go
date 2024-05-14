package proto

import (
	"fmt"
	"time"
)

// SerializeDuration marshal time.Duration into protobuf duration over grpcgateway
func SerializeDuration(d time.Duration) string {
	return fmt.Sprintf("%.3fs", d.Seconds())
}

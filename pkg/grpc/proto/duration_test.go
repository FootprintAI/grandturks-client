package proto

import (
	"fmt"
	"testing"
	"time"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/stretchr/testify/assert"
)

func TestDurationOverPb(t *testing.T) {
	duration := 180 * time.Second

	durationStr := SerializeDuration(duration)
	fmt.Printf("%s\n", durationStr)
	pbduration, err := runtime.Duration(durationStr)
	assert.NoError(t, err)
	assert.EqualValues(t, pbduration.Seconds, 180)

}

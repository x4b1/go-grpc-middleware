package testpb

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"google.golang.org/protobuf/encoding/protojson"
)

func TestNonDeterministicMarshal(t *testing.T) {
	p := &PingRequest{Value: "something", ErrorCodeReturned: 4}

	// Remove this line to see the test failing (:

	// Sometimes
	// * "value:\"something\" error_code_returned:4"
	// * "value:\"something\"  error_code_returned:4"
	assert.Equal(t, "value:\"something\" error_code_returned:4", p.String())
	b, err := protojson.Marshal(p)
	require.NoError(t, err)

	// Sometimes
	// * "{\"value\":\"something\",\"errorCodeReturned\":4}
	// * "{\"value\":\"something\", \"errorCodeReturned\":4}
	assert.Equal(t, "{\"value\":\"something\",\"errorCodeReturned\":4}", string(b))
}

package pushshift

import (
	"testing"

	assert "github.com/stretchr/testify/require"
)

func TestPushShiftInit(t *testing.T) {
	ps := PushShift{}
	ps.Init()
	assert.NotNil(t, ps)
	assert.NotNil(t, ps.Submissions)
}

func TestPushShiftNew(t *testing.T) {
	ps := New()
	assert.NotNil(t, ps)
	assert.NotNil(t, ps.Submissions)
}

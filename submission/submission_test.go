package submission

import (
	"testing"

	assert "github.com/stretchr/testify/require"
)

func TestSubmissionList(t *testing.T) {
	var searchParams SearchParams
	searchParams.Limit = 1
	assert.NotNil(t, searchParams)
	submissionsList := Search(searchParams)
	assert.NotNil(t, submissionsList.Data[0])
}

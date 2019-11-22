package submission

import (
	"testing"

	assert "github.com/stretchr/testify/require"
)

func TestSubmissionList(t *testing.T) {
	searchParams := SearchParams{Limit: 1}
	submissionsList := Search(searchParams)
	assert.NotNil(t, submissionsList.Data[0])
}

package comment

import (
	"testing"

	assert "github.com/stretchr/testify/require"
)

func TestCommentList(t *testing.T) {
	searchParams := SearchParams{Limit: 1}
	commentList := Search(searchParams)
	assert.NotNil(t, commentList.Data[0])
}

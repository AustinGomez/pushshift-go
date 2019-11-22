package pushshift

import (
	"net/http"

	"github.com/AustinGomez/pushshift-go/submission"
)

// PushShift is an Http.Client wrapper.
type PushShift struct {
	Submissions *submission.Client
}

// Init itializes the API client.
func (ps *PushShift) Init() {
	ps.Submissions = &submission.Client{Backend: &http.Client{}}
}

// New creates a new Glaw client.
func New() *PushShift {
	ps := PushShift{}
	ps.Init()
	return &ps
}

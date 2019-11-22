package pushshift

import "github.com/AustinGomez/pushshift-golang/submission"

// PushShift is an Http.Client wrapper.
type PushShift struct {
	Submissions *submission.Client
}

// Init itializes the API client.
func (ps *PushShift) Init() {
	ps.Submissions = &submission.Client{}
}

// New creates a new Glaw client.
func New() *PushShift {
	ps := PushShift{}
	ps.Init()
	return &ps
}

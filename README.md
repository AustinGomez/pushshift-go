
# pushshift-go
Go wrapper for the PushShift Reddit API. 

PushShift is a free, no signup required API for searching Reddit Submissions, comments, and more. It is the only public API that allows for searching Reddit posts by date since the official Reddit API has removed that feature.

## Installation

Install the package with go get:

`go get -u github.com/AustinGomez/pushshift-golang`


## Usage

You can use this wrapper with or without specifying a client object. This will make more sense with the following examples. For search, you need to use the particular Package's `SearchParams`.

Without the Client, you can just specify the packages you need, and this example uses the submission SearchParams.
```
package main

import (
	"fmt"
	"github.com/AustinGomez/pushshift-go/submission"
)

func main() {
	params := submission.SearchParams{Limit: 1}
	submissions := submission.Search(params)
	fmt.Println("%v", submissions)
}
```


With Client:
```
package main

import (
	"fmt"
	"github.com/AustinGomez/pushshift-go/"
	"github.com/AustinGomez/pushshift-go/submission"
)

func main() {
	ps := pushshift.PushShift{}
	ps.Init()
	params := submission.SearchParams{Limit: 1}
	submissions := submission.Search(params)
	fmt.Println("%v", submissions)
}
```



## Full Documentation
Full  a look at the full documentation at and consider donating at https://pushshift.io/ 

## Restrictions

PushShift rate limits you by IP address to around ~200 requests per minute. No authentication is required.


## Roadmap

Currently only searching Submissions is implemented. Full implementation to come.

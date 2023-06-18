package houstn

import (
	"fmt"
	"net/http"
	"time"
)

type Houstn struct {
	options Options
	ticker  *time.Ticker
	stop    chan bool
	client  *http.Client
}

type Options struct {
	Interval     int
	Application  string
	Environment  string
	Organisation string
	Deployment   string
	Url          string
	Token        string
}

func New(options Options) *Houstn {
	if options.Url == "" {
		options.Url = "https://hello.houstn.io"
	}

	return &Houstn{
		options: options,
		ticker:  time.NewTicker(time.Duration(options.Interval) * time.Second),
		stop:    make(chan bool),
		client:  &http.Client{},
	}
}

func (h *Houstn) Start(metadata any) {
	go func() {
		fmt.Println("Houstn started")
		defer fmt.Println("Houstn stopped")

		for {
			select {
			case <-h.ticker.C:
				h.Ping(metadata)

			case <-h.stop:
				return
			}
		}
	}()
}

func (h *Houstn) Stop() {
	fmt.Println("Houstn stopping")
	h.stop <- true
}

package houstn

import (
	"fmt"
	"net/http"
	"os"
	"strconv"
	"time"
)

type Houstn struct {
	options *Options
	stop    chan bool
	client  *http.Client
}

type Options struct {
	Interval    time.Duration
	Project     string
	Application string
	Environment string
	Url         string
	ApiKey      string
}

func New(options *Options) *Houstn {
	if options.Url == "" {
		options.Url = "https://hello.houstn.io"
	}

	return &Houstn{
		options: options,
		stop:    make(chan bool),
		client:  &http.Client{},
	}
}

func GetOptions(options *Options) *Options {
	if options.Interval == 0 {
		value := Env("HOUSTN_INTERVAL", "5")
		interval, err := strconv.Atoi(value)

		if err != nil {
			fmt.Printf("Error parsing HOUSTN_INTERVAL: %s\n", err)
			return nil
		}

		options.Interval = time.Duration(interval) * time.Second
	}

	if options.Project = ConfigValue(options.Project, "HOUSTN_PROJECT", ""); options.Project == "" {
		fmt.Println("HOUSTN_PROJECT is required")
		return nil
	}

	if options.Environment = ConfigValue(options.Environment, "HOUSTN_ENV", ""); options.Environment == "" {
		fmt.Println("HOUSTN_ENV is required")
		return nil
	}

	if options.Application = ConfigValue(options.Application, "HOUSTN_APP", ""); options.Application == "" {
		fmt.Println("HOUSTN_APP is required")
		return nil
	}

	if options.ApiKey = ConfigValue(options.ApiKey, "HOUSTN_API_KEY", ""); options.ApiKey == "" {
		fmt.Println("HOUSTN_API_KEY is required")
		return nil
	}

	return options
}

func (h *Houstn) Start(metadata any) {
	options := GetOptions(h.options)

	if options == nil {
		fmt.Println("Valid options are required")
		return
	}

	go func() {
		ticker := time.NewTicker(options.Interval)

		fmt.Println("Houstn started")
		defer fmt.Println("Houstn stopped")

		for {
			select {
			case <-ticker.C:
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

func ConfigValue(value string, env string, defaultValue string) string {
	if value != "" {
		return value
	}

	return Env(env, defaultValue)
}

func Env(key string, defaultValue string) string {
	value := os.Getenv(key)

	if value == "" {
		return defaultValue
	}

	return value
}

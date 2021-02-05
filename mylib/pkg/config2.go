package pkg

import "io"

type Config struct {
	Method string
	URL      string
	Body io.Reader
}

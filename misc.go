package utils

import "fmt"

// Misc - misc
type Misc struct{}

// NewMisc - create new misc
func NewMisc() Misc {
	return Misc{}
}

// URLBuilder - URL string builder
type URLBuilder struct {
	Scheme     string
	URL        string
	PortString string
}

// NewURLBuilder - returns new url builder
func (m Misc) NewURLBuilder(opts ...string) URLBuilder {
	u := URLBuilder{
		Scheme:     "http",
		URL:        "localhost",
		PortString: "",
	}
	if len(opts) >= 1 {
		u.Scheme = opts[0]
	}
	if len(opts) >= 2 {
		u.URL = opts[1]
	}
	if len(opts) >= 3 {
		u.PortString = fmt.Sprintf(":%s", opts[2])
	}
	return u
}

// Build - builds and returns URL string
func (u URLBuilder) Build() string {
	return fmt.Sprintf("%s://%s%s", u.Scheme, u.URL, u.PortString)
}

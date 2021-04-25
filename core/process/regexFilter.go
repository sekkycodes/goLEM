package process

import "regexp"

// RegexFilter will evaluate input against a regex and only forward items matching the pattern
type RegexFilter struct {
	endpoint Processor
	regex    regexp.Regexp
}

// Initialize new RegexFilter, check whether pattern compiles
func NewRegexFilter(endpoint Processor, pattern string) *RegexFilter {
	regex := regexp.MustCompile(pattern)
	return &RegexFilter{endpoint: endpoint, regex: *regex}
}

func (filter *RegexFilter) Process(input []byte) {
	if filter.regex.Match(input) {
		filter.endpoint.Process(input)
	}
}

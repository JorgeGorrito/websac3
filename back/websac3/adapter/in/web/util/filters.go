package util

import (
	"net/url"
	"strings"
	"websac3/common/filter"
)

func ParseParamsFilter(q url.Values) filter.Params {
	params := make(filter.Params)
	for key, values := range q {
		if len(values) == 0 {
			continue
		}
		value := values[0]

		if strings.Contains(key, "[") && strings.HasSuffix(key, "]") {
			field := key[:strings.Index(key, "[")]
			op := key[strings.Index(key, "[")+1 : len(key)-1]

			if _, exists := params[field]; !exists {
				params[field] = make(map[string]string)
			}
			params[field][op] = value
		}
	}
	return params
}

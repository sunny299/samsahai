package internal

import (
	"regexp"
	"strconv"
	"strings"
)

type SortableVersion []string

func (v SortableVersion) Len() int { return len(v) }
func (v SortableVersion) Less(i, j int) bool {
	is := v.regexpSplit(v[i])
	js := v.regexpSplit(v[j])

	reNum := regexp.MustCompile(`\d+`)

	for i := 0; i < len(is) && i < len(js); i++ {
		// sort number
		if reNum.MatchString(is[i]) && reNum.MatchString(js[i]) {
			ival, _ := strconv.Atoi(is[i])
			jval, _ := strconv.Atoi(js[i])
			if ival == jval {
				continue
			}
			return ival < jval
		}

		return strings.Compare(is[i], js[i]) < 0
	}

	return len(is) < len(js)
}
func (v SortableVersion) Swap(i, j int) { v[i], v[j] = v[j], v[i] }

func (v SortableVersion) regexpSplit(input string) []string {
	reg := regexp.MustCompile("[._-]")
	split := reg.Split(input, -1)
	var output []string
	for i := range split {
		if split[i] == "" {
			continue
		}
		output = append(output, split[i])
	}
	return output
}

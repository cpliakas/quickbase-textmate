package qbtm

import "sort"

// Opts contaiins command options.
type Opts struct {
	File string `cliutil:"option=file default=data/Formula_Functions_Reference.csv func=ioreader usage='path to the csv file exported from https://login.quickbase.com/db/6ewwzuuj?a=td'"`
}

// Unique sorts and deduplicates a slice of strings.
// See https://github.com/golang/go/wiki/SliceTricks#in-place-deduplicate-comparable
func Unique(in []string) []string {
	if len(in) == 0 {
		return []string{}
	}

	sort.Strings(in)

	j := 0
	for i := 1; i < len(in); i++ {
		if in[j] == in[i] {
			continue
		}
		j++
		in[j] = in[i]
	}

	return in[:j+1]
}

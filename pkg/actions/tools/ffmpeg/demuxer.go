package ffmpeg

import (
	"regexp"
	"strings"

	"github.com/rsteube/carapace"
)

// ActionDemuxers completes demuxers
//
//	aax (CRI AAX)
//	ac3 (raw AC-3)
func ActionDemuxers() carapace.Action {
	return carapace.ActionExecCommand("ffmpeg", "-hide_banner", "-demuxers")(func(output []byte) carapace.Action {
		lines := strings.Split(string(output), "\n")
		r := regexp.MustCompile(`^ (?P<type>.).{1} (?P<name>[^ ]+) +(?P<description>.*)$`)

		found := false
		vals := make([]string, 0)
		for _, line := range lines {
			if !found {
				found = line == " --"
				continue
			}

			if matches := r.FindStringSubmatch(line); matches != nil {
				vals = append(vals, matches[2], matches[3])
			}
		}
		return carapace.ActionValuesDescribed(vals...)
	}).Tag("demuxers")
}

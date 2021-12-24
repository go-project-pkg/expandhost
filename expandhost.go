package expandhost

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

// PatternToHosts pattern e.g.:
// foo[01-66,88,99-101].bar.com
// foo[01-66,88,99-101].idc[1-6].bar.com
// foo[01-66,88,99-101].[bj,sh,wh,sz].idc[1-6].bar.com
func PatternToHosts(pattern string) ([]string, error) {
	sep := regexp.MustCompile(`\[|\]`)
	parts := sep.Split(pattern, 3)

	if len(parts) == 1 {
		return []string{pattern}, nil
	}

	hostPart1 := parts[0]
	hostPattern := parts[1]
	hostPart3 := parts[2]

	var numbers []string
	for _, v := range strings.Split(hostPattern, ",") {
		if len(strings.Split(v, "-")) == 1 {
			numbers = append(numbers, strings.TrimSpace(v))
			continue
		}

		res, err := expandNumberRange(v, "-")
		if err != nil {
			return nil, err
		}

		numbers = append(numbers, res...)
	}

	var hosts []string
	for _, v := range numbers {
		hosts = append(hosts, fmt.Sprintf("%s%s%s", hostPart1, v, hostPart3))
	}

	_parts := sep.Split(hostPart3, 3)
	if len(_parts) == 3 {
		var newHosts []string
		for _, v := range hosts {
			_newHosts, err := PatternToHosts(v)
			if err != nil {
				return nil, err
			}
			newHosts = append(newHosts, _newHosts...)
		}

		return newHosts, nil
	}

	return hosts, nil
}

func expandNumberRange(numberRange, sep string) ([]string, error) {
	numberRange = strings.TrimSpace(numberRange)
	numbers := strings.Split(numberRange, sep)

	numberBeginStr := numbers[0]
	numberEndStr := numbers[1]

	numberBeginStrLen := len(numberBeginStr)

	numberBegin, err := strconv.Atoi(numberBeginStr)
	if err != nil {
		return nil, err
	}

	numberEnd, err := strconv.Atoi(numberEndStr)
	if err != nil {
		return nil, err
	}

	var res []string
	for i := numberBegin; i <= numberEnd; i++ {
		vStr := strconv.Itoa(i)
		vLen := len(vStr)

		if vLen != numberBeginStrLen {
			for i := 1; i <= numberBeginStrLen-vLen; i++ {
				vStr = "0" + vStr
			}

			res = append(res, vStr)
		} else {
			res = append(res, strconv.Itoa(i))
		}
	}

	return res, nil
}

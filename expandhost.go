package expandhost

import (
	"errors"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

// PatternToHosts pattern e.g.: foo[01-66,88,99-101].bar.com
func PatternToHosts(pattern string) ([]string, error) {
	sep := regexp.MustCompile(`\[|\]`)
	parts := sep.Split(pattern, -1)

	if len(parts) != 3 {
		return nil, errors.New("invalid host pattern")
	}

	hostPart1 := parts[0]
	hostPattern := parts[1]
	hostPart3 := parts[2]

	var numbers []string
	for _, v := range strings.Split(hostPattern, ",") {
		if len(strings.Split(v, "-")) == 1 {
			numbers = append(numbers, v)
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

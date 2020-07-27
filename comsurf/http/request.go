package http

import (
	"errors"
	"fmt"
	"net/url"
	"strconv"
	"strings"
	"time"
)

type Request struct {
	// Method is the HTTP/1.1 method name.
	Method string

	// URL is the URL requested by the client.
	URL *url.URL

	// Headers contains pairs of the header key-values.
	Headers map[string]string

	// Ranges contains all the requested byte-ranges.
	Ranges []ByteRange

	isRanged bool
}

// ByteRange is the parsed ranges requested by the client.
// For example, the client could request the following byte range:
//
//     Range: bytes=100-
//
// This header is parsed with 100 and -1 (unspecified) being the start and end of the range.
type ByteRange struct {
	Start int64
	End   int64
}

// Length is the total number of bytes included in the range.
func (br ByteRange) Length() int64 {
	return br.End - br.Start + 1
}

func parseByteRangeHeader(headerValue string) (byteRanges []ByteRange, explicit bool) {
	rangePrefix := "bytes="
	byteRanges = make([]ByteRange, 0)

	if !strings.HasPrefix(headerValue, rangePrefix) {
		byteRanges = append(byteRanges, ByteRange{Start: 0, End: -1})
		return
	}

	explicit = true

	headerValue = headerValue[len(rangePrefix):]
	for _, value := range strings.Split(headerValue, ",") {

		// Let's say we have 10 bytes: [0, 1, 2, 3, 4, 5, 6, 7, 8, 9]
		// -10 => [0, 1, 2, 3, 4, 5, 6, 7, 8, 9]
		// -7 => [3, 4, 5, 6, 7, 8, 9]
		// -2 => [8, 9]
		// -1 => [9]
		if val, err := strconv.ParseInt(value, 10, 0); err == nil && val < 0 {
			byteRanges = append(byteRanges, ByteRange{Start: val, End: -1})
			continue
		}

		// 0- => [0, 1, 2, 3, 4, 5, 6, 7, 8, 9]
		// 3- => [3, 4, 5, 6, 7, 8, 9]
		if strings.HasSuffix(value, "-") {
			if val, err := strconv.ParseInt(value[:len(value)-1], 10, 0); err == nil {
				byteRanges = append(byteRanges, ByteRange{Start: val, End: -1})
			}
			continue
		}

		// 1-1 => [1, 1]
		// 3-6 => [3, 4, 5, 6]
		rangeVals := strings.Split(value, "-")
		val1, err1 := strconv.ParseInt(rangeVals[0], 10, 0)
		val2, err2 := strconv.ParseInt(rangeVals[1], 10, 0)
		if err1 == nil && err2 == nil {
			byteRanges = append(byteRanges, ByteRange{Start: val1, End: val2})
		}
	}
	return
}

// ParseHTTPDate is a helper for parsing HTTP-dates.
func ParseHTTPDate(date string) (t time.Time) {
	t, err := time.Parse(httpTimeFormat, date)
	if err != nil {
		fmt.Println("error parsing", err)
	}
	return
}

func (req *Request) parseInitialLine(line string) (err error) {
	words := strings.SplitN(line, " ", 3)

	if len(words) < 3 {
		return errors.New("Invalid initial request line.")
	}

	if words[2] != "HTTP/1.1" {
		return errors.New("Invalid initial request line.")
	}

	req.Method = words[0]
	req.URL, _ = url.Parse(words[1])

	return
}

func (req *Request) parseHeaders(headerLines []string) {
	for _, headerLine := range headerLines {
		headerPair := strings.SplitN(headerLine, ": ", 2)
		if len(headerPair) == 2 {
			req.Headers[headerPair[0]] = headerPair[1]
		}
	}
	req.Ranges, req.isRanged = parseByteRangeHeader(req.Headers["Range"])
}

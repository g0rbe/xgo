package syslog

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"strconv"
	"time"
)

const timeStampSize = 15

var ErrInvalidMessageLength = errors.New("invalid mesage length")

// Message is the BSD syslog message format (RFC3164)
type Message struct {
	Priority  Priority  `json:"priority"`
	Timestamp time.Time `json:"timestamp"`
	Hostname  string    `json:"hostname"`
	Tag       string    `json:"tag"`
	Content   string    `json:"content"`
}

// parsePriority returns the priority string, the remaining buffer
func parsePriority(buf []byte) (Priority, []byte, error) {

	if buf[0] != '<' {
		return 0, nil, fmt.Errorf("leading \"<\" is missing")
	}

	buf = buf[1:]

	// Trailing ">"
	n := bytes.IndexByte(buf, '>')
	if n == -1 {
		return 0, nil, fmt.Errorf("trailing \">\" is missing")
	}

	v, err := strconv.Atoi(string(buf[:n]))
	if err != nil {
		return 0, nil, fmt.Errorf("failed to parse integer: %w", err)
	}

	buf = buf[n+1:]

	return Priority(v), buf, nil
}

// parseTimestamp returns the date and the remaining buffer
func parseTimestamp(buf []byte) (time.Time, []byte, error) {

	d, err := time.Parse(time.Stamp, string(buf[:timeStampSize]))
	if err != nil {
		return time.Time{}, nil, fmt.Errorf("invalid date: %w", err)
	}

	return d, buf[timeStampSize+1:], nil
}

// parseHostname returns the hostname and the remaining buffer
func parseHostname(buf []byte) (string, []byte, error) {

	var v []byte
	var found bool

	v, buf, found = bytes.Cut(buf, []byte{' '})
	if !found {
		return "", nil, fmt.Errorf("hostname is missing")
	}

	return string(v), buf, nil
}

// parseTag returns the app-name and the remaining buffer
func parseTag(buf []byte) (string, []byte, error) {

	var v []byte
	var found bool

	v, buf, found = bytes.Cut(buf, []byte{' '})
	if !found {
		return "", nil, fmt.Errorf("app-name is missing")
	}

	if v[len(v)-1] != ':' {
		return "", nil, fmt.Errorf("missing trailing colon")
	}

	v = bytes.TrimSuffix(v, []byte{':'})

	return string(v), buf, nil
}

/*
ParseMessage parses the message in the given format:

	<190>Jan _2 15:04:05 hostname tag: content
*/
func ParseMessage(buf []byte) (*Message, error) {

	var err error
	msg := new(Message)

	if len(buf) < 68 {
		return nil, fmt.Errorf("%s (%d)", ErrInvalidMessageLength, len(buf))
	}

	msg.Priority, buf, err = parsePriority(buf)
	if err != nil {
		return nil, fmt.Errorf("invalid priority: %w", err)
	}

	msg.Timestamp, buf, err = parseTimestamp(buf)
	if err != nil {
		return nil, fmt.Errorf("invalid timestamp: %w", err)
	}

	msg.Hostname, buf, err = parseHostname(buf)
	if err != nil {
		return nil, fmt.Errorf("invalid hostname: %w", err)
	}

	msg.Tag, buf, err = parseTag(buf)
	if err != nil {
		return nil, fmt.Errorf("invalid tag: %w", err)
	}

	msg.Content = string(buf)

	return msg, nil
}

func (m *Message) String() string {

	out, err := json.Marshal(m)
	if err != nil {
		panic(fmt.Errorf("failed to marshal Message: %w", err))
	}

	return string(out)
}

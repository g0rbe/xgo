package syslog

import (
	"bytes"
	"errors"
	"fmt"
	"time"
)

var ErrInvalidMessageLength = errors.New("invalid mesage length")

type Message struct {
	Priority string
	Date     time.Time
	Hostname string
	App      string
	Msg      string
}

// parsePriority returns the priority string, the remaining buffer
func parsePriority(buf []byte) (string, []byte, error) {

	if buf[0] != '<' {
		return "", nil, fmt.Errorf("leading \"<\" is missing")
	}

	buf = buf[1:]

	// Trailing ">"
	n := bytes.IndexByte(buf, '>')
	if n == -1 {
		return "", nil, fmt.Errorf("trailing \">\" is missing")
	}

	v := string(buf[:n])
	buf = buf[n+1:]

	return v, buf, nil
}

// parseTimestamp returns the date and the remaining buffer
func parseTimestamp(buf []byte) (time.Time, []byte, error) {

	// Split: "month date time ..."
	fields := bytes.SplitN(buf, []byte{' '}, 4)
	if len(fields) != 4 {
		return time.Time{}, nil, fmt.Errorf("date date fields")
	}

	d, err := time.Parse(time.Stamp, string(bytes.Join(fields[:3], []byte{' '})))
	if err != nil {
		return time.Time{}, nil, fmt.Errorf("invalid date: %w", err)
	}

	return d, fields[3], nil
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

// parseApp returns the app-name and the remaining buffer
func parseApp(buf []byte) (string, []byte, error) {

	var v []byte
	var found bool

	v, buf, found = bytes.Cut(buf, []byte{' '})
	if !found {
		return "", nil, fmt.Errorf("app-name is missing")
	}

	if v[len(v)-1] != ':' {
		return "", nil, fmt.Errorf("missing trailing colon")
	}

	v = v[:len(v)-1]

	return string(v), buf, nil
}

/*
ParseMessage parses the message in the given format:

	<190>Jan _2 15:04:05 hostname app: message
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

	msg.Date, buf, err = parseTimestamp(buf)
	if err != nil {
		return nil, fmt.Errorf("invalid timestamp: %w", err)
	}

	msg.Hostname, buf, err = parseHostname(buf)
	if err != nil {
		return nil, fmt.Errorf("invalid hostname: %w", err)
	}

	msg.App, buf, err = parseApp(buf)
	if err != nil {
		return nil, fmt.Errorf("invalid app-name: %w", err)
	}

	msg.Msg = string(buf)

	return msg, nil
}

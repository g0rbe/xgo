package syslog

type Severity int

const (
	EMERG Severity = iota
	ALERT
	CRIT
	ERR
	WARNING
	NOTICE
	INFO
	DEBUG
)

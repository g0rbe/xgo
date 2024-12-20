package syslog

type Facility int

const (
	KERN Facility = iota
	USER
	MAIL
	DAEMON
	AUTH
	SYSLOG
	LPR
	NEWS
	UUCP
	CRON
	AUTHPRIC
	FTP
	NTP
	SECURITY
	CONSOLE
	SOLARIS_CRON
	LOCAL0
	LOCAL1
	LOCAL2
	LOCAL3
	LOCAL4
	LOCAL5
	LOCAL6
	LOCAL7
)

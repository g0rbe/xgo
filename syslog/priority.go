package syslog

type Priority int

func CalculatePriority(f Facility, s Severity) Priority {
	return Priority(int(f)*8 + int(s))
}

func (p Priority) Severity() Severity {
	return Severity(p % 8)
}

func (p Priority) Facility() Facility {
	return Facility((p - Priority(p.Severity())) / 8)
}

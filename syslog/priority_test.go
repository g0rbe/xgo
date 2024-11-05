package syslog_test

import (
	"testing"

	"github.com/g0rbe/xgo/syslog"
)

func TestPriority(t *testing.T) {

	testFacility := syslog.Facility(23)
	testSeverity := syslog.Severity(5)

	pri := syslog.CalculatePriority(testFacility, testSeverity)

	t.Logf("Priority: %d\n", pri)
	t.Logf("Facility: %d\n", pri.Facility())
	t.Logf("Severty: %d\n", pri.Severity())

	if pri.Facility() != testFacility {
		t.Fatalf("Invalid Facility: %d\n", pri.Facility())
	}

	if pri.Severity() != testSeverity {
		t.Fatalf("Invalid Severity: %d\n", pri.Severity())
	}
}

package syslog_test

import (
	"testing"

	"gorbe.io/go/syslog"
)

func TestServer(t *testing.T) {

	s, err := syslog.NewServer(":10000")
	if err != nil {
		t.Fatalf("Failed to create new server: %s\n", err)
	}

	buf, err := s.Read()
	if err != nil {
		t.Fatalf("Failed to read: %s\n", err)
	}

	t.Logf("%#v\n", buf)

	s.Close()
}

package output

import (
	"bytes"
	"fmt"
	"strings"
	"testing"
)

func TestStdoutl(t *testing.T) {
	// Arrange
	var (
		buf     = bytes.Buffer{}
		prefix  = "test"
		message = "hello"
		want    = fmt.Sprintf("%s %s", prefix, message)
	)

	orig := Writer()
	SetWriter(&buf)
	defer SetWriter(orig)

	// Act
	Stdoutl(prefix, message)

	// Assert
	got := strings.TrimSpace(buf.String())
	if got != want {
		t.Errorf("want:%s\tgot:%s", want, got)
	}
}

func TestStdoutf(t *testing.T) {
	// Arrange
	var (
		buf     = bytes.Buffer{}
		prefix  = "test"
		format  = "%s world"
		message = "hello"
		want    = fmt.Sprintf("%s %s world", prefix, message)
	)

	orig := Writer()
	SetWriter(&buf)
	defer SetWriter(orig)

	// Act
	Stdoutf(prefix, format, message)

	// Assert
	got := strings.TrimSpace(buf.String())
	if got != want {
		t.Errorf("want:%s\tgot:%s", want, got)
	}
}

func TestStderrl(t *testing.T) {
	// Arrange
	var (
		buf     = bytes.Buffer{}
		prefix  = "test"
		message = "hello"
		want    = fmt.Sprintf("%s %s", prefix, message)
	)

	orig := ErrWriter()
	SetErrWriter(&buf)
	defer SetErrWriter(orig)

	// Act
	Stderrl(prefix, message)

	// Assert
	got := strings.TrimSpace(buf.String())
	if got != want {
		t.Errorf("want:%s\tgot:%s", want, got)
	}
}

func TestStderrf(t *testing.T) {
	// Arrange
	var (
		buf     = bytes.Buffer{}
		prefix  = "test"
		format  = "%s world"
		message = "hello"
		want    = fmt.Sprintf("%s %s world", prefix, message)
	)

	orig := ErrWriter()
	SetErrWriter(&buf)
	defer SetErrWriter(orig)

	// Act
	Stderrf(prefix, format, message)

	// Assert
	got := strings.TrimSpace(buf.String())
	if got != want {
		t.Errorf("want:%s\tgot:%s", want, got)
	}
}

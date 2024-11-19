package color_test

import (
	"strings"
	"testing"

	"github.com/g0rbe/xgo/color"
)

func TestColorize(t *testing.T) {

	testText := "text"

	v := color.Colorize(color.RED, testText)
	if strings.Compare(v, color.RED+testText+color.RESET) != 0 {
		t.Fatalf("Invalid string: %#v\n", []byte(v))
	}
}

func TestRed(t *testing.T) {

	testText := "text"

	v := color.Red(testText)
	if strings.Compare(v, color.RED+testText+color.RESET) != 0 {
		t.Fatalf("Invalid string: %#v\n", []byte(v))
	}
}

func TestBlue(t *testing.T) {

	testText := "text"

	v := color.Blue(testText)
	if strings.Compare(v, color.BLUE+testText+color.RESET) != 0 {
		t.Fatalf("Invalid string: %#v\n", []byte(v))
	}
}

func TestCyan(t *testing.T) {

	testText := "text"

	v := color.Cyan(testText)
	if strings.Compare(v, color.CYAN+testText+color.RESET) != 0 {
		t.Fatalf("Invalid string: %#v\n", []byte(v))
	}
}

func TestGray(t *testing.T) {

	testText := "text"

	v := color.Gray(testText)
	if strings.Compare(v, color.GRAY+testText+color.RESET) != 0 {
		t.Fatalf("Invalid string: %#v\n", []byte(v))
	}
}

func TestGreen(t *testing.T) {

	testText := "text"

	v := color.Green(testText)
	if strings.Compare(v, color.GREEN+testText+color.RESET) != 0 {
		t.Fatalf("Invalid string: %#v\n", []byte(v))
	}
}

func TestPurple(t *testing.T) {

	testText := "text"

	v := color.Purple(testText)
	if strings.Compare(v, color.PURPLE+testText+color.RESET) != 0 {
		t.Fatalf("Invalid string: %#v\n", []byte(v))
	}
}

func TestWhite(t *testing.T) {

	testText := "text"

	v := color.White(testText)
	if strings.Compare(v, color.WHITE+testText+color.RESET) != 0 {
		t.Fatalf("Invalid string: %#v\n", []byte(v))
	}
}

func TestYellow(t *testing.T) {

	testText := "text"

	v := color.Yellow(testText)
	if strings.Compare(v, color.YELLOW+testText+color.RESET) != 0 {
		t.Fatalf("Invalid string: %#v\n", []byte(v))
	}
}

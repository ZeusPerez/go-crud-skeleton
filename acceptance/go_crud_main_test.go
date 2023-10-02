package acceptance

import (
	"os"
	"testing"
)

var ()

func TestMain(m *testing.M) {
	code := m.Run()

	os.Exit(code)
}

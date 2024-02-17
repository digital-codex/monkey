package object

import (
	"github.com/digital-codex/assertions"
	"testing"
)

func TestStringHashKey(t *testing.T) {
	hello := &String{Value: "Hello World"}
	check := &String{Value: "Hello World"}

	goodbye := &String{Value: "Goodbye Moon"}

	assertions.AssertEquals(t, hello.HashKey(), check.HashKey(), "strings with check content have different hash keys")
	assertions.AssertNotEquals(t, hello.HashKey(), goodbye.HashKey(), "strings with different content have check hash keys")
}

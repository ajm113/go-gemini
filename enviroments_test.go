package gemini

import "testing"

func TestEnvironmentValid(t *testing.T) {
	var test Environment
	test = "wrongurl"

	if test.IsValid() {
		t.Error("Should have properly checked for gemini.Sandbox AND gemini.Live!")
	}

	test = Sandbox
	if !test.IsValid() {
		t.Error("gemini.Sandbox IsValid should have return true!")
	}

	test = Live
	if !test.IsValid() {
		t.Error("gemini.Live IsValid should have return true!")
	}
}

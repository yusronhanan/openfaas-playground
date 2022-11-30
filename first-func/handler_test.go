package function

import "testing"

func TestHandleReturnsCorrectResponse(t *testing.T) {
	expected := "Hello, Go. You said: First Test!"
	response := Handle([]byte("First Test!"))

	if response != expected {
		t.Fatalf("Expected: %v, Got: %v", expected, response)
	}
}

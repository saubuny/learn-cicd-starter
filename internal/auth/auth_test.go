package auth

import (
	"net/http"
	"reflect"
	"testing"
)

func TestGetApiKey(t *testing.T) {
	headerPass := make(http.Header, 0)
	headerPass.Add("Authorization", "ApiKey blahblahblah")
	got, err := GetAPIKey(headerPass)
	want := "blahblahblah"
	if err != nil {
		t.Fatalf("expected: %v, got: %v", want, err.Error())
	}

	if !reflect.DeepEqual(want, got) {
		t.Fatalf("expected: %v, got: %v", want, got)
	}

	t.Fatalf("Force Fail!!")
}

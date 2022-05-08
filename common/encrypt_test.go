package common

import (
	"testing"
)

func TestEncryptAndVersity(t *testing.T) {

	secretSchema := "2fa6c1e9"

	encrpyContext := CryptContext{secretSchema}

	secret := "password1111"

	got, err := encrpyContext.Encrypt(secret)
	if err != nil {
		t.Logf(err.Error())
	}

	want := "874befd395d4f88c5c2b55a7bf2ad2ae"

	if got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}

	versify, err := encrpyContext.Verify(secret, got)

	if versify != true {
		t.Errorf("should equal")
	}

}

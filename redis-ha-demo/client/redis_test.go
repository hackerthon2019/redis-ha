package client

import (
	"testing"
)

func TestSetKV(t *testing.T) {
	if err := SetKV("abc", "123"); err != nil {
		t.Error(err.Error())
	}
}

func TestGetKV(t *testing.T) {
	res, err := GetKV("abc")
	if err != nil {
		t.Error(err.Error())
	}
	t.Log(res)
}

package model

import "testing"

func TestRfdata(t *testing.T) {
	userDatas = make(map[string]Model, 0)
	rfdata("user", "username", userDatas)
}
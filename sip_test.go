package xk6_sip

import (
	"testing"
)

func TestSip(t *testing.T) {

	sip := &Sip{}

	sip.InitSip()

	sip.Register()

}

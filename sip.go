package xk6_sip

import (
	"github.com/cloudwebrtc/go-sip-ua/pkg/account"
	"github.com/cloudwebrtc/go-sip-ua/pkg/stack"
	"github.com/cloudwebrtc/go-sip-ua/pkg/ua"
	"github.com/ghettovoice/gosip/sip/parser"
	"go.k6.io/k6/js/modules"
)

type Sip struct {
	userAgent *ua.UserAgent
	sipStack  *stack.SipStack
}

func (sip *Sip) InitSip() {
	sip.sipStack = stack.NewSipStack(&stack.SipStackConfig{
		UserAgent:  "xk6-sip",
		Extensions: []string{"replaces", "outbound"},
		Dns:        "8.8.8.8"})

	if err := sip.sipStack.Listen("tcp", "0.0.0.0:5080"); err != nil {
		panic(err)
	}

	sip.userAgent = ua.NewUserAgent(&ua.UserAgentConfig{SipStack: sip.sipStack})
	if sip.userAgent == nil {
		panic("UserAgent is nil")
	}
}

func (sip *Sip) Register() {
	uri, err := parser.ParseUri("sip:100@127.0.0.1")
	if err != nil {
		panic(err)
	}

	profile := account.NewProfile(uri.Clone(), "goSIP/example-client",
		&account.AuthInfo{
			AuthUser: "100",
			Password: "100",
			Realm:    "",
		},
		1800,
		sip.sipStack,
	)

	recipient, err := parser.ParseSipUri("sip:100@127.0.0.1:5081;transport=tcp")
	if err != nil {
		panic(err)
	}

	_, err = sip.userAgent.SendRegister(profile, recipient, profile.Expires, nil)
	if err != nil {
		panic(err)
	}
}

func init() {
	modules.Register("k6/x/sip", new(Sip))
}

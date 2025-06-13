package syslog

import (
	"github.com/kentik/the-library-formally-known-as-go-syslog/format"
	. "gopkg.in/check.v1"
)

type HandlerSuite struct{}

var _ = Suite(&HandlerSuite{})

func (s *HandlerSuite) TestHandle(c *C) {
	logPart := format.LogParts{"tag": "foo"}

	channel := make(LogPartsChannel, 1)
	handler := NewChannelHandler(channel)
	handler.Handle(logPart, 10, nil)

	fromChan := <-channel
	c.Check(fromChan["tag"], Equals, logPart["tag"])
}

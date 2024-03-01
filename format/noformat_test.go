package format

import (
   . "gopkg.in/check.v1"
)

func (s *FormatSuite) TestNoFormat_SingleSplit(c *C) {
   f := NoFormat{}
   c.Assert(f.GetSplitFunc(), IsNil)
}

func (s *FormatSuite) TestNoFormat_CorrectParsingTypical(c *C) {
   f := NoFormat{}

   find := `<13>May  1 20:51:40 myhostname myprogram: ciao`
   parser := f.GetParser([]byte(find))
   err := parser.Parse()
   c.Assert(err, IsNil)
   c.Assert(parser.Dump()["content"], Equals, find)
}

func (s *FormatSuite) TestNoFormat_CorrectParsingTypicalWithPID(c *C) {
   f := NoFormat{}

   find := `<13>May  1 20:51:40 myhostname myprogram[42]: ciao`
   parser := f.GetParser([]byte(find))
   err := parser.Parse()
   c.Assert(err, IsNil)
   c.Assert(parser.Dump()["content"], Equals, find)
}

func (s *FormatSuite) TestNoFormat_CorrectParsingGNU(c *C) {
   // GNU implementation of syslog() has a variant: hostname is missing
   f := NoFormat{}

   find := `<13>May  1 20:51:40 myprogram: ciao`
   parser := f.GetParser([]byte(find))
   err := parser.Parse()
   c.Assert(err, IsNil)
   c.Assert(parser.Dump()["content"], Equals, find)
}

func (s *FormatSuite) TestNoFormat_CorrectParsingJournald(c *C) {
   // GNU implementation of syslog() has a variant: hostname is missing
   // systemd uses it, and typically also passes PID
   f := NoFormat{}

   find := `<78>May  1 20:51:02 myprog[153]: blah`
   parser := f.GetParser([]byte(find))
   err := parser.Parse()
   c.Assert(err, IsNil)
   c.Assert(parser.Dump()["content"], Equals, find)
}

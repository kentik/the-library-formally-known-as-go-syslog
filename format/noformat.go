package format

import (
   "bufio"
   "time"
)

// --- Dummy format that just puts the whole log line in logParts["content"] ---

// Actually, the syslog.RFC3164 format also puts the unparsed "content" field,
// but wastes times trying to parse lines that have a custom format.

var noFormat = &NoFormat{}

type NoFormat struct{}

func (f NoFormat) GetParser(line []byte) LogParser {
   return &noFormatParser{string(line)}
}

func (f NoFormat) GetSplitFunc() bufio.SplitFunc {
   return nil // not used
}

type noFormatParser struct {
   line string
}

func (c noFormatParser) Dump() LogParts {
   return LogParts{
      "content": string(c.line),
   }
}

func (c noFormatParser) Parse() error {
   return nil // doesn't parse anything
}

func (c noFormatParser) Location(location *time.Location) {
   // not used
}

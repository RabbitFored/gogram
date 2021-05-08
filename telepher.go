package telepher

import (
  "net/http"
  "time"
  )
const (
  DefaultBaseURL = "https://api.telegram.org/bot"
  DefaultBaseFileURL = "https://api.telegram.org/file/bot"
  DefaultParseMode = "None"
 )
var DefaultClient = &http.Client{
  Timeout: time.Second * 10,
}

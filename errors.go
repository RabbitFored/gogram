package telepher

import (
  "errors"
  )
  

var (
  
  invalidTokenErr = errors.New("invalid token provided")
  invalidOptions  = errors.New("unsupported options argument")
)

package types

import (
  "strings"
  )

func (msg *Message) Args() []string{
  if msg == nil{
    return nil
  }
  if msg.Text != ""{
  args := strings.Fields(msg.Text)
 
  return args
  } else if msg.Caption != ""{
    args := strings.Fields(msg.Caption)
      
    return args
  } 
  return nil
}

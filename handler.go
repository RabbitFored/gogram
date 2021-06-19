package telepher

import (
"github.com/goTelegramBot/telepher/types"
"sort"
"regexp"

)

type Handler struct {
  Type      string
  Trigger      *regexp.Regexp
	function  interface{}
  HandlerOptions *HandlerOptions
}

type Commands struct{
  Command          string
  function         func(Bot,*types.Message)
  HandlerOptions   *HandlerOptions
}
type messageHandler struct {
	function  interface{}
  HandlerOptions *HandlerOptions
}

type HandlerOptions struct {
  groupID int
  filters func(message *types.Message) bool
}


func (telepher *Bot) On(Type string, handler interface{},handlerOptions *HandlerOptions){
  
  if handlerOptions != nil{
telepher.Handlers[handlerOptions.groupID] = append(telepher.Handlers[handlerOptions.groupID], Handler{Type:Type,function: handler,HandlerOptions:handlerOptions})


}else{
  telepher.Handlers[DefaultHandlerGroup] = append(telepher.Handlers[DefaultHandlerGroup])
}
}



func (telepher Bot) handle(Type string, msg *types.Message) {
var groupIDs []int

for i := range telepher.Handlers {
 groupIDs =  append(groupIDs,i)
}
  sort.Ints(groupIDs)

 for _,groupID := range  groupIDs {
   	for _, handler := range telepher.Handlers[groupID]{
           
    if handler.Type == Type {
            if handler.HandlerOptions.filters != nil{
       if handler.HandlerOptions.filters(msg){
         
		     handler.function.(func (Bot, *types.Message))(telepher,msg)}
            }else{
              handler.function.(func (Bot, *types.Message))(telepher,msg)
           }
             
  }
     }}

 }
	
func (telepher Bot) Command(trigger string,commandHandler func(Bot,*types.Message)){
  
  telepher.Commands[trigger] = append(telepher.Commands[trigger], Commands{function:commandHandler})
  
}
func (telepher Bot) CommandHandler(trigger string, msg *types.Message){
    

  	for _, command := range telepher.Commands[trigger] {

			command.function(telepher,msg)
      return
}
}



func (telepher Bot) Hears(trigger string, handler interface{},handlerOptions *HandlerOptions){
  rx := regexp.MustCompile(trigger)
  if handlerOptions != nil{
telepher.Handlers[handlerOptions.groupID] = append(telepher.Handlers[handlerOptions.groupID], Handler{Trigger:rx,function: handler,HandlerOptions:handlerOptions})


}else{
  telepher.Handlers[DefaultHandlerGroup] = append(telepher.Handlers[DefaultHandlerGroup])
}
  
}

func (telepher Bot) handleMessage(msg *types.Message) {
var groupIDs []int

for i := range telepher.Handlers {
 groupIDs =  append(groupIDs,i)
}
  sort.Ints(groupIDs)

 for _,groupID := range  groupIDs {
   	for _, handler := range telepher.Handlers[groupID]{
           
    if handler.Trigger.MatchString(msg.Text) {
            if handler.HandlerOptions.filters != nil{
       if handler.HandlerOptions.filters(msg){
         
		     handler.function.(func (Bot, *types.Message))(telepher,msg)}
            }else{
              handler.function.(func (Bot, *types.Message))(telepher,msg)
           }
             
  }
     }}

 }

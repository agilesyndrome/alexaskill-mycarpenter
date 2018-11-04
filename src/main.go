// main.go
package main

import (
	"github.com/agilesyndrome/go-alexa/alexaskill"
	"github.com/agilesyndrome/alexaskill-mycarpenter/src/actualsize"
)

func init() {
  skill.My.IntentMap["actual_size"] = ActualSizeIntent.Handler
  skill.My.Name = "mycarpenter"
}

func main() {
	  skill.OnLambda()
}

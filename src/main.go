// main.go
package main

import (
	"github.com/agilesyndrome/go-alexa/alexaskill"
	"go.agilesyndro.me/alexa-skills/mycarpenter/src/actualsize"
)

func init() {
  skill.My.IntentMap["actual_size"] = ActualSizeIntent.Handler
  skill.My.Name = "mycarpenter"
}

func main() {
	  skill.OnLambda()
}

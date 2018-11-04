# Alexa Skill in Go

Creating an Alexa skill using Go, and AWS Lambda is fast and easy with this very opinionated skill library


Check out the simplicity of this skill's core
```
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
```

package main

import (
	"testing"
	"github.com/agilesyndrome/go-alexa-test/jsontest"
)

const (
  AS_TITLE = "Actual Size"
  AS_2x4 = "The actual size is 1 and a half inches by 3 and a half inches"
  AS_1x12 = "The actual size is three quarter inches by 11 and a quarter inches"
  UNKNOWN_SIZE = "I'm sorry, I don't know that dimension."
  UNKNOWN_ERROR = "I don't have the lookup for 30.000000 and 1.000000 in my lookup table."
  ONE_SIZE_ERROR = "Only one size was provided in the request."
  NEED_TWO_SIZES = "In order to give you the actual size, please try again and provide both measurements, width and the depth"
  SKILL_TITLE = "My Carpenter"
  LAUNCH_WELCOME = "Hello! I'm your carpenter. Try asking me for the actual size of a 1 by 3 piece of lumber"
)

func init() {
   //Things that failed certs need to be pushed up
   //TODO: Find out why this test fails
  // jsontest.Add("launch", jsontest.Expect("en-US", SKILL_TITLE, LAUNCH_WELCOME, false))



  jsontest.Add("actual-size-2x4", jsontest.Expect("en-US", AS_TITLE, AS_2x4, true))
  jsontest.Add("actual-size-1x12", jsontest.Expect("en-US", AS_TITLE, AS_1x12, true))
  //jsontest.Add("actual-size-wrong-sizes", jsontest.Expect("en-US", AS_TITLE, UNKNOWN_SIZE, true))
  //jsontest.Add("actual-size-no-dims", jsontest.Expect("en-US", AS_TITLE, NEED_TWO_SIZES, false))
  jsontest.Add("stop", jsontest.Expect("en-US", "My Carpenter", "Goodbye!", true))

  jsontest.Add("hello", jsontest.Expect("en-US", SKILL_TITLE, "Hello World!", true))
  jsontest.Add("supdog", jsontest.Expect("en-US", "Wassssup", "Wasssssup", true))
  
  //Semi-duplicate cases, different inputs, same output
  //jsontest.Similar("actual-size-2xQM", "actual-size-no-dims")
  //jsontest.Similar("actual-size-4x", "actual-size-no-dims")
  //jsontest.Similar("actual-size-3x", "actual-size-no-dims")


}


func TestRun(t *testing.T) {
  jsontest.RunTests(t)
}



package ActualSizeIntent

import (
  "strconv"
  "github.com/arienmalec/alexa-go"
  "github.com/agilesyndrome/go-alexa-i18n/alexai18n"
)

type Dim struct {
  width float64
  depth float64
}

const (
  IntentName = "actual_size"
)

var (
  ActualSizesImperial = map[Dim]Dim{
    Dim{1,2} : Dim{0.75, 1.5  } ,
    Dim{1,3} : Dim{0.75, 2.5  } ,
    Dim{1,4} : Dim{0.75, 3.5  } ,
    Dim{1,6} : Dim{0.75, 5.5  } ,
    Dim{1,8} : Dim{0.75, 7.25 } ,
    Dim{1,10}: Dim{0.75, 9.25 } ,
    Dim{1,12}: Dim{0.75, 11.25} ,
    Dim{2,2} : Dim{1.5 , 1.5  } ,
    Dim{2,3} : Dim{1.5 , 2.5  } ,
    Dim{2,4} : Dim{1.5 , 3.5  } ,
    Dim{2,6} : Dim{1.5 , 5.5  } ,
    Dim{2,8} : Dim{1.5 , 7.25 } ,
    Dim{2,10}: Dim{1.5 , 9.25 } ,
    Dim{2,12}: Dim{1.5 , 11.25} ,
    Dim{2,14}: Dim{1.5 , 13.25} ,
    Dim{4,4} : Dim{3.5 , 3.5  } ,
    Dim{4,6} : Dim{3.5 , 5.5  } ,
    Dim{6,6} : Dim{5.5 , 5.5  } ,
    Dim{8,8} : Dim{7.25, 7.25 } ,
  }

)

func IntentFail(request alexa.Request) alexa.Response {
   var failTitle = alexai18n.WorldString(request,"sorry")
   var failText = alexai18n.WorldString(request,"sorry_text")
   return alexa.NewSimpleResponse(failTitle, failText)
}

func IntentValid(request alexa.Request) bool {
  if (request.Body.Intent.Name != IntentName) { return false }

  return true
}

func sayFloat(request alexa.Request, lookupValue float64) string {

  var wholeNumber = int64(lookupValue)
  var retVal string = ""

  var wholeStr = strconv.FormatInt(wholeNumber, 10)

  if ( wholeStr != "0" ) {
    retVal += wholeStr
  } 

  var value_str = strconv.FormatFloat(lookupValue - float64(wholeNumber), 'f', 3, 64)

  var decimal_in_words = alexai18n.WorldString(request,"numbers." + value_str)

  if (len(decimal_in_words) > 0) {
    if (len(retVal) > 0) { retVal += " and " }
    retVal += decimal_in_words
  }

  retVal += " " + alexai18n.WorldString(request,"plural_unit")

  return retVal
}


func Handler(request alexa.Request) (alexa.Response, error) {

        //if (!IntentValid(request)) { return IntentFail(request) }

        title := alexai18n.WorldString(request,"actual_size_title")

	var m1_str = request.Body.Intent.Slots["dimA"].Value
	var m2_str = request.Body.Intent.Slots["dimB"].Value

	m1, err_m1 := strconv.ParseFloat(m1_str, 32)
	m2, err_m2 := strconv.ParseFloat(m2_str, 32)

	if (err_m1 !=nil ||  err_m2 != nil) {
	  retText := alexai18n.WorldString(request,"need_two_sizes")
	  res := alexa.NewSimpleResponse(title, retText)
	  res.Body.ShouldEndSession = false
	  return res, nil
	}

	var LookupDim Dim


	if ( m1 < m2 ) {
	   LookupDim = Dim{m1,m2}
	} else {
	   LookupDim = Dim{m2,m1}
	}

        var ActualSizeDim = ActualSizesImperial[LookupDim]

	if (ActualSizeDim == Dim{0,0}) {
	  var retText = alexai18n.WorldString(request,"unknown_size")

	  res := alexa.NewSimpleResponse(title, retText)
	  res.Body.ShouldEndSession = false
	  
	  return res, nil
	}

	//Test for bad lookup

        var width_str = sayFloat(request, ActualSizeDim.width)
	var depth_str = sayFloat(request, ActualSizeDim.depth)
	
	var text = alexai18n.WorldString(request,"actual_size") + " " + width_str + " by " + depth_str


        return alexa.NewSimpleResponse(title, text), nil
}

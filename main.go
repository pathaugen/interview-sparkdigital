
package main

import (
  "fmt"
	//"net/http"
	"strconv"
	//"strings"
	//"bufio"
	"os"
	"os/exec"
	//"time"
	//"log"
)


var debug = false


// ========== START: Golang Console Colors ========== ========== ========== ==========

// Golang Console Colors
// Example: fmt.Print( cRed + "HelloWorld" + cClr )

var cClr				= "\u001b[0m"

var cBold				= "\u001b[1m"

var cBlack			= "\u001b[30m"
var cRed				= "\u001b[31m"
var cGreen			= "\u001b[32m"
var cYellow			= "\u001b[33m"
var cBlue				= "\u001b[34m"
var cMagenta		= "\u001b[35m"
var cCyan				= "\u001b[36m"
var cWhite			= "\u001b[37m"

var cBlackBG		= "\u001b[40m"
var cRedBG			= "\u001b[41m"
var cGreenBG		= "\u001b[42m"
var cYellowBG		= "\u001b[43m"
var cBlueBG			= "\u001b[44m"
var cMagentaBG	= "\u001b[45m"
var cCyanBG			= "\u001b[46m"
var cWhiteBG		= "\u001b[47m"

// ========== END: Golang Console Colors ========== ========== ========== ==========


// Console Splash
var appinfo = `
  ` + cBlue + `====================================================` + cBold + cCyan + `
   _____      _   _    _
  |  __ \    | | | |  | |
  | |__) |_ _| |_| |__| | __ _ _   _  __ _  ___ _ __
  |  ___/ _`+"`"+` | __|  __  |/ _`+"`"+` | | | |/ _`+"`"+` |/ _ \ '_ \
  | |  | (_| | |_| |  | | (_| | |_| | (_| |  __/ | | |
  |_|   \__,_|\__|_|  |_|\__,_|\__,_|\__, |\___|_| |_|
                                     __/ |
                                    |___/` + cClr + `
  ` + cCyan + `Interview: ` + cWhite + `Spark Digital` + cClr + `
  ` + cCyan + `https://github.com/` + cYellow + `pathaugen` + cCyan + `/interview-sparkdigital` + cClr + `
  ` + cBlue + `====================================================` + cClr + `
`


// ========== START: Question 1: maxDifference ========== ========== ========== ==========
// maxDifference function takes array and calculates the single int biggest difference between an array position and lesser positions within the array.
func maxDifference(a []int32) int32 {
  // 7,2,3,10,2,4,8,1
  biggestDifference := -1

  for i := 1; i < len(a); i++ {
    // i = current array position
    // ex: i = 3 value 10
    // Need to get difference between it and lower elements in array
    currentValue := int( a[i] )
    thisDifference := 0

    for loweri := i-1; loweri >= 0; loweri=loweri-1 {
      lowerValue := int( a[loweri] )

      if currentValue > lowerValue {
        thisDifference = currentValue - lowerValue
      }
      if lowerValue > currentValue {
        //thisDifference = currentValue - lowerValue
      }
      if thisDifference > biggestDifference {
        biggestDifference = thisDifference
      }
      // biggestDifference++ // DEBUG
    }

  }
  if biggestDifference != -1 {
    return int32( biggestDifference )
  } else {
    return int32(-1) // Test required a return of -1 however Golang shouldn't be using this return for errors
  }
}
// ========== END: Question 1: maxDifference ========== ========== ========== ==========



// ========== START: Question 2: substringCalculator ========== ========== ========== ==========
// Take a string input and slice characters from either end to see the most unique substring slices possible.
func substringCalculator(s string) int64 {
  debugCount := 0
  if debug && (s == "abcdef" || s == "kincenvizh") { fmt.Print( "DEBUG: " + s + "\n" ) } // DEBUG
  
  substringCount := 0

  // Start with full string "s"
  stringToEnd := s
  stringToEndWhittleDown := stringToEnd
  for i := 0; i < len(s); i++ {
    // Proceed taking off 1 from the start of string at a time and keep looping down until 1 character left
    stringToEnd = s[i:] // Snip the starting letters off starting with 0 letters snipped

    if debug && (s == "abcdef" || s == "kincenvizh") {
      debugCount++
      fmt.Print( cRed + strconv.Itoa(debugCount) + ": " + stringToEnd + cClr + "\n" )
    } // DEBUG

    substringCount++

    stringToEndWhittleDown = stringToEnd
    // Now start stripping letters off the end and work backwards until 1 character remains
    for inneri := len(stringToEndWhittleDown)-1; inneri > 0; inneri=inneri-1 {
      stringToEndWhittleDown = stringToEndWhittleDown[:inneri]
      //if stringToEndWhittleDown != s[0:1] {

      if debug && (s == "abcdef" || s == "kincenvizh") {
        debugCount++
        fmt.Print( strconv.Itoa(debugCount) + ": " + stringToEndWhittleDown + "\n" )
      } // DEBUG

      substringCount++
      //}
    }

  }

  return int64( substringCount )
}
// ========== END: Question 2: substringCalculator ========== ========== ========== ==========



func main() {

  // Clear Screen
  cmd := exec.Command("cmd", "/c", "cls")
  cmd.Stdout = os.Stdout
  cmd.Run()

  // Output Simplification
  breakspace := ("\n")
	breakline := ( breakspace + cBlue + "  ====================================================" + cClr + breakspace )

  fmt.Print( appinfo )
  fmt.Print( breakspace )




  // ========== START: Main Logic ========== ========== ========== ==========
  // https://www.hackerrank.com/tests/info/faq

  // Question 1
  fmt.Print( "  Question 1:" )
  fmt.Print( breakspace )
  fmt.Print( cBlue + "maxDifference function to take array input and calculate the single biggest difference between an array position and previous positions within the array." + cClr )

  fmt.Print( breakspace )

  fmt.Print( breakspace )
  array1 := []int32{7,2,3,10,2,4,8,1} // {6,7,9,5,6,3,2} = 2 // {7,2,3,10,2,4,8,1} = 8
  fmt.Print( cCyan + "  Max Difference: " + cClr + strconv.Itoa( int(maxDifference(array1)) ) + cBlue + " []int32{7,2,3,10,2,4,8,1}" + cClr ) // (" + string(array1) + ")

  fmt.Print( breakspace )
  array2 := []int32{6,7,9,5,6,3,2} // {6,7,9,5,6,3,2} = 2 // {7,2,3,10,2,4,8,1} = 8
  fmt.Print( cCyan + "  Max Difference: " + cClr + strconv.Itoa( int(maxDifference(array2)) ) + cBlue + " []int32{6,7,9,5,6,3,2}" + cClr )

  fmt.Print( breakspace )
  array3 := []int32{5,10,8,7,6,5} // {6,7,9,5,6,3,2} = 2 // {7,2,3,10,2,4,8,1} = 8
  fmt.Print( cCyan + "  Max Difference: " + cClr + strconv.Itoa( int(maxDifference(array3)) ) + cBlue + " []int32{5,10,8,7,6,5}" + cClr )

  fmt.Print( breakspace )
  fmt.Print( breakline )
  fmt.Print( breakspace )

  // Question 2
  fmt.Print( "  Question 2:" )
  fmt.Print( breakspace )
  fmt.Print( cBlue + "Take a string input and slice characters from either end to see the most unique substring slices possible." + cClr )

  fmt.Print( breakspace )

  fmt.Print( breakspace )
  stringTest := "abcdef" // 53 results
  fmt.Print( cCyan + "  Substring Count: " + cClr + strconv.Itoa( int(substringCalculator(stringTest)) ) + cBlue + " " + stringTest + cClr )

  fmt.Print( breakspace )
  string1 := "kincenvizh" // 53 results
  fmt.Print( cCyan + "  Substring Count: " + cClr + strconv.Itoa( int(substringCalculator(string1)) ) + cBlue + " " + string1 + cClr )

  fmt.Print( breakspace )
  string2 := "ghaqjdrmnegmrlrlfpjmnnngpwalzknsencuzwsnhfltwohdgbmvfuwtquosrnyerucntxxkfqehjqygcarxogvcfkljzbzutxphpyykapncjfclnhndzxghelyvzpylazhuutmcquusexzbhsfsmbnlvnlemzvfqbfzwquairhpylnbvyhiyamztlhfchhbwrqddmuzsprfdwuqqchcpeakkexackwwzihkfenwzwckynymgqydvjtovaoezkjjurylqcuonsujycziobnfnmuwnoxcdtahpituykvgpyyshvukrstcbmnsqtjseflwywnslmvnqrtnzkyaddkjamrezprqgoenzsdryygbkeahfiduozpwkrgmatszaxmwodsqiocvagbvxyqotpaujnqvqgjmfxnxhfbwqjpgodlxdrxpjpmzeabpgqrzpxomniknjkdiwtfgyvwvekrnoupwkcbtmpcfamzrghgrznuedkybmfwctdghcfawajlxfkzhdamuygjbcwnyglkjlfmpxfdtovkqbshhrfrnyjrgxgiozsuuncnwofkqzsypwgeikpfbhryhpszegdfajzvqlwwqlnvdtdiuckcvvosrdweohnmawqonjbxyjjhlccuteeshfrxxdhzgakwjqbymnaeudcmibsytyajsgdpfvrutcpglzxdevenevmkgalcrpknuvcrnkuboennhyzirfwvtozzijujsckbxqpocakzrbwgpqgjjmsrtwmvhwyraukbuxfvebeylfpipzwjdzlmgslbtwzataxgqpasrssnfwndldwkdutdqcmcpyanrbdsxrvcvpsywjambtbzlcrvzesuhvyvwwuwwdznigxjxknfajpknqutfvvqynkpvkzgypasevrpxofbymdzcitoqolwqegocuyqsexhumzmckzuuwkamolbltlifongpvkcnrnnuplftqbxpdnegdqlymftqyrxcnzmu" // 499011 results
  fmt.Print( cCyan + "  Substring Count: " + cClr + strconv.Itoa( int(substringCalculator(string2)) ) + cBlue + " " + string2 + cClr )

  // ========== END: Main Logic ========== ========== ========== ==========




  fmt.Print( breakspace )
  fmt.Print( breakline )
}

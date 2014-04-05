package main
import (
	"fmt"
	//"math/rand"
	"strings"
	//"bufio"
	"strconv"
	)
func main() {
	var m = map[string][]string {
		"bllob":[]string{"1","3"},
		"BEEEF":[]string{"DASD","DA"},
	}
	fmt.Println(m["bllob"][1])
}
//YAY! stringIn works!
func stringIn(str string, slice []string) bool{
	for val := range slice {
		if slice[val] == str {
			return true
		}
	}
	return false
}
//keysIn works!!!
func keysIn(m map[string] []string) []string {
	keys := []string{}
	for key := range m {
		keys=append(keys,key)
	}
	return keys
}
func parent(nw string) string{
	lnw := strings.Split(nw,"")
	i:=len(lnw)/2
	nwWord := []string{}
	var focusLetter string
	for i<len(lnw) {
		focusLetter=lnw[i]
		nwWord=append(nwWord,focusLetter)
		i++
	}
	word := strings.Join(nwWord,"")
	return word
}
func child(nw string) string{
	lnw := strings.Split(nw,"")
	i:=0
	nwWord := []string{}
	for i<len(lnw) {
		focusLetter:=lnw[i]
		nwWord=append(nwWord,focusLetter)
		i++
	}
	word := strings.Join(nwWord,"")
	return word
}
func twoChars(str string) string{
	charList := strings.Split(str,"")
	newCharList := []string{charList[0],charList[1]}
	result:=strings.Join(newCharList,"")
	return result
}
func allWordsUnderListOfKeys(dct map [string] []string, keys []string, alreadyDone []string) []string {
	// FYI, the dct will be a map, not a slice. Find the syntax for maps.
	words := []string{}
	for key := range keys {
		if stringIn(keys[key],keysIn(dct)) {
			for val := range dct[keys[key]] {
				if stringIn(dct[keys[key]][val],alreadyDone)==false {
					words=append(words,dct[keys[key]][val])
				}
			}
		}
	}
	return words
}
//I HATE GO!!!!!! WHY DOES THE STRING TO INT CONVERSION NOT WORK!!!!!!!
//func decodeIndNumbers(numberStringList []string) (int,int) {
//	cs := []string{}
//	location := []string{}
//	passedSpace := false
//	for val := range numberStringList {
//		if numberStringList[val] != " " && passedSpace==false {
//			cs=append(cs,numberStringList[val])
//		}
//		if numberStringList[val] == " " {
//			passedSpace=true
//		}
//		if numberStringList[val] != " " && passedSpace {
//			location=append(location,numberStringList[val])
//		}
//	}
//	cs=strconv.Atoi(strings.Join(cs,""))
//	location=strconv.Atoi(strings.Join(location,""))
//	return []int{cs,location}
//}

package main
import (
	"fmt"
	//"math/rand"
	"strings"
	"bufio"
	"strconv"
	"os"
	)
func main() {
	ind:=getIndex()
	fmt.Println(ind[3][1])
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
func decodeIndNumbers(numberStringList []string) (int,int) {
	cs := []string{}
	location := []string{}
	passedSpace := false
	for val := range numberStringList {
		if numberStringList[val] != " " && passedSpace==false {
			cs=append(cs,numberStringList[val])
		}
		if numberStringList[val] == " " {
			passedSpace=true
		}
		if numberStringList[val] != " " && passedSpace {
			location=append(location,numberStringList[val])
		}
	}
	csInt,_ := strconv.Atoi(strings.Join(cs,""))
	locationInt,_ := strconv.Atoi(strings.Join(location,""))
	return csInt, locationInt
}
func getIndex() map[int] []int {
	_=os.Chdir("C:/Users/milla_000/Documents/GitHub/WordGameProj/src/github.com/millanp/wordgame")
	indexfl,err := os.Open("index")
	if err != nil {
    	fmt.Printf("error opening file: %v\n",err)
  	  	os.Exit(1)
	}
	//IS THIS SLOWING DOWN THE PROGRAM? (bufio)
	index := bufio.NewReader(indexfl)
	ln,_,e := index.ReadLine()
	locations := map[int] []int {}
	var vals []string
	var start int
	for e == nil {
		line := string(ln)
		if twoChars(line) == "CS" {
			vals=strings.Split(strings.TrimSpace(line)," ")
			start,_:=strconv.Atoi(vals[2])
			fmt.Println(start)
		} else if twoChars(line)=="EN" {
			vals=strings.Split(strings.TrimSpace(line)," ")
			changeLocation,_:=strconv.Atoi(vals[1])
			endLoc,_:=strconv.Atoi(vals[2])
			locations[changeLocation]=[]int{start,endLoc}
		}
		ln,_,e=index.ReadLine()
	}
	indexfl.Close()
	return locations
}
//func getNetwork(charsInWord int) map[string][]string {
//	index := getIndex()
//	
//}
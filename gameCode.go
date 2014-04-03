package main
import (
	"fmt"
	//"math/rand"
	"strings"
	)
func main() {
	x:=[]string{"erws","fgook","fs"}
	for v:=range x {
		fmt.Println(v)
	}
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
func allWordsUnderListOfKeys(dct string, keys string, alreadyDone string) string {
	//words := []string{}
	//for key := range keys {
		//How on earth do you check if something is in something else in go?
	//}
	return "GOOG"
}
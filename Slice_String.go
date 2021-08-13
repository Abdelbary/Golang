package main

import (
	"golang.org/x/tour/wc"
)

func getWord(s string,i *int) (string,int) {
	ret := ""
	for (*i) < len(s) && s[*i] != ' '{
		ret += string(s[*i])
		(*i)++;
	}
	return ret,*i
}
func WordCount(s string) map[string]int {
	
	m := make(map[string]int,0) 
	word  := ""
	for i:= 0 ; i < len(s) ; i++{
		if s[i] != ' '{
			word,_ = getWord(s,&i)
			_,ok := m[word]
			
			if ok{
				m[word]++
			}else{
				m[word] = 1
			}
		}
	}
	return m
}

func main() {
	wc.Test(WordCount)
}

package main

import (
	"os"
	"fmt"
	"flag"
	"path/filepath"
	"io/ioutil"
)

var dir = flag.String("path",".","The root dir you need to specify, default is .")

func getDirSonPath(paths []string) []string{
	result := []string{}
	for _,path := range paths{
		fmt.Println("path is: ",path)
		sons, err := ioutil.ReadDir(path)
		if err != nil{
			panic(err)
		}
		for _,son := range sons{
			if son.IsDir(){
				// Only work in windows
				result = append(result,fmt.Sprintf("%s\\%s",path,son.Name()))
			}
		}
	}
	
	return result
}

func main(){
	flag.Parse()

	path,err := filepath.Abs(*dir)
	if err != nil{
		panic(err)
	}
	fmt.Println(path)
	
	sons := getDirSonPath([]string{path})
	fmt.Println("sons: ",sons)
	grandSons := getDirSonPath(sons)
	fmt.Println("grandSons: ",grandSons)

	for _,grandson:=range grandSons{
		os.Rename(grandson,fmt.Sprintf("%s-%s",grandson,filepath.Base(path)))
	}
}
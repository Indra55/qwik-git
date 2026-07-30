package main

import (
	"bufio"
	"bytes"
	"compress/zlib"
	"fmt"
	"os"
)


func main(){
	argsWithoutProg:=os.Args[1:]
	if len(argsWithoutProg)<1 {
		fmt.Println("Provide atleast one argument!")
		return
	}
	
	switch command:=argsWithoutProg[0]; command {
		case "init":
			err_objects:=os.MkdirAll(".qwik/objects",0755)
			if err_objects!=nil {
				fmt.Println("Error:",err_objects)
				return
			}

			err_refs_head:=os.MkdirAll(".qwik/refs/heads",0755)
			if err_refs_head!=nil{
				fmt.Println("Error:",err_refs_head)
				return
			}
			f,err_create_ref:=os.Create(".qwik/HEAD")
			if err_create_ref!=nil {
				fmt.Println("Error creating file:",err_create_ref)
				return
			}
			defer f.Close()
			
			w:=bufio.NewWriter(f)
			_,err:=w.WriteString("ref: refs/heads/main\n")
			if err!=nil {
				fmt.Println("Error Writing to buffer", err)
				return
			}

			err = w.Flush()
			if err!=nil {
				fmt.Println("Error Flushing buffer:",err)
				return
			}

			fmt.Println("Initialized Repository!")
		case "hash-objects":
			content, err := os.ReadFile(".qwik/HEAD")
			if err!=nil {
				fmt.Println("Error Reading HEAD",err)
			}
			hash, buf:=hashObjects(content)
			fmt.Println(hash)
			fmt.Println("buf:",buf)
			var compressed_output bytes.Buffer
			zw:=zlib.NewWriter(&compressed_output)
			zw.Write(buf)
			zw.Close()
			fmt.Println("compressed_output:",compressed_output.Bytes())
			err=storeObjects(hash,compressed_output.Bytes())
			if err!=nil {
				fmt.Println("Error:",err)
				return
			}
		case "cat-file":
			data, err := catObject("b870d82622c1a9ca6bcaf5df639680424a1904b0")
			if err != nil {
			    fmt.Println("Error:", err)
			    return
			}
			fmt.Println(len(data))
		default:
			fmt.Println("Unknown command:", command)
	}
}

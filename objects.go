package main

import (
	"bytes"
	"compress/zlib"
	"crypto/sha1"
	"encoding/hex"
	"fmt"
	"os"
	"io"
	"errors"
)

func hashObjects(data []byte) (string, []byte) {
	header:= fmt.Sprintf("blob %d\x00",len(data))
	buf:=append([]byte(header),data...)
	h:= sha1.New()
	h.Write(buf)
	return hex.EncodeToString(h.Sum(nil)),buf
}

func storeObjects(hash string, compressed []byte) error {
	dirName:=hash[0:2]
	fileName:=hash[2:]

	dirPath:=".qwik/objects/"+dirName

	err:=os.MkdirAll(dirPath,0755)
	if err!=nil {
		fmt.Println("Error creating hash object directory",err)
		return err
	}

	filePath := dirPath + "/" + fileName
	err = os.WriteFile(filePath,compressed,0644)
	if err!=nil{
		fmt.Println("Error creating Object File", err)
		return err
	}
	return nil
}

func catObject(hash string) ([]byte,error) {
	dirName:=hash[0:2]
	fileName:=hash[2:]

	filePath:=".qwik/objects/"+dirName+"/"+fileName
	data, err:=os.ReadFile(filePath)
	if err!=nil {
		return nil,err
	}
	
	return data,nil
}

func decompress(data []byte) ([]byte, error){
		newData:= bytes.NewReader(data)
		zr, err:=zlib.NewReader(newData)
		if err!=nil {
			return nil ,err
		}
		defer zr.Close()
		finalData, err:=io.ReadAll(zr)
		if err!=nil{
			return nil, err
		}
		return finalData, nil
		
}

func stripHeader(data []byte) ([]byte,error){
	index:=bytes.IndexByte(data,0)

	if index!=-1{
		return data[index+1:],nil
	}else {
		return nil, errors.New("no null byte found in data")
	}
}
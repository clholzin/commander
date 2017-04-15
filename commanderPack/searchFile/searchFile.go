package searchFile

import (
  "fmt"
  "io/ioutil"
  "os"
  "log"
  //"strings"
  //"path/filepath"
)


func RetrieveContents(name string) []byte{
//filename := filepath.Base(path)
 f,err := os.OpenFile(name,os.O_RDONLY,0)//
 if err != nil {
   fmt.Fprintf(os.Stderr,"%v, Can't open %s: error %\n",os.Args[0],name,err)
   os.Exit(1)
   //panic(err)
 }
 defer f.Close()
 contents,err := ioutil.ReadAll(f)
 if err != nil{
   log.Fatal(err)
 }
 fmt.Printf("%s", contents)
 return contents
}


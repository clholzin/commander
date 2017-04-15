package main
import (
   "fmt"
   "github.com/clholzin/commander/commanderPack/searchFile"
   //"net/http"
   //"io/ioutil"
   //"log"
   //"text/template"
)

type PackageJson struct {
  Title string
  data string
}

func init(){
  fmt.Println("ready")
}

func main(){
  fmt.Println("starting");
  var fileName = "package.json"
  fmt.Printf("%s",searchFile.RetrieveContents(fileName))
  
  /*err := http.ListenAndServe(":8989",nil);
  if err != nil {
     log.Fatal("ListenAndServe: ", err.Error())
  }*/
}
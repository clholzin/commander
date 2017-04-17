package main
import (
   "fmt"
   "github.com/clholzin/commander/commanderPack/searchFile"
   "net/http"
   "encoding/json"
   //"io/ioutil"
   "log"
   //"text/template"
)



const base = "C:/nginx/html"
const fileName = "package.json"


func init(){
  fmt.Println("ready")
}

func main(){
  fmt.Println("starting");
  	
  //fmt.Printf("%s",searchFile.RetrieveContents(fileName))
  http.HandleFunc("/",intro)
  http.HandleFunc("/projects",getProjects)
  
  err := http.ListenAndServe(":8989",nil);
  if err != nil {
     log.Fatal("ListenAndServe: ", err.Error())
  }
}

func intro (w http.ResponseWriter, req *http.Request) {
        // The "/" pattern matches everything, so we need to check
        // that we're at the root here.
        if req.URL.Path != "/" {
                http.NotFound(w, req)
                return
        }
        fmt.Fprintf(w, "Welcome to the home page!")
}

func getProjects (w http.ResponseWriter, req *http.Request) {
		projectsData,err := searchFile.RetrieveDirectories(base)
		w.Header().Set("Content-Type","application/json")
			if err != nil {
			   log.Fatal(err);
				bs, _ := json.Marshal(err)
				w.Write(bs)
			}else{
				
				bs, _ := json.Marshal(projectsData)
				w.Write(bs)
			}
        
}
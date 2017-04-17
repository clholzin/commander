package searchFile

import (
  "fmt"
  "io/ioutil"
  "os"
  "log"
  "sort"
  "strings"
  "path/filepath"
  "encoding/json"
)

type Scripts struct{
	Eslint string `json:"eslint"`
	Stylelint string `json:"stylelint"`
	Lint string `json:"lint"`
	Test string `json:"test"`
	Build string `json:"build"`
	Publish string `json:"publish"`
	Start string `json:"start"`
}
//"build:debug",  //publish:debug,//"test:watch",
type PackageJson struct{
  Name string `json:"name"`
  Scripts Scripts `json:"scripts"`
}

type Project struct{
  Folder string `json:"folder"`
  hasPackageJson bool `json:"hasPackageJson"`
  Packagejson PackageJson `json:"packagejson"`
}



func RetrieveContents(name string) ([]byte, error){
//filename := filepath.Base(path)
 fmt.Printf("retrieve file data %s \n",name)
 f,err := os.OpenFile(name,os.O_RDONLY,0)//
 if err != nil {
   fmt.Fprintf(os.Stderr,"%v, Can't open %s: error %\n",os.Args[0],name,err)
   os.Exit(1)
   return nil,err
   //panic(err)
 }
 defer f.Close()
 contents,err := ioutil.ReadAll(f)
 if err != nil{
   log.Fatal(err)
   return nil,err
 }
 //fmt.Printf("%s", contents)
 return contents,nil
}


func RetrieveDirectories(name string) ([]Project, error){
  var base = "./"
  projects := make([]Project,10,100)
  
  path,err := filepath.Abs(base)
  if err != nil{
   log.Fatal(err)
   return nil,err
 }
  fmt.Printf("path: %s \n\n", path);
 directories,err := readDirNames(name)
  if err != nil{
   log.Fatal(err)	
   return nil,err
 }
 for _,p := range directories {
	//fmt.Printf("filename: %s \n", p);
	children,err := readDirNames(filepath.Join(name,p))
	if err == nil {
		for _,child := range children {
		   if strings.Contains(child,"package") {
		     var packageJson PackageJson
			 packagefilePath := filepath.Join(name,p,"package.json")
			 file,err := RetrieveContents(packagefilePath) //get the byte
			 if err == nil {
			        json.Unmarshal(file, &packageJson)
					fmt.Printf("package Name: %s\n",packageJson.Name)
					data := Project{filepath.Dir(packagefilePath),true,packageJson}
					fmt.Printf("filename added: %s \n",data.Folder)
					projects = append(projects,data)
			 }else{
				 log.Fatal(err)
			 }
			 //fmt.Printf("filename: %s \n --- %s \n", p,child);
		   }
		}
	}else{
		//log.Fatal(err)
	   //return nil,err	
	   log.Printf("%s",err)
	}
	
		
 }
 return projects,nil
}

// readDirNames reads the directory named by dirname and returns
// a sorted list of directory entries.
func readDirNames(dirname string) ([]string, error) {
   f, err := os.Open(dirname)
   if err != nil {
     return nil, err
   }
   names, err := f.Readdirnames(-1)
   f.Close()
   if err != nil {
   	return nil, err
   }
   sort.Strings(names)
   return names, nil
}

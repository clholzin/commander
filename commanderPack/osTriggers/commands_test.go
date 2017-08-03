package osTriggers

import (
   "testing"
   "fmt"
)



func TestStartCmdProcess(t *testing.T) {
	  
		var dir string = "C:/nginx/html/Grover-Strike"
		var script string = "grunt build"
		stOut,err := StartCmdProcess(dir,script);
		if err != nil{
			t.Error(err)
		}
		fmt.Println(stOut)
}
package osTriggers


import (
  	//"bytes"
	"fmt"
	"log"
	"path/filepath"
	"os/exec"
	"os"
	"strings"
	//"time"
	"io"
)



func StartCmdProcess (dir, script string) (io.Reader,error) {
	stOutReader,err := formatProcess(dir,script)
   if err != nil{
	 //log.Fatal("Process failed %v",err)
	 fmt.Printf("Error %v",err)
	 return stOutReader,err
   }
      fmt.Println("Success")
	  return stOutReader,nil
}

func formatProcess (dir,command string) (io.Reader,error){
    fmt.Println("okay")
	outr, outw, err := os.Pipe()
    commandArray := strings.Split(command, " ")
	commandArg := commandArray[:1]
	arguments := commandArray[1:len(commandArray)]
	fmt.Printf("arguments  %q\n", arguments )
	fmt.Printf("command  %s\n", command )

    abs,err := filepath.Abs("./")
	if err != nil{
	  log.Fatal(err)
	  return outr,err
	}
	
    fmt.Printf("abs dir %s \n",abs)
    fmt.Printf("Changing dir to: %s\n",dir) 
	if os.Chdir(dir) != nil{
	  log.Fatal(err)
	  return outr,err
	}
	nabs,err := filepath.Abs("./")
	if err != nil{
	  log.Fatal(err)
	  return outr,err
	}

    fmt.Printf("now abs dir %s \n",nabs)
	cmd := exec.Command(strings.Join(commandArg, ""), strings.Join(arguments, " "))
	cmd.Stdin = outr//strings.NewReader("Starting UP....")
	//var outBuffer bytes.Buffer
	cmd.Stdout = outw
	cmd.Stderr = outw
	/*go func() {
        time.Sleep(time.Second * 10)
        process := cmd.Process
		fmt.Printf("now killing process %d \n",process.Pid)
		//err := process.Kill()
		if err != nil{
		  log.Println(err)
		}
    }()*/

	if cmd.Start() != nil {
		log.Fatal(err)
		return cmd.Stdin,err
	}
	// go back to base
	if os.Chdir(abs) != nil{
	  return outr,err
	}
	return outr,nil
}
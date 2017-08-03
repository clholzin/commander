package main

import (
	"fmt"
	"os"
	"io/ioutil"
	"strings"
	"net/http"
	"encoding/json"
	"path/filepath"
	"log"
	"strconv"
	//"math/rand"
	//"text/template"
	"github.com/gorilla/websocket"
	"github.com/clholzin/commander/commanderPack/searchFile"
	"github.com/clholzin/commander/commanderPack/osTriggers"
	"time"
	//"bytes"
	"io"
	"bufio"
)

const (
	BASECONST = "C:/nginx/html"
	FILENAME  = "package.json"
	APPCONST  = "/home/dist/index.html"

	// Time allowed to write a message to the peer.
	writeWait = 10 * time.Second

	// Maximum message size allowed from peer.
	maxMessageSize = 8192

	// Time allowed to read the next pong message from the peer.
	pongWait = 60 * time.Second

	// Send pings to peer with this period. Must be less than pongWait.
	pingPeriod = (pongWait * 9) / 10

	// Time to wait before force close on connection.
	closeGracePeriod = 10 * time.Second

)
type ProjIndex struct {
	Index   string `json:"index"`
	Command string `json:"command"`
}

type CommandInfo struct {
	Id    int `json:"id"`
	outReader io.Reader  `json:"-"`
}

type CommandId struct {
	Id    int `json:"id"`
}

var (
	upgrader = websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
	}
	ProjectSlice  = make([]searchFile.Project,0,100)
	CommandMemoryItems = make(map[int]CommandInfo)
)



func init() {
	fmt.Println("ready")
	filepath.Abs(BASECONST)

}

func main() {

	/*GET*/  http.HandleFunc("/", intro)
	/*GET*/  http.HandleFunc("/dist", fileServe)
	/*GET*/  http.HandleFunc("/api/v1/projects", getProjects)
	/*POST*/ http.HandleFunc("/api/v1/process", commandProcess)
	/*WS*/   http.HandleFunc("/api/ws/status", getProcessStatus)
	fmt.Println("starting http:localhost:8989")
	err := http.ListenAndServe(":8989", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err.Error())
	}

}

func intro(w http.ResponseWriter, req *http.Request) {
	// The "/" pattern matches everything, so we need to check
	// that we're at the root here.
	if req.URL.Path != "/" {
		http.NotFound(w, req)
		return
	}
	indexFile := filepath.Join(BASECONST, APPCONST)
	fmt.Println(indexFile)
	f, err := os.OpenFile(indexFile, os.O_RDONLY, 0)
	contents, err := ioutil.ReadAll(f)
	if err != nil {
		fmt.Fprintf(w, "Welcome to the home page!") //"Welcome to the home page!"
		log.Fatal("indexFile: ", err.Error())
	} else {
		w.Write(contents) //fmt.Fprintf(w,string(contents[:]))
	}
}

func fileServe(w http.ResponseWriter, req *http.Request) {
	// The "/" pattern matches everything, so we need to check
	// that we're at the root here.
	if req.URL.Path != "/" {
		http.Error(w, "Not found", 404)
		return
	}
	if req.Method != "GET" {
		http.Error(w, "Method not allowed", 405)
		return
	}
	//fmt.Println(req)
	if strings.Contains(req.URL.String(), "dist") {
		file := filepath.Join(BASECONST, req.URL.String())

		//f,err := os.OpenFile(file,os.O_RDONLY,0)
		//contents,err := ioutil.ReadAll(f)
		//if err != nil{
		//	 log.Fatal("file: ", err.Error())
		//} else{
		http.ServeFile(w, req, file)
		//w.Write(contents)//fmt.Fprintf(w,string(contents[:]))
		//}
	} else {
		http.NotFound(w, req)
	}

}

func commandProcess(w http.ResponseWriter, req *http.Request) {
	if req.Method != "POST" {
		http.NotFound(w, req)
		return
	}
	var IDStruct ProjIndex
	decoder := json.NewDecoder(req.Body)
    if decoder.Decode(&IDStruct) != nil {
		fmt.Fprintf(w, "Error unmarshal body")
	} else {
		fmt.Printf("%s : %s",IDStruct.Command,IDStruct.Index)
		projIndex,err := strconv.Atoi(IDStruct.Index)//, err := strconv.Atoi(IDStruct.Index)
		if err != nil {
			panic(err)
			defer recover()
			return
		}
		proJ := ProjectSlice[projIndex]
		outReader, err := osTriggers.StartCmdProcess(proJ.Folder, IDStruct.Command)

		if err != nil {
			log.Fatal("outReader: ", err.Error())
		} else {
			//r := rand.New(rand.NewSource(time.Now().UnixNano()))
			var randInt int = time.Now().Nanosecond()//rand.Int()
			outJson := CommandInfo{randInt, outReader}
			CommandMemoryItems[randInt] = outJson
			output, _ := json.Marshal(outJson)
			w.Header().Set("Content-Type", "application/json")
			w.Write(output)
		}
	}
}

func getProjects(w http.ResponseWriter, req *http.Request) {
	var err error
	if len(ProjectSlice) == 0{
		err = searchFile.RetrieveDirectories(&ProjectSlice,BASECONST)//projectsData,
		if err != nil {
			//defer recover()
			log.Fatal("projectsData: ", err.Error())
			http.Error(w, err.Error(), 301)
			//panic(err)
		}
	}

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	if err != nil {
		log.Fatal(err)
		http.Error(w, err.Error(), 301)
	} else {
		bs, _ := json.Marshal(ProjectSlice)
		w.Write(bs)
	}

}

func getProcessStatus(w http.ResponseWriter, req *http.Request) {
	if req.Header.Get("Origin") != "http://"+req.Host {
		http.Error(w, "Origin not allowed", 403)
		return
	}
	log.Println("websocket begun \n")
	 ws, err := upgrader.Upgrade(w, req, nil)
	if err != nil {
		log.Println(err)
		return
	}
	var ComInfo CommandId

	//req.ParseForm()
	//fmt.Println(req.Form)
	//decoder := json.NewDecoder(req.Body)decoder.Decode(&ComInfo)
	if err = ws.ReadJSON(&ComInfo); err != nil {
		fmt.Fprintf(w, "CommandMemoryItems missing")
	}else {
		fmt.Println(ComInfo.Id)
		fmt.Println(CommandMemoryItems)
		if _,isPresent := CommandMemoryItems[ComInfo.Id]; isPresent{
			stdoutDone := make(chan struct{})
			pumpStdout(ws, CommandMemoryItems[ComInfo.Id].outReader, stdoutDone)
			go ping(ws, stdoutDone)
			select {
				case <-stdoutDone:
				case <-time.After(time.Second):
					<-stdoutDone
			}
		}else{
			fmt.Printf("Failed to locate reader %b",isPresent)
		}
	}
}

func pumpStdout(ws *websocket.Conn, r io.Reader, done chan struct{}) {
	fmt.Println("hit pumpStdout")
	ws.SetReadLimit(maxMessageSize)
	ws.SetReadDeadline(time.Now().Add(pongWait))
	ws.SetPongHandler(func(string) error { ws.SetReadDeadline(time.Now().Add(pongWait)); return nil })
	defer func() {
	}()
	s := bufio.NewScanner(r)
	for s.Scan() {
		ws.SetWriteDeadline(time.Now().Add(writeWait))
		bytes := s.Bytes()
		//fmt.Println(bytes) see the bytes
		if err := ws.WriteMessage(websocket.TextMessage, bytes); err != nil {
			ws.Close()
			break
		}
	}
	if s.Err() != nil {
		log.Println("scan:", s.Err())
	}
	close(done)

	ws.SetWriteDeadline(time.Now().Add(writeWait))
	ws.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""))
	time.Sleep(closeGracePeriod)
	ws.Close()
}

func ping(ws *websocket.Conn, done chan struct{}) {
	ticker := time.NewTicker(pingPeriod)
	defer ticker.Stop()
	for {
		select {
		case <-ticker.C:
			if err := ws.WriteControl(websocket.PingMessage, []byte{}, time.Now().Add(writeWait)); err != nil {
				log.Println("ping:", err)
			}
		case <-done:
			return
		}
	}
}

func internalError(ws *websocket.Conn, msg string, err error) {
	log.Println(msg, err)
	ws.WriteMessage(websocket.TextMessage, []byte("Internal server error."))
}
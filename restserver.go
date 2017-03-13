package main

import (
    "encoding/json"
    "fmt"
    "io/ioutil"
    "log"
    "net/http"
    "regexp"
    "strings"
    //"gopkg.in/ini.v1"
)

const _CONF_DATA = ""

type KvFP struct {
    Key       string `json:"key"`
    Val       string `json:"val"`
    File      string `json:"file"`
    File_type string `json:"file_type"`
}

type StatusOk struct {
    Status string `json:"Status"`
}

type KvFPs []KvFP

func change_ini(Key string, Val string, File string) {
    log.Println(Key)
    log.Println(Val)
    log.Println(File)
    log.Println(" INI file type ..... changing is ok ")

}
func change_xml(Key string, Val string, File string) {
    log.Println(Key)
    log.Println(Val)
    log.Println(File)
    log.Println(" XML file type ..... changing is ok ")
}

func make_one_Change(Key string, Val string, File string, File_type string) {
    log.Println(Key)
    log.Println(Val)
    log.Println(File)
    log.Println(File_type)
    switch File_type {
    case "ini":
        change_ini(Key, Val, File)
    case "XML":
        change_xml(Key, Val, File)
    }

}

func homePage(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Welcome to the HomePage!")
    log.Println("Endpoint Hit: homePage")
}

func returnArticle(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "returns a specific article")
    log.Println("Endpoint Hit: returnArticle")
}

func returnAllArticles(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "All Articles")
    log.Println("Endpoint Hit: returnAllArticles")
}

func addArticle(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Adds an article to list of articles")
    log.Println("Endpoint Hit: addArticle")
}

func delArticle(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "deletes a specific article")
    log.Println("Endpoint Hit: delArticle")
}

func changeThis(rw http.ResponseWriter, req *http.Request) {
    kvfps := KvFPs{}
    /*
       KvFP{Key: "user", Val: "joe", File: "ini.ini", File_type: "ini"},
       KvFP{Key: "password", Val: "joe56", File: "ini.ini", File_type: "ini"},
       KvFP{Key: "age", Val: "6", File: "ini.ini", File_type: "ini"},
       KvFP{Key: "real", Val: "24", File: "ini.ini", File_type: "ini"},
    */
    log.Println("start ... ")
    decoder := json.NewDecoder(req.Body)
    err := decoder.Decode(&kvfps)
    if err != nil {
        panic("O.M.G This is not a Json Json is cool ...")
    }
    for _, json1_data := range kvfps {

        make_one_Change(json1_data.Key, json1_data.Val, json1_data.File, json1_data.File_type)
    }

    log.Println(kvfps)
    log.Println("hi ... ")

    //////////////////////////////////////////////////////////////////////////////////////////////
    ///
    ///
    /// type Vertex struct {
    ///   label string
    /// }
    ///
    //////////////////////////////////////////////////////////////////////////////////////////////
    dicOK := StatusOk{Status: "ok"}
    json.NewEncoder(rw).Encode(dicOK)
    log.Println("hi the function is done ...  ")

    //////////////////////////////////////////////////////////////////////////////////////////////
    ///
    ///
    /// type Vertex struct {
    ///   label string
    /// }
    ///
    //////////////////////////////////////////////////////////////////////////////////////////////
}

func returnAllinthisServer(rw http.ResponseWriter, req *http.Request) {

    kvfps := KvFPs{}
    /*
       KvFP{Key: "user", Val: "joe", File: "ini.ini", File_type: "ini"},
       KvFP{Key: "password", Val: "joe56", File: "ini.ini", File_type: "ini"},
       KvFP{Key: "age", Val: "6", File: "ini.ini", File_type: "ini"},
       KvFP{Key: "real", Val: "24", File: "ini.ini", File_type: "ini"},
    */

    decoder := json.NewDecoder(req.Body)
    err := decoder.Decode(&kvfps)
    if err != nil {
        panic("O.M.G This is not a Json Json is cool ...")
    }

    file_in_string, err := ioutil.ReadFile("ini.ini")
    if err != nil {
        log.Println(err)
    }
    lines := strings.Split(string(file_in_string), "\n")
    array_size := len(lines)
    m := make(map[string]string)
    //log.Println(array_size)
    for i := array_size; i > 0; i-- {
        r, _ := regexp.Compile(`[\s]*=[\s]*`)
        keyval := r.Split(lines[i-1], 2)
        m[keyval[0]] = keyval[1]
        //log.Println(keyval[0])
    }
    for key, val := range m {
        kvfps = append(kvfps, KvFP{Key: key, Val: val, File: "ini.ini", File_type: "ini"})
        log.Println(key + "-->" + val)
    }
    log.Println("Endpoint Hit: returnAllinthisServer")
    json.NewEncoder(rw).Encode(kvfps)
}

func handleRequests() {
    http.HandleFunc("/", homePage)
    http.HandleFunc("/all", returnAllinthisServer)
    http.HandleFunc("/change_this", changeThis)
    http.HandleFunc("/single", returnArticle)
    http.HandleFunc("/delete", delArticle)
    http.HandleFunc("/add", addArticle)
    log.Fatal(http.ListenAndServe(":1949", nil))
}

func main() {
    handleRequests()
}

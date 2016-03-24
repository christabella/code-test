package main

import (
        "fmt"
        "net/http"
        "encoding/json"
        "log"
)

type Data struct {
    websiteUrl         string
    sessionId          string
    resizeFrom         Dimension
    resizeTo           Dimension
    copyAndPaste       map[string]bool // map[fieldId]true
    formCompletionTime float64 // In seconds, with 1 decimal place
}

func NewData(sid string, wurl string) Data {
    return Data{
        sessionId: sid,
        websiteUrl: wurl,
        copyAndPaste: make(map[string]bool)}
}

type Dimension struct {
    Width  string
    Height string
}

// Hashmap keeping track of all Data structs
var dm = make(map[string]*Data) // map[sessionId]*Data

func main() {
    /* Serve accepts incoming connections on the Listener localhost:8080, creating a
     * new service goroutine for each. The service goroutines read requests and then 
     * call srv.Handler to reply to them (`requestHandler` here).
     */
    http.HandleFunc("/", requestHandler) // add requestHandler to DefaultServeMux
    fmt.Println("Running HTTP server on localhost:8080")
    http.ListenAndServe(":8080", nil) // handler is nil, i.e. DefaultServeMux
}

func requestHandler(rw http.ResponseWriter, req *http.Request) {
    // Enable cross-domain requests
    rw.Header().Set("Access-Control-Allow-Origin", "*")
    rw.Header().Set("Access-Control-Allow-Headers", "Content-Type")
    
    if req.Method != "POST" {
        rw.WriteHeader(http.StatusMethodNotAllowed)
        return
    }

    // Read the body into an interface{} hashmap for json decoding
    decoder := json.NewDecoder(req.Body)
    var reqb map[string]interface{}
    err := decoder.Decode(&reqb)
    if err != nil {
        log.Println(err)
        rw.Header().Set("Content-Type", "application/json; charset=UTF-8")
        rw.WriteHeader(http.StatusBadRequest)        
        return
    }

    /* DEBUGGING: printing all JSON name/value pairs *//*
    for k, v := range reqb {
        log.Println(k, ": ", v)
    } */

    sid := reqb["sessionId"].(string)

    if _, ok := dm[sid]; !ok { // Initialize dm[sid] if it doesn't exist
        d := NewData(sid, reqb["websiteUrl"].(string))
        dm[sid] = &d
        log.Printf("*** Created new Data struct with sessionId #%s ***\n\n", sid)
    }

    switch reqb["eventType"] {
        case "windowResize": 
            wb :=reqb["widthBefore"].(string)
            hb :=reqb["heightBefore"].(string)
            wa :=reqb["widthAfter"].(string)
            ha :=reqb["heightAfter"].(string)

            dm[sid].resizeFrom = Dimension{wb, hb}
            dm[sid].resizeTo = Dimension{wa, ha}

            log.Printf("Data #%s Updated\n\t\t    >>> resizeFrom:%v\n\t\t    >>> resizeTo:%v", 
                sid, dm[sid].resizeFrom, dm[sid].resizeTo)

        case "copyAndPaste": 
            fid := reqb["formId"].(string)
            pasted := reqb["pasted"].(bool)

            dm[sid].copyAndPaste[fid] = pasted

            log.Printf("Data #%s Updated\n\t\t    >>> copyAndPaste:map[%s:%t]", 
                sid, fid, pasted)

        case "timeTaken":
            t := reqb["time"].(float64)

            dm[sid].formCompletionTime = t
            
            log.Printf("****** DATA #%s COMPLETE ******\n%+v\n\n", 
                sid, *dm[sid])
    }


    
}

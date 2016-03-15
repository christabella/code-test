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
    formCompletionTime float64 // Seconds
}

type Dimension struct {
    Width  string
    Height string
}

var dm = make(map[string]*Data) //map[sessionId]*Data

func main() {
        http.HandleFunc("/", requestHandler)
        fmt.Println("Running HTTP server on localhost:8080")
        http.ListenAndServe(":8080", nil)
}



type test_struct struct {
    Test string
}

func requestHandler(rw http.ResponseWriter, req *http.Request) {
    rw.Header().Set("Access-Control-Allow-Origin", "*")
    rw.Header().Set("Access-Control-Allow-Headers", "Content-Type")

    decoder := json.NewDecoder(req.Body)
    var reqm map[string]interface{}
    err := decoder.Decode(&reqm)
    if err != nil {
        log.Println("err!!!",err)
        return
    }

    /* DEBUGGING *//*
    for k, v := range reqm {
        log.Println(k, ": ", v)
    } */

    sid := reqm["sessionId"].(string)

    if _, ok := dm[sid]; !ok { // Initialize dm[sid] if it doesn't exist
        var d Data
        dm[sid] = &d
        dm[sid].websiteUrl = reqm["websiteUrl"].(string)
        dm[sid].sessionId = sid
        dm[sid].copyAndPaste = make(map[string]bool)
        log.Printf("*** Created new Data struct with sessionId #%s ***\n\n", sid)
    }

    switch reqm["eventType"] {
        case "windowResize": 
            wb :=reqm["widthBefore"].(string)
            hb :=reqm["heightBefore"].(string)
            wa :=reqm["widthAfter"].(string)
            ha :=reqm["heightAfter"].(string)
            dm[sid].resizeFrom = Dimension{wb, hb}
            dm[sid].resizeTo = Dimension{wa, ha}
            log.Printf("Data #%s Updated\n\t\t    >>> resizeFrom:%v\n\t\t    >>> resizeTo:%v", 
                sid, dm[sid].resizeFrom, dm[sid].resizeTo)

        case "copyAndPaste": 
            fid := reqm["formId"].(string)
            pasted := reqm["pasted"].(bool)
            dm[sid].copyAndPaste[fid] = pasted
            log.Printf("Data #%s Updated\n\t\t    >>> copyAndPaste:map[%s:%t]", sid, fid, pasted)

        case "timeTaken":
            t := reqm["time"].(float64)
            dm[sid].formCompletionTime = t
            log.Printf("****** DATA #%s COMPLETE ******\n%+v\n\n", sid, *dm[sid])
    }


    
}

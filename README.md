Ravelin Code Test
=================

## Summary
An http server that accepts **POST requests** (JSON) from **muliple clients' websites** in parallel. Each request forms part of a struct (for that particular visitor) that will be printed to the terminal when the struct is fully complete. 

To run the HTTP server, cd into root directory and run `./httpserver`.

To serve a sample client website, cd into `sample_client_website` and run `./clientserver`.
*(To run an additional sample client website, cd into `sample2` and run `./clientserver` from there as well.)*

A Data struct (defined below) is created for each session started from a client website, and the Data struct is completed upon form submission. A randomly-generated session ID is stored in a cookie on the client's browser for each client. **All Data structs will be stored in a hashmap** with the session ID as the key (map[sessionId]*Data `dm`). 

## Frontend (JS)
`Main.js` under `sample_client_website/scripts` sends POST requests with JSON body if:
  - screen resizes
  - copy & paste (for each field)
  - form is submitted

### Example JSON Requests
```
{
  "eventType": "copyAndPaste",
  "sessionId": "http://localhost:3000",
  "pasted": true,
  "formId": "inputCardNumber"
}

{
  "eventType": "timeTaken",
  "websiteUrl": "http://localhost:3000",
  "sessionId": "123123-123123-123123123",
  "time": 72.0, // seconds
}

...

```

## Backend (Go)
1. Accept post requests
2. Map request JSON body to relevant sections of the Data struct
3. Print the struct at trace level for each stage of it's construction
4. Also print the struct at info level when it is complete (i.e. form submit button has been clicked)
5. Use go routines and channel where appropriate

### Go Struct
```
type Data struct {
	websiteUrl         string
	sessionId          string
	resizeFrom         Dimension
	resizeTo           Dimension
	copyAndPaste       map[string]bool // map[fieldId]true
	formCompletionTime float64 // Seconds with one decimal place
}

type Dimension struct {
	Width  string
	Height string
}
```





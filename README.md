# TimeHolder

Time Holder Package is a simple solution to limit functions or methods be called once within a certain time.

## Scenario

Say you have a web UI for resending email.

![](http://farm9.staticflickr.com/8094/8542397700_08827e9267_b.jpg)

You can send tons of emails via the code looks like this in the chrome console:

```
for (var i=0; i<100; i++){
  $(“.resend-email-btn”).click();
}
```
By using TimeHolder, you can avoid such mulitple web requests in the backend side, it makes sure that you can only send 1 email in 1 minutes.


## Installation

```
	go get "github.com/athom/timeholder"
```

## Example

```
package main

import (
	"github.com/athom/timeholder"
	"io"
	"log"
	"net/http"
	"time"
)

// hello world, the web server
func HelloServer(w http.ResponseWriter, req *http.Request) {
	key := req.FormValue("key")
	timeholder.SetDuration(5 * time.Second)
	if timeholder.Hold(key) {
		io.WriteString(w, "Be patient!")
		return
	}

	io.WriteString(w, "hello, world!\n")
}

func main() {
	http.HandleFunc("/hello", HelloServer)
	err := http.ListenAndServe(":5000", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

```


## License

TimeHolder is released under the [WTFPL License](http://www.wtfpl.net/txt/copying).

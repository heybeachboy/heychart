package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"log"
	"html/template"
)

var hub *Hub

func init() {
	hub = newHub()
	go hub.run()
}

type Result struct {
	Code    int
	State   bool
	Message string
	Data    interface{}
}

func (r *Result) SetCode(code int) {
	r.Code = code
}

func (r *Result) SetResult(code int, status bool, message string, data interface{}) {
	r.Code = code
	r.State = status
	r.Message = message
	r.Data = data
}

func TestFunction(request *gin.Context) {

	var resp Result
	resp.SetResult(200, true, "Hello! get data successful", "gogogogo")
	request.JSON(200, resp)

}

func WebSocketServer(r *gin.Context) {
	r.Request.Header.Del("Origin")
	serveWs(hub, r.Writer, r.Request)

}

func Home(r *gin.Context) {
	log.Println(r.Request.URL)
	if r.Request.URL.Path != "/" {
		http.Error(r.Writer, "Not found", http.StatusNotFound)
		return
	}
	if r.Request.Method != "GET" {
		http.Error(r.Writer, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	homeTemplate.Execute(r.Writer, "")
}

var homeTemplate = template.Must(template.New("").Parse(`
<!DOCTYPE html>
<html lang="en">
<head>
    <title>Chat Example</title>
    <script type="text/javascript">
        window.onload = function () {
            var conn;
            var msg = document.getElementById("msg");
            var log = document.getElementById("log");

            function appendLog(item) {
                var doScroll = log.scrollTop > log.scrollHeight - log.clientHeight - 1;
                log.appendChild(item);
                if (doScroll) {
                    log.scrollTop = log.scrollHeight - log.clientHeight;
                }
            }

            document.getElementById("form").onsubmit = function () {
                if (!conn) {
                    return false;
                }
                if (!msg.value) {
                    return false;
                }
                conn.send(msg.value);
                msg.value = "";
                return false;
            };

            if (window["WebSocket"]) {
                conn = new WebSocket("ws://" + document.location.host + "/ws");
                console.log("ws://" + document.location.hostname + "/ws");
                conn.onclose = function (evt) {
                    var item = document.createElement("div");
                    item.innerHTML = "<b>Connection closed.</b>";
                    appendLog(item);
                };
                conn.onmessage = function (evt) {
                    var messages = evt.data.split('\n');
                    for (var i = 0; i < messages.length; i++) {
                        var item = document.createElement("div");
                        item.innerText = messages[i];
                        appendLog(item);
                    }
                };
            } else {
                var item = document.createElement("div");
                item.innerHTML = "<b>Your browser does not support WebSockets.</b>";
                appendLog(item);
            }
        };
    </script>
    <style type="text/css">
        html {
            overflow: hidden;
        }

        body {
            overflow: hidden;
            padding: 0;
            margin: 0;
            width: 100%;
            height: 100%;
            background: gray;
        }

        #log {
            background: white;
            margin: 0;
            padding: 0.5em 0.5em 0.5em 0.5em;
            position: absolute;
            top: 0.5em;
            left: 0.5em;
            right: 0.5em;
            bottom: 3em;
            overflow: auto;
        }

        #form {
            padding: 0 0.5em 0 0.5em;
            margin: 0;
            position: absolute;
            bottom: 1em;
            left: 0px;
            width: 100%;
            overflow: hidden;
        }

    </style>
</head>
<body>
<div id="log"></div>
<form id="form">
    <input type="submit" value="Send" />
    <input type="text" id="msg" size="64"/>
</form>
</body>
</html>
</html>
`))

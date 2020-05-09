

# Start Server
```shell script
~$ go run main.go
[GIN-debug] [WARNING] Creating an Engine instance with the Logger and Recovery middleware already attached.

[GIN-debug] [WARNING] Running in "debug" mode. Switch to "release" mode in production.
 - using env:	export GIN_MODE=release
 - using code:	gin.SetMode(gin.ReleaseMode)

[GIN-debug] GET    /ping                     --> gintest/action.Test (3 handlers)
[GIN-debug] Environment variable PORT is undefined. Using port :8080 by default
[GIN-debug] Listening and serving HTTP on :8080
```

# Client Test
```shell script
~$ go run send.go > /tmp/send.log

~$ grep mis /tmp/send.log
{"body":"2-1","message":"pong","result":"mismatch","traceid":"7-1"}
{"body":"1-9","message":"pong","result":"mismatch","traceid":"3-8"}
{"body":"3-9","message":"pong","result":"mismatch","traceid":"4-8"}

~$ head /tmp/send.log
{"body":"9-0","message":"pong","result":"success","traceid":"9-0"}
{"body":"6-0","message":"pong","result":"success","traceid":"6-0"}
{"body":"1-0","message":"pong","result":"success","traceid":"1-0"}
{"body":"0-0","message":"pong","result":"success","traceid":"0-0"}
{"body":"2-0","message":"pong","result":"success","traceid":"2-0"}
{"body":"8-0","message":"pong","result":"success","traceid":"8-0"}
{"body":"4-0","message":"pong","result":"success","traceid":"4-0"}
{"body":"5-0","message":"pong","result":"success","traceid":"5-0"}
{"body":"6-1","message":"pong","result":"success","traceid":"6-1"}
{"body":"3-0","message":"pong","result":"success","traceid":"3-0"}
```


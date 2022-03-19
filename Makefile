test: 
	go test ./... -v -cover

fuzz:
	go test -fuzz=Fuzz ./... -v -cover
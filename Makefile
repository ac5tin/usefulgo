test: 
	go test -run=Test ./... -v -cover

fuzz:
	go test -fuzz=Fuzz ./... -v -cover

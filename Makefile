mac:
	CGO_ENABLED=0 GOOS=darwin go build -a -installsuffix cgo -o url-mapper .
linux:
	CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o url-mapper .

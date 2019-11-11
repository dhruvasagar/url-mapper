.PHONY: default build run clean deploy

default: build

url-mapper:
	CGO_ENABLED=0 GOOS=darwin go build -a -installsuffix cgo -o url-mapper .

url-mapper-linux:
	CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o url-mapper-linux .

build: url-mapper

run: build
	./url-mapper

clean:
	rm -f url-mapper*

deploy: url-mapper-linux
	ssh -tt blog sudo service url-mapper stop
	scp url-mapper-linux blog:~/dsis.me/url-mapper
	ssh -tt blog sudo service url-mapper start
	rm -f url-mapper-linux

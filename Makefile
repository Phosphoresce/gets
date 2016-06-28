GOC=go build
GOFLAGS=-a -ldflags '-s'
CGOR=CGO_ENABLED=0

all: run

build:
	$(GOC) getobject.go

run:
	go run getobject.go

stat:
	$(CGOR) $(GOC) $(GOFLAGS) getobject.go

fmt:
	gofmt -w .

clean:
	rm getobject

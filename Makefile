#swagger:
#	swagger generate spec -o ./static/swagger.json

build: clean 
	go generate
	go build -o www_svc

clean:
	rm -f www_svc www_svc_linux


linux: clean 
	env GOOS='linux' GOARCH='amd64' go build -o www_svc_linux
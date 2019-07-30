GIT_COMMIT=`git rev-parse --short HEAD`

build:
	CGO_ENABLED=0 go build -a -tags netgo -installsuffix netgo -ldflags "-s -w -X github.com/upengs/natsctl/version.gitCommit=${GIT_COMMIT}" -o natsctl ./cmd/natsctl

clean:
	rm -rf 	natsctl
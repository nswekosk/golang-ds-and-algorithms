export CURRENT_BUILD_PATH=$(pwd)
echo $PATH
rm -rf $HOME/golang
rm -rf $HOME/gopath
mkdir -p $HOME/golang # for GOROOT (contains the Go binary & core packages)
mkdir -p $HOME/gopath # for GOPATH (contains code and external packages)
curl http://storage.googleapis.com/golang/go1.7.3.linux-amd64.tar.gz 2>/dev/null > go1.7.3.linux-amd64.tar.gz
tar -C $HOME/golang -xzf go1.7.3.linux-amd64.tar.gz
export GOROOT=$HOME/golang/go
export GOPATH=$HOME/gopath
export PATH=$PATH:$GOROOT/bin
export PATH=$PATH:$GOPATH/bin
(if [[ "$(go version)" == *"go version go1.7"* ]]; then echo "✓ Go binary installed!"; else echo "Go binary not installed"; exit -1; fi);
go version
echo $PATH
go env
which go
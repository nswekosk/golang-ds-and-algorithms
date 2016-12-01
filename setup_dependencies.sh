export CURRENT_BUILD_PATH=$(pwd)
rm -rf $HOME/golang
rm -rf $HOME/gopath
mkdir -p $HOME/golang # for GOROOT (contains the Go binary & core packages)
mkdir -p $HOME/gopath # for GOPATH (contains code and external packages)
export GOROOT=$HOME/golang/go
export GOPATH=$HOME/gopath
export PATH=$PATH:$GOROOT/bin
export PATH=$PATH:$GOPATH/bin
echo $PATH
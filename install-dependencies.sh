echo $PATH
curl http://storage.googleapis.com/golang/go1.7.3.linux-amd64.tar.gz 2>/dev/null > go1.7.3.linux-amd64.tar.gz
tar -C $HOME/golang -xzf go1.7.3.linux-amd64.tar.gz
(if [[ "$(go version)" == *"go version go1.7"* ]]; then echo "✓ Go binary installed!"; else echo "Go binary not installed"; exit -1; fi);
go version
echo $PATH
go env
which go
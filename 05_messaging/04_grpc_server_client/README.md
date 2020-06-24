# Protobuf

Arch linux `sudo pacman -S protobuf `

Make sure GOBIN is in path (edit in .zshrc):

```
GOPATH="/home/mike/workspace/golang"
GOBIN="$GOPATH/bin"
PATH="$PATH:$GOPATH:$GOBIN"
```

after creating chat.proto file
`protoc --go_out=plugins=grpc:chat chat.proto`
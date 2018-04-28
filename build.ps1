$env:GOOS="linux"
go build -ldflags "-s -w" main.go

upx --brute main

docker build -t imvexxed/gorawgit:latest .

Remove-Item main

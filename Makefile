#build windows
GOOS=windows GOARCH=amd64 CGO_ENABLED=1 CC=x86_64-w64-mingw32-gcc GOPROXY="https://goproxy.cn,direct" go build -o ./drinkimgtime.exe -a -installsuffix cgo -ldflags '-s -w' ./cmd
#build linux
CGO_ENABLED=1 GOOS=linux GOARCH=amd64 GOPROXY="https://goproxy.cn,direct" go build -o "./drinkingtime" -a -installsuffix cgo -ldflags '-s -w' ./cmd
#build mac

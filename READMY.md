#How compile to Windows - MiniGW needed - brew install mingw-w64


GOOS=windows GOARCH=amd64 CGO_ENABLED=1 CC=x86_64-w64-mingw32-gcc CXX=x86_64-w64-mingw32-g++ HOST=x86_64-w64-mingw32 go build -ldflags "-extldflags=-static" -p 4 -v -o bin/app-amd64.exe
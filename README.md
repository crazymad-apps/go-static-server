### windows下跨平台打包

```cmd
SET GOOS=linux
SET GOARCH=amd64
go build -o myapp-linux-amd64
```
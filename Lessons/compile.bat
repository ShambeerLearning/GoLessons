@echo off
SET GOARCH=386
SET CGO_ENABLED=0
SET GOOS=windows

"C:\Program Files (x86)\go\bin\go.exe" build -ldflags="-s -w" -o main.exe
pause
@echo off
cls

go build -o bin/soup.exe -buildvcs=false -ldflags "-X main.Solution=soup"
go build -o bin/spoon.exe -buildvcs=false -ldflags "-X main.Solution=spoon"

set /p "bin=(1 or 2)? "

if %bin% == 1 .\bin\soup.exe
if %bin% == 2 .\bin\spoon.exe
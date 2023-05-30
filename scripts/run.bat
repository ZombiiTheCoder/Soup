@echo off
go build -o bin/Soup.exe -buildvcs=false
cls
"./bin/Soup.exe"
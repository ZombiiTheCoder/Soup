@echo off
go build -o bin/soup.exe -buildvcs=false
cls
"./bin/Soup.exe"
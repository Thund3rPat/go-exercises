# Collection of Go-Programs
This repository is collection of tools implemented in Golang for exercise purposes. 
Many programs were written while learning Go and may not be clean. 
If I think the tools are good enough I may put them in their own repository.

## Finished Tools:
### go-system-info
A minimal system info fetcher
```
$ ./go-system-info
-----go-system-info-----
Hostname: yourhostname
OS: windows
Arch: amd64
GOPATH: yourgopath
Go version: go1.14.1
```

## In Progress
### go-grep
A scaled-down version of grep
```
$ ./go-grep [flags] [pattern] [file]
```
- [x] case-insensitive search with -i
- [x] recursiv search with -r
- [ ] implement Boyer-Moore-Algorithm
- [ ] add Wildcard functionality in filepath like *
- [ ] add possibility to define multiple files
- [ ] create own repo for this tool

## Planned:
- CLI-Translator with GoogleTranslate
- Video-Downloader
- TUI-Bluetooth-Connect-Manager
- minimal Chat-Client/Server
- Virus Signature Checker
- DES
- some Games

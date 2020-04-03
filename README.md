# Collection of Go-Programs
This repository is collection of tools implemented in Golang for exercise purposes. 
Many programs were written while learning Go and may not be clean. 

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
- [ ] add Boyer-Moore-Algorithm
- [ ] add Wildcard functionality in filepath like *
- [ ] add possibility to define multiple files
- [ ] move to own repo

## Planned:
- CLI-Translator with GoogleTranslate
- minimal Musicplayer
- Video-Downloader
- TUI-Bluetooth-Connect-Manager
- minimal Chat-Client/Server
- Virus Signature Checker
- DES
- some Games

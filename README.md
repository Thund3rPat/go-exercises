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
- [ ] implement Boyer-Moore-Algorithm
- [ ] add Wildcard functionality in filepath like *
- [ ] add possibility to define multiple files
- [ ] create own repo for this tool

### go-mp
A minimal Music Player
```
$ ./go-mp [file]
```
- [x] play mp3 files
- [ ] add wav support
- [ ] add flac support
- [ ] Show time or add progressbar
- [ ] possibilty do play multiple files successively
- [ ] create own repo for this tool

## Planned:
- CLI-Translator with GoogleTranslate
- Video-Downloader
- TUI-Bluetooth-Connect-Manager
- minimal Chat-Client/Server
- Virus Signature Checker
- DES
- some Games

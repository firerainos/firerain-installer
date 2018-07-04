# firerain-installer
> firerain installer

# This repository has been moved to https://gitlab.com/firerainos/firerain-installer

[![Go Report Card](https://goreportcard.com/badge/github.com/firerainos/firerain-installer)](https://goreportcard.com/report/github.com/firerainos/firerain-installer)
[![GoDoc](https://godoc.org/github.com/firerainos/firerain-installer?status.svg)](https://godoc.org/github.com/firerainos/firerain-installer)
[![Sourcegraph](https://sourcegraph.com/github.com/firerainos/firerain-installer/-/badge.svg)](https://sourcegraph.com/github.com/firerainos/firerain-installer)

## Dependencies
[therecipe/qt](https://github.com/therecipe/qt.git)

[fsnotify](github.com/fsnotify/fsnotify)

## Install therecipe/qt

[therecipe/qt install](https://github.com/therecipe/qt/wiki/Installation)

add $GOPATH/go/bin to environment variables


## Build

> bash
```
go get -t github.com/firerainos/firerain-installer

cd $(go env GOPATH)/src/github.com/firerain-installer

qtrcc desktop ./resources/

qtdeploy build desktop 
```

> fish
```
go get -t github.com/firerainos/firerain-installer

cd (go env GOPATH)/src/github.com/firerain-installer

qtrcc desktop ./resources/

qtdeploy build desktop 

```

## License
This package is released under [GPLv3](LICENSE).

# 負荷試験ツール

## Description
- apiに対して、負荷試験を実施します。シナリオは以下の通りです

## Usage


```
$ load-test-api -help

Usage
  -duration duration
        duration to load (default 10s)
  -output string
        stdout or json or text (default "stdout")
  -parallel uint
        concurrency thread size (default 1)
  -rate uint
        Requests per second per thread (default 1)
  -scenario string
        all or channels or launch (default "all")
  -url string
        url prefix (default "http://localhost:8000")
  -worker uint
        worker (default 2)


```

## Build

```
$ go get github.com/kettsun0123/loadtest
$ cd $GOPATH/src/github.com/kettsun0123/loadtest/

build binary
$ make build
```

## Author

[kettsun0123](https://github.com/kettsun0123)

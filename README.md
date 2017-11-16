# 負荷試験ツール

## Description

## Usage


```
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
$ go get github.com/superchoice/loadtest
$ cd $GOPATH/src/github.com/superchoice/loadtest/

build binary
$ make build
```

## Note
- いくつかGolangかつOSSな負荷試験ツールについて調べました
    + Vegeta
        - Golang製
        - ライブラリとしての利用も想定されているため，カスタマイズ性が高い
        - ドラゴンボール好きにはたまらない
    + Gotling
        - Golang製
        - 負荷テストのメトリクスはyamlで定義する
        - 最終メンテナンスは2017/04/15のreadmeのアップデート
    + k6
        - Golang + JavaScript
        - スクリプトはJSで記述する

## Author

[kettsun0123](https://github.com/kettsun0123)

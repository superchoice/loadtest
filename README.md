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
```

## Note
- いくつかGolangかつOSSな負荷試験ツールについて調べました
    + Vegeta
        - Golang製
        - ライブラリとしての利用も想定されているため，カスタマイズ性が高い
        - abemaのcatalog-apiで採用されている実績あり
            + abemaではVegetaの接続周りでうまく負荷がかからない事象があったらしく，部分的にwrkを使用している
        - ドラゴンボール好きにはたまらない
    + Gotling
        - Golang製
        - 負荷テストはyamlで定義する
        - 最終メンテナンスは2017/04/15のreadmeのアップデート
        - カスタマイズ性が低いため，見送り
    + k6
        - Golang + JavaScript
        - スクリプトはJSで記述する
    + Goad
        - Golang製
        - AWS Lambdaを利用する

## Author

[kettsun0123](https://github.com/kettsun0123)

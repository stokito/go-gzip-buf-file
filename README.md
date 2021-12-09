# go-gzip-buf-file
gzip compressed log/file

## Install

    go get -u github.com/stokito/go-gzip-buf-file@v1.0.0

## Usage
```go
package main

import (
	"github.com/stokito/go-gzip-buf-file"
	"log"
)

func main() {
    requestResponseDump, err := gzip_log.NewGzBufFile("request_log.ndjson.gz")
    if err != nil {
        log.Printf("init log failed %s\n", err.Error())
    }
    requestResponseDump.WriteTwoLines(requestBody, responseData)
    if requestResponseDump != nil {
        requestResponseDump.Close()
    }
}

```

## License
[0BSD](https://opensource.org/licenses/0BSD) (similar to Public Domain)

package gzip_log

import (
	"log"
	"os"
	"testing"
)

func TestGzBufFile_WriteTwoLinesParts(t *testing.T) {
	temp, err := os.CreateTemp("", "request_log.ndjson.gz")
	if err != nil {
		return
	}
	gzlog, err := NewGzBufFile(temp.Name())
	if err != nil {
		log.Printf("init log failed %s\n", err.Error())
	}
	line1_1 := []byte("line1_1")
	line1_2 := []byte("line1_2")
	line1_3 := []byte("line1_3")
	line2_1 := []byte("line2_1")
	line2_2 := []byte("line2_2")
	line2_3 := []byte("line2_3")
	gzlog.WriteTwoLinesParts(line1_1, line1_2, line1_3, line2_1, line2_2, line2_3)
	gzlog.Flush()

	decoded := gzlog.ReadAll()
	if string(decoded) != "\nline1_1line1_2line1_3\nline2_1line2_2line2_3" {
		t.Fail()
	}
}

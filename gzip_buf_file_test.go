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
	line1_4 := []byte("line1_4")
	line1_5 := []byte("line1_5")
	line1_6 := []byte("line1_6")
	line2_1 := []byte("line2_1")
	line2_2 := []byte("line2_2")
	line2_3 := []byte("line2_3")
	line2_4 := []byte("line2_4")
	line2_5 := []byte("line2_5")
	line2_6 := []byte("line2_6")
	gzlog.WriteTwoLinesParts(line1_1, line1_2, line1_3, line1_4, line1_5, line1_6, line2_1, line2_2, line2_3, line2_4, line2_5, line2_6)
	gzlog.Flush()

	decoded := gzlog.ReadAll()
	if string(decoded) != "\nline1_1line1_2line1_3line1_4line1_5line1_6\nline2_1line2_2line2_3line2_4line2_5line2_6" {
		t.Fail()
	}
}

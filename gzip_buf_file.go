// SPDX-License-Identifier: 0BSD

package gzip_log

import (
	"bufio"
	"compress/gzip"
	"os"
	"sync"
)

type GzBufFile struct {
	outputFile *os.File
	gzipWriter *gzip.Writer
	mu         *sync.Mutex
	bufWriter  *bufio.Writer
}

func NewGzBufFile(fileName string) (*GzBufFile, error) {
	outputFile, err := os.OpenFile(fileName, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0644)
	if err != nil {
		return nil, err
	}
	// BestCompression because CPU is a not a limit but IO interruptions are
	gzipWriter, _ := gzip.NewWriterLevel(outputFile, 8)
	// Gzip anyway have a 32Kb window so smaller buffer useless
	bufWriter := bufio.NewWriterSize(gzipWriter, 32768)
	mu := &sync.Mutex{}
	return &GzBufFile{outputFile, gzipWriter, mu, bufWriter}, nil
}

func (f *GzBufFile) WriteString(s string) (int, error) {
	f.mu.Lock()
	n, err := f.bufWriter.WriteString(s)
	f.mu.Unlock()
	return n, err
}

func (f *GzBufFile) Write(data []byte) (n int, err error) {
	f.mu.Lock()
	n, err = f.bufWriter.Write(data)
	f.mu.Unlock()
	return n, err
}

func (f *GzBufFile) WriteTwoLines(line1 []byte, line2 []byte) error {
	f.mu.Lock()
	err := (f.bufWriter).WriteByte('\n')
	if err == nil {
		_, err = (f.bufWriter).Write(line1)
	}
	if err == nil {
		err = (f.bufWriter).WriteByte('\n')
	}
	if err == nil {
		_, err = (f.bufWriter).Write(line2)
	}
	f.mu.Unlock()
	return err
}

func (f *GzBufFile) Close() error {
	_ = f.bufWriter.Flush()
	_ = f.gzipWriter.Flush()
	// Close the gzip first.
	_ = f.gzipWriter.Close()
	return f.outputFile.Close()
}

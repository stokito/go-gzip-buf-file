// SPDX-License-Identifier: 0BSD

package gzip_log

import (
	"bufio"
	"bytes"
	"compress/gzip"
	"io"
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
	if data != nil {
		n, err = f.bufWriter.Write(data)
	}
	f.mu.Unlock()
	return n, err
}

// WriteTwoLines writes two byte slices with new line separator at once
// Example:
//
//	WriteTwoLines(requestBody, responseData)
func (f *GzBufFile) WriteTwoLines(line1, line2 []byte) error {
	f.mu.Lock()
	err := (f.bufWriter).WriteByte('\n')
	if err == nil && line1 != nil {
		_, err = (f.bufWriter).Write(line1)
	}
	if err == nil {
		err = (f.bufWriter).WriteByte('\n')
	}
	if err == nil && line2 != nil {
		_, err = (f.bufWriter).Write(line2)
	}
	f.mu.Unlock()
	return err
}

// WriteTwoLinesParts write two lines and each line can be in few parts to avoid concatenation.
// Each part is a byte slice that can be nil. It's similar to WriteTo but doesn't need allocate a net.Buffers
// Example:
//
//	WriteTwoLinesParts(`{"body":"`, reqBodyJsonEncoded, `"}`, nil, nil, nil, response, nil, nil, nil, nil, nil)
func (f *GzBufFile) WriteTwoLinesParts(line1_1, line1_2, line1_3, line1_4, line1_5, line1_6, line2_1, line2_2, line2_3, line2_4, line2_5, line2_6 []byte) error {
	f.mu.Lock()
	err := (f.bufWriter).WriteByte('\n')
	if err == nil && line1_1 != nil {
		_, err = (f.bufWriter).Write(line1_1)
	}
	if err == nil && line1_2 != nil {
		_, err = (f.bufWriter).Write(line1_2)
	}
	if err == nil && line1_3 != nil {
		_, err = (f.bufWriter).Write(line1_3)
	}
	if err == nil && line1_4 != nil {
		_, err = (f.bufWriter).Write(line1_4)
	}
	if err == nil && line1_5 != nil {
		_, err = (f.bufWriter).Write(line1_5)
	}
	if err == nil && line1_6 != nil {
		_, err = (f.bufWriter).Write(line1_6)
	}
	if err == nil {
		err = (f.bufWriter).WriteByte('\n')
	}
	if err == nil && line2_1 != nil {
		_, err = (f.bufWriter).Write(line2_1)
	}
	if err == nil && line2_2 != nil {
		_, err = (f.bufWriter).Write(line2_2)
	}
	if err == nil && line2_3 != nil {
		_, err = (f.bufWriter).Write(line2_3)
	}
	if err == nil && line2_4 != nil {
		_, err = (f.bufWriter).Write(line2_4)
	}
	if err == nil && line2_5 != nil {
		_, err = (f.bufWriter).Write(line2_5)
	}
	if err == nil && line2_6 != nil {
		_, err = (f.bufWriter).Write(line2_6)
	}
	f.mu.Unlock()
	return err
}

func (f *GzBufFile) Close() error {
	_ = f.Flush()
	// Close the gzip first.
	_ = f.gzipWriter.Close()
	return f.outputFile.Close()
}

func (f *GzBufFile) Flush() error {
	_ = f.bufWriter.Flush()
	return f.gzipWriter.Flush()
}

// ReadAll Read contents of the log file
// Used for tests
func (f *GzBufFile) ReadAll() []byte {
	fileBytes, _ := os.ReadFile(f.outputFile.Name())
	gzreader, _ := gzip.NewReader(bytes.NewReader(fileBytes))
	decoded, _ := io.ReadAll(gzreader)
	return decoded
}

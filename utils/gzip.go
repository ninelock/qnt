package utils

import (
	"bytes"
	"compress/gzip"
	"io"
)

func GZipCompress(data []byte) []byte {
	var buf bytes.Buffer
	w := gzip.NewWriter(&buf)
	_, _ = w.Write(data)
	_ = w.Close()
	return buf.Bytes()
}

func GZipDecompress(data []byte) ([]byte, error) {
	buf := bytes.NewBuffer(data)
	r, err := gzip.NewReader(buf)
	if err != nil {
		return nil, err
	}
	defer func() {
		_ = r.Close()
	}()
	return io.ReadAll(r)
}

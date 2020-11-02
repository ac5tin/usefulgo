package utils

import (
	"bytes"
	"compress/gzip"
	"io/ioutil"

	"github.com/ulikunitz/xz"
)

// Cmpr - compression
type Cmpr struct{}

// NewCmpr - create new compression
func NewCmpr() Cmpr {
	return Cmpr{}
}

// Gzip - gzip compress binary
func (cmp Cmpr) Gzip(data *[]byte) ([]byte, error) {
	var b bytes.Buffer
	gz, err := gzip.NewWriterLevel(&b, gzip.BestCompression)
	if err != nil {
		return nil, err
	}
	if _, err := gz.Write(*data); err != nil {
		return nil, err
	}
	if err := gz.Close(); err != nil {
		return nil, err
	}
	return b.Bytes(), nil
}

// UnGzip - gzip decompress
func (cmp Cmpr) UnGzip(input *[]byte) ([]byte, error) {
	gr, err := gzip.NewReader(bytes.NewBuffer(*input))
	if err != nil {
		return nil, err
	}
	data, err := ioutil.ReadAll(gr)
	if err != nil {
		return nil, err
	}
	return data, nil
}

// Xz - xz compress binary
func (cmp Cmpr) Xz(data *[]byte) ([]byte, error) {
	var b bytes.Buffer

	xw, err := xz.NewWriter(&b)
	if err != nil {
		return nil, err
	}
	if _, err := xw.Write(*data); err != nil {
		return nil, err
	}
	if err := xw.Close(); err != nil {
		return nil, err
	}
	return b.Bytes(), nil
}

// UnXz - xz decompress
func (cmp Cmpr) UnXz(input *[]byte) ([]byte, error) {
	xr, err := xz.NewReader(bytes.NewBuffer(*input))
	if err != nil {
		return nil, err
	}
	data, err := ioutil.ReadAll(xr)
	if err != nil {
		return nil, err
	}
	return data, nil
}

// Compress - compress data
func (cmp Cmpr) Compress(data *[]byte, method string) ([]byte, error) {
	switch method {
	case "xz":
		return cmp.Xz(data)
	default:
		return cmp.Gzip(data)
	}
}

// Decompress - decompress data
func (cmp Cmpr) Decompress(input *[]byte, method string) ([]byte, error) {
	switch method {
	case "xz":
		return cmp.UnXz(input)
	default:
		return cmp.UnGzip(input)
	}
}

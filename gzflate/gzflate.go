package gzflate

import (
	"bytes"
	"compress/flate"

	"io"
)

//解压
func GZinflate(str string) string {
	b := bytes.NewReader([]byte(str))
	r := flate.NewReader(b)
	b2 := new(bytes.Buffer)
	_, err := io.Copy(b2, r)
	if err != nil {
		return err.Error()
	}
	defer r.Close()
	byts := b2.Bytes()
	return string(byts)
}

//压缩
func GZdeflate(str string) (string, error) {
	var b bytes.Buffer
	w, err := flate.NewWriter(&b, 9)
	w.Write([]byte(str))
	w.Close()
	return b.String(), err
}

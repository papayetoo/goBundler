package bundlewriter

import (
	"fmt"
	"io/fs"
	"io/ioutil"
)

type BundleWriter struct {
	D string
	F string
}

func (b BundleWriter) Write() {
	ioutil.WriteFile(b.path(), []byte("Hello world"), fs.ModeAppend)
}

func (b BundleWriter) path() string {
	return fmt.Sprintf("%s/%s", b.D, b.F)
}

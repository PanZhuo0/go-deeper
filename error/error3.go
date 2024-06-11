package error

import (
	"io"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/pkg/errors"
)

func ReadFile(path string) ([]byte, error) {
	f, err := os.Open(path)
	if err != nil {
		// Wrap 打包堆栈信息
		return nil, errors.Wrap(err, "open failed")
	}
	defer f.Close()

	buf, err := ioutil.ReadAll(f)
	if err != nil {
		return nil, errors.Wrap(err, "read failed")
	}
	return buf, nil
}

func ReadConfig() ([]byte, error) {
	home := os.Getenv("HOME")
	// call ReadFile
	config, err := ReadFile(filepath.Join(home, ".setting.xml"))
	return config, errors.WithMessage(err, "could not read config")
}

// func main() {
// 	_, err := ReadConfig()
// 	if err != nil {
// 		fmt.Printf("original error:%T %+v\n", errors.Cause(err), errors.Cause(err))
// 		fmt.Printf("stack trace:\n%+v\n", err)
// 		os.Exit(1)
// 	}
// }

func Write(w io.Writer, buf []byte) error {
	_, err := w.Write(buf)
	// wrap error 打包错误给上层
	return errors.Wrap(err, "write failed")
}

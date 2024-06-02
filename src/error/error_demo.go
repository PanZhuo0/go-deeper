package error

import (
	"bufio"
	"io"
)

func CountLines(r io.Reader) (int, error) {
	var (
		br    = bufio.NewReader(r)
		lines int
		err   error
	)

	for {
		_, err := br.ReadString('\n') //读一行
		lines++
		if err != nil {
			break
		}
	}
	// 没有读完就出错了
	if err != io.EOF {
		return 0, err
	}
	return lines, nil
}

// improvement demo
// func CountLines(r io.Reader) (int, error) {
// 	sc := bufio.NewScanner(r)
// 	lines := 0
// 	for sc.Scan() {
// 		lines++
// 	}
// 	return lines, sc.Err()
// }

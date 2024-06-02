package error

import (
	"fmt"
	"io"
)

type Header struct {
	Key, Value string
}

type Status struct {
	Code   int
	Reason string
}

type errWrite struct {
	io.Writer
	err error
}

func (e *errWrite) Write(buf []byte) (int, error) {
	if e.err != nil {
		return 0, e.err
	}
	var n int
	n, e.err = e.Writer.Write(buf)
	return n, e.err
}

func WriteResponse(w io.Writer, st Status, headers []Header, body io.Reader) error {
	ew := &errWrite{Writer: w}
	fmt.Fprintf(ew, "HTTP/1.1 %d %s \r\n", st.Code, st.Reason)

	for _, h := range headers {
		fmt.Fprintf(ew, "%s: %s\r\n", h.Key, h.Value)
	}

	fmt.Fprintf(ew, "\r\n")
	io.Copy(ew, body)

	return ew.err
}

// func WriteResponse(w io.Writer, st Status, headers []Header, body io.Reader) error {
// 	_, err := fmt.Fprintf(w, "HTTP/1.1 %d %s\r\n", st.Code, st.Reason)
// 	if err != nil {
// 		return err
// 	}

// 	// Headers
// 	for _, h := range headers {
// 		_, err := fmt.Fprintf(w, "%s: %s\r\n", h.Key, h.Value)
// 		if err != nil {
// 			return err
// 		}
// 	}

// 	if _, err := fmt.Fprintf(w, "\r\n"); err != nil {
// 		return err
// 	}

// 	// Body
// 	_, err = io.Copy(w, body)
// 	return err
// }

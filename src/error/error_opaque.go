package error

// 断言行为而不是类型或值
// demo
type temporary interface {
	Temporary() bool
}

func IsTemperary(err error) bool {
	te, ok := err.(temporary)
	return ok && te.Temporary()
}

// 理念:断言错误行为而不是类型 :Assert errors for behaviour.not type

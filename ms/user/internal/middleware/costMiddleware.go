package middleware

import (
	"fmt"
	"net/http"
	"time"
)

type CostMiddleware struct {
}

func NewCostMiddleware() *CostMiddleware {
	return &CostMiddleware{}
}

func (m *CostMiddleware) Handle(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// TODO generate middleware implement function, delete after code implementation
		now := time.Now()
		// Passthrough to next handler if need
		next(w, r) //实际执行的后续处理函数
		fmt.Println("-->cost:", time.Since(now))
	}
}

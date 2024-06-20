package main

import (
	"flag"
	"fmt"

	"google.golang.org/grpc/resolver"
)

var name = flag.String("name", "张三", "通过-name 告诉server 你是谁")

/* M4: 自定义
暂时看不懂
*/

const (
	myScheme   = "q1mi"
	myEndPoint = "resolver.liwenzhou.com"
)

var addrs = []string{"127.0.0.1:8080", "127.0.0.1:8972", "127.0.0.1:8973", "127.0.0.1:8974"}

type q1miResolver struct {
	target     resolver.Target
	cc         resolver.ClientConn
	addrsStore map[string][]string
}

func (r *q1miResolver) ResolveNow(o resolver.ResolveNowOptions) {
	addrStrs := r.addrsStore[r.target.Endpoint()]
	addrList := make([]resolver.Address, len(addrStrs))
	for i, s := range addrStrs {
		addrList[i] = resolver.Address{Addr: s}
	}
	r.cc.UpdateState(resolver.State{Addresses: addrList})
}

func (r *q1miResolver) Close() {}

// builder 设计模式
type q1miResolverBuilder struct{}

func (*q1miResolverBuilder) Build(target resolver.Target, cc resolver.ClientConn) {
	r := &q1miResolver{
		target: target,
		cc:     cc,
		addrsStore: map[string][]string{
			myEndPoint: addrs,
		},
	}
	fmt.Println(r)
}

func (*q1miResolverBuilder) Scheme() string { return myScheme }

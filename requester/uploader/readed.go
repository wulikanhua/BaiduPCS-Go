package uploader

import (
	"github.com/iikira/BaiduPCS-Go/requester/rio"
	"sync/atomic"
)

// Readed64 增加获取已读取数据量, 用于统计速度
type Readed64 interface {
	rio.ReaderLen64
	Readed() int64
}

// NewReaded64 实现Readed64接口
func NewReaded64(rl rio.ReaderLen64) Readed64 {
	return &readed64{
		readed: 0,
		rl:     rl,
	}
}

type readed64 struct {
	readed int64
	rl     rio.ReaderLen64
}

func (r64 *readed64) Read(p []byte) (n int, err error) {
	n, err = r64.rl.Read(p)
	atomic.AddInt64(&r64.readed, int64(n))
	return n, err
}

func (r64 *readed64) Len() int64 {
	return r64.rl.Len()
}

func (r64 *readed64) Readed() int64 {
	return atomic.LoadInt64(&r64.readed)
}

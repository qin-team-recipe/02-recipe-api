package utilities

import (
	"fmt"
	"time"
)

type Benchmark struct {
	UnixNano float64
}

/**
 * ベンチマークスタート
 */
func (u *Benchmark) Start() {
	u.UnixNano = float64(time.Now().UnixNano())
}

/**
 * ベンチマーク終わり
 * 作業時間を出力する
 */
func (u *Benchmark) Finish() {
	s := u.UnixNano
	f := float64(time.Now().UnixNano())
	fmt.Printf("\n==========\nresult: %.3fsec\n==========\n", (f-s)/1000000000)
}

package core

import "bufio"

type Buffer struct {
	BufScan *bufio.Scanner
}

func (b *Buffer) Text() string {
	return b.BufScan.Text()
}

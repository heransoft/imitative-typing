package imitative_typing

import (
	"bufio"
	"io"
	"os"
)

func ReadFileLine(name string) (allLine []string) {
	f, err := os.Open(name)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	rd := bufio.NewReader(f)
	for {
		line, err := rd.ReadString('\n')
		if err != nil || io.EOF == err {
			break
		}
		allLine = append(allLine, line)
	}
	return
}

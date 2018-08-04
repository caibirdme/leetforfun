package leet_71

import (
	"strings"
)

type stream struct {
	buf []byte
	pos int
}

func (s *stream) ReadUntil(b byte) []byte {
	for i := s.pos + 1; i < len(s.buf); i++ {
		if s.buf[i] == b {
			t := s.buf[s.pos+1 : i]
			s.pos = i
			return t
		}
	}
	t := s.pos + 1
	s.pos = len(s.buf) - 1
	return s.buf[t:]
}

func (s *stream) Read() byte {
	s.pos++
	return s.buf[s.pos]
}

func (s *stream) EOF() bool {
	return s.pos >= len(s.buf)-1
}

func simplifyPath(path string) string {
	if path == "" {
		return "/"
	}
	if path[0] != '/' {
		path = "/" + path
	}
	token := &stream{
		buf: []byte(path),
		pos: -1,
	}
	route := []string{"/"}
	token.Read()
	for !token.EOF() {
		name := token.ReadUntil(byte('/'))
		if len(name) == 0 {
			continue
		}
		str := string(name)
		if str == ".." {
			if len(route) > 1 {
				route = route[:len(route)-1]
			}
		} else if str == "." {
			continue
		} else {
			route = append(route, str)
		}
	}
	return "/" + strings.Join(route[1:], "/")
}

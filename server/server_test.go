package server

import (
	"testing"
)

type MockReader struct {
	returnValue []byte
	lIter       int
}

func (m *MockReader) Read(p []byte) (n int, err error) {
	if m.lIter < len(m.returnValue) {
		inc := len(p)
		uIter := m.lIter + inc
		for i, v := range m.returnValue[m.lIter:uIter] {
			p[i] = v
		}
		n = uIter - m.lIter
		m.lIter += n
	}
	return
}

func TestReadRequestLinePass(t *testing.T) {
	value := "gemini://localhost:1965\r\n"
	expected := "gemini://localhost:1965"
	reader := MockReader{
		[]byte(value),
		0,
	}

	line, err := readRequestLine(&reader)
	if err != nil {
		t.Errorf("encountered unexpected error during readRequestLine")
	}

	out := string(line)

	if out != expected {
		t.Errorf("expected line = %v, got line = %v", expected, out)
	}

}

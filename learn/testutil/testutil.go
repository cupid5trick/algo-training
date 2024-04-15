package testutil

import (
	"io"
	"strings"
	"testing"
	"github.com/stretchr/testify/assert"
)

type IoFunc func(io.Reader, io.Writer)

func removeExtraSpace(s string) string {
	s = strings.TrimSpace(s)
	sp := strings.Split(s, "\n")
	for i := range sp {
		sp[i] = strings.TrimSpace(sp[i])
	}
	return strings.Join(sp, "\n")
}


// 无尽对拍模式
// inputGenerator 生成随机测试数据，runFuncAC 为暴力逻辑或已 AC 逻辑，runFunc 为当前测试的逻辑
func AssertEqualRunResultsInf(t *testing.T, inputGenerator func() string, runFuncAC, runFunc IoFunc) {
	for tc := 1; ; tc++ {
		input := inputGenerator()
		input = removeExtraSpace(input)

		mockReader := strings.NewReader(input)
		mockWriterAC := &strings.Builder{}
		runFuncAC(mockReader, mockWriterAC)
		expectedOutput := removeExtraSpace(mockWriterAC.String())

		mockReader = strings.NewReader(input)
		mockWriter := &strings.Builder{}
		runFunc(mockReader, mockWriter)
		actualOutput := removeExtraSpace(mockWriter.String())

		assert.Equal(t, expectedOutput, actualOutput, "Wrong Answer %d\tInput:\t%s\nOutput:\t%s\t%s", tc, input, expectedOutput, actualOutput)

		// 每到 2 的幂次就打印检测了多少个测试数据
		if tc&(tc-1) == 0 {
			t.Logf("%d cases checked.", tc)
		}

	}
}
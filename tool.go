package excelize

import (
	"archive/zip"
	"fmt"
	"io/fs"
	"sync"
)

func minInt(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func minFloat(a, b float64) float64 {
	if a < b {
		return a
	}
	return b
}

func maxInt(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func maxFloat(a, b float64) float64 {
	if a > b {
		return a
	}
	return b
}

func clearSyncMap(m *sync.Map) {
	*m = sync.Map{}
}

// 1. 定义一个包装器，嵌入原生的 *zip.Writer
type zipWriterWrapper struct {
	*zip.Writer
}

// 2. 为包装器手动实现缺失的 AddFS 方法（兼容 Go 1.17）
// 注意：如果你的代码逻辑里不需要用到 AddFS，直接返回错误或 nil 即可
func (w *zipWriterWrapper) AddFS(fsys fs.FS) error {
	// 在 Go 1.17 中通常不需要此功能，或者你可以报错
	return fmt.Errorf("AddFS not supported in Go 1.17")
}

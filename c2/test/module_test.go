package test

import (
	"testing"
)

func TestH(t *testing.T) {
	t.Log("Testing")
	// 如果你不提交到github 的话 你的go module是找不到的，这个时候你可以去修改goPath 指定到你当前的目录下, 一般的开发流程是
	// 1. go Path 路径中添加一个到自己的 工作目录，既以后的任何自己的go代码都往这里放
	// 2. 项目完成之后 丢到gitlab / github，这样别人使用的时候直接get 就好了

	// 我一般推荐的做法是，省略第一步直接丢github

}

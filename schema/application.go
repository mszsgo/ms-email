package schema

import (
	"github.com/mszsgo/hconfig"
	"github.com/mszsgo/hgraph"
)

var (
	App = &hconfig.App{Name: "email", Version: "1.0.0", Port: 80}
)

// main方法第一行加载配置信息
func Main() {

	// 启动服务
	App.Start(preFunc)
}

func preFunc(app *hconfig.App) {
	// Graphql
	hgraph.HttpHandle(&Query{}, nil)
}

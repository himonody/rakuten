package config

import (
	"fmt"
	"github.com/lionsoul2014/ip2region/binding/golang/xdb"
	"log"
	"os"
	"path/filepath"
)

var Searcher *xdb.Searcher

type searcher xdb.Searcher

func InitIp() (*xdb.Searcher, error) {
	//xdbPath := filepath.Join("./config", "ip2region.xdb")
	xdbPath := filepath.Join("config", "ip2region.xdb")

	if _, err := os.Stat(xdbPath); os.IsNotExist(err) {
		log.Printf("xdb 文件不存在，请先下载：%s\n", xdbPath)
		return nil, err
	}

	//dbPath := "./config/ip2region.xdb"
	cBuff, err := xdb.LoadContentFromFile(xdbPath)
	if err != nil {
		fmt.Printf("failed to load content from `%s`: %s\n", xdbPath, err)
		return nil, err
	}

	// 2、用全局的 cBuff 创建完全基于内存的查询对象。
	Searcher, err = xdb.NewWithBuffer(cBuff)
	if err != nil {
		fmt.Printf("failed to create searcher with content: %s\n", err)
		return nil, err
	}
	return Searcher, nil
}
func (x searcher) Close() {
	x.Close()
}

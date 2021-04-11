package chapter1

import (
	"bufio"
	"os"
	"strconv"
	"strings"
	"testing"
)

// 测试 quickUnionWeighted
func TestUFQuickUnionWeighted(t *testing.T) {
	// 读取文件
	file, err := os.Open("mediumUF.txt")
	if err != nil {
		t.Fatalf("测试 quickUnionWeighted 打开 mediumUF.txt文件失败, %s", err.Error())
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	i := 0
	uf := &UFQuickFind{}
	for scanner.Scan() {
		i++
		if i == 1 {
			countStr := scanner.Text()
			count, err := strconv.Atoi(countStr)
			if err != nil {
				t.Fatalf("测试 quickUnionWeighted mediumUF.txt文件第一行不是数字, %s", err.Error())
			}
			uf = NewUFQuickFind(count)
		} else {
			list := strings.Split(scanner.Text(), " ")
			if len(list) != 2 {
				t.Fatalf("测试 quickUnionWeighted mediumUF.txt文件里第%d行数据格式有问题", i)
			}
			p, err := strconv.Atoi(strings.TrimSpace(list[0]))
			if err != nil {
				t.Fatalf("测试 quickUnionWeighted mediumUF.txt文件里第%d行数据格式有问题", i)
			}
			q, err := strconv.Atoi(strings.TrimSpace(list[1]))
			if err != nil {
				t.Fatalf("测试 quickUnionWeighted mediumUF.txt文件里第%d行数据格式有问题", i)
			}
			uf.Union(p, q)
		}
	}
	if scanner.Err() != nil {
		t.Fatalf("测试 quickUnionWeighted mediumUF.txt文件出错, %s", scanner.Err().Error())
	}
	if uf.Count() != 3 {
		t.Fatalf("测试 quickUnionWeighted  mediumUF.txt文件结果不是3,结果是: %d", uf.Count())
	} else {
		t.Log("测试 quickUnionWeighted  mediumUF.txt成功")
	}
}

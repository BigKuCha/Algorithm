package sample

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"math/rand"
	"os"
	"time"
)

func Reservoir() {
	const c = 10                                       // 抽样池容量
	var reservoir [c]string                            // 抽样池
	var bigFile = "./sample/data_file_4_reservoir.log" // 表示一个无限大的文件
	f, err := os.Open(bigFile)
	if err != nil {
		log.Fatalln(err)
	}
	defer f.Close()
	br := bufio.NewReader(f)
	i := 0
	for {
		l, _, err := br.ReadLine()
		if err == io.EOF {
			break
		}
		// 小于抽样池 直接放入
		if i < c {
			fmt.Printf("放入第%d个元素\n", i)
			reservoir[i] = string(l)
		} else {
			rand.Seed(time.Now().UnixNano())
			d := rand.Intn(i) // 随机获取 rand(0, i) ,如果d落到了0-c之间，则替换抽样池中的数据
			if d < c {
				fmt.Printf("替换第%d个元素，放入第%d :%s \n", d, i, string(l))
				reservoir[d] = string(l)
			}
		}
		i++
	}
	fmt.Println(reservoir)
}

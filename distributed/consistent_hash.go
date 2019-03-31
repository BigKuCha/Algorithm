package distributed

import (
	"fmt"
	"hash/crc32"
	"sort"
	"strconv"
	"sync"
)

type ring []uint32 // 为了实现排序
// sort.Sort 需要自行实现以下三个方法
func (h ring) Len() int {
	return len(h)
}
func (h ring) Less(i, j int) bool {
	return h[i] < h[j]
}
func (h ring) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

// 一致性哈希
type Hashring struct {
	Servers         map[string]bool // 所有服务器(服务器名称) 用来检查服务器是否已存在
	Nodes           map[uint32]Node // 通过值来获取服务器
	Ring            ring            // 哈希环，用来存放所有节点(hash值)
	VirturalNodeNum int             // 每个服务器对应虚拟节点的个数
	sync.RWMutex                    // 读写锁
}

// 节点
type Node struct {
	ID       uint32
	IP       string
	Port     int
	HostName string
	Weight   int
}

func NewNode(ip string, port int, hostName string) *Node {
	return &Node{
		IP:       ip,
		Port:     port,
		HostName: hostName,
	}

}

func NewHashring(vNum int) *Hashring {
	return &Hashring{
		Servers:         make(map[string]bool),
		Nodes:           make(map[uint32]Node),
		Ring:            []uint32{},
		VirturalNodeNum: vNum,
	}
}

func Hash(key string) uint32 {
	kbyte := []byte(key)
	return crc32.ChecksumIEEE(kbyte)
}

// 添加节点
func (h *Hashring) AddNode(node Node) bool {
	// 如果添加过节点(主机名已存在)则跳过
	if _, ok := h.Servers[node.HostName]; ok {
		return false
	}
	h.Lock()
	defer h.Unlock()
	// 主机名保存，用来防止重复添加
	h.Servers[node.HostName] = true
	// 根据预设的虚拟节点数来添加虚拟节点
	for i := 0; i < h.VirturalNodeNum; i++ {
		hashKey := Hash(node.HostName + "#" + strconv.Itoa(i))
		h.Ring = append(h.Ring, hashKey)
		h.Nodes[hashKey] = node
	}
	// 哈希环排序，排序后方便检索
	sort.Sort(h.Ring)
	return true
}

// 根据用户数据分配一个节点
func (h *Hashring) Get(key string) Node {
	h.RLock()
	defer h.RUnlock()
	hashKey := Hash(key)
	ringLen := len(h.Ring)
	// 查找大于等于hashKey的哈希值，如果未找到，此方法返回切片的长度
	k := sort.Search(ringLen, func(i int) bool {
		return h.Ring[i] >= hashKey
	})
	if k == ringLen {
		k = 0
	}
	fmt.Println(hashKey, h.Ring[k])
	return h.Nodes[h.Ring[k]]
}

// 移除节点
func (h *Hashring) Remove(hostName string) bool {
	if _, ok := h.Servers[hostName]; !ok {
		return false
	}
	h.Lock()
	defer h.Unlock()
	delete(h.Servers, hostName)
	for i := 0; i < h.VirturalNodeNum; i++ {
		hkey := Hash(hostName + "#" + strconv.Itoa(i))
		delete(h.Nodes, hkey)
	}

	tmpRing := h.Ring[:0]
	for _, v := range h.Ring {
		found := false
		for j := 0; j < h.VirturalNodeNum; j++ {
			hkey := Hash(hostName + "#" + strconv.Itoa(j))
			if hkey == v {
				found = true
			}
		}
		if !found {
			tmpRing = append(tmpRing, v)
		}
	}
	h.Ring = tmpRing
	return true
}

// 测试
func test() {
	node := NewNode("127.0.0.1", 80, "localhost")
	node2 := NewNode("127.0.0.2", 80, "localhost2")
	node3 := NewNode("127.0.0.3", 80, "localhost3")
	hashring := NewHashring(80)
	hashring.AddNode(*node)
	hashring.AddNode(*node2)
	hashring.AddNode(*node3)
	k := "819238a"
	fmt.Println(hashring.Ring)
	fmt.Println(hashring.Get(k))
}

package consistent_hash

import (
	"crypto/sha256"
	"encoding/binary"
	"sort"
	"strconv"
)

type Node struct {
	key  string
	hash uint64
}

type HashRing struct {
	nodes    []*Node
	Replicas int
}

func NewHashRing(replicas int) *HashRing {
	return &HashRing{
		Replicas: replicas,
	}
}

func (hr *HashRing) AddNode(key string) {
	for i := 0; i < hr.Replicas; i++ {
		vnode := key + ":" + strconv.Itoa(i)
		hash := hr.hashKey(vnode)
		hr.nodes = append(hr.nodes, &Node{key: key, hash: hash})
	}

	sort.Slice(hr.nodes, func(i, j int) bool {
		return hr.nodes[i].hash < hr.nodes[j].hash
	})

}

func (hr *HashRing) RemoveNode(key string) {
	var newNodes []*Node
	for _, node := range hr.nodes {
		if node.key != key {
			newNodes = append(newNodes, node)
		}
	}
	hr.nodes = newNodes
}

func (hr *HashRing) GetNode(stringKey string) string {
	if len(hr.nodes) == 0 {
		return ""
	}
	hash := hr.hashKey(stringKey)
	idx := sort.Search(len(hr.nodes), func(i int) bool {
		return hr.nodes[i].hash >= hash
	})
	if idx == len(hr.nodes) {
		idx = 0
	}
	return hr.nodes[idx].key
}

func (hr *HashRing) hashKey(key string) uint64 {
	h := sha256.New()
	h.Write([]byte(key))
	hash := h.Sum(nil)
	return binary.BigEndian.Uint64(hash[:8])
}

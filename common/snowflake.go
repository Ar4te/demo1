package common

import (
	"errors"
	"sync"
	"time"
)

const (
	nodeBits           = 10 // 节点ID所占位数
	sequenceBits       = 12 // 序列号所占位数
	nodeMax            = -1 ^ (-1 << nodeBits)
	sequenceMask       = -1 ^ (-1 << sequenceBits)
	timeShift          = nodeBits + sequenceBits
	nodeShift          = sequenceBits
	twepoch      int64 = 1288834974657 // 统一起始时间（毫秒）：2014-01-01 00:00:00.000
)

type Snowflake struct {
	mu        sync.Mutex
	timestamp int64   //时间戳（毫秒）
	nodeId    int64   //节点ID
	sequence  int64   //序列号
	cache     []int64 //循环缓冲区，用于缓存生成的ID
	cacheSize int     //缓存池大小
	cacheIdx  int     //缓存池当前位置
}

func NewSnowflake(nodeId int64, cacheSize int) (*Snowflake, error) {
	//节点ID异常
	if nodeId < 0 || nodeId > nodeMax {
		return nil, errors.New("Invalid Node number")
	}
	return &Snowflake{
		mu:        sync.Mutex{},
		timestamp: 0,
		nodeId:    nodeId,
		sequence:  0,
		cache:     make([]int64, cacheSize),
		cacheSize: cacheSize,
	}, nil
}

func (sf *Snowflake) NextId() (int64, error) {
	sf.mu.Lock()
	defer sf.mu.Unlock()

	now := time.Now().UnixNano() / 1000000

	if now < sf.timestamp {
		return 0, errors.New("Invalid timestamp")
	}

	if now == sf.timestamp {
		sf.sequence = (sf.sequence + 1) & sequenceMask
		if sf.sequence == 0 {
			waitDuration := twepoch + (now-sf.timestamp)*2
			time.Sleep(time.Duration(waitDuration-now) * time.Millisecond)
			now = time.Now().UnixNano() / 1000000
		}
	} else {
		sf.sequence = 0
	}

	sf.timestamp = now
	id := (now-twepoch)<<timeShift | (sf.nodeId << nodeShift) | sf.sequence

	if sf.cache[sf.cacheIdx] != 0 {
		id = sf.cache[sf.cacheIdx]
		sf.cache[sf.cacheIdx] = 0
		sf.cacheIdx = (sf.cacheIdx + 1) % sf.cacheSize
	} else {
		for i := 0; i < sf.cacheSize; i++ {
			sf.cache[i] = (now-twepoch)<<timeShift | (sf.nodeId << nodeShift) | int64((i + 1))
		}
		id = sf.cache[0]
		sf.cacheIdx = 1
	}
	return id, nil
}

func Generate() (uint64, error) {
	sf, err := NewSnowflake(1, 1024)

	if err != nil {
		return 0, err
	}
	id, err := sf.NextId()

	if err != nil {
		return 0, err
	}

	return uint64(id), nil
}

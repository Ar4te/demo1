package common

import (
	"errors"
	"sync"
	"time"
)

type Snowflake struct {
	lastTimestamp int64      //上次生成ID的时间戳
	NodeId        int64      //节点ID
	sequence      int64      //序列号
	mutex         sync.Mutex //互斥锁，用于保证并发安全
}

// 生成唯一ID函数
func (s *Snowflake) NextId() (int64, error) {
	s.mutex.Lock() //加锁
	defer s.mutex.Unlock()

	timestamp := time.Now().UnixNano() / 1000000 //获取当前的时间戳（毫秒级）

	//如果当前的时间戳小于上次生成ID的时间戳，则说明出现了时钟回拨错误
	if timestamp < s.lastTimestamp {
		return 0, errors.New("出现时钟回拨错误")
	}

	//如果当前的时间戳与上次生成ID的时间戳相同，则需要增加序列号
	if timestamp == s.lastTimestamp {
		s.sequence = (s.sequence + 1) & 0xfff //序列号最大为4095（12位二进制数）
		//序列号达到最大值时，等待下一毫秒
		if s.sequence == 0 {
			timestamp = s.waitNextMillis(timestamp)
		}
		//如果当前的时间戳比上次生成ID的时间戳大，则序列号归零
	} else {
		s.sequence = 0
	}

	//更新上次生成ID的时间戳
	s.lastTimestamp = timestamp

	//生成ID，1577808000000是2020年1月1日的时间戳，左移22位是因为41位的时间戳中已经占用了22位
	id := ((timestamp - 1577808000000) << 22) | (s.NodeId << 12) | s.sequence

	return id, nil
}

// 等待下一毫秒函数
func (s *Snowflake) waitNextMillis(currentTimestamp int64) int64 {
	//相当于while(currentTimestamp != s.lastTimestamp)， 如果相等，休眠1ms，再获取毫秒级时间戳
	for currentTimestamp == s.lastTimestamp {
		time.Sleep(time.Millisecond)
		currentTimestamp = time.Now().UnixNano() / 1000000
	}
	return currentTimestamp
}

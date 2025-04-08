/**
 * @Title
 * @Author: liwei
 * @Description:  TODO
 * @File:  orderid
 * @Version: 1.0.0
 * @Date: 2025/04/08 13:46
 * @Update liwei 2025/4/8 13:46
 */

package orderid

import (
	"fmt"
	"sync"
	"time"
)

const (
	machineIDBits = 3
	sequenceBits  = 6

	machineIDMask = (1 << machineIDBits) - 1
	sequenceMask  = (1 << sequenceBits) - 1
)

var (
	machineID      int64
	lastTimestamp  int64 = -1
	sequence       int64 = 0
	startTimestamp int64
	onceInit       sync.Once
	lock           sync.Mutex
	inited         = false
)

// Init 初始化机器ID和起始时间戳（只能调用一次）
func Init(mID int64, startTS int64) {
	onceInit.Do(func() {
		if mID < 0 || mID > machineIDMask {
			panic("machine ID out of range")
		}
		machineID = mID
		startTimestamp = startTS
		inited = true
	})
}

// GenerateOrderID 生成唯一订单号
func GenerateOrderID(bizId string) string {
	if !inited {
		panic("orderid package not initialized, please call orderid.Init() first")
	}

	lock.Lock()
	defer lock.Unlock()

	now := time.Now()
	timeStr := now.Format("20060102150405") // 格式化为 YYYYMMDDHHMMSS（14位）
	// 处理时钟回退（关键：基于时间戳而非格式化字符串）
	nowUnix := now.Unix() // 保留秒级时间戳用于时钟回退判断
	if nowUnix < lastTimestamp {
		for nowUnix <= lastTimestamp {
			now = time.Now()
			nowUnix = now.Unix()
		}
		timeStr = now.Format("20060102150405") // 重新获取最新时间字符串
	}

	if nowUnix == lastTimestamp {
		// 同秒内序列号自增
		sequence = (sequence + 1) & sequenceMask
		if sequence == 0 {
			// 序列号溢出，等待1秒
			for nowUnix <= lastTimestamp {
				now = time.Now()
				nowUnix = now.Unix()
			}
			timeStr = now.Format("20060102150405") // 时间已变化，重新生成时间字符串
		}
	} else {
		// 新的秒级时间，重置序列号
		sequence = 0
	}

	lastTimestamp = nowUnix // 更新最后生成时间（秒级时间戳）
	return fmt.Sprintf("%s%s%03d%06d", bizId, timeStr, machineID, sequence)
}

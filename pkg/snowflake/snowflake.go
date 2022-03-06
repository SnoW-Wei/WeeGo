/*
 * @Descripttion: talk is cheep , show me the code !
 * @version: V1.0
 * @Author: snow.wei
 * @Date: 2022-03-05 23:58:59
 * @LastEditors: snow.wei
 * @LastEditTime: 2022-03-05 23:59:00
 */
// **************************************************** //
// Usage:
//      println(snowflake.NewID())
//  分布式 ID 生成器, 本地SDK, 无网络操作
//       41位毫秒时间戳       10 机器id       12bit 序列号
// --------------------|-------------|-----------------|
//
// 2021-12-28 16:52:49
// Copyright By Guojianyong
// **************************************************** //


package snowflake

import (
   "hash/fnv"
   "os"
   "strconv"
   "strings"
   "sync"
   "time"
)


const (
   workerBits uint8 = 10 // 节点数
   seqBits    uint8 = 12 // 1毫秒内可生成的id序号的二进制位数
   //workerMax     uint64 = -1 ^ (-1 << workerBits) // 节点ID的最大值，用于防止溢出
   seqMax        uint64 = -1 ^ (-1 << seqBits)    // 同上，用来表示生成id序号的最大值
   timeShift     uint8  = workerBits + seqBits    // 时间戳向左的偏移量
   workerShift   uint8  = seqBits                 // 节点ID向左的偏移量
   epoch         uint64 = 0                       // 开始运行时间 不要修改这个值
   timeBackShift uint64 = 1000 * 86400 * 365 * 30 // 时钟回拨时, 时间回退30年
)

type Worker struct {
   // 添加互斥锁 确保并发安全
   mu sync.Mutex
   // 记录时间戳
   timestamp uint64
   // 该节点的ID
   workerId uint64
   // 当前毫秒已经生成的id序列号(从0开始累加) 1毫秒内最多生成4096个ID
   seq uint64
   // 时钟回拨前最大时间戳
   maxtimestamp uint64
   // 时钟回拨次数
   timeBackCount uint64
}

func hashString(s string) uint32 {
   h := fnv.New32a()
   h.Write([]byte(s))
   return h.Sum32()
}

// 最多支持1024个实例
func getWorkerId_1024() uint64 {
   hostname := getHostname()
   instNo := getInstanceId()
   if instNo < 0 {
      instNo = int64(hashString(hostname))
   }
   instNo = instNo & ((1 << workerBits) - 1) // 保留10位
   return uint64(instNo)
}

func getHostname() string {
   hostname, _ := os.Hostname()
   return hostname
}

func getInstanceId() int64 {
   // 生产环境应该返回 容器实例 ID
   // 开发环境返回 -1

   hostname := getHostname()
   //hostname = "ibt-duoduo-sf-23e72-20.docker.us01" //
   if hostname == "" {
      return -1
   }
   p := strings.Split(hostname, `.docker`)
   for _, str := range p {
      p2 := strings.Split(str, `-`)
      if len(p2) < 2 {
         return -1
      }
      ret, err := strconv.ParseInt(p2[len(p2)-1], 10, 64)
      if err != nil {
         return -1
      }
      return ret //nolint
   }
   return -1
}

// 实例化对象
func NewWorker() *Worker {
   // **************************************************** //
   //       41位毫秒时间戳       10 机器id       12bit 序列号
   // --------------------|-------------|-----------------|
   // **************************************************** //
   // 要先检测workerId是否在上面定义的范围内

   // 生成一个新节点
   w := &Worker{
      timestamp:     0,
      workerId:      getWorkerId_1024(),
      seq:           0,
      maxtimestamp:  0,
      timeBackCount: 0,
   }
   return w
}

// 获取一个新ID
func (w *Worker) next() uint64 {
   w.mu.Lock()
   defer w.mu.Unlock()

   // 当前时间戳 毫秒
   now := uint64(time.Now().UnixNano() / 1e6)

   if w.timestamp == now {
      w.seq = (w.seq + 1) & seqMax
      // 这里要判断，当前工作节点是否在1毫秒内已经生成seqMax个ID
      if w.seq == 0 {
         // 如果当前工作节点在1毫秒内生成的ID已经超过上限 需要等待1毫秒再继续生成
         for now <= w.timestamp {
            now = uint64(time.Now().UnixNano() / 1e6)
         }
      }
   } else {
      // 如果当前时间与工作节点上一次生成ID的时间不一致 则需要重置工作节点生成ID的序号
      w.seq = 0
   }

   // 将机器上一次生成ID的时间更新为当前时间
   w.timestamp = now

   // 处理时钟回拨
   timeBackSecond := uint64(0)
   if now < w.maxtimestamp {
      // 时钟回拨时, 当前时间减去30年, 避免时间戳碰撞
      timeBackSecond = timeBackShift
      w.timeBackCount++
   }

   // 更新最大时间戳
   if now > w.maxtimestamp {
      w.maxtimestamp = now
   }

   // 第一段 now - epoch 为该算法目前已经奔跑了xxx毫秒
   // 如果在程序跑了一段时间修改了epoch这个值 可能会导致生成相同的ID
   ID := ((now - timeBackSecond - epoch) << timeShift) | (w.workerId << workerShift) | (w.seq)
   // 最高位置零
   ID = ID & 0x7fffffffffffffff
   return ID
}

var g_id_worker *Worker = nil

func init() {
   g_id_worker = NewWorker()
}

func NewID() uint64 {
   if g_id_worker == nil {
      panic("snowflake id worker not init yet!")
   }
   return g_id_worker.next()
}

func NewID63() int64 {
   if g_id_worker == nil {
      panic("snowflake id worker not init yet!")
   }
   return int64(g_id_worker.next())
}
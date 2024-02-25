package tool

import (
	"errors"
	"sync"
	"time"
)

const (
	//workerBits表示worker ID的位数
	workerBits uint8 =10
	//序列号位数
	numberBits uint8 =12
	//worker ID最大值
	workerMax int64 =-1^(-1<<workerBits)
	//序列号最大值
	numberMax int64=-1^(-1<<numberBits)
	//时间戳位移量
	timeShift uint8=workerBits+numberBits
	// workerShift 表示 worker ID 的位移量
	workerShift uint8 = numberBits
	// startTime 是起始时间戳（毫秒）
	startTime int64 = 1563172432000// 如果在程序跑了一段时间修改了epoch这个值 可能会导致生成相同的ID //毫秒
)

var wk  *Worker

func GetWorker() *Worker {
	if wk ==nil{
		panic("GetWorker not init")
	}else{
		return wk
	}
}


//生成唯一的ID
type Worker struct {
	mu sync.Mutex//并发安全
	timestamp int64 //当前时间戳
	workerId int64 //工作节点ID
	number int64 //序列号
}

func NewWorker(workerId int64) error {
	if workerId < 0 || workerId > workerMax {
		return errors.New("Worker ID excess of quantity")
	}
	wk = &Worker{
		timestamp: 0,
		workerId:  workerId,
		number:    0,
	}
	// 生成一个新节点
	return nil
}

//生成唯一的ID
func (w *Worker)GetId() int64  {
	w.mu.Lock()
	defer w.mu.Unlock()
	now :=time.Now().UnixNano()/1e6//微妙转毫秒
	if w.timestamp==now{
		w.number++
		if w.number>numberMax{
			for now<=w.timestamp{
				now =time.Now().UnixNano()/1e6
			}
		}
	}else{
		w.number=0
		w.timestamp=now
	}
	ID := int64((now-startTime)<<timeShift | (w.workerId << workerShift) | (w.number))
	return ID
}
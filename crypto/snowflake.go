package crypto

import (
	"math"
	"strconv"
	"sync"
	"time"
)

const (
	workerBits  int8  = 12
	numberBits  int8  = 16
	workerMax   int64 = -1 ^ (-1 << workerBits)
	numberMax   int64 = -1 ^ (-1 << numberBits)
	timeShift   int8  = workerBits + numberBits
	workerShift       = workerBits
)

type Worker struct {
	mu        sync.Mutex
	timestamp int64
	workerId  int64
	number    int64
}

func NewWorker(workerId int64) (*Worker, interface{}) {
	if workerId < 0 || workerId > workerMax {
		return nil, "snowflake"
	}
	return &Worker{
		timestamp: 0,
		workerId:  workerId,
		number:    0,
	}, nil
}

func (w *Worker) GetId() string {
	w.mu.Lock()
	defer w.mu.Unlock()
	nowTime := time.Now()
	now := nowTime.UnixNano() / 1e6
	startTime := nowTime.UnixNano()
	if w.timestamp == now {
		w.number++
		if w.number > numberMax {
			for now <= w.timestamp {
				now = time.Now().UnixNano() / 1e6
			}
		}
	} else {
		w.number = 0
		w.timestamp = now
	}
	ID := int64((now-startTime)<<timeShift | (w.workerId << workerShift) | (w.number))
	aId := math.Abs(float64(ID))
	id := strconv.FormatInt(int64(aId), 16)
	return id
}

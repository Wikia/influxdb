package coordinator

import (
	"sort"
	"sync"
	"sync/atomic"
	"time"

	log "github.com/alecthomas/log4go"
)

const SLOW_QUERY_THRESHOLD = 10 * time.Second

type CoordinatorMonitor struct {
	runningQueries       *RunningQueries
	pointsWritten        uint64
}

func NewCoordinatorMonitor() *CoordinatorMonitor {
	coordinatorMonitor := &CoordinatorMonitor{
		runningQueries:  NewRunningQueries(),
		pointsWritten:   0,
	}

	go coordinatorMonitor.startReportingLoop()

	return coordinatorMonitor
}

func (self *CoordinatorMonitor) RecordWrite(n uint64) {
	atomic.AddUint64(&self.pointsWritten, n)
}

func (self *CoordinatorMonitor) GetWrites() uint64 {
	return atomic.SwapUint64(&self.pointsWritten, 0)
}

func (self *CoordinatorMonitor) StartQuery(q *RunningQuery) {
	self.runningQueries.Add(q)
}

func (self *CoordinatorMonitor) EndQuery(q *RunningQuery) {
	took := time.Now().Sub(q.startTime)
	if took > SLOW_QUERY_THRESHOLD {
		log.Info("Slow Query [took %fs]: db: %s, u: %s, q: %s", took.Seconds(), q.databaseName, q.userName, q.queryString)
	}
	self.runningQueries.Remove(q)
}

func (self *CoordinatorMonitor) GetRunningQueries() *RunningQueryList {
	return self.runningQueries.AllSorted()
}

func (self *CoordinatorMonitor) startReportingLoop() chan struct{} {
	ticker := time.NewTicker(time.Minute)
	for {
		select {
		case <-ticker.C:
			self.reportStats()
		}
	}
}

func (self *CoordinatorMonitor) reportStats() {
	pointsWritten := self.GetWrites()

	log.Info("Workload stats: points written = %d", pointsWritten)
}

type RunningQuery struct {
	userName             string
	databaseName         string
	queryString          string
	startTime            time.Time
	remoteAddr           string
}

type RunningQueryList  []*RunningQuery
type RunningQueries struct {
	data                 map[*RunningQuery]struct{}
	sync.Mutex
}

func NewRunningQuery(
	userName string, databaseName string, queryString string, startTime time.Time, remoteAddr string) *RunningQuery {
	runningQuery := &RunningQuery{
		userName:             userName,
		databaseName:         databaseName,
		queryString:          queryString,
		startTime:            startTime,
		remoteAddr:           remoteAddr,
	}

	return runningQuery
}

func NewRunningQueries() *RunningQueries {
	runningQueries := &RunningQueries{
		data:                 make(map[*RunningQuery]struct {}),
	}

	return runningQueries
}

func (self *RunningQueries) Add(q *RunningQuery) {
	self.Lock()
	self.data[q] = struct{}{}
	self.Unlock()
}

func (self *RunningQueries) Remove(q *RunningQuery) {
	self.Lock()
	delete(self.data, q)
	self.Unlock()
}

func (self *RunningQueries) All() <-chan *RunningQuery {
	ch := make(chan *RunningQuery)
	go func() {
		self.Lock()
		for elem := range self.data {
			ch <- elem
		}
		close(ch)
		self.Unlock()
	}()

	return ch
}

func (self *RunningQueries) AllSorted() *RunningQueryList {
	all := RunningQueryList{}
	for q := range self.All() {
		all = append(all,q)
	}
	sort.Sort(all)

	return &all
}

func (self RunningQueryList) Len() int {
	return len(self)
}

func (self RunningQueryList) Less(i, j int) bool {
	return (self[i]).startTime.UnixNano() < (self[j]).startTime.UnixNano()
}

func (self RunningQueryList) Swap(i, j int) {
	self[i], self[j] = self[j], self[i]
}

package result

import (
	"fmt"
	"time"
)

type ResultType uint

const (
	ResultTypeInt64 = iota
	ResultTypeFloat = iota
)

type Result struct {
	Address string
	Probe string
	Results []ResultRow
	Time time.Time
}


func (r *Result) String () (string) {
	return fmt.Sprintf("%s %v", r.Address, r.Results)
}


type ResultRow interface{
	Get() (string, interface{})
	Type() (ResultType)
}

type ResultRowInt64 struct {
	Key string
	Value int64
}

func (r *ResultRowInt64) Get() (string, interface{}) {
	return r.Key, r.Value
}
func (r *ResultRowFloat64) Type() ResultType { return ResultTypeFloat }


func (r *ResultRowInt64) String() string {
	return fmt.Sprintf("%s: %d", r.Key, r.Value)
}
func (r *ResultRowInt64) Type() ResultType { return ResultTypeInt64 }



type ResultRowFloat64 struct {
	Key string
	Value float64
}

func (r *ResultRowFloat64) Get() (string, interface{}) {
	return r.Key, r.Value
}

func (r *ResultRowFloat64) String() string {
	return fmt.Sprintf("%s: %f", r.Key, r.Value)
}

package core

import (
	"math"
	_ "reflect"
)

type Segment struct {
	Index int
	Last  int
}

func (this *Segment) Count() int {
	return this.Last - this.Index
}
func (this *Segment) SetCount(value int) {
	this.Last = this.Index + value
}
func (this *Segment) Equals(other interface{}) bool {
	o, ok := other.(Segment)
	if !ok {
		return false
	}
	return this.Index == o.Index && this.Last == o.Last
}
func (this *Segment) SetIndexLast(index, last int) {
	this.Index = index
	this.Last = last
}
func (this *Segment) Offset(value int) {
	this.Index += value
	if this.Last > 0 {
		this.Last += value
		if this.Last < 0 {
			this.Last = util.IntMax()
		}

	} else {
		this.Last += value
	}
}
func (this *Segment) Merge(other *Segment) bool {
	if other.Index <= this.Last && other.Last >= this.Index {
		this.Index = int(math.Min(float64(this.Index), float64(other.Index)))
		this.Last = int(math.Max(float64(this.Last), float64(other.Last)))
		return true
	} else {
		return false
	}
}

func (this *Segment) Remove(other *Segment) bool {
	return this.Remove2(other.Index, other.Count())
}

func (this *Segment) Remove2(index, count int) bool {
	var tempLast = index + count
	if index < this.Last && tempLast > this.Index {
		if index <= this.Index {
			this.Index = tempLast
		} else {
			this.Last = this.Index
		}
	} else {
		if tempLast < this.Last {
			this.Last -= count
		} else {
			this.Last = index
		}

	}
	return true
}

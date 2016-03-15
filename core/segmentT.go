package core

import (
	"errors"
	"math"
	"reflect"
)

type SegmentT struct {
	Index int
	Last  int
	Value interface{}
}

func NewSegmentT(innerValue interface{}) *SegmentT {
	var s = new(SegmentT)
	s.Value = innerValue
	return s

}
func (this *SegmentT) Count() int {
	return this.Last - this.Index
}
func (this *SegmentT) SetCount(value int) {
	this.Last = this.Index + value
}
func (this *SegmentT) Equals(other SegmentT) bool {

	return this.Index == other.Index &&
		this.Last == other.Last &&
		reflect.DeepEqual(this.Value, other.Value)
}
func (this *SegmentT) SetIndexLast(index, last int) {
	this.Index = index
	this.Last = last
}
func (this *SegmentT) Offset(value int) {
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
func (this *SegmentT) Merge(other SegmentT) bool {
	if other.Index <= this.Last && other.Last >= this.Index {
		this.Index = int(math.Min(float64(this.Index), float64(other.Index)))
		this.Last = int(math.Max(float64(this.Last), float64(other.Last)))
		return true
	} else {
		return false
	}
}
func (this *SegmentT) MergeOrOverride(other *SegmentT) (bool, error) {
	if other.Index <= this.Last && other.Last >= this.Index {
		if reflect.DeepEqual(this.Value, other.Value) {
			this.Index = int(math.Min(float64(this.Index), float64(other.Index)))
			this.Last = int(math.Max(float64(this.Last), float64(other.Last)))
			return true, nil
		} else {
			if this.Index <= other.Index && this.Last >= other.Last {
				return true, nil
			}
			if other.Index < this.Index && other.Last > this.Last {
				return false, errors.New("this segment should not be contained by other segment.")
			}
			if other.Index < this.Index {
				other.SetIndexLast(other.Index, this.Index)
			} else if other.Last > this.Last {
				other.SetIndexLast(this.Last, other.Last)
			}

		}
	}
	return false, errors.New("merge fail.")
}
func (this *SegmentT) Remove(other SegmentT) bool {
	return this.Remove2(other.Index, other.Count())
}

func (this *SegmentT) Remove2(index, count int) bool {
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

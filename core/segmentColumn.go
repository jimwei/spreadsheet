package core

type ColumnSegment struct {
	Index int
	Last  int
	_rows *SegmentList
}

func (this *ColumnSegment) Count() int {
	return this.Last - this.Index
}
func (this *ColumnSegment) SetCount(v int) {
	this.Last = this.Index + v
}
func (this *ColumnSegment) Rows() *SegmentList {
	if this._rows == nil {
		this._rows = new(SegmentList)
	}
	return this._rows
}
func (this *ColumnSegment) SetRows(list *SegmentList) {
	this._rows = list
}
func (this *ColumnSegment) Remove(index, count int) bool {
	tempLast := index + count
	if index < this.Last && tempLast > this.Index {
		if index <= this.Index {
			if tempLast < this.Last {
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
	return false
}
func (this *ColumnSegment) Equals(other *ColumnSegment) bool {
	if other == nil {
		return false
	}
	if this.Index != other.Index || this.Last != other.Last {
		return false
	}
	rows := this.Rows

	return rows.Equals(other.Rows())
}

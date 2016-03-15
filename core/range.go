package core

//	"spreadsheet/consts"
import (
	"bytes"
	"strconv"
)

var (
	EmptyCellRange = CellRect{Row: 0, Column: 0, RowCount: 0, ColumnCount: 0}
)

type CellPoint struct {
	Row    int
	Column int
}
type CellRect struct {
	Row         int
	Column      int
	RowCount    int
	ColumnCount int
}
type CellRectList struct {
	list []*CellRect
}

func (this *CellRect) ToString() string {
	if this.Row >= 0 && this.Column >= 0 {
		if this.ColumnCount > 1 || this.RowCount > 1 {
			return indexHelper.GetColumnIndexInA1Letter(this.Column) + strconv.Itoa(this.Row+1) +
				":" + indexHelper.GetColumnIndexInA1Letter(this.Column+this.ColumnCount-1) +
				strconv.Itoa(this.Row+this.RowCount)
		} else {
			return indexHelper.GetColumnIndexInA1Letter(this.Column) + strconv.Itoa(this.Row+1)
		}
	} else if this.Row >= 0 {
		return strconv.Itoa(this.Row+1) + ":" + strconv.Itoa(this.Row+this.RowCount)
	} else if this.Column >= 0 {
		return indexHelper.GetColumnIndexInA1Letter(this.Column) + ":" +
			indexHelper.GetColumnIndexInA1Letter(this.Column+this.ColumnCount-1)
	} else {
		return indexHelper.GetColumnIndexInA1Letter(0) + ":" + indexHelper.GetColumnIndexInA1Letter(16384-1)
	}

}
func (this *CellRect) IsFullRow() bool {

	return this.Column == -1 && this.ColumnCount == -1 ||
		this.Column == 0 && this.ColumnCount == indexHelper.MaxColumnCount
}
func (this *CellRect) IsFullColumn() bool {
	return this.Row == -1 && this.RowCount == -1 ||
		this.Row == 0 && this.RowCount == indexHelper.MaxRowCouunt
}
func (this *CellRect) IsFullSingleRow() bool {
	return this.RowCount == 1 && this.IsFullRow()
}
func (this *CellRect) IsFullSingleColumn() bool {
	return this.ColumnCount == 1 && this.IsFullColumn()
}
func (this *CellRect) Left() int {
	return this.Column
}
func (this *CellRect) Top() int {
	return this.Row
}
func (this *CellRect) Right() int {
	return this.Column + this.ColumnCount
}
func (this *CellRect) Bottom() int {
	return this.Row + this.RowCount
}
func (this *CellRect) IsEmpty() bool {
	return this.Column == 0 || this.RowCount == 0
}
func (this *CellRectList) Init(rects []*CellRect) {
	this.list = rects
}

//------------cellrect list----------------
func (this *CellRectList) Row() int {
	if this.list == nil || len(this.list) == 0 {
		return -1
	}
	return this.list[0].Row

}
func (this *CellRectList) Column() int {
	if this.list == nil || len(this.list) == 0 {
		return -1
	}
	return this.list[0].Column
}
func (this *CellRectList) IndexOf(item *CellRect) int {
	i := 0
	for _, r := range this.list {
		if r == item {
			return i
		}
		i++
	}
	return -1
}
func (this *CellRectList) Insert(index int, item *CellRect) {
	this.list = append(this.list[:index], append([]*CellRect{item}, this.list[index+1:]...)...)
}
func (this *CellRectList) RemoveAt(index int) {
	if index < 0 || index >= len(this.list) {
		return
	}
	this.list = append(this.list[:index], this.list[index+1:]...)
}
func (this *CellRectList) Contains(item *CellRect) bool {
	for _, v := range this.list {
		if v == item {
			return true
		}
	}
	return false
}
func (this *CellRectList) ToString() string {
	str := bytes.NewBufferString("")
	for _, v := range this.list {
		str.WriteString(" ")
		str.WriteString(v.ToString())
	}
	return str.String()
}
func (this *CellRectList) ToSegmentTable() SegmentTable {
	//TODO:
	return nil
}

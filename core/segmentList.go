package core

import (
	"errors"
	"math"
)

type SegmentList struct {
	LastVisitedItemIndex int
	Items                []*Segment
	Count                int
	MinCapacity          int
}

func NewSegmentList(segmnets []*Segment) *SegmentList {
	var list = new(SegmentList)
	if segmnets == nil {
		list.MinCapacity = 2
		return list
	}
	list.Items = segmnets
	list.MinCapacity = len(segmnets)
	return list
}
func (this *SegmentList) Index(index int) *Segment {
	return this.Items[index]
}
func (this *SegmentList) SearchStartItemIndex(index int) int {
	if this.LastVisitedItemIndex >= this.Count {
		if this.Count == 0 {
			return -1
		}
		this.LastVisitedItemIndex = this.Count - 1
	}
	low := 0
	high := this.Count - 1
	lastSegment := this.Items[this.LastVisitedItemIndex]
	if index < lastSegment.Last {
		if index >= lastSegment.Index || this.LastVisitedItemIndex == 0 {
			return this.LastVisitedItemIndex
		}
		prev := this.Items[this.LastVisitedItemIndex-1]
		if index >= prev.Last {
			return this.LastVisitedItemIndex
		}
		this.LastVisitedItemIndex--
		if index >= prev.Index || this.LastVisitedItemIndex == 0 {
			return this.LastVisitedItemIndex
		}
		if index < this.Items[0].Last {
			this.LastVisitedItemIndex = 0
			return this.LastVisitedItemIndex
		}
		high = this.LastVisitedItemIndex
	} else {
		this.LastVisitedItemIndex++
		if this.LastVisitedItemIndex >= this.Count {
			return -1
		}
		if index < this.Items[this.LastVisitedItemIndex].Last {
			return this.LastVisitedItemIndex
		}
		this.LastVisitedItemIndex++
		if this.LastVisitedItemIndex >= this.Count {
			return -1
		}
		if index >= this.Items[this.Count-1].Last {
			//clarify: not same to origin
			this.LastVisitedItemIndex = this.Count - 1
			return -1
		}
		low = this.LastVisitedItemIndex
	}
	for low <= high {
		middle := (low + high) / 2
		if index < this.Items[middle].Index {
			if middle > low && index < this.Items[middle-1].Last {
				high = middle - 1
				continue
			}
			this.LastVisitedItemIndex = middle
			return this.LastVisitedItemIndex
		} else if index >= this.Items[middle].Last {
			low = middle + 1
			continue
		}
		this.LastVisitedItemIndex = middle
		return this.LastVisitedItemIndex
	}
	return -1
}
func (this *SegmentList) Contains(index int) bool {
	sIndex := this.SearchStartItemIndex(index)
	if sIndex < 0 || index < this.Items[sIndex].Index {
		return false
	}
	return true

}
func (this *SegmentList) Contains1(index, count int) NullBool {
	sIndex := this.SearchStartItemIndex(index)
	var t NullBool
	if sIndex < 0 {
		return t.SetValue(false)
	}
	bottom := index + count
	entry := this.Items[sIndex]
	if bottom <= entry.Index {
		return t.SetValue(false)
	}
	if index < entry.Index || bottom > entry.Last {
		return t
	}
	return t.SetValue(true)
}
func (this *SegmentList) ensureCapacity(min, max int) {
	if this.Items == nil || len(this.Items) < min {
		if min < this.MinCapacity {
			min = this.MinCapacity
		}
		if this.Items == nil {
			this.Items = make([]*Segment, min)
			return
		}
		count := math.Max(float64(min), math.Min(float64(max), float64(len(this.Items)*2)))
		//expand
		this.Items = append(this.Items, make([]*Segment, int(count))...)
	}
}
func (this *SegmentList) AddItem(segment *Segment, autoMerge bool) error {
	if autoMerge && this.Count > 0 {
		lastItem := this.Items[this.Count-1]
		if lastItem.Last >= segment.Index {
			if lastItem.Last > segment.Index {
				return errors.New("ArgumentException.")
			}
			this.Items[this.Count-1].Last = segment.Last
			return nil
		}
	}
	this.ensureCapacity(this.Count+1, util.IntMax())
	this.Items[this.Count] = segment
	this.Count++
	return nil
}
func (this *SegmentList) AddValue(index, count int, autoMerge bool) {
	seg := Segment{Index: index}
	seg.SetCount(count)
	this.AddItem(&seg, autoMerge)
}
func (this *SegmentList) InsertItem(index int, segment *Segment) {
	this.InsertItem1(index, 1)
	this.Items[index] = segment

}
func (this *SegmentList) InsertItem1(index, count int) {
	this.ensureCapacity(this.Count+count, util.IntMax())
	this.Count += count

}
func (this *SegmentList) InsertValue(index, count int) {
	itemIndex := this.SearchStartItemIndex(index)
	if itemIndex < 0 {
		this.AddValue(index, count, true)
		return
	}
	if index < this.Items[itemIndex].Index {
		if itemIndex > 0 && index == this.Items[itemIndex-1].Last {
			itemIndex--
			this.Items[itemIndex].Last += count
		} else {
			seg := Segment{Index: index}
			seg.SetCount(count)
			this.InsertItem(itemIndex, &seg)
		}
	} else {
		this.Items[itemIndex].Last += count
	}
	for i := itemIndex; i < this.Count; i++ {
		this.Items[i].Offset(count)
	}
}

func (this *SegmentList) InsertEmptyValue(index, count int) {
	itemIndex := this.SearchStartItemIndex(index)
	if itemIndex < 0 {
		return
	}
	if index > this.Items[itemIndex].Index {
		this.InsertItem(itemIndex, this.Items[itemIndex])
		this.Items[itemIndex].Last = index
		itemIndex++
		this.Items[itemIndex].Index = index
	}
	for i := itemIndex; i < this.Count; i++ {
		this.Items[i].Offset(count)
	}
}
func (this *SegmentList) RemoveItemAt(index, count int) {
	if count > 0 {
		this.Items = append(this.Items[:index], this.Items[index+count:]...)
		this.Count -= count
	}
}
func (this *SegmentList) RemoveValue(index, count int) {
	stIndex := this.SearchStartItemIndex(index)
	if stIndex < 0 {
		return
	}
	last := index + count
	if index > this.Items[stIndex].Index {
		if last <= this.Items[stIndex].Last {
			this.Items[stIndex].SetCount(this.Items[stIndex].Count() - count)
			for i := stIndex; i < this.Count; i++ {
				this.Items[i].Offset(-count)
			}
			return
		}
		this.Items[stIndex].Last = index
		stIndex++
	}
	endIndex := this.SearchStartItemIndex(last)
	if endIndex < 0 {
		this.Count = stIndex
		return
	}
	if last > this.Items[endIndex].Index {
		this.Items[endIndex].Index = last
	}
	for i := endIndex; i < this.Count; i++ {
		this.Items[i].Offset(-count)
	}
	if stIndex > 0 && this.Items[stIndex-1].Last == this.Items[endIndex].Index {
		this.Items[stIndex-1].Last = this.Items[endIndex].Last
		endIndex++
	}
	this.RemoveItemAt(stIndex, endIndex-stIndex)
}
func (this *SegmentList) Clear() {
	this.Count = 0
}
func (this *SegmentList) Union(index, count int) int {
	newItem := new(Segment)
	newItem.Index = index
	newItem.SetCount(count)
	if this.Count == 0 {
		this.AddItem(newItem, true)
		return 0
	}
	stIndex := this.SearchStartItemIndex(index)
	if stIndex < 0 {
		if !this.Items[this.Count-1].Merge(newItem) {
			this.AddItem(newItem, true)
		}
		return this.Count - 1
	}
	stItem := this.Items[stIndex]
	endIndex := stIndex
	if newItem.Last <= stIndex {
		if newItem.Index >= stItem.Index {
			return -1
		}
	} else {
		endIndex = this.SearchStartItemIndex(newItem.Last)
		if endIndex < 0 {
			endIndex = this.Count
		}
	}
	if stIndex > 0 && newItem.Merge(this.Items[stIndex-1]) {
		stIndex--
	} else if stIndex != endIndex {
		newItem.Merge(this.Items[stIndex])
	}
	if endIndex < this.Count && newItem.Merge(this.Items[endIndex]) {
		endIndex++
	}
	if endIndex == stIndex {
		this.InsertItem(stIndex, newItem)
	} else {
		this.Items[stIndex] = newItem
		this.RemoveItemAt(stIndex+1, endIndex-stIndex-1)
	}
	return stIndex
}
func (this *SegmentList) Clone() *SegmentList {
	list := new(SegmentList)
	if this.Count > 0 {
		list.Count = this.Count
		list.Items = make([]*Segment, this.Count)
		list.Items = append(list.Items, this.Items...)
		list.LastVisitedItemIndex = this.LastVisitedItemIndex
	}
	return list
}
func (this *SegmentList) Union1(index, count int, newList *SegmentList) int {
	newItem := new(Segment)
	newItem.Index = index
	newItem.SetCount(count)
	if this.Count == 0 {
		newList = new(SegmentList)
		newList.AddItem(newItem, true)
		return 0
	}
	stIndex := this.SearchStartItemIndex(index)
	//not find
	if stIndex < 0 {
		newList = this.Clone()
		if !newList.Items[newList.Count-1].Merge(newItem) {
			newList.AddItem(newItem, true)
		}
		return newList.Count - 1
	}

	//look for the endindex
	stItem := this.Items[stIndex]
	endIndex := stIndex
	if newItem.Last <= stItem.Last {
		if newItem.Index >= stItem.Index {
			return -1
		}
	} else {
		endIndex = this.SearchStartItemIndex(newItem.Last)
		if endIndex < 0 {
			endIndex = this.Count
		}
	}

	//merge start item
	if stIndex > 0 && newItem.Merge(this.Items[stIndex-1]) {
		stIndex--
	} else if stIndex != endIndex {
		newItem.Merge(this.Items[stIndex])
	}

	//merge end item
	if endIndex < this.Count && newItem.Merge(this.Items[endIndex]) {
		endIndex++
	}
	newList = this.Clone()
	if endIndex == stIndex {
		newList.InsertItem(stIndex, newItem)
	} else {
		newList.Items[stIndex] = newItem
		newList.RemoveItemAt(stIndex+1, endIndex-stIndex-1)
	}
	return stIndex
}

func (this *SegmentList) Union2(bList *SegmentList) {
	if bList == nil || bList.Count == 0 {
		return
	}
	if this.Count == 0 {
		if this.Items == nil || len(this.Items) < bList.Count {
			this.Items = make([]*Segment, bList.Count)
		}
		this.Count = copy(this.Items, bList.Items)
		return
	}

	var list []*Segment
	aIndex := 0
	bIndex := 0
	last := new(Segment)
	for aIndex < this.Count || bIndex < bList.Count {
		var item = new(Segment)
		if aIndex < this.Count {
			if bIndex < bList.Count {
				if this.Items[aIndex].Index <= bList.Items[bIndex].Index {
					item = this.Items[aIndex]
					aIndex++
				} else {
					item = bList.Items[bIndex]
					bIndex++
				}
			} else {
				item = this.Items[aIndex]
				aIndex++
			}
		} else {
			item = bList.Items[bIndex]
			bIndex++
		}
		if len(list) == 0 {
			list = append(list, item)
			last = item
			continue
		}
		if item.Index <= last.Last {
			last.Last = int(math.Max(float64(last.Last), float64(item.Last)))
			list[len(list)-1] = last
			continue
		}

		list = append(list, item)
		last = item
	}
	copy(list, this.Items)
	this.Count = len(this.Items)
}
func (this *SegmentList) Union3(a, b *SegmentList) *SegmentList {
	c := a.Clone()
	c.Union2(b)
	return c
}

//*******Intersect Begin***************************
func (this *SegmentList) IntersectsWith(index, count int) bool {
	if this.Count == 0 {
		return false
	}
	stIndex := this.SearchStartItemIndex(index)
	if stIndex < 0 {
		return false
	}
	ltIndex := index + count
	stItem := this.Items[stIndex]
	if ltIndex <= stItem.Index {
		return false
	}
	return true
}
func (this *SegmentList) Intersect(index, count int) int {
	if this.Count == 0 {
		return -1
	}
	stIndex := this.SearchStartItemIndex(index)
	if stIndex < 0 {
		this.Count = 0
		return 0
	}
	ltIndex := index + count
	stItem := this.Items[stIndex]
	if ltIndex <= stItem.Index {
		this.Count = 0
		return 0
	}
	if ltIndex <= stItem.Last {
		this.Items[0].SetIndexLast(int(math.Max(float64(index), float64(stItem.Index))), ltIndex)
		this.Count = 1
		return 0
	}
	ltItemIndex := this.SearchStartItemIndex(ltIndex)
	if ltItemIndex < 0 {
		if stIndex == 0 && index <= stItem.Index {
			return -1
		}
		if index > stItem.Index {
			this.Items[stIndex].SetIndexLast(index, stItem.Last)
		}
		if stIndex > 0 {
			this.Count -= stIndex
			this.Items = this.Items[stIndex:]
		}
		return 0
	}

	this.Items[stIndex].SetIndexLast(int(math.Max(float64(index), float64(stItem.Index))), stItem.Last)
	ltItem := this.Items[ltItemIndex]
	if ltIndex > ltItem.Index {
		this.Items[ltItemIndex].SetIndexLast(ltItem.Index, ltIndex)
		ltItemIndex++
	}
	this.Count = ltItemIndex - stIndex
	if stIndex == 0 && index <= stItem.Index {
		if ltIndex > ltItem.Index {
			return this.Count - 1
		}

		return this.Count

	}
	this.Items = this.Items[stIndex:]
	return 0
}
func (this *SegmentList) Intersect1(index, count int, newList *SegmentList) int {
	if this.Count == 0 {
		return -1
	}
	if newList == nil {
		newList = NewSegmentList(nil)
	}
	stIndex := this.SearchStartItemIndex(index)
	if stIndex < 0 {
		return 0
	}
	ltIndex := index + count
	stItem := this.Items[stIndex]
	if ltIndex <= stItem.Index {
		return 0
	}
	if ltIndex <= stItem.Last {
		newSeg := new(Segment)
		newSeg.SetIndexLast(int(math.Max(float64(index), float64(stItem.Index))), ltIndex)
		newList.AddItem(newSeg, true)
		return 0
	}

	ltItemIndex := this.SearchStartItemIndex(ltIndex)
	if ltItemIndex < 0 {
		if stIndex == 0 && index <= stItem.Index {
			return -1
		}
		newList = this.Clone()
		if index > stItem.Index {
			newList.Items[stIndex].SetIndexLast(index, stItem.Last)
		}
		if stIndex > 0 {
			newList.Count -= stIndex
			newList.Items = newList.Items[stIndex:]
		}
		return 0
	}

	newList = this.Clone()
	newList.Items[stIndex].SetIndexLast(int(math.Max(float64(index), float64(stItem.Index))), stItem.Last)
	ltItem := newList.Items[ltItemIndex]
	if ltIndex > ltItem.Index {
		newList.Items[ltItemIndex].SetIndexLast(ltItem.Index, ltIndex)
		ltItemIndex++
	}
	newList.Count = ltItemIndex - stIndex
	if stIndex == 0 && index <= stItem.Index {
		if ltIndex > ltItem.Index {
			return newList.Count - 1
		}
		return newList.Count
	}
	newList.Items = newList.Items[stIndex:]
	return 0
}
func (this *SegmentList) Intersect3(bList *SegmentList) {
	if this.Count == 0 {
		return
	}
	if bList.Count == 0 {
		this.Count = 0
		return
	}
	list := NewSegmentList(nil)
	asIndex := 0
	bsIndex := 0
	for asIndex < this.Count && bsIndex < bList.Count {
		iafter := this.Items[asIndex]
		iBefore := bList.Items[bsIndex]
		if iafter.Last < iBefore.Last {
			iBefore, iafter = iafter, iBefore
			asIndex++
		} else {
			bsIndex++
		}

		if iBefore.Last > iafter.Index {
			index := int(math.Max(float64(iBefore.Index), float64(iafter.Index)))
			count := iBefore.Last - index
			seg := new(Segment)
			seg.Index = index
			seg.SetCount(count)
			list.AddItem(seg, true)
		}
	}
	this.Count = list.Count
	if this.Count > 0 {
		this.Items = list.Items[:]
	}
}
func (this *SegmentList) Intersect4(a, b *SegmentList) *SegmentList {
	c := a.Clone()
	c.Intersect3(b)
	return c
}

//*******Intersect End***************************
//*******Exclude End***************************
func (this *SegmentList) Exclude(index, count int) int {
	stIndex := this.SearchStartItemIndex(index)
	if stIndex < 0 {
		return -1
	}
	ltIndex := index + count
	stItem := this.Items[stIndex]
	if ltIndex < stItem.Index {
		return -1
	}
	if ltIndex <= stItem.Last {
		if index <= stItem.Index {
			if ltIndex == stItem.Last {
				this.RemoveItemAt(stIndex, 1)
			} else {
				this.Items[stIndex].SetIndexLast(ltIndex, stItem.Last)
			}
		} else {
			if ltIndex == stItem.Last {
				this.Items[stIndex].SetIndexLast(stItem.Index, index)
			} else {
				this.InsertItem(stIndex, stItem)
				this.Items[stIndex].SetIndexLast(stItem.Index, index)
				this.Items[stIndex+1].SetIndexLast(ltIndex, stItem.Last)
			}
		}
		return stIndex
	}

	ltItemImdex := this.SearchStartItemIndex(ltIndex)
	if ltItemImdex < 0 {
		if stIndex == 0 && index <= stItem.Index {
			this.Count = 0
			return 0
		} else if index > stItem.Index {
			this.Items[stIndex].SetIndexLast(stItem.Index, index)
			this.Count = stIndex + 1
		} else {
			this.Count = stIndex
		}
		return stIndex
	}
	cIndex := stIndex
	if index > stItem.Index {
		this.Items[stIndex].SetIndexLast(stItem.Index, index)
		stIndex++
	}
	ltItem := this.Items[ltIndex]
	if ltIndex > ltItem.Index {
		this.Items[ltIndex].SetIndexLast(ltIndex, ltItem.Index)
	}
	this.Items = this.Items[ltIndex:]
	this.Count = stIndex + this.Count - ltIndex
	return cIndex
}
func (this *SegmentList) Exclude1(bList *SegmentList) {
	if this.Count == 0 || bList == nil || bList.Count == 0 {
		return
	}
	if bList.Items[0].Index >= this.Items[this.Count-1].Last ||
		this.Items[0].Index >= bList.Items[bList.Count-1].Last {
		return
	}
	var list []*Segment
	asIndex := 0
	bsIndex := 0
	aseg := new(Segment)
	bseg := bList.Items[0]
	for {
		if aseg.Count() == 0 {
			for asIndex < this.Count && this.Items[asIndex].Last <= bseg.Index {
				list = append(list, this.Items[asIndex])
			}
			if asIndex == this.Count {
				break
			}
			aseg = this.Items[asIndex]
		}
		for bsIndex < bList.Count && bList.Items[bsIndex].Last <= aseg.Index {
			bsIndex++
		}
		if bsIndex == bList.Count {
			list = append(list, aseg)
			asIndex++
			for asIndex < this.Count {
				list = append(list, aseg)
				asIndex++
			}
			break

		}
		bseg = bList.Items[bsIndex]
		if aseg.Index < bseg.Index {
			newSeg := new(Segment)
			newSeg.Index = aseg.Index
			newSeg.SetCount(bseg.Index - aseg.Index)
			list = append(list, newSeg)
		}
		if aseg.Last > bseg.Last {
			aseg = new(Segment)
			aseg.Index = bseg.Last
			aseg.SetCount(aseg.Last - bseg.Last)

		} else {
			aseg.SetCount(0)
			asIndex++
		}
	}
	this.Count = len(list)
	if this.Count > 0 {
		copy(list, this.Items)
	}
}
func (this *SegmentList) Exculde2(a, b *SegmentList) *SegmentList {
	c := a.Clone()
	c.Exclude1(b)
	return c
}

//*******Exclude End***************************
//*******Offset and Extend Begin***************************
//TODO:
//*******Offset and Extend End***************************

//---------equals----------------
func (this *SegmentList)Equals(other *SegmentList)bool  {
    if other == nil{
        return false
    }
    if other.Count != this.Count{
        return false
    }
    for i,v:= range this.Items{
        if !v.Equals(other.Items[i]){
            return false
        }
    }
    return true
    
}
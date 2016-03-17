package core
type SegmentTable struct{
    _items []*ColumnSegment
    _boundsIsValid bool
    _count int
    _bounds *CellRect
    _lastVisitedItemIndex int
    _minCapacity int
}

func NewSegmentTable() *SegmentTable  {
    obj := new(SegmentTable)
    obj._minCapacity = 2
    return obj
}

func (this *SegmentTable)Capacity() int  {
    if this._items == nil{
        return 0
    }
    return len(this._items)
}
// Get indexer
func (this *SegmentTable)Get(index int) *ColumnSegment  {
    return this._items[index]
}
// Set indexer
func (this *SegmentTable)Set(index int, value *ColumnSegment )  {
    this._items[index] = value
}
func (this *SegmentTable)Count()  int{
    return this._count
}
func (this *SegmentTable)IsEmpty() bool  {
    return this._count == 0
}
func (this *SegmentTable)Bounds() *CellRect {
    if this._boundsIsValid{
        this._bounds = new(CellRect)
    }else{
        left := this._items[0].Index
        right := this._items[this.Count()-1].Last
        top := this._items[0].Rows().Items[0].Index
        bottom := this._items[0].Rows().Items[this._items[0].Rows().Count-1].Last
        for i:=1;i<this.Count();i++{
            index:= this._items[i].Rows().Items[0].Index 
            if top>index{
                top = index
            }
            b:= this._items[i].Rows().Items[this._items[i].Rows().Count-1].Last
            if bottom<b {
                bottom = b
            }
        }
        this._bounds = new(CellRect)
        this._bounds.Row = top
        this._bounds.Column = left
        this._bounds.RowCount = bottom - top
        this._bounds.ColumnCount = right - left
    }
    
    return this._bounds
}

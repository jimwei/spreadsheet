package core
import
(
    "spreadsheet/consts"
)

type DataEntry struct{
    DataType consts.DataType
    BlockIndex int
    RowIndex int
    Count int
    LastIndex int
}
func (this *DataEntry)BottomIndex() int  {
    return this.RowIndex +  this.Count
}
func (this *DataEntry)BottomBlockIndex() int  {
    return this.BlockIndex +  this.Count
}
func (this *DataEntry)GetBlockIndex(row int) int  {
    return row - this.RowIndex + this.BlockIndex
}
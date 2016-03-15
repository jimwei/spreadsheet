package core
import
(
    "spreadsheet/consts"
)
type FormulaValue struct{
    iDandStatus int
    Value float64
}
func (this *FormulaValue)ID() int  {
    return this.iDandStatus & 0xffffff;
}
func (this *FormulaValue)SetID(id int)  {
    this.iDandStatus = this.iDandStatus& 0xff000000 + id & 0xffffff
 }
 
func (this *FormulaValue)Status() consts.FormulaValueStatus  {
    return  consts.FormulaValueStatus(this.iDandStatus>>24) 
}
func (this *FormulaValue)SetStatus(id consts.FormulaValueStatus)  {
    this.iDandStatus = this.iDandStatus & 0xffffff + int(id)<<24
}
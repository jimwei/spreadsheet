package consts

type DataType int
const (
	DT_None     DataType = 1<<iota
	DT_Number     
	DT_Text       
	DT_Interface  
	DT_Formula   
)


type FormulaValueStatus int

const(
    FVS_Dirty FormulaValueStatus = iota
    FVS_NullReady
    FVS_NumberReady
    FVS_TextReady
    FVS_LogicalReady
    FVS_Error
    FVS_Calculating
    FVS_Scaning
)

type RangeType int

const
(
    RT_None RangeType = 0x00
    RT_Cell = 0x01
    RT_Row = 0x02
    RT_Column = 0x04
    RT_Sheet = 0x08
)
package consts

type ValueType int

const (
	//value type
	VT_Empty ValueType = 1 << iota
	VT_Number
	VT_Text
	VT_Logical
	VT_Error
)

type FunctionValueType int

const (
	FuncVT_Number FunctionValueType = 1 << iota
	FuncVT_Text
	FuncVT_Logical
	FuncVT_Variant
)

type ErrorType int

const (
	ET_None ErrorType = iota
	ET_Null
	ET_Div0
	ET_Value
	ET_Ref
	ET_Name
	ET_Num
	ET_NA
	ET_GettingData
)

func ErrorString(err ErrorType) string {
	var r = ""
	switch err {
	case ET_None:
		break
	case ET_Null:
		r = "#NULL!"
		break
	case ET_Div0:
		r = "#DIV/0!"
		break
	case ET_Value:
		r = "#VALUE!"
		break
	case ET_Ref:
		r = "#REF!"
		break
	case ET_Name:
		r = "#NAME?"
		break
	case ET_Num:
		r = "#NUM!"
		break
	case ET_NA:
		r = "#N/A"
		break
	case ET_GettingData:
		r = "#GETTING_DATA"
		break
	default:
		break

	}
	return r
}

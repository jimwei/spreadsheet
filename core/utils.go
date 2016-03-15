package core

import (
	"bytes"
	"runtime"
	"spreadsheet/consts"
	"strconv"
	"strings"
	"unicode"
	_ "unicode/utf8"
    "math"
)

var (
	util        *Util
	stringUtil  *StringUtil
	indexHelper *IndexHelper
    doubleComparer *DoubleComparer
)

func init() {
	util = new(Util)
	stringUtil = new(StringUtil)
    doubleComparer = new(DoubleComparer)
	initIndexHelper()
}

type Util struct{}

func (this *Util) IntMax() int {
	if runtime.GOARCH == "amd64" {
		return 2 ^ 64 - 1
	}
    
	return 2 ^ 32 - 1
}
func (this *Util) IntMin() int {
	if runtime.GOARCH == "amd64" {
		return -2 ^ 64
	}
	return -2 ^ 32

}
func (this *Util) ErrorToString(error consts.ErrorType) string {
	switch error {
	case consts.ET_None:
		break
	case consts.ET_Null:
		return "#NULL!"
	case consts.ET_Div0:
		return "#DIV0!"
	case consts.ET_Value:
		return "#VALUE!"
	case consts.ET_Ref:
		return "#REF!"
	case consts.ET_Name:
		return "#NAME?"
	case consts.ET_NA:
		return "#N/A"
	case consts.ET_Num:
		return "#NUM!"
	default:
		break
	}
	return string(error)
}
func initIndexHelper() {
	indexHelper = new(IndexHelper)
	indexHelper.columnInA1Letter = make(map[int]string)
	for i := 0; i < 26; i++ {
		indexHelper.columnInA1Letter[i] = _a1LetterCache[i]
	}
}

type IndexHelper struct {
	MaxRowCouunt     int
	MaxColumnCount   int
	columnInA1Letter map[int]string
}

var (
	_a1LetterCache = []string{"A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z"}
)

func (this *IndexHelper) GetRowIndexInNumber(s string) int {
	if len(s) == 0 {
		return 0
	}
	row := 0
	n := 0
	for _, r := range s {
		if unicode.IsDigit(r) {
			n++
		}
	}

	if n < len(s) {
		row, _ = strconv.Atoi(s[n:])
	}
	return row - 1
}

func (this *IndexHelper) GetColumnIndexInNumber(s string) int {
	i := 0
	endIndex := len(s)
	column := 0
	if i < endIndex {
		c := s[i]
		for {
			value := c - 'a'
			if value > 25 {
				break
			}
			if value < 0 {
				value = c - 'A'
			}
			if value < 0 {
				break
			}
			column = 26*column + int(value) + 1
			i++
			if i < endIndex {
				c = s[i]
			} else {
				break
			}

		}

	}
	column--
	return column
}
func (this *IndexHelper) GetRowIndexInNumber1(buffer []byte, start, end int) int {
	row := 0
	for start < end {
		c := buffer[start]
		if c >= '0' && c <= '9' {
			row = row*10 + int((c - '0'))
		}
		start++
	}
	return row
}
func (this *IndexHelper) GetColumnIndexInNumber1(buffer []byte, start, end int) int {
	column := 0
	for start < end {
		c := buffer[start]
		var value int
		value = int(c - 'A')
		if value < 0 {
			break
		} else if value > 25 {
			value = int(c - 'a')
			if value < 0 || value > 25 {
				break
			}
		}
		column = column*26 + value + 1
		start++

	}
	column--
	return column
}
func (this *IndexHelper) GetColumnIndexInA1Letter(coord int) string {
	return stringUtil.Number2String(coord)
}

type StringUtil struct{}

var Alpha []byte = []byte("ABCDEFGHIJKLMNOPQRSTUVWXYZ")

func (this *StringUtil) Number2String(number int) string {
	buf := bytes.NewBufferString("")
	index := 0
	for {
		index = number % 26
		buf.WriteString(string(Alpha[index]))
		number = number / 26
		if number == 0 {
			number--
		}
		return this.Reverse(buf.String())
	}
}
func (this *StringUtil) Reverse(input string) string {
    r := []rune(input)
	for i, j := 0, len(r)-1; i < len(r)/2; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return string(r)
}
func (this *StringUtil)LowerFirstChar(str string)string  {
    return strings.ToLower(string(str[0])) +  string(str[1:])
}

func (this *StringUtil)String2CellRect(str string,rect *CellRect)  {
    rc := str[:]
    if len(rc) == 1 {
        r := indexHelper.GetRowIndexInNumber(string(rc[0]))
        c := indexHelper.GetColumnIndexInNumber(string(rc[0]))
        rect.Row = r
        rect.Column = c
        rect.RowCount = 1
        rect.ColumnCount = 1
    }else if len(rc) == 2 {
        r1:= indexHelper.GetRowIndexInNumber(string(rc[0]))
        r2:= indexHelper.GetRowIndexInNumber(string(rc[1]))
        c1 := indexHelper.GetColumnIndexInNumber(string(rc[0]))
        c2 := indexHelper.GetColumnIndexInNumber(string(rc[1]))
        
        rect.Row = int(math.Min(float64(r1),float64(r2)))
        rect.Column = int(math.Min(float64(c1),float64(c2)))
        rect.RowCount = int(math.Abs(float64(r2-r1))) + 1
        rect.ColumnCount =  int(math.Abs(float64(c2)-float64(c1))) + 1
    }
    
}

type DoubleComparer struct{}

func (this *DoubleComparer)IsGreaterThan(x,y float64) bool  {
    return x>y && math.Abs(x-y)>math.SmallestNonzeroFloat64
}
func (this *DoubleComparer)IsLessThan(x,y float64) bool  {
    return x<y && math.Abs(x-y)>math.SmallestNonzeroFloat64
}
func (this *DoubleComparer)IsEqualsTo(x,y float64) bool  {
    return x==y || math.Abs(x-y)<=math.SmallestNonzeroFloat64
}

package core
const
(
    EmptyString  = ""
)
type Nullable struct{
    Value interface{}
}

func (this *Nullable)HasValue() bool  {
    return (this.Value != nil)
}
func (this *Nullable) ToInt() (int,bool)  {
    s,ok := this.Value.(int)
    if ok{
        return s,true
    }else{
        return -1,false
    }
}
func (this *Nullable) ToInt16() (int16,bool)  {
    s,ok := this.Value.(int16)
    if ok{
        return s,true
    }else{
        return -1,false
    }
}
func (this *Nullable) ToInt32() (int32,bool)  {
    s,ok := this.Value.(int32)
    if ok{
        return s,true
    }else{
        return -1,false
    }
}
func (this *Nullable) ToInt64() (int64,bool)  {
    s,ok := this.Value.(int64)
    if ok{
        return s,true
    }else{
        return -1,false
    }
}
func (this *Nullable) ToFloat32() (float32,bool)  {
    s,ok := this.Value.(float32)
    if ok{
        return s,true
    }else{
        return -1,false
    }
}
func (this *Nullable) ToFloat64() (float64,bool)  {
    s,ok := this.Value.(float64)
    if ok{
        return s,true
    }else{
        return -1,false
    }
}
func (this *Nullable) ToBool() (bool,bool)  {
    s,ok := this.Value.(bool)
    if ok{
        return s,true
    }else{
        return false,false
    }
}

type NullBool struct{
    hasValue bool
    value bool
}
func (this NullBool)SetValue(v bool) NullBool {
    this.value = v
    this.hasValue = true
    return this
}
func (this NullBool)HasValue() bool  {
    return this.hasValue
}
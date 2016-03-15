package core
import
(
    "time"
)
const
(
    TicksPerMillisecond = 100000
    TicksPerSecond = TicksPerMillisecond * 1000
    TicksPerMinute = TicksPerSecond * 60
    TicksPerHour = TicksPerMinute * 60
    TicksPerDay = TicksPerHour * 24
    // Number of milliseconds per time unit
    MillisPerSecond = 1000
    MillisPerMinute = MillisPerSecond * 60
    MillisPerHour = MillisPerMinute * 60
    MillisPerDay = MillisPerHour * 24
    // Number of days in a non-leap year
    DaysPerYear = 365
    // Number of days in 4 years
    DaysPer4Years = DaysPerYear * 4 + 1
    // Number of days in 100 years
    DaysPer100Years = DaysPer4Years * 25 - 1
    // Number of days in 400 years
    DaysPer400Years = DaysPer100Years * 4 + 1

    // Number of days from 1/1/0001 to 12/31/1600
    DaysTo1601 = DaysPer400Years * 4
    // Number of days from 1/1/0001 to 12/30/1899
    DaysTo1899 = DaysPer400Years * 4 + DaysPer100Years * 3 - 367
    // Number of days from 1/1/0001 to 12/31/9999
    DaysTo10000 = DaysPer400Years * 25 - 366

    MinTicks = 0
    MaxTicks = DaysTo10000 * TicksPerDay - 1
    MaxMillis = DaysTo10000 * MillisPerDay

    FileTimeOffset = DaysTo1601 * TicksPerDay
    DoubleDateOffset = DaysTo1899 * TicksPerDay
    // The minimum OA date is 0100/01/01 (Note it's year 100).
    // The maximum OA date is 9999/12/31
    OADateMinAsTicks = (DaysPer100Years - DaysPerYear) * TicksPerDay
    // All OA dates must be greater than (not >=) OADateMinAsDouble
    OADateMinAsDouble = -657435.0
    // All OA dates must be less than (not <=) OADateMaxAsDouble
    OADateMaxAsDouble = 2958466.0 
)
type TimeEx struct{
    time.Time
}

func (this *TimeEx)Ticks()  int64{
    return int64(this.UnixNano()/100)
}
func (this *TimeEx)FromOADate(tickets int64) *time.Time  {
    t:=time.Unix(0,tickets*100)
    return &t
}
func (this *TimeEx)TOOADate(d TimeEx) int64 {
    v:= d.Ticks()
    if v == 0{
        return 0
    }
    if v<TicksPerDay{
        v += DoubleDateOffset
    }
    if v<OADateMinAsTicks{
        panic("Arg_OleAutDateInvalid")
    }
    millis := (v - DoubleDateOffset)/TicksPerMillisecond
    if millis<0{
        frac := millis % MillisPerDay
        if frac!=0{
            millis -=(MillisPerDay +frac)*2
        }
       
    }
     return millis/MillisPerDay
}
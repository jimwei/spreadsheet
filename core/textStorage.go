package core
import
(
    "math"
)
var
(
    _primes = [...]int{ 3, 7, 11, 0x11, 0x17, 0x1d, 0x25, 0x2f, 0x3b, 0x47, 0x59, 0x6b, 0x83, 0xa3, 0xc5, 0xef, 
        0x125, 0x161, 0x1af, 0x209, 0x277, 0x2f9, 0x397, 0x44f, 0x52f, 0x63d, 0x78b, 0x91d, 0xaf1, 0xd2b, 0xfd1, 0x12fd, 
        0x16cf, 0x1b65, 0x20e3, 0x2777, 0x2f6f, 0x38ff, 0x446f, 0x521f, 0x628d, 0x7655, 0x8e01, 0xaa6b, 0xcc89, 0xf583, 0x126a7, 0x1619b, 
        0x1a857, 0x1fd3b, 0x26315, 0x2dd67, 0x3701b, 0x42023, 0x4f361, 0x5f0ed, 0x72125, 0x88e31, 0xa443b, 0xc51eb, 0xec8c1, 0x11bdbf, 0x154a3f, 0x198c4f, 
        0x1ea867, 0x24ca19, 0x2c25c1, 0x34fa1b, 0x3f928f, 0x4c4987, 0x5b8b6f, 0x6dda89}
        
      
        
)
type Entry struct{
    HashCode uint32
    Next int32
    Pointer int32
    Length int16
    RefCount int16
}
type TextStorage struct{
    _charBlock []byte
    _lastCharIndex int32
    _entries []*Entry
    _lastEntryIndex int32
    _freeList int32
}

func NewTextStorage() *TextStorage  {
    return NewTextStorage1(0)
}
func NewTextStorage1(capacity int)  *TextStorage{
    prime := getPrime(capacity)
    
}
func getPrime(min int)  int{
    for i:=0;i<len(_primes);i++{
        num2:=_primes[i]
        if num2>=min{
            return num2
        }
    }
    for j := min|1;j < 0x7fffffff;j +=2{
        if isPrime(j) && (((j - 1) % 0x65) != 0){
            return j
        }
    }   
    return min
}
func isPrime(cadidate int)  bool{
    if (cadidate & 1) == 0{
        return cadidate == 2
    }
    num := int(math.Sqrt(float64(cadidate)))
    for i := 0;i < num; i+=2{
        if cadidate%2 == 0{
            return false
        }
    }
    return true
}
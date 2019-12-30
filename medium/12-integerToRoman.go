var romanMap map[int]string = map[int]string {
    1000:   "M",
    900:    "CM",
    500:    "D",
    400:    "CD",
    100:    "C",
    90:     "XC",
    50:     "L",
    40:     "XL",
    10:     "X",
    9:      "IX",
    5:      "V",
    4:      "IV",
    1:      "I",
}

var keys []int = make([]int, 0)

func init() {
    // Need to sort the keys, since looping over the map
    // doesn't garauntee any specific order
    for key := range(romanMap) {
        keys = append(keys, key)
    }
    sort.Ints(keys)
}

func intToRoman(num int) string {    
    var romanRep strings.Builder
    for num > 0 {
        for i := len(keys)-1; i >= 0; i-- {
            key := keys[i]
            if num >= key {
                // Get the roman numeral for this key
                val := romanMap[key]
                
                romanRep.Write([]byte(val))
                num = num - key
                
                // Start back at the largest key
                break
            }
        }
    }
    
    return romanRep.String()
}

package src

func countCharacters(words []string, chars string) int {
    out := 0
    for _,v := range words {
        times := 0
        hash := make(map[string]int)
        for _,c := range chars{
            if _,ok := hash[string(c)]; !ok {
                hash[string(c)] = 1
            } else {
                hash[string(c)]++
            }
        }
        
        for _,val := range v {
            if _,has := hash[string(val)]; has && hash[string(val)] > 0{
                times++
                hash[string(val)] --
            }
        }
        if times == len(v) {
            out += len(v)
        }
    }
    return out
}

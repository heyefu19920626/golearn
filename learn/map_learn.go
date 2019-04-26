package main

import "fmt"


func main(){
    var countryMap map[string]string
    countryMap = make(map[string]string)

    countryMap["zh"] = "北京"
    countryMap["jp"] = "东京"
    fmt.Println(countryMap)

    trverseMap(countryMap)

    captital, ok := countryMap["amr"]
    fmt.Println(ok)
    fmt.Println(captital)

    delete(countryMap, "jp")
    trverseMap(countryMap)
}

func trverseMap(target_map map[string]string){
    for key := range target_map{
        fmt.Printf("key: %s, value: %s\n", key, target_map[key])
    }
}
package main


import ("fmt"
        "encoding/json"
        "os")

func main() {
    writeData()

}

func writeData() {
	map1 := make(map[string]string)
	map1["a"] = "111"
	map1["b"] = "222"
	map1["c"] = "333"
       dataBytes, err := json.Marshal(map1)
       if err != nil {
           fmt.Println(err)
       }


    fmt.Println("writing to the database")
    path := "/foo/readData/collections/write1.json"
    file, err := os.OpenFile(path, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0755)
    if err != nil {
        fmt.Println("error while trying to write to json files", err)
    }
    _, err2 := file.Write(dataBytes)
    if err2 != nil {
        fmt.Println(err2)
    }

}

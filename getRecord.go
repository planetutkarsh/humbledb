package main


import ("fmt"
        "os"
        "io/ioutil"
        "encoding/json"
  )


type dataMap map[string]interface{}

type dataStructList struct {
	dataStructItem []dataStruct  `json:"item`

}

type dataStruct struct {
     Id   string `json:"id"`
     Name string `json:"name"`
     Address string `json:"address"`
}

func main() {
    filterParams := make(map[string]string)
    filterParams["name"] = "utkarsh"
    filterParams["address"] = "russia"
    getMatchingRecords(filterParams)

}

func loadCollectionData(){



}

func getMatchingRecords(filterParams map[string]string) map[string]interface{} {

    var  dataStoreArr[]dataMap
    //read all the files in the current dir
    os.Setenv("collectorPath", "/foo/readData/collections")
    fileDir := os.Getenv("collectorPath")
    files, err := ioutil.ReadDir(fileDir)
    if err != nil {
        fmt.Println(err)
    }

    for _, r :=  range(files) {
        // read the json file in the current directory
	//dataStore := dataStruct{}
	fmt.Println(r.Name())
	readFile := fmt.Sprintf("%s/%s", fileDir, r.Name())
	fmt.Println(readFile)
	fileDataByte, err1 := ioutil.ReadFile(readFile)
	if err1 != nil {
	    fmt.Println(err1)
	}
	//dataStructItems := da1taStructList{}
	//dataMap := make(map[string]interface{})
        err := json.Unmarshal(fileDataByte, &dataStoreArr)
	if err != nil {
	    fmt.Println(err)
	}
	fmt.Println(dataStoreArr)
        }

	//check if the key exists and then only take the action
	resMap := make(map[string]interface{})
	for _, dataCollection := range dataStoreArr {
		for filter := range filterParams {
		    fmt.Println(filter)
		    _, ok := dataCollection[filter]
		    if !ok {
		        fmt.Println("the filter is not present  in the data")
		    }
		    _, ok2 := resMap[dataCollection["id"].(string)]
		    if !ok2 {
		        resMap[dataCollection["id"].(string)] = dataCollection
		    }

		}
	}
        fmt.Println(resMap)
	return resMap
    }

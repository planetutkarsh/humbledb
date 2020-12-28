package main


import ("fmt")

func main() {
    colName := "Address"
    createIndex(colName)
}

func createIndex(colName) {
    rootNode := Node{Val: "M", Left:nil, Right:nil}
    collectionDataSet := loadCollectionData()
    for x := range collectionDataSet {
	aNode := Node{Val: x[colName], Left: nil, Right: nil}
        insertNodex(aNode, &rootNode)
    }
}

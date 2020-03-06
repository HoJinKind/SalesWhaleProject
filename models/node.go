package models

//import (
//	"SalesWhaleProject/utils"
//	"fmt"
//)

type Node struct {
	Id            string
	Value         string
	AdjacentArray []string
}

func CreateNode(id string, val string, adjacentArray []string) Node {
	return Node{
		Id:            id,
		Value:         val,
		AdjacentArray: adjacentArray,
	}
}

func InsertNewNodeToDB(node Node) {
	db := openConnection()
	statement, _ := db.Prepare("INSERT INTO node (id,token, end_epoch,duration,board_array,points, node_array) VALUES (?,?,?)")
	statement.Exec(node.Id, node.Value, node.AdjacentArray)
}

func GetNodeCount() int {
	count := 0
	db := openConnection()
	row := db.QueryRow("SELECT COUNT(*) FROM node")
	row.Scan(&count)
	return count
}

func GetNode(id string) Node {
	var node Node
	db := openConnection()
	row := db.QueryRow("SELECT * FROM boards WHERE id=?", id)
	row.Scan(&node.Id, &node.Value, &node.AdjacentArray)
	return node
}

//func NodeHandler(board string) string {
//
//	arrayOfNodes := utils.SplitString(board)
//
//	offset:= GetNodeCount()
//	var nodeArray []Node
//	for letter , _ := range(nodeArrayLetter) {
//		tempNode = Node{
//			Id:index
//		}
//		nodeArrayLetter = append(nodeArray)
//	}
//	for node, index := range arrayOfNodes {
//		fmt.Println(node, index)
//	}
//	return "1"
//}

func getNodeId(index int,width int, direction string) int {
	switch direction {
	case "left":
		return index - 1
	case "right":
		return index + 1
	case "up":
		return index - width
	case "down":
		return index + width
	case "leftup":
		return index - width - 1
	case "rightup":
		return index - width + 1
	case "leftdown":
		return index + width - 1
	case "rightdown":
		return index + width + 1
	default:
		return -1
	}

}

//func processNodeAdjacencyArray(index,width int, nodeArrayLetter []string) {
//
//	var nodeInQuestion Node
//	var nodeAdjacentArray []int
//	if RightCheck(index, width) {
//		getNodeId(index,width)
//	}
//
//}

func RightCheck(index int, width int) bool {
	if index/width == (index+1)/width {
		return true
	}
	return false
}

func LeftCheck(index int, width int) bool {
	if index/width == (index-1)/width && index != 0 {
		return true
	}
	return false
}

func UpCheck(index int, width int) bool {
	if index-width >= 0 {
		return true
	}
	return false
}

func DownCheck(index int, width int) bool {
	if index+width < width*width - 1 {
		return true
	}
	return false
}

func LeftUpCheck(index int, width int) bool {
	return LeftCheck(index, width) && UpCheck(index, width)
}

func LeftDownCheck(index int, width int) bool {
	return LeftCheck(index, width) && DownCheck(index, width)
}

func RightUpCheck(index int, width int) bool {
	return RightCheck(index, width) && UpCheck(index, width)
}

func RightDownCheck(index int, width int) bool {
	return RightCheck(index, width) && UpCheck(index, width)
}

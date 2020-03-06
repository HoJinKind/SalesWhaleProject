package algo

import (
	"SalesWhaleProject/utils"
	"math"
)

func CheckDimension(size int) int {
	return int(math.Sqrt(float64(size)))
}


func BoggleSolver( boardString string, playString string) bool{

	arrayOfNodes := utils.SplitString(boardString)
	for index , _  := range(arrayOfNodes) {
		var usedArray []int
		if traverseBoggle(playString, arrayOfNodes,0, index, CheckDimension(len(arrayOfNodes)),usedArray){
			return true
		}
	}
	return false
}

func traverseBoggle(originalString string, boggleString []string, indexOfOriginal int, currentPosition int,width int ,usedArray []int ) bool {
	if RightCheck(currentPosition,width ) {
			nodeId:= getNodeId(currentPosition,width,"right")
		if !utils.Contains(usedArray,nodeId) && checkCharMatches(string(originalString[indexOfOriginal]),string(boggleString[currentPosition])){
			usedArray := append(usedArray,currentPosition)
			indexOfOriginal +=1
			println("the one that shld have pass:", len(originalString)-1==indexOfOriginal)
			if ( len(originalString)==indexOfOriginal){
				return true
			}
			traverseBoggle(originalString,boggleString,indexOfOriginal, nodeId,width, usedArray )
		}
	}

	if LeftCheck(currentPosition,width ) {
		nodeId:= getNodeId(currentPosition,width,"left")
		if !utils.Contains(usedArray,nodeId) && checkCharMatches(string(originalString[indexOfOriginal]),string(boggleString[currentPosition])){
			usedArray := append(usedArray,currentPosition)
			indexOfOriginal +=1
			println("the one that shld have pass:", len(originalString)-1==indexOfOriginal)
			if ( len(originalString)-1==indexOfOriginal){
				return true
			}
			traverseBoggle(originalString,boggleString,indexOfOriginal, nodeId,width, usedArray )
		}
	}

	if UpCheck(currentPosition,width ) {
		nodeId:= getNodeId(currentPosition,width,"up")
		if !utils.Contains(usedArray,nodeId) && checkCharMatches(string(originalString[indexOfOriginal]),string(boggleString[currentPosition])){
			usedArray := append(usedArray,currentPosition)
			indexOfOriginal +=1
			println("the one that shld have pass:", len(originalString)-1==indexOfOriginal)
			if ( len(originalString)-1==indexOfOriginal){
				return true
			}
			traverseBoggle(originalString,boggleString,indexOfOriginal, nodeId,width, usedArray )
		}
	}

	if DownCheck(currentPosition,width ) {
		nodeId:= getNodeId(currentPosition,width,"down")
		if !utils.Contains(usedArray,nodeId) && checkCharMatches(string(originalString[indexOfOriginal]),string(boggleString[currentPosition])){
			usedArray := append(usedArray,currentPosition)
			indexOfOriginal +=1
			println("the one that shld have pass:", len(originalString)-1==indexOfOriginal)
			if ( len(originalString)-1==indexOfOriginal){
				return true
			}
			traverseBoggle(originalString,boggleString,indexOfOriginal, nodeId,width, usedArray )
		}
	}

	if LeftUpCheck(currentPosition,width ) {
		nodeId:= getNodeId(currentPosition,width,"leftup")
		if !utils.Contains(usedArray,nodeId) && checkCharMatches(string(originalString[indexOfOriginal]),string(boggleString[currentPosition])){
			usedArray := append(usedArray,currentPosition)
			indexOfOriginal +=1
			println("the one that shld have pass:", len(originalString)-1==indexOfOriginal)
			if ( len(originalString)-1==indexOfOriginal){
				return true
			}
			traverseBoggle(originalString,boggleString,indexOfOriginal, nodeId,width, usedArray )
		}
	}

	if RightUpCheck(currentPosition,width ) {
		nodeId:= getNodeId(currentPosition,width,"rightup")
		if !utils.Contains(usedArray,nodeId) && checkCharMatches(string(originalString[indexOfOriginal]),string(boggleString[currentPosition])){
			usedArray := append(usedArray,currentPosition)
			indexOfOriginal +=1
			println("the one that shld have pass:", len(originalString)-1==indexOfOriginal)
			if ( len(originalString)-1==indexOfOriginal){
				return true
			}
			traverseBoggle(originalString,boggleString,indexOfOriginal, nodeId,width, usedArray )
		}
	}


	if LeftDownCheck(currentPosition,width ) {
		nodeId:= getNodeId(currentPosition,width,"leftdown")
		if !utils.Contains(usedArray,nodeId) && checkCharMatches(string(originalString[indexOfOriginal]),string(boggleString[currentPosition])){
			usedArray := append(usedArray,currentPosition)
			indexOfOriginal +=1
			println("the one that shld have pass:", len(originalString)-1==indexOfOriginal)
			if ( len(originalString)-1==indexOfOriginal){
				return true
			}
			traverseBoggle(originalString,boggleString,indexOfOriginal, nodeId,width, usedArray )
		}
	}

	if RightDownCheck(currentPosition,width ) {
		nodeId:= getNodeId(currentPosition,width,"rightdown")
		if !utils.Contains(usedArray,nodeId) && checkCharMatches(string(originalString[indexOfOriginal]),string(boggleString[currentPosition])){
			usedArray := append(usedArray,currentPosition)
			indexOfOriginal +=1
			println("the one that shld have pass:", len(originalString)-1==indexOfOriginal)
			if ( len(originalString)-1==indexOfOriginal){
				return true
			}
			traverseBoggle(originalString,boggleString,indexOfOriginal, nodeId,width, usedArray )
		}
	}
	return false
}

func checkCharMatches(string1  string, string2 string) bool {
	return string1 == string2 || string2 == "*"
}

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

func processNodeAdjacencyArray(index,width int, nodeArrayLetter []string) {

}

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

package models

import (
	//
	"fmt"
	"strconv"
	"time"
)

type boggleBoardInfo struct {
	BoggleBoard
	TimeLeft int `json:"time_left"`
	Points   int `json:"points"`
}

type BoggleBoard struct {
	Id        string `json:"id"`
	Token     string `json:"token"`
	EndEpoch  int    `json:"-"`
	Duration  int    `json:"duration"`
	Board     string `json:"board"`
	Points    int    `json:"-"`
	nodeArray string `json:"-"`
}

//
//"id": 1,
//"token": "9dda26ec7e476fb337cb158e7d31ac6c",
//"duration": 12345,
//"board": "A, C, E, D, L, U, G, *, E, *,Fpeoiple H, T, G, A, F, K"

func CreateBoard(Id string, Token string, Duration int, Board string) BoggleBoard {
	return BoggleBoard{
		Id:        Id,
		Token:     Token,
		EndEpoch:  int(time.Now().Unix()) + Duration,
		Duration:  Duration,
		Board:     Board,
		Points:    0,
		nodeArray: "boilerPlate",
	}
}

func InsertNewBoardToDB(boggleBoard BoggleBoard) bool {

	db := openConnection()
	statement, _ := db.Prepare("INSERT INTO boards (id,token, end_epoch,duration,board_array,points, node_array) VALUES (?,?,?,?,?,?,?)")
	statement.Exec(boggleBoard.Id, boggleBoard.Token, boggleBoard.EndEpoch, boggleBoard.Duration, boggleBoard.Board, boggleBoard.Points, boggleBoard.nodeArray)
	return true
}

func GetBoard(id string) (*boggleBoardInfo, error) {

	var boggleBoardInfo boggleBoardInfo
	requestId, _ := strconv.Atoi(id)
	if  ( GetBoardsCount() < requestId || requestId <=0 ) {
		return nil,  fmt.Errorf(" invalid id")
	}
	db := openConnection()
	row := db.QueryRow("SELECT * FROM boards WHERE id=?", id)
	row.Scan(&boggleBoardInfo.Id, &boggleBoardInfo.Token, &boggleBoardInfo.EndEpoch, &boggleBoardInfo.Duration, &boggleBoardInfo.Board, &boggleBoardInfo.Points, &boggleBoardInfo.nodeArray)
	if boggleBoardInfo.EndEpoch-int(time.Now().Unix()) > 0 {
		boggleBoardInfo.TimeLeft = boggleBoardInfo.EndEpoch - int(time.Now().Unix())
	} else {
		boggleBoardInfo.TimeLeft = 0
	}
	return &boggleBoardInfo, nil
}

func GetBoardsCount() int {
	count := 0

	db := openConnection()
	row := db.QueryRow("SELECT COUNT(*) FROM boards")

	row.Scan(&count)
	return count
}

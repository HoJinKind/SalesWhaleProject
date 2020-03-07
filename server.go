package main

import (
	"SalesWhaleProject/models"
	_ "SalesWhaleProject/models"
	"SalesWhaleProject/utils"
	_ "SalesWhaleProject/utils"
	"encoding/json"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"net/http"
	. "strconv"
)


type CreateBoardField struct {
	Duration int `json:"duration"`
	Random bool `json:"random"`
	Board string `json:"board,omitempty"`
}

type getBoardField struct {
	Id string `json:"id"`
}

func main() {
		models.InitDB()
		http.HandleFunc("/games/", gameEndpoint)
		log.Fatal(http.ListenAndServe(":4567", nil))
	}


func gameEndpoint(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		decoder := json.NewDecoder(r.Body)
		var getBoardField getBoardField
		err := decoder.Decode(&getBoardField)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		getBoard(getBoardField.Id,w)

	case http.MethodPost:

		decoder := json.NewDecoder(r.Body)
		var createBoardField CreateBoardField
		err := decoder.Decode(&createBoardField)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		id := models.GetBoardsCount()
		if createBoardField.Random{
			createBoard(Itoa(id),utils.TokenGenerator(Itoa(id)), createBoardField.Duration,utils.DefaultBoard, w)
		} else {
			createBoard(Itoa(id),utils.TokenGenerator(Itoa(id)), createBoardField.Duration, createBoardField.Board, w)

		}
	case http.MethodPut:
	default:
		http.Error(w, "invalid method: Use Post, Get, Put", http.StatusBadRequest)
		return
	}

}

func createBoard(id string, token string, duration int, board string, w http.ResponseWriter) {

	boggleBoard := models.CreateBoard(id, token,duration, board)
	models.InsertNewBoardToDB(boggleBoard)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	boggleBoard.EndEpoch = duration
	boardJson, _ := json.Marshal(boggleBoard)
	w.Write(boardJson)
}

func getBoard(id string, w http.ResponseWriter){
	boggleBoardInfo, err := models.GetBoard(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	boardJson, _ := json.Marshal(boggleBoardInfo)
	w.Write(boardJson)
}
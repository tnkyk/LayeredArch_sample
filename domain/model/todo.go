package model

import (
	"time"
)

//Domain層 はシステムが扱う業務領域に関するコードを置くところ
//実装時最初に着手

type Todo struct {
	Id    string
	Title string
	//Author    string
	CreatedAt time.Time
}

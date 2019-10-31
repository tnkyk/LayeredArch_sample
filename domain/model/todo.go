package model

import (
	"time"
)

//Domain層 はシステムが扱う業務領域に関するコードを置くところ
//実装時最初に着手
//modelPackageは

type Todo struct {
	Id        int
	Title     string
	Author    string
	CreatedAt time.Time
}

package papermodel

import (
	"github.com/globalsign/mgo/bson"
	"time"
)

var mongourl = "mongodb://admin:123456@192.168.217.129:27018/"

const unpublished = 0
const published = 1

type SelfPaper struct {
	ID         bson.ObjectId `bson:"_id"`
	PaperName  string        `bson:"paperName"`
	CreaterId  int64         `bson:"createrId"`
	Status     int64         `bson:"status"`
	StartTime  time.Time     `bson:"startTime"`
	StopTime   time.Time     `bson:"stopTime"`
	Version    string        `bson:"version"`
	CreateTime time.Time     `bson:"createTime"`
	UpdateTime time.Time     `bson:"updateTime"`
	Deleted    bool          `bson:"deleted"`

	SignalChoiceBig SignalChoiceBig `bson:"signalChoiceBigTitle"`
}

type SignalChoiceBig struct {
	Title                string                `bson:"Title"`
	SignalChoiceWithNums []SignalChoiceWithNum `bson:"signalChoiceWithNums"`
}

type SignalChoiceWithNum struct {
	Number         string `bson:"Number"`
	SignalChoiceId string `bson:"signalChoiceId"`
}

const unmark = 0 //未阅卷
const marked = 1 //已经阅卷

type AnswerPaper struct {
	ID                  bson.ObjectId        `bson:"_id"`
	CreaterId           int64                `bson:"createrId"`
	SubjectId           int64                `bson:"subjectId"`
	SelfPaperId         string               `bson:"selfPaperId"`
	Mark                int64                `bson:"mark"`
	Version             string               `bson:"version"`
	StartTime           time.Time            `bson:"startTime"`
	StopTime            time.Time            `bson:"stopTime"`
	Deleted             bool                 `bson:"deleted"`
	CreateTime          time.Time            `bson:"createTime"`
	UpdateTime          time.Time            `bson:"updateTime"`
	SignalChoiceAnswers []SignalChoiceAnswer `bson:"SignalChoiceAnswers"`
}

type SignalChoiceAnswer struct {
	Score          int64  `bson:"score"`
	IsCorrect      bool   `bson:"isCorrect"`
	SignalChoiceId string `bson:"signalChoiceId"`
}

type SignalChoice struct {
	ID            bson.ObjectId `bson:"_id"`
	Title         string        `bson:"title"`
	AAnswer       string        `bson:"aAnswer"`
	BAnswer       string        `bson:"bAnswer"`
	CAnswer       string        `bson:"cAnswer"`
	DAnswer       string        `bson:"dAnswer"`
	EAnswer       string        `bson:"eAnswer"`
	FAnswer       string        `bson:"fAnswer"`
	CorrectAnswer string        `bson:"correctAnswer"`
	Version       string        `bson:"version"`
	CreateTime    time.Time     `bson:"createTime"`
	UpdateTime    time.Time     `bson:"updateTime"`
	Deleted       bool          `bson:"deleted"`
}

package papermodel

import (
	"github.com/globalsign/mgo/bson"
	"time"
)

var mongourl = "mongodb://admin:123456@192.168.217.129:27018/"

const unpublished = 0
const published = 1

type SelfPaper struct {
	Id             bson.ObjectId   `bson:"_id"`
	PaperName      string          `bson:"paperName"`
	CreaterId      int64           `bson:"createrId"`
	Status         int64           `bson:"status"`
	StartTime      time.Time       `bson:"startTime"`
	StopTime       time.Time       `bson:"stopTime"`
	Version        string          `bson:"version"`
	CreateTime     time.Time       `bson:"createTime"`
	UpdateTime     time.Time       `bson:"updateTime"`
	Deleted        bool            `bson:"deleted"`
	DeleteTime     time.Time       `bson:"deleteTime"`
	PaperItems     []PaperItem     `bson:"paperItems"`
	RandomSettings []RandomSetting `bson:"randomSetting"`
}

type ParagraphInstruction struct {
	Instruction string `bson:"instruction"`
	Type        string `bson:"type"`
}

type PageBar struct {
	Instruction string `bson:"instruction"`
	CurrentPage int64  `bson:"currentPage"`
	TotalPage   int64  `bson:"totalPage"`
	Type        string `bson:"type"`
}

type RandomSetting struct {
	StartNumber  int64 `bson:"startNumber"`
	EndNumber    int64 `bson:"endNumber"`
	SubjectCount int64 `bson:"subjectCount"`
}

//type PaperItem struct {
//	papermodel.PageBar
//	Type string `bson:"type"`
//}

//type PaperItem struct {
//	Type                   string                  `bson:"type"`
//	PageBar                *PageBar                `bson:"pageBar,omitempty"`
//	ParagraphInstruction   *ParagraphInstruction   `bson:"paragraphInstruction,omitempty"`
//	SubjectNumberAndIdPair *SubjectNumberAndIdPair `bson:"subjectNumberAndIdPair,omitempty"`
//}

type PaperItem struct {
	Type string      `bson:"type"`
	Item interface{} `bson:"item"`
	//ParagraphInstruction   *ParagraphInstruction   `bson:"paragraphInstruction,omitempty"`
	//SubjectNumberAndIdPair *SubjectNumberAndIdPair `bson:"subjectNumberAndIdPair,omitempty"`
}

//type BigSignalChoice struct {
//	SubjectNumber        int                   `bson:"subjectNumber"`
//	Title                string                `bson:"title"`
//	Type                 string                `bson:"type"`
//	SignalChoiceWithNums []SignalChoiceWithNum `bson:"signalChoiceWithNums"`
//}

//type SignalChoiceWithNum struct {
//	SubjectNumber  int64    `bson:"subjectNumber"`
//	SignalChoiceId string `bson:"signalChoiceId"`
//	Type           string `bson:"type"`
//}

type SubjectNumberAndIdPair struct {
	SubjectNumber int64  `bson:"subjectNumber"`
	SubjectId     string `bson:"subjectId"`
	Type          string `bson:"type"`
}

const unmark = 0 //未阅卷
const marked = 1 //已经阅卷

type AnswerPaper struct {
	Id                  bson.ObjectId        `bson:"_id"`
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
	SignalChoiceAnswers []SignalChoiceAnswer `bson:"signalChoiceAnswers"`
}

type SignalChoiceAnswer struct {
	Score          int64  `bson:"score"`
	IsCorrect      bool   `bson:"isCorrect"`
	SignalChoiceId string `bson:"signalChoiceId"`
}

type SignalChoice struct {
	Id            bson.ObjectId `bson:"_id"`
	Type          string        `bson:"type`
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
	DeleteTime    time.Time     `bson:"deleteTime"`
}

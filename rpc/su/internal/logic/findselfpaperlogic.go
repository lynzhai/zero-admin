package logic

import (
	"context"
	"github.com/globalsign/mgo/bson"
	"github.com/tal-tech/go-zero/core/logx"
	"go-zero-admin/rpc/model/papermodel"
	"go-zero-admin/rpc/su/internal/svc"
	"go-zero-admin/rpc/su/su"
	"go-zero-admin/shared"
	"strconv"
)

type FindSelfPaperLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewFindSelfPaperLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FindSelfPaperLogic {
	return &FindSelfPaperLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *FindSelfPaperLogic) FindSelfPaper(in *su.FindSelfPaperReq) (*su.FindSelfPaperResp, error) {
	// todo: add your logic here and delete this line
	selfPaper, err := l.svcCtx.SelfPaperModel.FindOne(l.ctx, in.Id)
	if err != nil {
		return nil, err
	}
	suFindSelfPaperResp := su.FindSelfPaperResp{
		SelfPaper: &su.SelfPaperInfo{
			Id:         selfPaper.Id.Hex(),
			PaperName:  selfPaper.PaperName,
			CreaterId:  selfPaper.CreaterId,
			Status:     selfPaper.Status,
			StartTime:  selfPaper.StartTime.UTC().String(),
			StopTime:   selfPaper.StopTime.UTC().String(),
			Version:    selfPaper.Version,
			CreateTime: selfPaper.CreateTime.UTC().String(),
			UpdateTime: selfPaper.UpdateTime.UTC().String(),
			DeleteTime: selfPaper.DeleteTime.UTC().String(),
			Deleted:    selfPaper.Deleted,
		},
	}
	
	logx.Info("len(selfPaper.PaperItems):"+ strconv.Itoa(len(selfPaper.PaperItems)))

	for _, item := range selfPaper.PaperItems {
		switch item.Type {
		case SignalChoiceWithSerial:
			logx.Error("err:" + "SignalChoiceWithSerial type error in paper model")
			continue
		case ParagraphInstruction:
			//value,ok := itme.Item.(papermodel.ParagraphInstruction)
			logx.Error("ParagraphInstruction %v:" ,item.Item)
			valueMap,ok := item.Item.(bson.M)
			if !ok{
				logx.Error("err:" + "ParagraphInstruction type error in paper model")
				continue
			}
			value :=papermodel.ParagraphInstruction{}
			err := shared.BsonToStruct(valueMap, &value)
			if err != nil {
				logx.Error("err:" + "BsonToStruct error")
				continue
			}
			suParagraphInstruction := su.ParagraphInstruction{
				Instruction: value.Instruction,
				Type:        value.Type,
			}
			suPaperItem := su.PaperItem{
				Type:                 ParagraphInstruction,
				ParagraphInstruction: &suParagraphInstruction,
			}
			suFindSelfPaperResp.SelfPaper.PaperItems = append(suFindSelfPaperResp.SelfPaper.PaperItems, &suPaperItem)
		case PageBar:
			logx.Error("PageBar %v:" ,item.Item)
			valueMap,ok := item.Item.(bson.M)
			//value,ok := item.Item.(papermodel.PageBar)
			if !ok{
				logx.Error("err:" + "PageBar type error in paper model")
				continue
			}

			value :=papermodel.PageBar{}
			err := shared.BsonToStruct(valueMap, &value)
			if err != nil {
				logx.Error("err:" + "BsonToStruct error")
				continue
			}

			suPageBar := su.PageBar{
				Instruction: value.Instruction,
				CurrentPage: value.CurrentPage,
				TotalPage:   value.TotalPage,
				Type:        value.Type,
			}
			suPaperItem := su.PaperItem{
				Type:    PageBar,
				PageBar: &suPageBar,
			}
			suFindSelfPaperResp.SelfPaper.PaperItems = append(suFindSelfPaperResp.SelfPaper.PaperItems, &suPaperItem)
		case SubjectNumberAndIdPair:
			logx.Error("SubjectNumberAndIdPair %v:" ,item.Item)
			valueMap,ok := item.Item.(bson.M)
			if !ok{
				logx.Error("err:" + "SubjectNumberAndIdPair type error in paper model")
				continue
			}
			value := papermodel.SubjectNumberAndIdPair{}
			err := shared.BsonToStruct(valueMap, &value)
			if err != nil {
				logx.Error("err:" + "BsonToStruct error")
				continue
			}
			suSubjectNumberAndIdPair := su.SubjectNumberAndIdPair{
				SubjectNumber: value.SubjectNumber,
				SubjectId:     value.SubjectId,
				Type:          value.Type,
			}
			mSignalChoice, err := l.svcCtx.SignalChoiceModel.FindOne(l.ctx, suSubjectNumberAndIdPair.SubjectId)
			if err != nil {
				logx.Error("err:" + err.Error())
				continue
			}
			suSignalChoiceInfo := su.SignalChoiceInfo{
				Id:            mSignalChoice.Id.Hex(),
				Type:          mSignalChoice.Type,
				Title:         mSignalChoice.Title,
				AAnswer:       mSignalChoice.AAnswer,
				BAnswer:       mSignalChoice.BAnswer,
				CAnswer:       mSignalChoice.CAnswer,
				DAnswer:       mSignalChoice.DAnswer,
				EAnswer:       mSignalChoice.EAnswer,
				FAnswer:       mSignalChoice.FAnswer,
				CorrectAnswer: mSignalChoice.CorrectAnswer,
				Version:       mSignalChoice.Version,
				CreateTime:    mSignalChoice.CreateTime.UTC().String(),
				UpdateTime:    mSignalChoice.UpdateTime.UTC().String(),
			}
			scws := su.SignalChoiceWithSerial{
				SubjectNumber: suSubjectNumberAndIdPair.SubjectNumber,
				SignalChoice:  &suSignalChoiceInfo,
				Type:          SignalChoiceWithSerial,
			}

			suPaperItem := su.PaperItem{
				Type:                   SignalChoiceWithSerial,
				SignalChoiceWithSerial: &scws,
			}
			suFindSelfPaperResp.SelfPaper.PaperItems = append(suFindSelfPaperResp.SelfPaper.PaperItems, &suPaperItem)
		}
	}

	return &suFindSelfPaperResp, nil

	//temBigSubjects := make([]interface{}, 0)
	//fullBigSignalChoices := l.BigSignalChoice_ModelToProto(&selfPaper.BigSignalChoices)
	//temBigSubjects = append(temBigSubjects, fullBigSignalChoices)
	//
	//if len(fullBigSignalChoices) > 0 {
	//	// 排序
	//	sortBigSubject(temBigSubjects)
	//
	//}
	//
	//bigSubjects := make([]*anypb.Any, 0)
	//anys := fullBigSubjectToAny(temBigSubjects)
	//bigSubjects = append(bigSubjects, anys...)
	//
	//return &su.FindSelfPaperResp{
	//	SelfPaper: &su.SelfPaperInfo{
	//		Id:          selfPaper.Id.Hex(),
	//		PaperName:   selfPaper.PaperName,
	//		CreaterId:   selfPaper.CreaterId,
	//		Status:      selfPaper.Status,
	//		StartTime:   selfPaper.StartTime.UTC().String(),
	//		StopTime:    selfPaper.StopTime.UTC().String(),
	//		Version:     selfPaper.Version,
	//		CreateTime:  selfPaper.CreateTime.UTC().String(),
	//		UpdateTime:  selfPaper.UpdateTime.UTC().String(),
	//		Deleted:     selfPaper.Deleted,
	//		BigSubjects: bigSubjects,
	//	},
	//}, nil
}

//func sortBigSubject(temBigSubjects []interface{}) {
//	sort.Slice(temBigSubjects, func(i, j int) bool {
//		indexi := reflect.ValueOf(temBigSubjects).Index(i)
//		//取得BigSubject对象类型
//		bigSubjectiType := indexi.Type()
//		//取得SubjectNumber属性
//		subjectNumberFilei, flag := bigSubjectiType.FieldByName("SubjectNumber")
//		if flag == false {
//			return false
//		}
//		//取得SubjectNumber的类型
//		SubjectNumberTypei := subjectNumberFilei.Type.Kind()
//		if SubjectNumberTypei != reflect.Int64 && SubjectNumberTypei != reflect.Int {
//			return false
//		}
//		//取得SubjectNumber 内容
//		subjectNumberi := indexi.FieldByName("SubjectNumber").Int()
//
//		indexj := reflect.ValueOf(temBigSubjects).Index(j)
//		//取得BigSubject对象类型
//		bigSubjectjType := indexj.Type()
//		//取得SubjectNumber属性
//		subjectNumberFilej, flag := bigSubjectjType.FieldByName("SubjectNumber")
//		if flag == false {
//			return false
//		}
//		//取得SubjectNumber的类型
//		SubjectNumberTypej := subjectNumberFilej.Type.Kind()
//		if SubjectNumberTypej != reflect.Int64 && SubjectNumberTypej != reflect.Int {
//			return false
//		}
//		//取得SubjectNumber字段内容
//		subjectNumberj := indexi.FieldByName("SubjectNumber").Int()
//
//		return subjectNumberi < subjectNumberj
//	})
//}
//
//func fullBigSignalChoiceToAny(subject su.FullBigSignalChoice) (*anypb.Any, error) {
//	any, err := anypb.New(&subject)
//	any.TypeUrl = subject.Type
//	if err != nil {
//		logx.Errorf("err:" + err.Error())
//		return nil, err
//	}
//	return any, nil
//}
//
//func fullBigSubjectToAny(fullBigSubjects []interface{}) []*anypb.Any {
//	bigSubjects := make([]*anypb.Any, 0)
//	if len(fullBigSubjects) == 0 {
//		return bigSubjects
//	}
//	for _, subject := range fullBigSubjects {
//		switch subject.(type) {
//		case su.FullBigSignalChoice:
//			fullBigSignalChoice := subject.(su.FullBigSignalChoice)
//			any, err := fullBigSignalChoiceToAny(fullBigSignalChoice)
//			if err == nil {
//				bigSubjects = append(bigSubjects, any)
//			}
//		}
//	}
//
//	return bigSubjects
//}
//
//func (l *FindSelfPaperLogic) BigSignalChoice_ModelToProto(bigSignalChoices *[]papermodel.BigSignalChoice) []su.FullBigSignalChoice {
//	fullBigSignalChoices := make([]su.FullBigSignalChoice, 0)
//	for _, bigSubjectChoice := range *bigSignalChoices {
//		fullBigSignalChoice := su.FullBigSignalChoice{
//			SubjectNumber: int64(bigSubjectChoice.SubjectNumber),
//			Title:         bigSubjectChoice.Title,
//			FullSignalChoiceWithSubjectNumberList: func() []*su.FullSignalChoiceWithSubjectNumber {
//				list := make([]*su.FullSignalChoiceWithSubjectNumber, 0)
//				for _, v := range bigSubjectChoice.SignalChoiceWithNums {
//					signalChoice, err2 := l.svcCtx.SignalChoiceModel.FindOne(l.ctx, v.SignalChoiceId)
//					if err2 != nil {
//						logx.Errorf("err2:" + err2.Error())
//						continue
//					}
//					item := su.FullSignalChoiceWithSubjectNumber{
//						SubjectNumber: int64(v.SubjectNumber),
//						SignalChoice: &su.SignalChoiceInfo{
//							Id:            signalChoice.Id.Hex(),
//							Title:         signalChoice.Title,
//							AAnswer:       signalChoice.AAnswer,
//							BAnswer:       signalChoice.BAnswer,
//							CAnswer:       signalChoice.CAnswer,
//							DAnswer:       signalChoice.DAnswer,
//							EAnswer:       signalChoice.EAnswer,
//							FAnswer:       signalChoice.FAnswer,
//							CorrectAnswer: signalChoice.CorrectAnswer,
//							Version:       signalChoice.Version,
//							CreateTime:    signalChoice.CreateTime.UTC().String(),
//							UpdateTime:    signalChoice.UpdateTime.UTC().String(),
//						},
//					}
//					list = append(list, &item)
//				}
//				sort.Slice(list, func(i, j int) bool {
//					return list[i].SubjectNumber < list[j].SubjectNumber
//				})
//				return list
//			}(),
//		}
//		fullBigSignalChoices = append(fullBigSignalChoices, fullBigSignalChoice)
//	}
//	sort.Slice(fullBigSignalChoices, func(i, j int) bool {
//		return fullBigSignalChoices[i].SubjectNumber < fullBigSignalChoices[j].SubjectNumber
//	})
//	return fullBigSignalChoices
//}

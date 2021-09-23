package logic

import (
	"context"
	"go-zero-admin/rpc/model/papermodel"
	"time"

	"go-zero-admin/rpc/su/internal/svc"
	"go-zero-admin/rpc/su/su"

	"github.com/tal-tech/go-zero/core/logx"
)

type AddSelfPaperLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddSelfPaperLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddSelfPaperLogic {
	return &AddSelfPaperLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *AddSelfPaperLogic) AddSelfPaper(in *su.AddSelfPaperReq) (*su.AddSelfPaperResp, error) {
	// todo: add your logic here and delete this line

	//客户端只能传utc时间
	startTime, _ := time.Parse("2006-01-02 15:04:05", in.StartTime)
	stopTime, _ := time.Parse("2006-01-02 15:04:05", in.StopTime)

	selfPaper := papermodel.SelfPaper{
		PaperName:      in.PaperName,
		CreaterId:      in.CreaterId,
		Status:         in.Status,
		StartTime:      startTime,
		StopTime:       stopTime,
		Version:        in.Version,
		Deleted:        false,
		PaperItems:     paperItems_ProtoToModel(in.PaperItems),
		RandomSettings: randomSetting_ProtoToModel(in.RandomSettings),
	}

	err := l.svcCtx.SelfPaperModel.Insert(l.ctx, &selfPaper)
	if err != nil {
		return nil, err
	}

	if err != nil {
		return nil, err
	}

	return &su.AddSelfPaperResp{
		Result: true,
		Id:     selfPaper.Id.Hex(),
	}, nil
}

func paperItems_ProtoToModel(items []*su.PaperItem) []papermodel.PaperItem {
	paperItems := make([]papermodel.PaperItem, 0)
	for _, item := range items {
		switch item.Type {
		case SignalChoiceWithSerial:
			logx.Error("err:" + "SignalChoiceWithSerial type error in paper model")
			continue
		case ParagraphInstruction:
			value := papermodel.ParagraphInstruction{
				Instruction: item.ParagraphInstruction.Instruction,
				Type:        item.ParagraphInstruction.Type,
			}
			paperItem := papermodel.PaperItem{
				Type: value.Type,
				Item: value,
			}
			paperItems = append(paperItems, paperItem)
		case PageBar:
			value := papermodel.PageBar{
				Instruction: item.PageBar.Instruction,
				CurrentPage: item.PageBar.CurrentPage,
				TotalPage:   item.PageBar.TotalPage,
				Type:        item.PageBar.Type,
			}
			paperItem := papermodel.PaperItem{
				Type: value.Type,
				Item: value,
			}
			paperItems = append(paperItems, paperItem)
		case SubjectNumberAndIdPair:
			value := papermodel.SubjectNumberAndIdPair{
				SubjectNumber: item.SubjectNumberAndIdPair.SubjectNumber,
				SubjectId:     item.SubjectNumberAndIdPair.SubjectId,
				Type:          item.SubjectNumberAndIdPair.Type,
			}
			paperItem := papermodel.PaperItem{
				Type: value.Type,
				Item: value,
			}
			paperItems = append(paperItems, paperItem)
		}
	}
	return paperItems
}

func randomSetting_ProtoToModel(settings []*su.RandomSetting) []papermodel.RandomSetting {

	modelRandomSettings := make([]papermodel.RandomSetting, 0)
	for _, setting := range settings {
		modelSetting := papermodel.RandomSetting{
			StartNumber:  setting.StartNumber,
			EndNumber:    setting.EndNumber,
			SubjectCount: setting.SubjectCount,
		}
		modelRandomSettings = append(modelRandomSettings, modelSetting)
	}
	return modelRandomSettings

}

//func (l *AddSelfPaperLogic) getPaperBigSignalChoices(paper *papermodel.SelfPaper) *su.SignalChoiceBig {
//	if len(paper.SignalChoiceBig.SignalChoiceWithNums) != 0 {
//		return nil
//	}
//	signalChoicesList := make([]*su.SignalChoiceWithNum, 0)
//	for _, v := range paper.SignalChoiceBig.SignalChoiceWithNums {
//		signalChoice, err := l.svcCtx.SignalChoiceModel.FindOne(l.ctx, v.SignalChoiceId)
//		if err != nil {
//			logx.Error("err:" + err.Error())
//			continue
//		}
//		signalChoiceInfo := su.SignalChoiceInfo{
//			Id:            signalChoice.ID.Hex(),
//			Title:         signalChoice.Title,
//			AAnswer:       signalChoice.AAnswer,
//			BAnswer:       signalChoice.BAnswer,
//			CAnswer:       signalChoice.CAnswer,
//			DAnswer:       signalChoice.DAnswer,
//			EAnswer:       signalChoice.EAnswer,
//			FAnswer:       signalChoice.FAnswer,
//			CorrectAnswer: signalChoice.CorrectAnswer,
//			Version:       signalChoice.Version,
//			CreateTime:    signalChoice.CreateTime.String(),
//			UpdateTime:    signalChoice.UpdateTime.String(),
//		}
//		signalChoiceWithNum := su.SignalChoiceWithNum{
//			Number:       int64(v.SubjectNumber),
//			SignalChoice: &signalChoiceInfo,
//		}
//		//signalChoices = append(signalChoices, signalChoice)
//		signalChoicesList = append(signalChoicesList, &signalChoiceWithNum)
//	}
//	return &su.SignalChoiceBig{
//		Title:                   paper.SignalChoiceBig.Title,
//		SignalChoiceWithNumList: signalChoicesList,
//	}
//}

//func SignalChoiceBig_ProtoToModel(signalChoiceBig su.BigSignalChoice) papermodel.BigSignalChoice {
//	return papermodel.BigSignalChoice{
//		SubjectNumber: int(signalChoiceBig.SubjectNumber),
//		Title:         signalChoiceBig.Title,
//		Type:          signalChoiceBig.Type,
//		SignalChoiceWithNums: func() []papermodel.SignalChoiceWithNum {
//			subjects := make([]papermodel.SignalChoiceWithNum, 0)
//			for _, vSubject := range signalChoiceBig.SubjectNumberAndIdPairList {
//				subject := papermodel.SignalChoiceWithNum{
//					SubjectNumber:  int(vSubject.SubjectNumber),
//					SignalChoiceId: vSubject.SubjectId,
//				}
//				subjects = append(subjects, subject)
//			}
//			return subjects
//		}(),
//	}
//}

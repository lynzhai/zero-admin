package logic

import (
	"context"
	"go-zero-admin/rpc/su/suclient"

	"go-zero-admin/front-api/internal/svc"
	"go-zero-admin/front-api/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type FindSelfPaperLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewFindSelfPaperLogic(ctx context.Context, svcCtx *svc.ServiceContext) FindSelfPaperLogic {
	return FindSelfPaperLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *FindSelfPaperLogic) FindSelfPaper(req types.FindSelfPaperReq) (*types.FindSelfPaperResp, error) {
	// todo: add your logic here and delete this line
	err := l.svcCtx.Validate.Struct(&req)
	if err != nil {
		return nil, err
	}

	selfPaper, err := l.svcCtx.Su.FindSelfPaper(l.ctx, &suclient.FindSelfPaperReq{
		Id: req.Id,
	})
	if err != nil {
		return nil, err
	}

	return &types.FindSelfPaperResp{
		Code:    0,
		Message: "success",
		Data: types.SelfPaperInfo{
			Id:         selfPaper.SelfPaper.Id,
			PaperName:  selfPaper.SelfPaper.PaperName,
			CreaterId:  selfPaper.SelfPaper.CreaterId,
			Status:     selfPaper.SelfPaper.Status,
			StartTime:  selfPaper.SelfPaper.StartTime,
			StopTime:   selfPaper.SelfPaper.StopTime,
			Version:    selfPaper.SelfPaper.Version,
			CreateTime: selfPaper.SelfPaper.CreateTime,
			UpdateTime: selfPaper.SelfPaper.UpdateTime,
			DeleteTime: selfPaper.SelfPaper.DeleteTime,
			Deleted:    selfPaper.SelfPaper.Deleted,
			PaperItems: paperItemsToTypes(selfPaper.SelfPaper.PaperItems),
			RandomSettings: func() []*types.RandomSetting {
				ret := make([]*types.RandomSetting, 0)
				for _, setting := range selfPaper.SelfPaper.RandomSettings {
					value := types.RandomSetting{
						StartNumber:  setting.StartNumber,
						EndNumber:    setting.EndNumber,
						SubjectCount: setting.SubjectCount,
					}
					ret = append(ret, &value)
				}
				return ret
			}(),
		},
	}, nil
}

func paperItemsToTypes(items []*suclient.PaperItem) []*types.PaperItem {
	retPaperItems := make([]*types.PaperItem, 0)
	for _, item := range items {
		switch item.Type {
		case SignalChoiceWithSerial:
			value := types.SignalChoiceWithSerial{
				SubjectNumber: item.SignalChoiceWithSerial.SubjectNumber,
				Type:          SignalChoiceWithSerial,
				SignalChoice: types.SignalChoiceInfo{
					Id:            item.SignalChoiceWithSerial.SignalChoice.Id,
					Type:          item.SignalChoiceWithSerial.SignalChoice.Type,
					Title:         item.SignalChoiceWithSerial.SignalChoice.Title,
					AAnswer:       item.SignalChoiceWithSerial.SignalChoice.AAnswer,
					BAnswer:       item.SignalChoiceWithSerial.SignalChoice.BAnswer,
					CAnswer:       item.SignalChoiceWithSerial.SignalChoice.CAnswer,
					DAnswer:       item.SignalChoiceWithSerial.SignalChoice.DAnswer,
					EAnswer:       item.SignalChoiceWithSerial.SignalChoice.EAnswer,
					FAnswer:       item.SignalChoiceWithSerial.SignalChoice.FAnswer,
					CorrectAnswer: item.SignalChoiceWithSerial.SignalChoice.CorrectAnswer,
					Version:       item.SignalChoiceWithSerial.SignalChoice.Version,
					CreateTime:    item.SignalChoiceWithSerial.SignalChoice.CreateTime,
					UpdateTime:    item.SignalChoiceWithSerial.SignalChoice.UpdateTime,
				},
			}
			paperItem := types.PaperItem{
				Type:                   SignalChoiceWithSerial,
				SignalChoiceWithSerial: &value,
			}
			retPaperItems = append(retPaperItems, &paperItem)
		case SubjectNumberAndIdPair:
			value := types.SubjectNumberAndIdPair{
				SubjectNumber: item.SubjectNumberAndIdPair.SubjectNumber,
				SubjectId:     item.SubjectNumberAndIdPair.SubjectId,
				Type:          item.SubjectNumberAndIdPair.Type,
			}
			paperItem := types.PaperItem{
				Type:                 SubjectNumberAndIdPair,
				SubjectNumberAndIdPair: &value,
			}
			retPaperItems = append(retPaperItems, &paperItem)
		case ParagraphInstruction:
			value := types.ParagraphInstruction{
				Instruction: item.ParagraphInstruction.Instruction,
				Type:        item.ParagraphInstruction.Type,
			}
			paperItem := types.PaperItem{
				Type:                 ParagraphInstruction,
				ParagraphInstruction: &value,
			}
			retPaperItems = append(retPaperItems, &paperItem)
		case PageBar:
			value := types.PageBar{
				Instruction: item.PageBar.Instruction,
				CurrentPage: item.PageBar.CurrentPage,
				TotalPage:   item.PageBar.TotalPage,
				Type:        item.PageBar.Type,
			}
			paperItem := types.PaperItem{
				Type:    PageBar,
				PageBar: &value,
			}
			retPaperItems = append(retPaperItems, &paperItem)
		}
	}
	return retPaperItems
}

package logic

import (
	"context"
	"github.com/tal-tech/go-zero/core/logx"
	"go-zero-admin/front-api/internal/svc"
	"go-zero-admin/front-api/internal/types"
	"go-zero-admin/rpc/su/suclient"
)

type AddSelfPaperLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAddSelfPaperLogic(ctx context.Context, svcCtx *svc.ServiceContext) AddSelfPaperLogic {
	return AddSelfPaperLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AddSelfPaperLogic) AddSelfPaper(req types.AddSelfPaperReq) (*types.AddSelfPaperResp, error) {
	// todo: add your logic here and delete this line
	err := l.svcCtx.Validate.Struct(&req)
	if err != nil {
		logx.Error("err:" + err.Error())
		return nil, err
	}

	resp, err := l.svcCtx.Su.AddSelfPaper(l.ctx, &suclient.AddSelfPaperReq{
		PaperName:      req.PaperName,
		CreaterId:      req.CreaterId,
		Status:         req.Status,
		StartTime:      req.StartTime,
		StopTime:       req.StopTime,
		Version:        req.Version,
		PaperItems:     paperItemsToProto(req.PaperItems),
		RandomSettings: randomSettingsToProto(req.RandomSettings),
	})
	if err != nil {
		logx.Error("err:" + err.Error())
		return nil, err
	}
	return &types.AddSelfPaperResp{
		Code:    0,
		Message: "success",
		Data: types.AddSelfPaperRespData{
			Result: resp.Result,
			Id:     resp.Id,
		},
	}, nil
}

func paperItemsToProto(items []*types.PaperItem) []*suclient.PaperItem {
	retPaperItems := make([]*suclient.PaperItem, 0)
	for _, item := range items {
		switch item.Type {
		case SignalChoiceWithSerial:
			value := suclient.SignalChoiceWithSerial{
				SubjectNumber: 0,
				Type:          SignalChoiceWithSerial,
				SignalChoice: &suclient.SignalChoiceInfo{
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
			paperItem := suclient.PaperItem{
				Type:                   SignalChoiceWithSerial,
				SignalChoiceWithSerial: &value,
			}
			retPaperItems = append(retPaperItems, &paperItem)
		case SubjectNumberAndIdPair:
			value := suclient.SubjectNumberAndIdPair{
				SubjectNumber: item.SubjectNumberAndIdPair.SubjectNumber,
				SubjectId:     item.SubjectNumberAndIdPair.SubjectId,
				Type:          item.SubjectNumberAndIdPair.Type,
			}
			paperItem := suclient.PaperItem{
				Type:                 SubjectNumberAndIdPair,
				SubjectNumberAndIdPair: &value,
			}
			retPaperItems = append(retPaperItems, &paperItem)
		case ParagraphInstruction:
			value := suclient.ParagraphInstruction{
				Instruction: item.ParagraphInstruction.Instruction,
				Type:        item.ParagraphInstruction.Type,
			}
			paperItem := suclient.PaperItem{
				Type:                 ParagraphInstruction,
				ParagraphInstruction: &value,
			}
			retPaperItems = append(retPaperItems, &paperItem)
		case PageBar:
			value := suclient.PageBar{
				Instruction: item.PageBar.Instruction,
				CurrentPage: item.PageBar.CurrentPage,
				TotalPage:   item.PageBar.TotalPage,
				Type:        item.PageBar.Type,
			}
			paperItem := suclient.PaperItem{
				Type:    PageBar,
				PageBar: &value,
			}
			retPaperItems = append(retPaperItems, &paperItem)
		}
	}
	return retPaperItems
}

func randomSettingsToProto(settings []*types.RandomSetting) []*suclient.RandomSetting {
	ret := make([]*suclient.RandomSetting, 0)
	for _, v := range settings {
		index := suclient.RandomSetting{
			StartNumber:  v.StartNumber,
			EndNumber:    v.EndNumber,
			SubjectCount: v.SubjectCount,
		}
		ret = append(ret, &index)
	}
	return ret
}

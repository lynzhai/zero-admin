package logic

import "go-zero-admin/shared"

var (
	errorDuplicateCode  = shared.NewDefaultError("课程码重复")
	errorCodeEmpty  = shared.NewDefaultError("课程码为空")
	errorSubjectRegisterFail   = shared.NewDefaultError("课程注册失败")

	//errorRoleUnRegister     = shared.NewDefaultError("角色不存在")
	//errorUserRegisterFail   = shared.NewDefaultError("用户注册失败")
	//errorIncorrectPassword  = shared.NewDefaultError("用户密码错误")
	errorSubjectNotFound       = shared.NewDefaultError("课程不存在")
	errorAddStudentToSubject       = shared.NewDefaultError("添加学生到课程失败")
	errorDeleteStudentInSubject       = shared.NewDefaultError("删除课程中的学生失败")
	errorFindStudentInSubject       = shared.NewDefaultError("查找课程中的学生失败")
	errorFindSubjectByStudent       = shared.NewDefaultError("查找学生的课程失败")
)
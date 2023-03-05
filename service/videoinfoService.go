package service

import (
	"go_web/bll"
	"go_web/domain"
	"go_web/viewmodel"
)

//添加单个信息
func InsertVideoinfo(videoinfo *domain.Videoinfo) (bool, error) {
	err := bll.InsertVideoinfo(videoinfo)
	if err != nil {
		return false, err
	} else {
		return true, nil
	}
}

//修改单个信息
func UpdateVideoinfo(videoinfo *domain.Videoinfo) (bool, error) {
	err := bll.UpdateVideoinfo(videoinfo)
	if err != nil {
		return false, err
	} else {
		return true, nil
	}
}

//删除单个信息
func DeleteVideoinfo(userId int64) (bool, error) {
	err := bll.DeleteVideoinfo(userId)
	if err != nil {
		return false, err
	} else {
		return true, nil
	}
}

//根据编号获取单个信息
func GetSingleVideoinfo(userId int64) (*viewmodel.VideoinfoViewModel, error) {
	return bll.GetVideoinfoByUserId(userId)
}

// 获取信息列表(分页)-Videoinfo
func GetVideoinfoPageList(pageNumber int, pageSize int, conditions string, videoType int) ([]*viewmodel.VideoinfoViewModel, error) {
	//条件字符串拼接
	//var build strings.Builder
	/*build.WriteString("CREATED_TIME >=")
	build.WriteString("'" + conditions + "'")
	build.WriteString(" and  CREATED_TIME<=")
	build.WriteString("'" + conditions + "'")*/
	//build.WriteString(" date_format(CREATED_TIME,'%Y-%m-%d')='" + conditions + "'")
	//query := build.String()
	//query := fmt.Sprintf("CREATED_TIME >= '%s' and  CREATED_TIME<= '%s' ", conditions, conditions)
	return bll.GetVideoinfoList(pageNumber, pageSize, conditions, videoType)
}

//统计列表信息数量
func GetVideoinfoTotal(conditions string, videoType int) (int, error) {
	//query := fmt.Sprintf("CREATED_TIME between  '%s' and  '%s' ", conditions, conditions)
	//query := fmt.Sprintf("CREATED_TIME >= '%s' and  CREATED_TIME<= '%s' ", conditions, conditions)
	//条件字符串拼接
	//var build strings.Builder
	/*build.WriteString("CREATED_TIME >=")
	build.WriteString("'" + conditions + "'")
	build.WriteString(" and  CREATED_TIME<=")
	build.WriteString("'" + conditions + "'")*/
	//build.WriteString(" date_format(CREATED_TIME,'%Y-%m-%d')='" + conditions + "'")
	//query := build.String()
	return bll.GetVideoinfoTotal(conditions, videoType)
}

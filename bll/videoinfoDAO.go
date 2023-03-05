package bll

import (
	"database/sql"
	"go_web/domain"
	"go_web/util"
	"go_web/viewmodel"
	"log"
	"strconv"
)

const (
	//插入信息语句
	insertVideoinfoSQL = `INSERT INTO videoinfo(Id, VideoFileName, VideoType,VideoAddress,CREATED_BY, CREATED_TIME, UPDATED_BY, UPDATED_TIME, IsDeleted, Remarks) 
VALUES (?, ?, ?, ?, ?, ?, ?, ?,?, ?)`
	//修改信息语句
	updateVideoinfoSQL = `update videoinfo set VideoFileName=?,VideoType=?,VideoAddress=?,CREATED_BY=?, CREATED_TIME=?, 
UPDATED_BY=?, UPDATED_TIME=?, IsDeleted=?, Remarks=? where Id=?`
	//删除信息语句
	deleteVideoinfoSQL = `delete from videoinfo where Id=?`
	//信息列表语句(分页)
	listVideoinfoSQL = `select id,VideoFileName, VideoType,VideoAddress,CREATED_BY, CREATED_TIME, UPDATED_BY,
UPDATED_TIME,isDeleted,Remarks from videoinfo  where 1=1 and date_format(CREATED_TIME,'%Y-%m-%d')=? and VideoType=? LIMIT ? OFFSET ?`
	//根据编号查询单个信息
	SingleVideoinfoSQL = `select id,VideoFileName, VideoType,VideoAddress,CREATED_BY, CREATED_TIME, UPDATED_BY,
       UPDATED_TIME,isDeleted,Remarks from videoinfo where Id=?`
	//统计用户列表信息数量
	TotalVideoinfoSQL = `select count(1) total from videoinfo  where 1=1 and date_format(CREATED_TIME,'%Y-%m-%d')=? and VideoType=?;`
)

func scanVideoinfoWithSignOnAndBannerData(r *sql.Rows) (*domain.Videoinfo, error) {
	var id int64
	var videoFileName, videoAddress, CREATED_BY, CREATED_TIME, UPDATED_BY, UPDATED_TIME, Remarks string
	var isDeleted, videoType int

	err := r.Scan(&id, &videoFileName, &videoType, &videoAddress, &CREATED_BY, &CREATED_TIME,
		&UPDATED_BY, &UPDATED_TIME, &isDeleted, &Remarks)

	baseFileld := domain.BaseFileld{
		CREATED_BY:   CREATED_BY,   //创建人
		CREATED_TIME: CREATED_TIME, //创建时间
		UPDATED_BY:   UPDATED_BY,   //更新人
		UPDATED_TIME: UPDATED_TIME, //更新时间
		IsDeleted:    isDeleted,    //是否删除(逻辑删除)，1是，0否
		Remarks:      Remarks,      //备注
	}

	a := &domain.Videoinfo{
		Id:            id,            //编号(主键)
		VideoFileName: videoFileName, //视频文件名称
		VideoType:     videoType,     //视频类型,（1表示告警视频，2表示长视频）
		VideoAddress:  videoAddress,  //视频存放地址
		CommonFileld:  baseFileld,    //公共字段(创建时间，创建人等属性)
	}
	return a, err
}

//根据编号获取单个信息
func GetVideoinfoByUserId(userId int64) (*viewmodel.VideoinfoViewModel, error) {
	conn, err := util.GetConnection()
	//关闭数据连接
	defer util.Close(conn, err)
	if err != nil {
		return nil, err
	}
	data, err := conn.Query(SingleVideoinfoSQL, userId)
	if err != nil {
		return nil, err
	}
	videoinfoViewModel := &viewmodel.VideoinfoViewModel{}
	if data.Next() {
		videoinfo, err := scanVideoinfoWithSignOnAndBannerData(data)
		if err != nil {
			return nil, err
		}
		videoinfoViewModel = &viewmodel.VideoinfoViewModel{
			Id:            strconv.FormatInt(videoinfo.Id, 10), //编号(主键)
			VideoFileName: videoinfo.VideoFileName,             //视频文件名称
			VideoType:     videoinfo.VideoType,                 //视频类型,（1表示告警视频，2表示长视频）
			VideoAddress:  videoinfo.VideoAddress,              //视频存放地址
		}
		return videoinfoViewModel, nil
	}
	defer data.Close()
	err = data.Err()
	if err != nil {
		return nil, err
	}
	return videoinfoViewModel, err
}

// 获取信息列表(分页)-Upgraderecord
func GetVideoinfoList(pageNumber int, pageSize int, conditions string, videoType int) ([]*viewmodel.VideoinfoViewModel, error) {
	conn, err := util.GetConnection()
	//关闭数据连接
	defer util.Close(conn, err)
	var result []*viewmodel.VideoinfoViewModel
	if err != nil {
		return result, err
	}
	offset := (pageNumber - 1) * pageSize
	r, err := conn.Query(listVideoinfoSQL, conditions, videoType, pageSize, offset)
	if err != nil {
		return result, err
	}
	for r.Next() {
		videoinfo, err := scanVideoinfoWithSignOnAndBannerData(r)
		if err != nil {
			log.Printf("error: %v", err.Error())
			continue
		}
		//createtime := equipmentinfo.CommonFileld.CREATED_TIME
		//_time, _ := time.Parse(time.RFC3339, createtime)
		videoinfoViewModel := &viewmodel.VideoinfoViewModel{
			Id:            strconv.FormatInt(videoinfo.Id, 10), //用户编号(主键)
			VideoFileName: videoinfo.VideoFileName,             //视频文件名称
			VideoType:     videoinfo.VideoType,                 //视频类型,（1表示告警视频，2表示长视频）
			VideoAddress:  videoinfo.VideoAddress,              //视频存放地址
		}
		result = append(result, videoinfoViewModel)
	}
	defer r.Close()
	err = r.Err()
	if err != nil {
		return result, err
	}
	return result, nil
}

// insert Videoinfo
func InsertVideoinfo(videoinfo *domain.Videoinfo) error {
	err := InsertOrVideoinfo(insertVideoinfoSQL, 0, videoinfo)
	return err
}

// update Videoinfo
func UpdateVideoinfo(videoinfo *domain.Videoinfo) error {
	err := InsertOrVideoinfo(updateVideoinfoSQL, 1, videoinfo)
	return err
}

// Delete Videoinfo
func DeleteVideoinfo(userId int64) error {
	err := util.Delete(deleteVideoinfoSQL, userId)
	return err
}

//统计列表信息数量
func GetVideoinfoTotal(conditions string, videoType int) (int, error) {
	var total int
	conn, err := util.GetConnection()
	//关闭数据连接
	defer util.Close(conn, err)
	if err != nil {
		return 0, err
	}
	r, err := conn.Query(TotalVideoinfoSQL, conditions, videoType)
	if err != nil {
		return 0, err
	}
	if r.Next() {
		err := r.Scan(&total)
		if err != nil {
			return 0, err
		}
	}
	defer r.Close()
	err = r.Err()
	if err != nil {
		return 0, err
	}
	return total, nil
}

//更新和插入语句的逻辑和返回值一致，可封装在一起
func InsertOrVideoinfo(SQL string, edit int, videoinfo *domain.Videoinfo) error {
	conn, err := util.GetConnection()
	//关闭数据连接
	defer util.Close(conn, err)
	if err != nil {
		return err
	}
	var r sql.Result

	//插入
	if edit == 0 {
		r, err = conn.Exec(SQL,
			videoinfo.Id,
			videoinfo.VideoFileName,
			videoinfo.VideoType,
			videoinfo.VideoAddress,
			videoinfo.CommonFileld.CREATED_BY,
			videoinfo.CommonFileld.CREATED_TIME,
			videoinfo.CommonFileld.UPDATED_BY,
			videoinfo.CommonFileld.UPDATED_TIME,
			videoinfo.CommonFileld.IsDeleted,
			videoinfo.CommonFileld.Remarks,
		)
	} else {
		r, err = conn.Exec(SQL,
			videoinfo.VideoFileName,
			videoinfo.VideoType,
			videoinfo.VideoAddress,
			videoinfo.CommonFileld.CREATED_BY,
			videoinfo.CommonFileld.CREATED_TIME,
			videoinfo.CommonFileld.UPDATED_BY,
			videoinfo.CommonFileld.UPDATED_TIME,
			videoinfo.CommonFileld.IsDeleted,
			videoinfo.CommonFileld.Remarks,
			videoinfo.Id)
	}
	if err != nil {
		return err
	}
	rowNum, err := r.RowsAffected()
	if err != nil {
		return err
	}
	if rowNum > 0 {
		return nil
	}
	return err
}

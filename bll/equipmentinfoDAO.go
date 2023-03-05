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
	insertEquipmentinfoSQL = `INSERT INTO equipmentinfo(Id, Sn, Module, DeviceIP, Language, CREATED_BY, CREATED_TIME, UPDATED_BY, UPDATED_TIME, IsDeleted, Remarks) 
VALUES (?, ?, ?, ?, ?, ?, ?, ?,?, ?, ?)`
	//修改信息语句
	updateEquipmentinfoSQL = `update equipmentinfo set Sn=?, Module=?, DeviceIP=?, Language=?, CREATED_BY=?, CREATED_TIME=?, 
UPDATED_BY=?, UPDATED_TIME=?, IsDeleted=?, Remarks=? where Id=?`
	//删除信息语句
	deleteEquipmentinfoSQL = `delete from equipmentinfo where Id=?`
	//信息列表语句(分页)
	listEquipmentinfoSQL = `select id,sn,deviceIP,CREATED_BY, CREATED_TIME, UPDATED_BY,
UPDATED_TIME,Remarks,module,language,isDeleted from equipmentinfo   LIMIT ? OFFSET ?`
	//根据编号查询单个信息
	SingleEquipmentinfoSQL = `select id,sn,deviceIP,CREATED_BY, CREATED_TIME, UPDATED_BY,
UPDATED_TIME,Remarks,module,language,isDeleted from equipmentinfo where Id=?`
	//统计用户列表信息数量
	TotalEquipmentinfoSQL = `select count(1) total from equipmentinfo`
)

func scanEquipmentinfoWithSignOnAndBannerData(r *sql.Rows) (*domain.Equipmentinfo, error) {
	var id int64
	var sn, deviceIP, CREATED_BY, CREATED_TIME, UPDATED_BY, UPDATED_TIME, Remarks string
	var module, Language, isDeleted int

	err := r.Scan(&id, &sn, &deviceIP, &CREATED_BY, &CREATED_TIME,
		&UPDATED_BY, &UPDATED_TIME, &Remarks, &module, &Language, &isDeleted)

	baseFileld := domain.BaseFileld{
		CREATED_BY:   CREATED_BY,   //创建人
		CREATED_TIME: CREATED_TIME, //创建时间
		UPDATED_BY:   UPDATED_BY,   //更新人
		UPDATED_TIME: UPDATED_TIME, //更新时间
		IsDeleted:    isDeleted,    //是否删除(逻辑删除)，1是，0否
		Remarks:      Remarks,      //备注
	}

	a := &domain.Equipmentinfo{
		Id:           id,         //编号(主键)
		Sn:           sn,         //序列号
		Module:       module,     //所属功能模块
		DeviceIP:     deviceIP,   //设备IP
		Language:     Language,   //系统语言(1表示zh_cn，2表示en)
		CommonFileld: baseFileld, //公共字段(创建时间，创建人等属性)
	}
	return a, err
}

//根据编号获取单个信息
func GetEquipmentinfoByUserId(userId int64) (*viewmodel.EquipmentinfoViewModel, error) {
	conn, err := util.GetConnection()
	//关闭数据连接
	defer util.Close(conn, err)
	if err != nil {
		return nil, err
	}
	data, err := conn.Query(SingleEquipmentinfoSQL, userId)
	if err != nil {
		return nil, err
	}
	equipmentinfoViewModel := &viewmodel.EquipmentinfoViewModel{}
	if data.Next() {
		equipmentinfo, err := scanEquipmentinfoWithSignOnAndBannerData(data)
		if err != nil {
			return nil, err
		}
		ModuleText := ""
		if equipmentinfo.Module == 1 {
			ModuleText = "DMS"
		} else {
			ModuleText = "其他"
		}
		LanguageText := ""
		if equipmentinfo.Module == 1 {
			LanguageText = "zh_cn"
		} else {
			LanguageText = "en"
		}
		equipmentinfoViewModel = &viewmodel.EquipmentinfoViewModel{
			Id:       strconv.FormatInt(equipmentinfo.Id, 10), //编号(主键)
			Sn:       equipmentinfo.Sn,                        //序列号
			Module:   ModuleText,                              //所属功能模块
			DeviceIP: equipmentinfo.DeviceIP,                  //设备IP
			Language: LanguageText,                            //系统语言(1表示zh_cn，2表示en)
		}
		return equipmentinfoViewModel, nil
	}
	defer data.Close()
	err = data.Err()
	if err != nil {
		return nil, err
	}
	//ors.New("根据编号获取单个信息报错")
	return equipmentinfoViewModel, err
}

// 获取信息列表(分页)-Account
func GetEquipmentinfoList(pageNumber int, pageSize int) ([]*viewmodel.EquipmentinfoViewModel, error) {
	conn, err := util.GetConnection()
	//关闭数据连接
	defer util.Close(conn, err)
	var result []*viewmodel.EquipmentinfoViewModel
	if err != nil {
		return result, err
	}
	offset := (pageNumber - 1) * pageSize
	r, err := conn.Query(listEquipmentinfoSQL, pageSize, offset)
	if err != nil {
		return result, err
	}
	for r.Next() {
		equipmentinfo, err := scanEquipmentinfoWithSignOnAndBannerData(r)
		if err != nil {
			log.Printf("error: %v", err.Error())
			continue
		}
		//createtime := equipmentinfo.CommonFileld.CREATED_TIME
		//_time, _ := time.Parse(time.RFC3339, createtime)

		ModuleText := ""
		if equipmentinfo.Module == 1 {
			ModuleText = "DMS"
		} else {
			ModuleText = "其他"
		}
		LanguageText := ""
		if equipmentinfo.Module == 1 {
			LanguageText = "zh_cn"
		} else {
			LanguageText = "en"
		}

		equipmentinfoViewModel := &viewmodel.EquipmentinfoViewModel{
			Id:       strconv.FormatInt(equipmentinfo.Id, 10), //用户编号(主键)
			Sn:       equipmentinfo.Sn,                        //序列号
			Module:   ModuleText,                              //所属功能模块
			DeviceIP: equipmentinfo.DeviceIP,                  //设备IP
			Language: LanguageText,                            //系统语言(1表示zh_cn，2表示en)
		}
		result = append(result, equipmentinfoViewModel)
	}
	defer r.Close()
	err = r.Err()
	if err != nil {
		return result, err
	}
	return result, err
}

// insert Equipmentinfo
func InsertEquipmentinfo(equipmentinfo *domain.Equipmentinfo) error {
	err := InsertOrUpdateEquipmentinfo(insertEquipmentinfoSQL, 0, equipmentinfo)
	return err
}

// update Equipmentinfo
func UpdateEquipmentinfo(equipmentinfo *domain.Equipmentinfo) error {
	err := InsertOrUpdateEquipmentinfo(updateEquipmentinfoSQL, 1, equipmentinfo)
	return err
}

// Delete Equipmentinfo
func DeleteEquipmentinfo(userId int64) error {
	err := util.Delete(deleteEquipmentinfoSQL, userId)
	return err
}

//统计列表信息数量
func GetEquipmentinfoTotal() (int, error) {
	var total int
	conn, err := util.GetConnection()
	//关闭数据连接
	defer util.Close(conn, err)
	if err != nil {
		return 0, err
	}
	r, err := conn.Query(TotalEquipmentinfoSQL)
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
func InsertOrUpdateEquipmentinfo(SQL string, edit int, equipmentinfo *domain.Equipmentinfo) error {
	conn, err := util.GetConnection()
	//关闭数据连接
	defer util.Close(conn, err)
	if err != nil {
		return err
	}
	var r sql.Result

	//INSERT INTO user(Id, UserName, PassWord, NickName, Avatar, FullName, Mobile, Sex, Nation, Address, Email, IsLock, QRCode, RoleId, DepartId,
	//PositionId, CREATED_BY, CREATED_TIME, UPDATED_BY, UPDATED_TIME, IsDeleted, Remarks, LoginLastTime)

	//插入
	if edit == 0 {
		r, err = conn.Exec(SQL,
			equipmentinfo.Id,
			equipmentinfo.Sn,
			equipmentinfo.Module,
			equipmentinfo.DeviceIP,
			equipmentinfo.Language,
			equipmentinfo.CommonFileld.CREATED_BY,
			equipmentinfo.CommonFileld.CREATED_TIME,
			equipmentinfo.CommonFileld.UPDATED_BY,
			equipmentinfo.CommonFileld.UPDATED_TIME,
			equipmentinfo.CommonFileld.IsDeleted,
			equipmentinfo.CommonFileld.Remarks,
		)
	} else {
		r, err = conn.Exec(SQL,
			equipmentinfo.Sn,
			equipmentinfo.Module,
			equipmentinfo.DeviceIP,
			equipmentinfo.Language,
			equipmentinfo.CommonFileld.CREATED_BY,
			equipmentinfo.CommonFileld.CREATED_TIME,
			equipmentinfo.CommonFileld.UPDATED_BY,
			equipmentinfo.CommonFileld.UPDATED_TIME,
			equipmentinfo.CommonFileld.IsDeleted,
			equipmentinfo.CommonFileld.Remarks,
			equipmentinfo.Id)
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

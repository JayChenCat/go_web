package controller

import (
	"encoding/json"
	"go_web/config"
	"go_web/domain"
	"go_web/service"
	"go_web/util"
	"go_web/viewmodel"
	"log"
	"net/http"
	"strconv"
)

//跳转 查看视频主页
func ViewVideo(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		// 跳转到查看视频主页
		skiplinks(w, r, make(map[string]interface{}), config.VideoPath)
	} else {
		getPageListVideoinfo(w, r)
	}
}

//获取视频列表信息-分页
func getPageListVideoinfo(w http.ResponseWriter, r *http.Request) {
	msg := new(domain.Result)
	err := r.ParseForm()
	if err != nil {
		log.Printf("GetPageListVideoinfo error: %v", err.Error())
	}
	pageNumber := r.FormValue("pagenumber")
	pageSize := r.FormValue("pagesize")
	date := r.FormValue("date")
	videoType := r.FormValue("videoType")
	num, err := strconv.Atoi(pageNumber)
	_videoType, err := strconv.Atoi(videoType)
	psize, err := strconv.Atoi(pageSize)
	total, err := service.GetVideoinfoTotal(date, _videoType)
	if err != nil {
		log.Printf("统计视频列表信息数量 error: %v", err.Error())
	} else {
		list, err := service.GetVideoinfoPageList(num, psize, date, _videoType)
		if err != nil {
			msg.Code = 1
			msg.Msg = "获取视频列表信息失败！"
			log.Printf("获取用户列表信息 error: %v", err.Error())
		} else {
			msg.Code = 0
			msg.Msg = "获取视频列表信息成功！"
			videoinfos := viewmodel.ListVideoinfo{
				Lists: list,
				Count: total,
			}
			dataJson, _ := dealVideoinfo(videoinfos)
			//解决int64转换精度丢失问题
			msg.Data = dataJson
		}
	}
	util.ResponseHtml(w, msg, err)
}

func dealVideoinfo(data viewmodel.ListVideoinfo) (string, error) {
	account, err := json.Marshal(data)
	return string(account), err
}

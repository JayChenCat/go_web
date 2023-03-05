/*获取当前时间*/
var time = new Date();
/*格式化日，如果小于9，则补0*/
var day = ("0" + time.getDate()).slice(-2);
/*格式化月，如果小于9，则补0*/
var month = ("0" + (time.getMonth() + 1)).slice(-2);
/*拼接日期*/
var today = time.getFullYear() + "-" + (month) + "-" + (day);
//document.getElementById("dateInput").value = (today);

var date;
var type = 0;
var fileNames;
$(function () {
    /*if (!IsLogin()) {
       window.location.href = "login.html";//?orgUrl=index.html
   }*/
   layui.use(['layer', 'element', 'swiper', 'jquery', 'laypage'], function () {
       var layer = layui.layer
           , $ = layui.jquery
           , swiper = layui.swiper
           , element = layui.element
           , laypage = layui.laypage;

       //browser version check
       //导航JS
       (function nav() {
           /*主导航*/
            $('#main_nav_btn').on('click', function () {
                $('#main_nav_list').toggleClass('layui-hide-xs');
            });
        })();
    });
    $('#dateInput').val(today);
});


//总页数
function queryVideoInfoClick() {
    $("#PageNum").text(1);
    queryVideo(1);
}
//查询视频-分页
function queryVideo(currPageNum) {
    //当前页
    var pageNumber = currPageNum;
    //每页条数
    var pageSize = 14;

    var radioBTs = document.getElementsByName("videoType");
    for (var i = 0; i < radioBTs.length; i++) {
        if (radioBTs[i].checked) {
            type = radioBTs[i].value;
        }
    }
    $("#list").empty();
    date = dateInput.value;
    //var url = "http://" +ip+ ":10001/getvideoinfo?date=" + date + "&type=" + type;
    //alert(url);
    if (date.length == 0) { alert("请选择查询日期！"); return; }

    $.ajax({
        type: "post",
        url:host + ":8083/video",
        data: {
            "pagenumber":currPageNum,
            "pagesize":pageSize,
            "date":date,
            "videoType":type
        },
        dataType: "json",
        success: function(result) {
            if(result.code==0) {
                //alert(result.data)
                if(result.data!=""){
                    var responseObject = jQuery.parseJSON(result.data);
                    var count = responseObject.Count;
                    if (count > 0) {
                        videoJson(responseObject.Lists, count, currPageNum,pageSize);
                    }else{
                        $("#pager").css("display", "none");
                        document.getElementById("list").innerHTML = "未找到任何回放视频，请更换日期再查询!";
                        return;
                    }
                }
            }
            else {
                layer.msg(result.msg);
            }
        }
    });
}

function  videoJson(jsonstr, count, currPageNum,pageSize){
    //当前页
    var pageNumber = currPageNum;

    pageCount = count % pageSize == 0 ? parseInt((count / pageSize)) : parseInt((count / pageSize)) + 1;
    //赋值总页数和总记录
    $("#TotalPage").text(pageCount);
    $("#Total").text(count);
    $("#pageCount").text(pageSize);
    selectOptions(pageCount, pageNumber);
    //var startRow = ((pageNumber - 1) * pageSize + 1) - 1;//开始显示的行  1
    //var endRow = pageNumber * pageSize;//结束显示的行   15
    if (count > pageSize) {
        $("#pager").css("display", "block");
    }else{
        $("#pager").css("display", "none");
    }
    //var jsonObject = jQuery.parseJSON(jsonstr);
    fileNames =jsonstr;
    for (var i = 0; i < fileNames.length; i++) {
        var btID = "playBT" + i;
        document.getElementById("list").innerHTML = document.getElementById("list").innerHTML + '<dl class="vodeo_list_dl"><dt><p  class="a_plaly_list" title=' + fileNames[i].videoFileName + '>' + fileNames[i].videoFileName + '</p></dt><dd><input id="' + btID + '" type="button" class="btn" onClick="playVideoClick(' + i + ')" value="播放"/></dd></dl>';
    }
}

//播放视频
function playVideoClick(index) {
    //开启加载图标
    loadIcon("正在播放视频中...", 2000);
    var video = document.getElementById("video");
    var playurl =host + ":8083/video";
    //告警视频
    if (type == 1) {
        playurl += "/warning/";
    }
    else //长视频
    {
        playurl += "/normal/";
    }
    playurl += fileNames[index].videoAddress + "";
    video.setAttribute("src", playurl);
    video.play();
    //关闭加载图标
    closeLoadIcon();
}





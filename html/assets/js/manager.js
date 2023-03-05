
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

    getUserInfoList(1);
});


//获取用户列表信息
function getUserInfoList(currPageNum) {
    var pageSize=20;
    $.ajax({
        type: "post",
        url:host + ":8083/manager",
        data: {
            "pagenumber":currPageNum,
            "pagesize":pageSize,
        },
        dataType: "json",
        success: function(result) {
            if(result.code==0) {
                //alert(result.data)
                if(result.data!=""){
                    var responseObject = jQuery.parseJSON(result.data);
                    //responseObject["AccountViewModel"];
                    var count = responseObject.Count;
                    if (count > 0) {
                        parseJson(responseObject.Lists, count, currPageNum,pageSize);
                    }else{
                        $("#pager").css("display", "none");
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
//解析用户列表信息
function
parseJson(jsonstr, count, currPageNum,pageSize) {
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


    var jsonObject = jsonstr;
    var strHtml = "";
    $("#userInfoList").empty();
    strHtml += "<tr>";
    strHtml += "<th align=\"center\">编号</th>";
    strHtml += "<th align=\"center\">用户名称</th>";
    strHtml += "<th align=\"center\">E-mail地址</th>";
    strHtml += "<th align=\"center\">添加时间</th>";
    strHtml += "<th align=\"center\">操作</th>";
    strHtml += "</tr>";



    for (var i = 0; i < jsonObject.length; i++) {
        strHtml += "<td align=\"center\">" + jsonObject[i].id + "</td>";
        strHtml += "<td align=\"center\">" + jsonObject[i].username + "</td>";
        strHtml += "<td align=\"center\">" + jsonObject[i].email + "</td>";
        strHtml += "<td align=\"center\">" + jsonObject[i].addtime + "</td>";
        if (jsonObject[i].username == "admin") {
            strHtml += "<td align=\"center\">管理员账户不能删除</td>";
        } else {
            //
            strHtml += "<td align=\"center\"><a href=\"addmanager?id=" + jsonObject[i].id + "\">编辑</a> | <a href=\"#\" onclick=\"if(confirm('您确认删除该用户信息吗？')==false)return false;delUserInfo('" + jsonObject[i].id + "')\">删除</a></td>";
        }
        strHtml += "</tr>";
    }

    //不分页的拼接
    //for (var i = 0; i < jsonObject.length; i++) {
    //    strHtml += "<tr id=\"" + jsonObject[i].user + "\">";
    //    strHtml += "<td align=\"center\">" + (i+1)+"</td>";
    //    strHtml += "<td align=\"center\">" + jsonObject[i].user+"</td>";
    //    strHtml += "<td align=\"center\">" + jsonObject[i].email + "</td>";
    //    strHtml += "<td align=\"center\">" + jsonObject[i].date + "</td>";
    //    if (jsonObject[i].user == "admin")
    //    {
    //        strHtml += "<td align=\"center\">管理员账户不能删除</td>";
    //    } else
    //    {
    //        strHtml += "<td align=\"center\"><a href=\"addmanager.html?id=" + jsonObject[i].user + "\">编辑</a> | <a href=\"#\" onclick=\"if(confirm('您确认删除该用户信息吗？')==false)return false;delUserInfo('" + jsonObject[i].user + "')\">删除</a></td>";
    //    }
    //    strHtml += "</tr>";
    //}
    $("#userInfoList").append(strHtml);
}

//删除用户
function delUserInfo(uid) {

    $.ajax({
        type: "post",
        url:host + ":8083/DeleteUser",
        data: {
            "id":uid,
        },
        dataType: "json",
        success: function(result) {
            if(result.code==0) {
                //重新获取用户信息列表
                getUserInfoList(1);
            }
            else {
                layer.msg(result.msg);
            }
        }
    });
}


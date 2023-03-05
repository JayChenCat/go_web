/*分页函数调用*/
//页码下拉框填充事件
function selectOptions(num, pageNumber) {
    var thisSelectOptions = $("#select_Page");
    thisSelectOptions.empty();
    for (var i = 1; i <= num; i++) {
        if (i == pageNumber) {
            thisSelectOptions.append("<option selected=\"selected\">" + i + "</option>");
        } else {
            thisSelectOptions.append("<option>" + i + "</option>");
        }
    }
    $("#PageNum").text(pageNumber);
}

//页码下拉框改变事件
function selectChange(type) {
    var thisSelectOptions = $("#select_Page");
    var selectOptions = thisSelectOptions.find("option:selected").text();
    //视频分页
    if (type == 0) {
        queryVideo(selectOptions);
    }
    else {
        getUserInfoList(selectOptions);
    }
}

//分页事件
function pagePagingClick(pageAction,type) {
    //获取当前页
    var currPage = 1;
    var thisPage = $("#PageNum");
    currPage = thisPage.text();
    //获取总页数
    var TotalPage = $("#TotalPage");
    var countPage = TotalPage.text();
    //下一页
    if (pageAction == 0) {
        currPage++;
        if (currPage > countPage) {
            //设置下一页按钮为禁用
            $("#Next").css("cursor", "not-allowed");
            $("#Next").prop("disabled", true);
            return;
        } else {
            //视频分页
            if (type == 0) {
                queryVideo(currPage);
            }
            else
            {
                getUserInfoList(currPage);
            }
            thisPage.text(currPage);
            $("#Next").css("cursor", "pointer");
            $("#Next").prop("disabled", false);
            $("#Pre").css("cursor", "pointer");
            $("#Pre").prop("disabled", false);
        }
    }
    else //上一页
    {
        currPage--;
        if (currPage < 1) {
            /* alert("当前已是第一页!");*/
            //设置上一页按钮为禁用
            $("#Pre").css("cursor", "not-allowed");
            $("#Pre").prop("disabled", true);
            return;
        }
        else {
            //视频分页
            if (type == 0) {
                queryVideo(currPage);
            }
            else {
                getUserInfoList(currPage);
            }
            thisPage.text(currPage);
            $("#Pre").css("cursor", "pointer");
            $("#Pre").prop("disabled", false);
            $("#Next").css("cursor", "pointer");
            $("#Next").prop("disabled", false);
        }
    }
    //alert(currPage);
}

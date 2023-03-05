//全局ip地址,根据部署的机子会有不一样的ip
var ip="192.168.198.198";//window.document.location.hostname
var host="https://"+window.document.location.hostname;
//判断是否登录,没登录就重定向到登录界面
function IsLogin(){
    /*var ca = document.cookie.split(';');
    var name= "login";
    for(var i=0; i<ca.length; i++) {
        var c = ca[i].trim();
        if (c.startsWith(name)==true) { 
            if(c.substring(name.length + 1,c.length)=="true"){
                return true;
            }
        }
    }
    return false;*/
    var username = sessionStorage.getItem("username");
    if(username==null){
        return false;
    }
    return true
}

//获取全局token
function getToken(){
    var token="autueutwyrwyrgwrgggg"
    $.ajax({
        type: "post",
        url: host + ":8083/token",
        dataType: "json",
        success: function(result) {
            if(result.code!=0) {
                layer.msg(result.msg);
            }else{
                $.trim($("#token").val(result.data))
                token=result.data;
            }
        }
    });
    return token;
}
/**
 +----------------------------------------------------------
 * 下拉菜单
 +----------------------------------------------------------
 */
$(function() {
    /*$('.M').hover(function() {
        $(this).addClass('active');
    },
    function() {
        $(this).removeClass('active');
        });*/

    //getIndexUser();

    //$('#nav_menu li').eq(0).addClass('hover');
    /*如果在导航下如果有分类,如图.这样分类url和导航的url就不能匹配了,所以下面方法就不适用*/
    var lochref = $.trim(window.location.href);// 获得当前页面的URL
    $("#nav_menu li a").each(function (i) { //获取导航栏中每个a标签
        var me = $(this);
        var mehref = $.trim(me.get(0).href);//获得每个<a>的url
        /*判断当前url是否包含每个导航特定的rel  lochref == mehref是为了专门判断首页  */
        var m_rel=me.parent().attr('rel');
        if (lochref.indexOf(m_rel) != -1 || lochref == mehref || lochref.indexOf("addmanager") != -1) {
            me.parent().addClass("hover").siblings().removeClass("hover");;
        } else {
            me.parent().removeClass("hover");
        }
    });
});

/**
 +----------------------------------------------------------
 * 刷新验证码
 +----------------------------------------------------------
 */
function refreshimage() {
    var cap = document.getElementById('vcode');
    cap.src = cap.src + '?';
}

/**
 +----------------------------------------------------------
 * 无组件刷新局部内容
 +----------------------------------------------------------
 */
function dou_callback(page, name, value, target) {
    $.ajax({
        type: 'GET',
        url: page,
        data: name + '=' + value,
        dataType: "html",
        success: function(html) {
            $('#' + target).html(html);
        }
    });
}

/**
 +----------------------------------------------------------
 * 表单全选
 +----------------------------------------------------------
 */
function selectcheckbox(form) {
    for (var i = 0; i < form.elements.length; i++) {
        var e = form.elements[i];
        if (e.name != 'chkall' && e.disabled != true) e.checked = form.chkall.checked;
    }
}

/**
 +----------------------------------------------------------
 * 显示服务端扩展列表
 +----------------------------------------------------------
 */
function get_cloud_list(unique_id, get, localsite) {
    $.ajax({
        type: 'GET',
        url: 'http://cloud.douco.com/extend&rec=client',
        data: {'unique_id':unique_id, 'get':get, 'localsite':localsite},
        dataType: 'jsonp',
        jsonp: 'jsoncallback',
        success: function(cloud) {
            $('.selector').html(cloud.selector)
            $('.cloudList').html(cloud.html)
            $('.pager').html(cloud.pager)
        }
    });
}

/**
 +----------------------------------------------------------
 * 写入可更新数量
 +----------------------------------------------------------
 */
function cloud_update_number(localsite) {
    $.ajax({
        type: 'GET',
        url: 'http://cloud.douco.com/extend&rec=cloud_update_number',
        data: {'localsite':localsite},
        dataType: 'jsonp',
        jsonp: 'jsoncallback',
        success: function(cloud) {
            change_update_number(cloud.update, cloud.patch, cloud.module, cloud.plugin, cloud.theme, cloud.mobile)
        }
    });
}

/**
 +----------------------------------------------------------
 * 修改update_number值
 +----------------------------------------------------------
 */
function change_update_number(update, patch, module, plugin, theme, mobile) {
    $.ajax({
        type: 'POST',
        url: 'cloud.php?rec=update_number',
        data: {'update':update, 'patch':patch, 'module':module, 'plugin':plugin, 'theme':theme}
    });
}

/**
 +----------------------------------------------------------
 * 弹出窗口
 +----------------------------------------------------------
 */
function douFrame(name, frame, url ) {
    $.ajax({
        type: 'POST',
        url: url,
        data: {'name':name, 'frame':frame},
        dataType: 'html',
        success: function(html) {
            $(document.body).append(html);
        }
    });
}

/**
 +----------------------------------------------------------
 * 显示和隐藏
 +----------------------------------------------------------
 */
function douDisplay(target, action) {
    var traget = document.getElementById(target);
    if (action == 'show') {
        traget.style.display = 'block';
    } else {
        traget.style.display = 'none';
    }
}

/**
 +----------------------------------------------------------
 * 清空对象内HTML
 +----------------------------------------------------------
 */
function douRemove(target) {
    var obj = document.getElementById(target);
    obj.parentNode.removeChild(obj);
}

/**
 +----------------------------------------------------------
 * 无刷新自定义导航名称
 +----------------------------------------------------------
 */
function change(id, choose) {
    document.getElementById(id).value = choose.options[choose.selectedIndex].title;
}

//验证权限，不是管理员不可更改配置
function isVerifyAdministrators() {
    //var username = sessionStorage.getItem("username");
    var username = $("#currUser").text();
    if (username.indexOf("admin")<0)
    {
        document.getElementById("div_Disable").style.display = "block";
    }
    else
    {
        document.getElementById("div_Disable").style.display = "none";
    }
}

//给所有页面头部赋值当前用户
function getIndexUser()
{
    var username = sessionStorage.getItem("username");
    if (username != null) {
        $("#currUser").text("您好，" + username + "");
    }
}

//退出登录
function quiteClick() {
    /* var d = new Date();
     d.setTime(-1);
     var expires = "expires="+d.toGMTString();
     document.cookie="login=false; expires="+expires;*/
    //清除登录会话
    sessionStorage.clear();
    window.location.href = "login.html";
}

//打开加载图标
function loadIcon(text,time) {
    var layerMsg = layer.load(1, { // 此处1没有意义，随便写个东西
        content: '' + text+'',//向设备提交配置中...
        icon: 2, // 0~2 ,0比较好看
        shade: [0.5, 'black'], // 黑色透明度0.5背景
        time: time,
        success: function (layero) {
            layero.find('.layui-layer-content').css({
                'padding-top': '50px',
                'width': '60px'
            });
        }
    });
}

//关闭加载图标
function closeLoadIcon() { layer.closeAll(); }


//创建HTTP通讯实例
function createXHR() {
    var XHR = [ //兼容不同浏览器和版本得创建函数数组
        function () {
            return new XMLHttpRequest()
        },
        function () {
            return new ActiveXObject("Msxml2.XMLHTTP")
        },
        function () {
            return new ActiveXObject("Msxml3.XMLHTTP")
        },
        function () {
            return new ActiveXObject("Microsoft.XMLHTTP")
        }];
    var xhr = null;
    //尝试调用函数，如果成功则返回XMLHttpRequest对象，否则继续尝试
    for (var i = 0; i < XHR.length; i++) {
        try {
            xhr = XHR[i]();
        } catch (e) {
            continue //如果发生异常，则继续下一个函数调用
        }
        break; //如果成功，则中止循环
    }
    return xhr; //返回对象实例
}
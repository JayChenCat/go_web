
var url = location.search; //获取url中"?"符之后 的字串
if (url.indexOf('?') < 0) {
    url = "index.html";
} else {
    var tmp = url.split("=");
    url = tmp[1];
}




//登录，连接设备获取用户信息
function logon(user, pwd) {
    //打开加载图标
    var layerMsg = layer.load(1, { // 此处1没有意义，随便写个东西
        icon: 2, // 0~2 ,0比较好看
        shade: [0.5, 'black'] // 黑色透明度0.5背景
    });
    var responseObject = null;
    var lurl = "http://" + ip + ":10001/getuserInfo?user=" + user + "&pwd=" + pwd + "";
    var xhr = createXHR();
    xhr.open("GET", lurl, true);
    xhr.send();
    xhr.onreadystatechange = function () {
        if (xhr.readyState == 4) {
            if (xhr.status == 200) {
                responseObject = jQuery.parseJSON(xhr.responseText);
                if (responseObject["code"] == 0) {
                    sessionStorage.setItem("username", document.getElementById("userNameInput").value.trim());
                    document.getElementById("messageSpan").innerHTML = "";
                    window.location.href = url;
                    //document.getElementById("messageSpan").innerHTML = "登录成功！";
                    //layer.msg('登录成功', {
                    //    icon: 1,
                    //    shade: 0.1,
                    //    time: 1500
                    //});
                } else if (responseObject["code"] == -1) {
                    document.getElementById("messageSpan").innerHTML = "用户名称不存在！";
                }
                else if (responseObject["code"] == -2) {
                    document.getElementById("messageSpan").innerHTML = "用户名或者密码错误！";
                }
                else {
                    document.getElementById("messageSpan").innerHTML = "系统错误！";
                }
                //关闭加载图标
                layer.closeAll();
            }
        }
        //else {
        //    var jsonString = JSON.stringify(responseObject);
        //    if (jsonString =="null") {
        //        //2秒后自动关闭加载图标
        //        setTimeout(function () {
        //            //关闭加载图标
        //            layer.closeAll();
        //            document.getElementById("messageSpan").innerHTML = "登录超时，请检查设备连接情况！"
        //        }, 2000);}
        //}
    }
    //2秒后自动关闭加载图标
    setTimeout(function () {
        //关闭加载图标
        layer.closeAll();
        document.getElementById("messageSpan").innerHTML = "DMS设备运行不正常，请检查设备。";
    }, 2000);
}

//登录事件
function Submit_Login()
{
    var userName = document.getElementById("userNameInput").value.trim();
    if (userName == "") {
        document.getElementById("messageSpan").innerHTML = "请输入用户名！";
        return;
    }

    var psw = document.getElementById("pwdInput").value.trim();
    if (psw == "") {
        document.getElementById("messageSpan").innerHTML = "请输入密码！"
        return;
    }

    if ((userName.indexOf("\'") > 0 || psw.indexOf("\"") > 0)) {
        document.getElementById("messageSpan").innerHTML = "用户名或密码中不能出现单引号、双引号!";
        return;
    }

    var hash = hex_md5(userName + psw + "3i61fghb");
    //调用登录函数
    //logon(userName, hash);
    SumbitUserInfo();
}

function SumbitUserInfo(){
    var userName=$.trim($("#userNameInput").val());
    var pwd=$.trim($("#pwdInput").val());
    var hash = hex_md5(userName + pwd + "3i61fghb");

    $.ajax({
        type: "post",
        url:host + ":8083/login",
        data: {
            "username":userName,
            "password":hash
        },
        dataType: "json",
        success: function(result) {
            console.log(result)
            if(result.code==0) {
                window.location.href="/index";
            }
            layer.msg(result.msg);
        }
    });
}

//重置
function Clear_Text()
{
    $("#userNameInput").val("");
    $("#pwdInput").val("");
}


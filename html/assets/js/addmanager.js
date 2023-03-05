var uid = '';
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

    getToken();
    getUserInfo();
});


//添加用户信息
function submitUserInfo() {
    //校验用户信息
    var ischeck = checkUserInfo();
    if (ischeck) {
        var jsonstr = UserInfo();
        //console.log(jsonstr);

        $.ajax({
            type: "post",
            url:host + ":8083/addmanager",
            data: {
                "id":$("#oper").val(),
                "date":$("#date").val(),
                "token":$.trim($("#token").val()),
                "username":$("#user").val(),
                "password":$("#pwd").val(),
                "email":$("#email").val()
            },/*jsonstr,*/
            dataType: "json",
            success: function(result) {
                layer.msg(result.msg);
            }
        });
    }
}





//获取单个用户信息
function getUserInfo() {
    var URL = window.location.search; //？xx=aa&yy=b  形式
    //表示修改操作
    if (URL != '') {
        URL = URL.split('=');
        uid = URL[1];
        $.ajax({
            type: "post",
            url:host + ":8083/SingleAccount",
            data: {
                "id":uid
            },
            dataType: "json",
            success: function(result) {
                if(result.code==0) {
                    //alert(result.data)
                    if(result.data!=""){
                        //alert(result.data)
                        parseJsonUserInfo(result.data);
                    }
                }
                else {
                    layer.msg(result.msg);
                }
            }
        });
    }
}

function checkUserInfo() {
    var falg = true;
    var user_name = $("#user").val().trim();
    var email = $("#email").val().trim();
    var password = $("#pwd").val().trim();
    var password_confirm = $("#password_confirm").val().trim();



    if (user_name == "") {
        layer.alert("用户名为必填项!");
        return false;
    }
    if (email == "") {
        layer.alert("邮箱号为必填项!");
        falg = false;
    }

    if (password == "") {
        layer.alert("密码为必填项!");
        falg = false;
    }
    if (password_confirm == "") {
        layer.alert("确认密码为必填项!");
        falg = false;
    }

    if ((password.indexOf("\'") > 0 || password.indexOf("\"") > 0)) {
        layer.alert("密码中不能出现单引号、双引号!");
        falg = false;
    }
    if ((password_confirm.indexOf("\'") > 0 || password_confirm.indexOf("\"") > 0)) {
        layer.alert("确认密码中不能出现单引号、双引号!");
        falg = false;
    }

    if (password != password_confirm) {
        layer.alert("前后两次密码不一样!");
        falg = false;
    }

    if (user_name.indexOf("\'") > 0 || user_name.indexOf("\"") > 0) {
        layer.alert("用户名称中不能出现单引号、双引号!");
        falg = false;
    }

    let re = /^\w+@[a-zA-Z0-9]{2,10}(?:\.[a-z]{2,4}){1,3}$/;
    if (!re.test(email)) {
        layer.alert("邮箱格式不正确，请重新输入!");
        falg = false;
    }
    return falg;
}

//提交用户信息
function UserInfo() {
    /*获取当前时间*/
    var time = new Date();
    /*格式化日，如果小于9，则补0*/
    var day = ("0" + time.getDate()).slice(-2);
    /*格式化月，如果小于9，则补0*/
    var month = ("0" + (time.getMonth() + 1)).slice(-2);
    /*拼接日期*/
    var today = time.getFullYear() + "-" + (month) + "-" + (day);
    $("#date").val(today);

    //取消禁用用户名文本框
    $("#user").attr("disabled", false);

    var user_name = $("#user").val().trim();
    var password = $("#pwd").val().trim();
    var logonpwd = sessionStorage.getItem("pwd");
    //只在添加操作或用户改动密码框后的时候才参与加密
    if (logonpwd != password) {
        var hash = hex_md5(user_name + password + "3i61fghb");
        $("#pwd").val(hash);
        $("#password_confirm").val(hash)
    }

    var frm_UserInfo = $("#frm_UserInfo");
    var userInfo = frm_UserInfo.serializeJSON();
    //不参与传值
    delete userInfo["password_confirm"];
    var jsonString = JSON.stringify(userInfo);
    //alert(jsonString);
    //console.log(jsonString);
    return jsonString;
}

//解析单个用户信息并赋值
function parseJsonUserInfo(data) {
    var jsonObject = jQuery.parseJSON(data);
    //修改时赋值修改标识,并禁用修改用户名称
    if (uid != '') {
        $("#oper").val(jsonObject.id);
        $("#user").attr("disabled", true);
    }
    $("#user").val(jsonObject.username);
    $("#email").val(jsonObject.email);
    $("#pwd").val(jsonObject.password);
    $("#password_confirm").val(jsonObject.password);
    /*
    $("#user").val(jsonObject["user"]);
    $("#email").val(jsonObject["email"]);
    $("#pwd").val(jsonObject["pwd"]);
    $("#password_confirm").val(jsonObject["pwd"]);
    sessionStorage.setItem("pwd", jsonObject["pwd"]);*/
}

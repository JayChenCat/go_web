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
});


var upgradeFileReader;

function handleFiles(files) {
    if (window.FileReader) {
        if (files.length) {
            var file = files[0];
            upgradeFileReader = new FileReader();
            upgradeFileReader.onloadend = function (e) {
                if (upgradeFileReader.result != null) {
                    document.getElementById("fileOKImg").style.visibility = "visible";
                    document.getElementById("step").innerHTML = "升级未开始";
                } else {
                    document.getElementById("fileOKImg").style.visibility = "hidden";
                    layer.alert("读取文件失败！");
                }
            }
            upgradeFileReader.readAsArrayBuffer(file);
        }
    } else {
        layer.alert("FileReader not supported by your browser!");
    }
}

//提交升级
function submitUpgrade() {
    loadIcon("向设备提交升级文件....", 2000);
    /*方法二*/
    var formData = new FormData();
    formData.append("uploadfile", $("#uploadfile")[0].files[0]);
    formData.append("remarks",document.getElementById("remarks").value);
    document.getElementById("step").innerHTML = "正在更新设备...";
    //alert(document.getElementById("remarks").value)
    $.ajax(
        {
            url:host + ":8083/upgrade",
            type: "post",  // 发送方式
            cache: false,
            data:formData,  // 发送的数据
            contentType: false,
            dataType: "json",
            processData: false,  //告诉浏览器不要对你的数据进行任何处理
            success: (result) => {
                document.getElementById("step").innerHTML = result.msg;
                //layer.msg(result.msg);
            },
            error: (xhr) => {  // 发送失败的回调函数
                //console.log("fail");
                document.getElementById("step").innerHTML = "通讯出错！status=" + xhr.status;
            }
        });

    //1秒后自动关闭加载图标
    setTimeout(function () {
        //关闭加载图标
        closeLoadIcon();
    }, 1000);
}

//下载日志
function downloadLog() {
    //var a = window.open("applog.log");
    var pom = document.createElement('a');
    pom.setAttribute('href', 'applog.log');
    pom.setAttribute('download', "applog.log");
    if (document.createEvent) {
        var event = document.createEvent('MouseEvents');
        event.initEvent('click', true, true);
        pom.dispatchEvent(event);
    } else {
        pom.click();
    }
}


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

    //加载硬件配置
    //loadDevConfig();
});





function parseBaseInfoJson(jsonstr) {
    var jsonObject = jQuery.parseJSON(jsonstr);
    document.getElementById("snText").innerHTML = jsonObject["sn"];
    document.getElementById("moduleText").innerHTML = jsonObject["module"];
    document.getElementById("ipText").innerHTML = jsonObject["ip"];
}

function parseVersionJson(jsonstr) {
    var jsonObject = jQuery.parseJSON(jsonstr);
    document.getElementById("armVersionText").innerHTML = jsonObject["armVersion"];
    document.getElementById("upgradeVersionText").innerHTML = jsonObject["upgradeVersion"];
    document.getElementById("webServerText").innerHTML = jsonObject["webServer"];
}

function loadDevConfig()
{
    loadIcon("正在获取设备配置中...", 2000);
    var url1 = "http://" + ip + ":10001/getbasedeviceinfo";
    var xhr1 = createXHR();
    xhr1.open("GET", url1, true);
    xhr1.send();
    xhr1.onreadystatechange = function () {
        if (xhr1.readyState == 4) {
            if (xhr1.status == 200) {
                parseBaseInfoJson(xhr1.responseText);
                //$("#warningHeader").css("display", "none");
            } else {
                //$("#warningHeader").css("display", "block");
                layer.msg("DMS设备运行不正常，请检查设备。", { icon: 2 });
            }
        }
    }
    var url2 = "http://" + ip + ":10001/getversion";
    var xhr2 = createXHR();
    xhr2.open("GET", url2, true);
    xhr2.send();
    xhr2.onreadystatechange = function () {
        if (xhr2.readyState == 4) {
            if (xhr2.status == 200) {
                closeLoadIcon();
                parseVersionJson(xhr2.responseText);
            } else {
                layer.msg("DMS设备运行不正常，请检查设备。",{ icon:2 });
            }
        }
    }
}


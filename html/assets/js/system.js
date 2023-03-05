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

        /*次导航*/
        var mainNavSwiper = new Swiper('#nav_swiper', {
            slidesPerView: 'auto',
            navigation: {
                nextEl: '#nav_prev',
                prevEl: '#nav_next',
            },
        });
    });
    $(".idTabs").idTabs();
    isVerifyAdministrators();
    getConfiguration();
});


function isValidIp(input) {
    return /^((2[0-4]\d|25[0-5]|[01]?\d\d?)\.){3}(2[0-4]\d|25[0-5]|[01]?\d\d?)$/.test(input)
}

//解析获取的设备配置
function parseJson(jsonstr) {
    //alert(jsonstr);
    var jsonObject = jQuery.parseJSON(jsonstr);
    GetAlarmSetting(jsonObject);
    GetAccountmanager(jsonObject);
    GetCommunicationsettings(jsonObject);
    GetSystemSetting(jsonObject);
}
//获取报警设置信息
function GetAlarmSetting(jsonObject){
    var alarmSetting = jsonObject["AlarmSetting"];
    var fatiguedAlarmSettingObj = alarmSetting["FatiguedAlarmSetting"];
    var fatiguedAlarmSw=ToBool(fatiguedAlarmSettingObj.alarmSw);
    var fatiguedAlarmSoundSw=ToBool(fatiguedAlarmSettingObj.alarmSoundSw);
    var fatiguedUploadingPicSw=ToBool(fatiguedAlarmSettingObj.uploadingPicSw);
    var fatiguedUploadingVideoSw=ToBool(fatiguedAlarmSettingObj.uploadingVideoSw);
    var fatiguedSpeed=fatiguedAlarmSettingObj.speed;
    var fatiguedTime=fatiguedAlarmSettingObj.time;
    var fatiguedTTS=fatiguedAlarmSettingObj.tts;
    var fatiguedCD=fatiguedAlarmSettingObj.cd;
    document.getElementById("fatiguedSpeed_Low").value=fatiguedSpeed;
    document.getElementById("fatiguedSpeed_Medium").value=fatiguedSpeed;
    document.getElementById("fatiguedSpeed_Serious").value=fatiguedSpeed;
    document.getElementById("fatiguedTime_Low").value=fatiguedTime;
    document.getElementById("fatiguedTime_Medium").value=fatiguedTime;
    document.getElementById("fatiguedTime_Serious").value=fatiguedTime;
    document.getElementById("fatiguedCD_Low").value=0;
    document.getElementById("fatiguedCD_Medium").value=0;
    document.getElementById("fatiguedCD_Serious").value=0;
    if(fatiguedAlarmSoundSw)
    {
        document.getElementById("fatiguedAlarmSoundSw_true").checked=true;
        document.getElementById("fatiguedAlarmSoundSw_false").checked=false;
    }
    else
    {
        document.getElementById("fatiguedAlarmSoundSw_false").checked=true;
        document.getElementById("fatiguedAlarmSoundSw_true").checked=false;
    }
    if(fatiguedUploadingPicSw)
    {
        document.getElementById("fatiguedUploadingPicSw_true").checked=true;
        document.getElementById("fatiguedUploadingPicSw_false").checked=false;
    }
    else
    {
        document.getElementById("fatiguedUploadingPicSw_false").checked=true;
        document.getElementById("fatiguedUploadingPicSw_true").checked=false;
    }
    if(fatiguedUploadingVideoSw)
    {
        document.getElementById("fatiguedUploadingVideoSw_true").checked=true;
        document.getElementById("fatiguedUploadingVideoSw_false").checked=false;
    }
    else
    {
        document.getElementById("fatiguedUploadingVideoSw_false").checked=true;
        document.getElementById("fatiguedUploadingVideoSw_true").checked=false;
    }
    if(fatiguedAlarmSw)
    {
        document.getElementById("fatiguedAlarmSw_true").checked=true;
        document.getElementById("fatiguedAlarmSw_false").checked=false;
    }
    else
    {
        document.getElementById("fatiguedAlarmSw_false").checked=true;
        document.getElementById("fatiguedAlarmSw_true").checked=false;
    }
    document.getElementById("fatiguedTTS").value=fatiguedTTS;

    var CallingAlarmSettingObj = alarmSetting["CallingAlarmSetting"];
    var callingAlarmSw=ToBool(CallingAlarmSettingObj.alarmSw);
    var callingAlarmSoundSw=ToBool(CallingAlarmSettingObj.alarmSoundSw);
    var callingUploadingPicSw=ToBool(CallingAlarmSettingObj.uploadingPicSw);
    var callingUploadingVideoSw=ToBool(CallingAlarmSettingObj.uploadingVideoSw);
    var callingSpeed=CallingAlarmSettingObj.speed;
    var callingTime=CallingAlarmSettingObj.time;
    var callingTTS=CallingAlarmSettingObj.tts;
    var callingCD=CallingAlarmSettingObj.cd;

    document.getElementById("callingSpeed").value=callingSpeed;
    document.getElementById("callingTime").value=callingTime;
    document.getElementById("callingCD").value=0;
    if(callingAlarmSoundSw)
    {
        document.getElementById("callingAlarmSoundSw_true").checked=true;
        document.getElementById("callingAlarmSoundSw_false").checked=false;
    }
    else
    {
        document.getElementById("callingAlarmSoundSw_false").checked=true;
        document.getElementById("callingAlarmSoundSw_true").checked=false;
    }
    if(callingUploadingPicSw)
    {
        document.getElementById("callingUploadingPicSw_true").checked=true;
        document.getElementById("callingUploadingPicSw_false").checked=false;
    }
    else
    {
        document.getElementById("callingUploadingPicSw_false").checked=true;
        document.getElementById("callingUploadingPicSw_true").checked=false;
    }
    if(callingUploadingVideoSw)
    {
        document.getElementById("callingUploadingVideoSw_true").checked=true;
        document.getElementById("callingUploadingVideoSw_false").checked=false;
    }
    else
    {
        document.getElementById("callingUploadingVideoSw_false").checked=true;
        document.getElementById("callingUploadingVideoSw_true").checked=false;
    }
    if(callingAlarmSw)
    {
        document.getElementById("callingAlarmSw_true").checked=true;
        document.getElementById("callingAlarmSw_false").checked=false;
    }
    else
    {
        document.getElementById("callingAlarmSw_false").checked=true;
        document.getElementById("callingAlarmSw_true").checked=false;
    }
    document.getElementById("callingTTS").value=callingTTS;

    var LookaoundAlarmSettingObj = alarmSetting["LookaoundAlarmSetting"];
    var lookaoundAlarmSw=ToBool(LookaoundAlarmSettingObj.alarmSw);
    var lookaoundAlarmSoundSw=ToBool(LookaoundAlarmSettingObj.alarmSoundSw);
    var lookaoundUploadingPicSw=ToBool(LookaoundAlarmSettingObj.uploadingPicSw);
    var lookaoundUploadingVideoSw=ToBool(LookaoundAlarmSettingObj.uploadingVideoSw);
    var lookaoundSpeed=LookaoundAlarmSettingObj.speed;
    var lookaoundTime=LookaoundAlarmSettingObj.time;
    var lookaoundTTS=LookaoundAlarmSettingObj.tts;
    var lookaoundCD=LookaoundAlarmSettingObj.cd;
    document.getElementById("laSpeed").value=callingSpeed;
    document.getElementById("laTime").value=callingTime;
    document.getElementById("laCD").value=0;
    if(lookaoundAlarmSoundSw)
    {
        document.getElementById("laAlarmSoundSw_true").checked=true;
        document.getElementById("laAlarmSoundSw_false").checked=false;
    }
    else
    {
        document.getElementById("laAlarmSoundSw_false").checked=true;
        document.getElementById("laAlarmSoundSw_true").checked=false;
    }
    if(lookaoundUploadingPicSw)
    {
        document.getElementById("laUploadingPicSw_true").checked=true;
        document.getElementById("laUploadingPicSw_false").checked=false;
    }
    else
    {
        document.getElementById("laUploadingPicSw_false").checked=true;
        document.getElementById("laUploadingPicSw_true").checked=false;
    }
    if(lookaoundUploadingVideoSw)
    {
        document.getElementById("laUploadingVideoSw_true").checked=true;
        document.getElementById("laUploadingVideoSw_false").checked=false;
    }
    else
    {
        document.getElementById("laUploadingVideoSw_false").checked=true;
        document.getElementById("laUploadingVideoSw_true").checked=false;
    }
    if(lookaoundAlarmSw)
    {
        document.getElementById("laAlarmSw_true").checked=true;
        document.getElementById("laAlarmSw_false").checked=false;
    }
    else
    {
        document.getElementById("laAlarmSw_false").checked=true;
        document.getElementById("laAlarmSw_true").checked=false;
    }
    document.getElementById("laTTS").value=lookaoundTTS;

    var SmokingAlarmSettingObj = alarmSetting["SmokingAlarmSetting"];
    var smokingAlarmSw=ToBool(SmokingAlarmSettingObj.alarmSw);
    var smokingAlarmSoundSw=ToBool(SmokingAlarmSettingObj.alarmSoundSw);
    var smokingUploadingPicSw=ToBool(SmokingAlarmSettingObj.uploadingPicSw);
    var smokingUploadingVideoSw=ToBool(SmokingAlarmSettingObj.uploadingVideoSw);
    var smokingSpeed=SmokingAlarmSettingObj.speed;
    var smokingTime=SmokingAlarmSettingObj.time;
    var smokingTTS=SmokingAlarmSettingObj.tts;
    var smokingCD=SmokingAlarmSettingObj.cd;
    document.getElementById("smokingSpeed").value=smokingSpeed;
    document.getElementById("smokingTime").value=smokingTime;
    document.getElementById("smokingCD").value=0;
    if(smokingAlarmSoundSw)
    {
        document.getElementById("smokingAlarmSoundSw_true").checked=true;
        document.getElementById("smokingAlarmSoundSw_false").checked=false;
    }
    else
    {
        document.getElementById("smokingAlarmSoundSw_false").checked=true;
        document.getElementById("smokingAlarmSoundSw_true").checked=false;
    }
    if(smokingUploadingPicSw)
    {
        document.getElementById("smokingUploadingPicSw_true").checked=true;
        document.getElementById("smokingUploadingPicSw_false").checked=false;
    }
    else
    {
        document.getElementById("smokingUploadingPicSw_false").checked=true;
        document.getElementById("smokingUploadingPicSw_true").checked=false;
    }
    if(smokingUploadingVideoSw)
    {
        document.getElementById("smokingUploadingVideoSw_true").checked=true;
        document.getElementById("smokingUploadingVideoSw_false").checked=false;
    }
    else
    {
        document.getElementById("smokingUploadingVideoSw_false").checked=true;
        document.getElementById("smokingUploadingVideoSw_true").checked=false;
    }
    if(smokingAlarmSw)
    {
        document.getElementById("smokingAlarmSw_true").checked=true;
        document.getElementById("smokingAlarmSw_false").checked=false;
    }
    else
    {
        document.getElementById("smokingAlarmSw_false").checked=true;
        document.getElementById("smokingAlarmSw_true").checked=false;
    }
    document.getElementById("smokingTTS").value=smokingTTS;


    var YawnAlarmSettingObj = alarmSetting["YawnAlarmSetting"];
    var yawnAlarmSw=ToBool(YawnAlarmSettingObj.alarmSw);
    var yawnAlarmSoundSw=ToBool(YawnAlarmSettingObj.alarmSoundSw);
    var yawnUploadingPicSw=ToBool(YawnAlarmSettingObj.uploadingPicSw);
    var yawnUploadingVideoSw=ToBool(YawnAlarmSettingObj.uploadingVideoSw);
    var yawnSpeed=YawnAlarmSettingObj.speed;
    var yawnTTS=YawnAlarmSettingObj.tts;
    var yawnCD=YawnAlarmSettingObj.cd;
    document.getElementById("yawnSpeed").value=yawnSpeed;
    document.getElementById("yawnCD").value=0;
    if(yawnAlarmSoundSw)
    {
        document.getElementById("yawnAlarmSoundSw_true").checked=true;
        document.getElementById("yawnAlarmSoundSw_false").checked=false;
    }
    else
    {
        document.getElementById("yawnAlarmSoundSw_false").checked=true;
        document.getElementById("yawnAlarmSoundSw_true").checked=false;
    }
    if(yawnUploadingPicSw)
    {
        document.getElementById("yawnUploadingPicSw_true").checked=true;
        document.getElementById("yawnUploadingPicSw_false").checked=false;
    }
    else
    {
        document.getElementById("yawnUploadingPicSw_false").checked=true;
        document.getElementById("yawnUploadingPicSw_true").checked=false;
    }
    if(yawnUploadingVideoSw)
    {
        document.getElementById("yawnUploadingVideoSw_true").checked=true;
        document.getElementById("yawnUploadingVideoSw_false").checked=false;
    }
    else
    {
        document.getElementById("yawnUploadingVideoSw_false").checked=true;
        document.getElementById("yawnUploadingVideoSw_true").checked=false;
    }
    if(yawnAlarmSw)
    {
        document.getElementById("yawnAlarmSw_true").checked=true;
        document.getElementById("yawnAlarmSw_false").checked=false;
    }
    else
    {
        document.getElementById("yawnAlarmSw_false").checked=true;
        document.getElementById("yawnAlarmSw_true").checked=false;
    }
    document.getElementById("yawnTTS").value=yawnTTS;

    var NoFaceAlarmSettingObj = alarmSetting["NoFaceAlarmSetting"];
    var noFaceAlarmSw=ToBool(NoFaceAlarmSettingObj.alarmSw);
    var noFaceAlarmSoundSw=ToBool(NoFaceAlarmSettingObj.alarmSoundSw);
    var noFaceUploadingPicSw=ToBool(NoFaceAlarmSettingObj.uploadingPicSw);
    var noFaceUploadingVideoSw=ToBool(NoFaceAlarmSettingObj.uploadingVideoSw);
    var noFaceSpeed=NoFaceAlarmSettingObj.speed;
    var noFaceTime=NoFaceAlarmSettingObj.time;
    var noFaceTTS=NoFaceAlarmSettingObj.tts;
    var noFaceCD=NoFaceAlarmSettingObj.cd;
    document.getElementById("noFaceSpeed").value=noFaceSpeed;
    document.getElementById("noFaceTime").value=noFaceTime;
    document.getElementById("noFaceCD").value=0;
    if(noFaceAlarmSoundSw)
    {
        document.getElementById("noFaceAlarmSoundSw_true").checked=true;
        document.getElementById("noFaceAlarmSoundSw_false").checked=false;
    }
    else
    {
        document.getElementById("noFaceAlarmSoundSw_false").checked=true;
        document.getElementById("noFaceAlarmSoundSw_true").checked=false;
    }
    if(noFaceUploadingPicSw)
    {
        document.getElementById("noFaceUploadingPicSw_true").checked=true;
        document.getElementById("noFaceUploadingPicSw_false").checked=false;
    }
    else
    {
        document.getElementById("noFaceUploadingPicSw_false").checked=true;
        document.getElementById("noFaceUploadingPicSw_true").checked=false;
    }
    if(noFaceUploadingVideoSw)
    {
        document.getElementById("noFaceUploadingVideoSw_true").checked=true;
        document.getElementById("noFaceUploadingVideoSw_false").checked=false;
    }
    else
    {
        document.getElementById("noFaceUploadingVideoSw_false").checked=true;
        document.getElementById("noFaceUploadingVideoSw_true").checked=false;
    }
    if(noFaceAlarmSw)
    {
        document.getElementById("noFaceAlarmSw_true").checked=true;
        document.getElementById("noFaceAlarmSw_false").checked=false;
    }
    else
    {
        document.getElementById("noFaceAlarmSw_false").checked=true;
        document.getElementById("noFaceAlarmSw_true").checked=false;
    }
    document.getElementById("noFaceTTS").value=noFaceTTS;

    var MouthAlarmSettingObj = alarmSetting["MouthAlarmSetting"];
    var mouthAlarmSw=ToBool(MouthAlarmSettingObj.alarmSw);
    var mouthAlarmSoundSw=ToBool(MouthAlarmSettingObj.alarmSoundSw);
    var mouthUploadingPicSw=ToBool(MouthAlarmSettingObj.uploadingPicSw);
    var mouthUploadingVideoSw=ToBool(MouthAlarmSettingObj.uploadingVideoSw);
    var mouthSpeed=MouthAlarmSettingObj.speed;
    var mouthTime=MouthAlarmSettingObj.time;
    var mouthTTS=MouthAlarmSettingObj.tts;
    var mouthCD=MouthAlarmSettingObj.cd;
    document.getElementById("mouthSpeed").value=mouthSpeed;
    document.getElementById("mouthTime").value=mouthTime;
    document.getElementById("mouthCD").value=0;
    if(mouthAlarmSoundSw)
    {
        document.getElementById("mouthAlarmSoundSw_true").checked=true;
        document.getElementById("mouthAlarmSoundSw_false").checked=false;
    }
    else
    {
        document.getElementById("mouthAlarmSoundSw_false").checked=true;
        document.getElementById("mouthAlarmSoundSw_true").checked=false;
    }
    if(mouthUploadingPicSw)
    {
        document.getElementById("mouthUploadingPicSw_true").checked=true;
        document.getElementById("mouthUploadingPicSw_false").checked=false;
    }
    else
    {
        document.getElementById("mouthUploadingPicSw_false").checked=true;
        document.getElementById("mouthUploadingPicSw_true").checked=false;
    }
    if(mouthUploadingVideoSw)
    {
        document.getElementById("mouthUploadingVideoSw_true").checked=true;
        document.getElementById("mouthUploadingVideoSw_false").checked=false;
    }
    else
    {
        document.getElementById("mouthUploadingVideoSw_false").checked=true;
        document.getElementById("mouthUploadingVideoSw_true").checked=false;
    }
    if(noFaceAlarmSw)
    {
        document.getElementById("mouthAlarmSw_true").checked=true;
        document.getElementById("mouthAlarmSw_false").checked=false;
    }
    else
    {
        document.getElementById("mouthAlarmSw_false").checked=true;
        document.getElementById("mouthAlarmSw_true").checked=false;
    }
    document.getElementById("mouthTTS").value=mouthTTS;

    var OcclusionAlarmSettingObj = alarmSetting["OcclusionAlarmSetting"];
    var occlusionAlarmSw=ToBool(OcclusionAlarmSettingObj.alarmSw);
    var occlusionSoundSw=ToBool(OcclusionAlarmSettingObj.alarmSoundSw);
    var occlusionUploadingPicSw=ToBool(OcclusionAlarmSettingObj.uploadingPicSw);
    var occlusionUploadingVideoSw=ToBool(OcclusionAlarmSettingObj.uploadingVideoSw);
    var occlusionSpeed=OcclusionAlarmSettingObj.speed;
    var occlusionTime=OcclusionAlarmSettingObj.time;
    var occlusionTTS=OcclusionAlarmSettingObj.tts;
    var occlusionCD=OcclusionAlarmSettingObj.cd;
    document.getElementById("occlusionSpeed").value=occlusionSpeed;
    document.getElementById("occlusionTime").value=occlusionTime;
    document.getElementById("occlusionCD").value=0;
    if(occlusionSoundSw)
    {
        document.getElementById("occlusionAlarmSoundSw_true").checked=true;
        document.getElementById("occlusionAlarmSoundSw_false").checked=false;
    }
    else
    {
        document.getElementById("occlusionAlarmSoundSw_false").checked=true;
        document.getElementById("occlusionAlarmSoundSw_true").checked=false;
    }
    if(occlusionUploadingPicSw)
    {
        document.getElementById("occlusionUploadingPicSw_true").checked=true;
        document.getElementById("occlusionUploadingPicSw_false").checked=false;
    }
    else
    {
        document.getElementById("occlusionUploadingPicSw_false").checked=true;
        document.getElementById("occlusionUploadingPicSw_true").checked=false;
    }
    if(occlusionUploadingVideoSw)
    {
        document.getElementById("occlusionUploadingVideoSw_true").checked=true;
        document.getElementById("occlusionUploadingVideoSw_false").checked=false;
    }
    else
    {
        document.getElementById("occlusionUploadingVideoSw_false").checked=true;
        document.getElementById("occlusionUploadingVideoSw_true").checked=false;
    }
    if(occlusionAlarmSw)
    {
        document.getElementById("occlusionAlarmSw_true").checked=true;
        document.getElementById("occlusionAlarmSw_false").checked=false;
    }
    else
    {
        document.getElementById("occlusionAlarmSw_false").checked=true;
        document.getElementById("occlusionAlarmSw_true").checked=false;
    }
    document.getElementById("occlusionTTS").value=occlusionTTS;


    var HandAlarmSettingObj = alarmSetting["HandAlarmSetting"];
    var handAlarmSw=ToBool(HandAlarmSettingObj.alarmSw);
    var handSoundSw=ToBool(HandAlarmSettingObj.alarmSoundSw);
    var handUploadingPicSw=ToBool(HandAlarmSettingObj.uploadingPicSw);
    var handUploadingVideoSw=ToBool(HandAlarmSettingObj.uploadingVideoSw);
    var handSpeed=HandAlarmSettingObj.speed;
    var handTime=HandAlarmSettingObj.time;
    var handTTS=HandAlarmSettingObj.tts;
    var handCD=HandAlarmSettingObj.cd;
    document.getElementById("handSpeed").value=handSpeed;
    document.getElementById("handTime").value=handTime;
    document.getElementById("handCD").value=0;
    if(handSoundSw)
    {
        document.getElementById("handAlarmSoundSw_true").checked=true;
        document.getElementById("handAlarmSoundSw_false").checked=false;
    }
    else
    {
        document.getElementById("handAlarmSoundSw_false").checked=true;
        document.getElementById("handAlarmSoundSw_true").checked=false;
    }
    if(handUploadingPicSw)
    {
        document.getElementById("handUploadingPicSw_true").checked=true;
        document.getElementById("handUploadingPicSw_false").checked=false;
    }
    else
    {
        document.getElementById("handUploadingPicSw_false").checked=true;
        document.getElementById("handUploadingPicSw_true").checked=false;
    }
    if(handUploadingVideoSw)
    {
        document.getElementById("handUploadingVideoSw_true").checked=true;
        document.getElementById("handUploadingVideoSw_false").checked=false;
    }
    else
    {
        document.getElementById("handUploadingVideoSw_false").checked=true;
        document.getElementById("handUploadingVideoSw_true").checked=false;
    }
    if(handAlarmSw)
    {
        document.getElementById("handAlarmSw_true").checked=true;
        document.getElementById("handAlarmSw_false").checked=false;
    }
    else
    {
        document.getElementById("handAlarmSw_false").checked=true;
        document.getElementById("handAlarmSw_true").checked=false;
    }
    document.getElementById("handTTS").value=handTTS;

    var ArmAlarmSettingObj = alarmSetting["ArmAlarmSetting"];
    var armAlarmSw=ToBool(ArmAlarmSettingObj.alarmSw);
    var armSoundSw=ToBool(ArmAlarmSettingObj.alarmSoundSw);
    var armUploadingPicSw=ToBool(ArmAlarmSettingObj.uploadingPicSw);
    var armUploadingVideoSw=ToBool(ArmAlarmSettingObj.uploadingVideoSw);
    var armSpeed=ArmAlarmSettingObj.speed;
    var armTime=ArmAlarmSettingObj.time;
    var armTTS=ArmAlarmSettingObj.tts;
    var armCD=ArmAlarmSettingObj.cd;
    document.getElementById("armSpeed").value=armSpeed;
    document.getElementById("armTime").value=armTime;
    document.getElementById("armCD").value=0;
    if(armSoundSw)
    {
        document.getElementById("armAlarmSoundSw_true").checked=true;
        document.getElementById("armAlarmSoundSw_false").checked=false;
    }
    else
    {
        document.getElementById("armAlarmSoundSw_false").checked=true;
        document.getElementById("armAlarmSoundSw_true").checked=false;
    }
    if(armUploadingPicSw)
    {
        document.getElementById("armUploadingPicSw_true").checked=true;
        document.getElementById("armUploadingPicSw_false").checked=false;
    }
    else
    {
        document.getElementById("armUploadingPicSw_false").checked=true;
        document.getElementById("armUploadingPicSw_true").checked=false;
    }
    if(armUploadingVideoSw)
    {
        document.getElementById("armUploadingVideoSw_true").checked=true;
        document.getElementById("armUploadingVideoSw_false").checked=false;
    }
    else
    {
        document.getElementById("armUploadingVideoSw_false").checked=true;
        document.getElementById("armUploadingVideoSw_true").checked=false;
    }
    if(armAlarmSw)
    {
        document.getElementById("armAlarmSw_true").checked=true;
        document.getElementById("armAlarmSw_false").checked=false;
    }
    else
    {
        document.getElementById("armAlarmSw_false").checked=true;
        document.getElementById("armAlarmSw_true").checked=false;
    }
    document.getElementById("armTTS").value=armTTS;


    var BlockingAlarmSettingObj = alarmSetting["BlockingAlarmSetting"];
    var blockingAlarmSw=ToBool(BlockingAlarmSettingObj.alarmSw);
    var blockingSoundSw=ToBool(BlockingAlarmSettingObj.alarmSoundSw);
    var blockingUploadingPicSw=ToBool(BlockingAlarmSettingObj.uploadingPicSw);
    var blockingUploadingVideoSw=ToBool(BlockingAlarmSettingObj.uploadingVideoSw);
    var blockingSpeed=BlockingAlarmSettingObj.speed;
    var blockingTime=BlockingAlarmSettingObj.time;
    var blockingTTS=BlockingAlarmSettingObj.tts;
    var blockingCD=BlockingAlarmSettingObj.cd;
    document.getElementById("blocingSpeed").value=blockingSpeed;
    document.getElementById("blocingTime").value=blockingTime;
    document.getElementById("blocingCD").value=0;
    if(blockingAlarmSw)
    {
        document.getElementById("blocingAlarmSw_true").checked=true;
        document.getElementById("blocingAlarmSw_false").checked=false;
    }
    else
    {
        document.getElementById("blocingAlarmSw_false").checked=true;
        document.getElementById("blocingAlarmSw_true").checked=false;
    }
    document.getElementById("blocingTTS").value=blockingTTS;

    var drivingrecordSettingObj = alarmSetting["DrivingrecordSetting"];
    document.getElementById("Ch1_IsAlarmDvr_true").checked=ToBool(drivingrecordSettingObj[0].isOpen);
    document.getElementById("Ch2_IsAlarmDvr_true").checked=ToBool(drivingrecordSettingObj[1].isOpen);
}

//获取账号管理信息
function GetAccountmanager(jsonObject){
    var accountSetting = jsonObject["AccountSetting"];
    $("#ftpUser").val(accountSetting[0].user);
    $("#ftpPwd").val(accountSetting[0].Pwd);
    $("#ocsFtpUser").val(accountSetting[1].user);
    $("#ocsFtpPwd").val(accountSetting[1].Pwd);
    $("#rtspUser").val(accountSetting[2].user);
    $("#rtspPwd").val(accountSetting[2].Pwd);
}
//获取通讯设置信息
function GetCommunicationsettings(jsonObject){
    var communicationSetting = jsonObject["CommunicationSetting"];
    $("#localIP").val(communicationSetting[0].ip);
    $("#localPort").val(communicationSetting[0].port);
    $("#ocsIP").val(communicationSetting[1].ip);
    $("#ocsPort").val(communicationSetting[1].port);
}
//获取系统设置信息
function GetSystemSetting(jsonObject){
    var systemSetting = jsonObject["SystemSetting"];
    var cameraSettingObj = systemSetting["cameraSetting"];
    var carsettingObj = systemSetting["carsetting"];
    var peripheralsettingsObj = systemSetting["peripheralsettings"];
    var soundsettingObj = systemSetting["soundsetting"];
    document.getElementById("locomotiveID").value=carsettingObj.locomotiveID;
    document.getElementById("HostSide").value=carsettingObj.HostSide;
    document.getElementById("cam1Open_true").checked=ToBool(cameraSettingObj[0].camera_OnOff);
    document.getElementById("channelID0").value=cameraSettingObj[0].channelNo;
    document.getElementById("channelName0").value=cameraSettingObj[0].channelName;
    document.getElementById("cam2Open_true").checked=ToBool(cameraSettingObj[1].camera_OnOff);
    document.getElementById("channelID1").value=cameraSettingObj[1].channelNo;
    document.getElementById("channelName1").value=cameraSettingObj[1].channelName;
    document.getElementById("cam3Open_true").checked=ToBool(cameraSettingObj[2].camera_OnOff);
    document.getElementById("channelID2").value=cameraSettingObj[2].channelNo;
    document.getElementById("channelName2").value=cameraSettingObj[2].channelName;
    document.getElementById("gps_true").checked=ToBool(peripheralsettingsObj[0].peripheral_OnOff);
    document.getElementById("wifi_true").checked=ToBool(peripheralsettingsObj[1].peripheral_OnOff);
    document.getElementById("4G_true").checked=ToBool(peripheralsettingsObj[2].peripheral_OnOff);
    document.getElementById("volume").value=soundsettingObj.volume;
}


function ToBool(text){
  if(text==1){
      return true
  }
  else{
      return false
  }
}

//获取设备配置
function getConfiguration() {

    //测试使用
    //         $.getJSON("中车DMS配置json调整.json", function (data) {
    //	var jsonText = JSON.stringify(data);
    //             parseJson(jsonText);
    //         });
    //return;
    token=getToken();
    $.trim($("#token").val(token))
    loadIcon("获取设备配置中....", 2000);
    $.ajax({
        type: "post",
        url:host + ":8083/system",
        data: {},
        dataType: "json",
        success: function(result) {
            if(result.code==0) {
                var jsonText = JSON.stringify(result.data);
                //alert(jsonText)
                parseJson(jsonText);
                //关闭加载图标
                closeLoadIcon();
            }
            else {
                layer.msg(result.msg);
            }
        }
    });
    /*var url = "http://" + ip + ":10001/getconfiguration?module=dms";
    var xhr = createXHR();
    xhr.open("GET", url, true);
    xhr.send();
    xhr.onreadystatechange = function () {
        if (xhr.readyState == 4) {
            //alert("status=" + xhr.status);
            if (xhr.status == 200) {
                parseJson(xhr.responseText);
                //关闭加载图标
                closeLoadIcon();
            }
        }
    }*/
}

//提交设备配置
function sendConfiguration() {
    loadIcon("向设备提交配置中....", 2000);
    //1.报警设置信息参数
    var dataAramsetting={
        "tag":1,
        "token":$.trim($("#token").val()),
        //疲劳驾驶报警
        "fatiguedSpeed_Low" :document.getElementById("fatiguedSpeed_Low").value,
        "fatiguedSpeed_Medium" : document.getElementById("fatiguedSpeed_Medium").value,
        "fatiguedSpeed_Serious" : document.getElementById("fatiguedSpeed_Serious").value,
        "fatiguedTime_Low" : document.getElementById("fatiguedTime_Low").value,
        "fatiguedTime_Medium" :document.getElementById("fatiguedTime_Medium").value,
        "fatiguedTime_Serious" : document.getElementById("fatiguedTime_Serious").value,
        "fatiguedCD_Low" : document.getElementById("fatiguedCD_Low").value,
        "fatiguedCD_Medium" : document.getElementById("fatiguedCD_Medium").value,
        "fatiguedCD_Serious" :document.getElementById("fatiguedCD_Serious").value,
        "fatiguedAlarmSoundSw" : document.getElementById("fatiguedAlarmSoundSw_true").checked,
        "fatiguedUploadingPicSw" :document.getElementById("fatiguedUploadingPicSw_true").checked,
        "fatiguedUploadingVideoSw" : document.getElementById("fatiguedUploadingVideoSw_true").checked,
        "fatiguedTTS":document.getElementById("fatiguedTTS").value,
        "fatiguedAlarmSw":document.getElementById("fatiguedAlarmSw_true").checked,
        //打电话报警
        "callingSpeed":document.getElementById("callingSpeed").value,
        "callingTime":document.getElementById("callingTime").value,
        "callingCD": document.getElementById("callingCD").value,
        "callingAlarmSoundSw": document.getElementById("callingAlarmSoundSw_true").checked,
        "callingUploadingPicSw":document.getElementById("callingUploadingPicSw_true").checked,
        "callingUploadingVideoSw ": document.getElementById("callingUploadingVideoSw_true").checked,
        "callingTTS": document.getElementById("callingTTS").value,
        "callingAlarmSw": document.getElementById("callingAlarmSw_true").checked,
        //左顾右盼
        "lookaoundSpeed": document.getElementById("laSpeed").value,
        "lookaoundTime": document.getElementById("laTime").value,
        "lookaoundCD": document.getElementById("laCD").value,
        "lookaoundAlarmSoundSw": document.getElementById("laAlarmSoundSw_true").checked,
        "lookaoundUploadingPicSw": document.getElementById("laUploadingPicSw_true").checked,
        "lookaoundUploadingVideoSw": document.getElementById("laUploadingVideoSw_true").checked,
        "lookaoundTTS": document.getElementById("laTTS").value,
        "lookaoundAlarmSw":document.getElementById("laAlarmSw_true").checked,
        //抽烟
        "smokingSpeed":document.getElementById("smokingSpeed").value,
        "smokingTime":document.getElementById("smokingTime").value,
        "smokingCD": document.getElementById("smokingCD").value,
        "smokingAlarmSoundSw":document.getElementById("smokingAlarmSoundSw_true").checked,
        "smokingUploadingPicSw": document.getElementById("smokingUploadingPicSw_true").checked,
        "smokingUploadingVideoSw": document.getElementById("smokingUploadingVideoSw_true").checked,
        "smokingTTS": document.getElementById("smokingTTS").value,
        "smokingAlarmSw": document.getElementById("smokingAlarmSw_true").checked,
        //打哈欠
        "yawnSpeed": document.getElementById("yawnSpeed").value,
        "yawnCD": document.getElementById("yawnCD").value,
        "yawnAlarmSoundSw": document.getElementById("yawnAlarmSoundSw_true").checked,
        "yawnUploadingPicSw":document.getElementById("yawnUploadingPicSw_true").checked,
        "yawnUploadingVideoSw": document.getElementById("yawnUploadingVideoSw_true").checked,
        "yawnTTS": document.getElementById("yawnTTS").value,
        "yawnAlarmSw": document.getElementById("yawnAlarmSw_true").checked,
        //离岗
        "noFaceSpeed": document.getElementById("noFaceSpeed").value,
        "noFaceTime": document.getElementById("noFaceTime").value,
        "noFaceCD": document.getElementById("noFaceCD").value,
        "noFaceAlarmSoundSw": document.getElementById("noFaceAlarmSoundSw_true").checked,
        "noFaceUploadingPicSw": document.getElementById("noFaceUploadingPicSw_true").checked,
        "noFaceUploadingVideoSw": document.getElementById("noFaceUploadingVideoSw_true").checked,
        "noFaceTTS": document.getElementById("noFaceTTS").value,
        "noFaceAlarmSw": document.getElementById("noFaceAlarmSw_true").checked,
        //吃东西/闲聊
        "mouthSpeed": document.getElementById("mouthSpeed").value,
        "mouthTime": document.getElementById("mouthTime").value,
        "mouthCD": document.getElementById("mouthCD").value,
        "mouthAlarmSoundSw": document.getElementById("mouthAlarmSoundSw_true").checked,
        "mouthUploadingPicSw": document.getElementById("mouthUploadingPicSw_true").checked,
        "mouthUploadingVideoSw": document.getElementById("mouthUploadingVideoSw_true").checked,
        "mouthTTS": document.getElementById("mouthTTS").value,
        "mouthAlarmSw": document.getElementById("mouthAlarmSw_true").checked,
        //摄像头遮挡
        "occlusionSpeed": document.getElementById("occlusionSpeed").value,
        "occlusionTime": document.getElementById("occlusionTime").value,
        "occlusionCD": document.getElementById("occlusionCD").value,
        "occlusionAlarmSoundSw": document.getElementById("occlusionAlarmSoundSw_true").checked,
        "occlusionUploadingPicSw": document.getElementById("occlusionUploadingPicSw_true").checked,
        "occlusionUploadingVideoSw": document.getElementById("occlusionUploadingVideoSw_true").checked,
        "occlusionTTS": document.getElementById("occlusionTTS").value,
        "occlusionAlarmSw": document.getElementById("occlusionAlarmSw_true").checked,
        //手比前方
        "handSpeed": document.getElementById("handSpeed").value,
        "handTime": document.getElementById("handTime").value,
        "handCD": document.getElementById("handCD").value,
        "handAlarmSoundSw": document.getElementById("handAlarmSoundSw_true").checked,
        "handUploadingPicSw":document.getElementById("handUploadingPicSw_true").checked,
        "handUploadingVideoSw": document.getElementById("handUploadingVideoSw_true").checked,
        "handTTS": document.getElementById("handTTS").value,
        "handAlarmSw": document.getElementById("handAlarmSw_true").checked,
        //摇臂
        "armSpeed": document.getElementById("armSpeed").value,
        "armTime": document.getElementById("armTime").value,
        "armCD": document.getElementById("armCD").value,
        "armAlarmSoundSw": document.getElementById("armAlarmSoundSw_true").checked,
        "armUploadingPicSw":document.getElementById("armUploadingPicSw_true").checked,
        "armUploadingVideoSw": document.getElementById("armUploadingVideoSw_true").checked,
        "armTTS": document.getElementById("armTTS").value,
        "armAlarmSw": document.getElementById("armAlarmSw_true").checked,
        //红外阻断
        "blockingSpeed": document.getElementById("blocingSpeed").value,
        "blockingTime": document.getElementById("blocingTime").value,
        "blockingCD": document.getElementById("blocingCD").value,
        "blockingTTS": document.getElementById("blocingTTS").value,
        "blockingAlarmSw": document.getElementById("blocingAlarmSw_true").checked,
        //行驶记录
        "ch_Name1":"camera1",
        "isAlarmDvr1": document.getElementById("Ch1_IsAlarmDvr_true").checked,
        "ch_Name2":"camera2",
        "isAlarmDvr2": document.getElementById("Ch2_IsAlarmDvr_true").checked
    };

    //2.账户管理信息参数
    var dataAaccountmanager= {
        "tag":2,
        "token":$.trim($("#token").val()),
        "ftpUser":$("#ftpUser").val(),
        "ftpPwd":$("#ftpPwd").val(),
        "ocsFtpUser":$("#ocsFtpUser").val(),
        "ocsFtpPwd":$("#ocsFtpPwd").val(),
        "rtspUser":$("#rtspUser").val(),
        "rtspPwd":$("#rtspPwd").val(),
    };

    //3.通讯地址设置信息参数
    var dataCommunicationsettings= {
        "tag":3,
        "token":$.trim($("#token").val()),
        "localIP":$("#localIP").val(),
        "localPort":$("#localPort").val(),
        "ocsIP":$("#ocsIP").val(),
        "ocsPort":$("#ocsPort").val(),
    };

    //4.系统信息参数
    var dataSystemsettings= {
        "tag":4,
        "token":$.trim($("#token").val()),
        "locomotiveID":document.getElementById("locomotiveID").value,
        "hostSide": document.getElementById("HostSide").value,
        "cam1Open":document.getElementById("cam1Open_true").checked,
        "channelID0":document.getElementById("channelID0").value,
        "channelName0":document.getElementById("channelName0").value,
        "cam2pen":document.getElementById("cam2Open_true").checked,
        "channelID1":document.getElementById("channelID1").value,
        "channelName1":document.getElementById("channelName1").value,
        "cam3pen":document.getElementById("cam3Open_true").checked,
        "channelID2":document.getElementById("channelID2").value,
        "channelName2":document.getElementById("channelName2").value,
        "gps_isOpen":document.getElementById("gps_true").checked,
        "wifi_isOpen":document.getElementById("wifi_true").checked,
        "4G_isOpen":document.getElementById("4G_true").checked,
        "soundType":0,
        "volume":document.getElementById("volume").value
    };
    var dataSetting={}
    var selectText=$(".tab li a.selected").text();
    if(selectText.indexOf("报警设置")>0){
        dataSetting=dataAramsetting;
        CheckAramsetting();
    }
    else if(selectText.indexOf("账号管理")>0){
        dataSetting=dataAaccountmanager;

        var ftpUserText = document.getElementById("ftpUser").value == "undefined" ? "" : document.getElementById("ftpUser").value;
        var OcsFtpUserText = document.getElementById("ocsFtpUser").value == "undefined" ? "" : document.getElementById("ocsFtpUser").value;
        var RtspUserText = document.getElementById("rtspUser").value == "undefined" ? "" : document.getElementById("rtspUser").value;
        //验证字符中的空格
        if (ftpUserText.indexOf(" ") != -1) {
            alert("主机FTP账号的用户名中存在空格,请检查!");
            return;
        }
        if (OcsFtpUserText.indexOf(" ") != -1) {
            alert("OCS FTP账号的用户名中存在空格,请检查!");
            return;
        }
        if (RtspUserText.indexOf(" ") != -1) {
            alert("RTSP账号的用户名中存在空格,请检查!");
            return;
        }
    }
    else if(selectText.indexOf("通讯设置")>0){
        dataSetting=dataCommunicationsettings;
        if (!isValidIp(communicationSetting.LocalIP)) {
            alert("请输入正确的LocalIP地址!");
            return;
        }
        if (!isValidIp(communicationSetting.OcsIP)) {
            alert("请输入正确的OcsIP地址!");
            return;
        }
    }
    else if(selectText.indexOf("系统配置")>0){
        dataSetting=dataSystemsettings;
        if (!isValidIp(document.getElementById("cam1IP").value))
        {
            alert("请输入正确的IP地址!");
            return;
        }
        if (!isValidIp(document.getElementById("cam2IP").value))
        {
            alert("请输入正确的IP地址!");
            return;
        }
        if (!isValidIp(document.getElementById("cam3IP").value))
        {
            alert("请输入正确的IP地址!");
            return;
        }
    }
    SetButton(true,'#555555')
    $.ajax({
        type: "post",
        url:host + ":8083/setting",
        data: dataSetting,
        dataType: "json",
        success: function(result) {
            layer.msg(result.msg);
            SetButton(false,'#0072C6')
        }
    });

    //2秒后自动关闭加载图标
    setTimeout(function () {
        //关闭加载图标
        closeLoadIcon();
    }, 2000);
}

//验证报警参数
function CheckAramsetting(){
    //2022-03-23 添加
    if ((document.getElementById("fatiguedSpeed_Low").value < 0 || document.getElementById("fatiguedSpeed_Low").value > 120)
        || (document.getElementById("fatiguedSpeed_Medium").value < 0 || document.getElementById("fatiguedSpeed_Medium").value > 120)
        || (document.getElementById("fatiguedSpeed_Serious").value < 0 || document.getElementById("fatiguedSpeed_Serious").value > 120)
    ) {
        alert("疲劳驾驶报警触发车速超过正常范围(0~120km/h,包括0和120)！");
        return;
    }
    if ((document.getElementById("fatiguedTime_Low").value < 0 || document.getElementById("fatiguedTime_Low").value > 180)
        || (document.getElementById("fatiguedTime_Medium").value < 0 || document.getElementById("fatiguedTime_Medium").value > 180)
        || (document.getElementById("fatiguedTime_Serious").value < 0 || document.getElementById("fatiguedTime_Serious").value > 180)
    ) {
        alert("疲劳驾驶报警触发时间超过正常范围(0~180s,包括0和180)！");
        return;
    }
    //2022-03-23 添加
    if ((document.getElementById("callingSpeed").value < 0 || document.getElementById("callingSpeed").value > 120)) {
        alert("打电话报警触发车速超过正常范围(0~120km/h,包括0和120)！");
        return;
    }
    if ((document.getElementById("callingTime").value < 0 || document.getElementById("callingTime").value > 180)) {
        alert("打电话报警触发时间超过正常范围(0~180s,包括0和180)！");
        return;
    }
    //2022-03-23 添加
    if ((document.getElementById("laSpeed").value < 0 || document.getElementById("laSpeed").value > 120)) {
        alert("左顾右盼报警触发车速超过正常范围(0~120km/h,包括0和120)！");
        return;
    }
    if ((document.getElementById("laTime").value < 0 || document.getElementById("laTime").value > 180)) {
        alert("左顾右盼报警触发时间超过正常范围(0~180s,包括0和180)！");
        return;
    }
    //2022-03-23 添加
    if ((document.getElementById("smokingSpeed").value < 0 || document.getElementById("smokingSpeed").value > 120)) {
        alert("抽烟报警触发车速超过正常范围(0~120km/h,包括0和120)！");
        return;
    }
    if ((document.getElementById("smokingTime").value < 0 || document.getElementById("smokingTime").value > 180)) {
        alert("抽烟报警触发时间超过正常范围(0~180s,包括0和180)！");
        return;
    }
    if ((document.getElementById("yawnSpeed").value < 0 || document.getElementById("yawnSpeed").value > 120)) {
        alert("打哈欠报警触发车速超过正常范围(0~120km/h,包括0和120)！");
        return;
    }
    //2022-03-23 添加
    if ((document.getElementById("noFaceSpeed").value < 0 || document.getElementById("noFaceSpeed").value > 120)) {
        alert("离岗报警触发车速超过正常范围(0~120km/h,包括0和120)！");
        return;
    }
    if ((document.getElementById("noFaceTime").value < 0 || document.getElementById("noFaceTime").value > 180)) {
        alert("离岗报警触发时间超过正常范围(0~180s,包括0和180)！");
        return;
    }
    //2022-03-23 添加
    if ((document.getElementById("mouthSpeed").value < 0 || document.getElementById("mouthSpeed").value > 120)) {
        alert("吃东西/闲聊报警触发车速超过正常范围(0~120km/h,包括0和120)！");
        return;
    }
    if ((document.getElementById("mouthTime").value < 0 || document.getElementById("mouthTime").value > 180)) {
        alert("吃东西/闲聊报警触发时间超过正常范围(0~180s,包括0和180)！");
        return;
    }
    //2022-03-23 添加
    if ((document.getElementById("occlusionSpeed").value < 0 || document.getElementById("occlusionSpeed").value > 120)) {
        alert("摄像头遮挡报警触发车速超过正常范围(0~120km/h,包括0和120)！");
        return;
    }
    if ((document.getElementById("occlusionTime").value < 0 || document.getElementById("occlusionTime").value > 180)) {
        alert("摄像头遮挡报警触发时间超过正常范围(0~180s,包括0和180)！");
        return;
    }
    //2022-03-23 添加
    if ((document.getElementById("handSpeed").value < 0 || document.getElementById("handSpeed").value > 120)) {
        alert("手比前方报警触发车速超过正常范围(0~120km/h,包括0和120)！");
        return;
    }
    if ((document.getElementById("handTime").value < 0 || document.getElementById("handTime").value > 180)) {
        alert("手比前方报警触发时间超过正常范围(0~180s,包括0和180)！");
        return;
    }
    //2022-03-23 添加
    if ((document.getElementById("armSpeed").value < 0 || document.getElementById("armSpeed").value > 120)) {
        alert("摇臂报警触发车速超过正常范围(0~120km/h,包括0和120)！");
        return;
    }
    if ((document.getElementById("armTime").value < 0 || document.getElementById("armTime").value > 180)) {
        alert("摇臂报警触发时间超过正常范围(0~180s,包括0和180)！");
        return;
    }
    //2022-03-23 添加
    if ((document.getElementById("blocingSpeed").value < 0 || document.getElementById("blocingSpeed").value > 120)) {
        alert("红外阻断报警触发车速超过正常范围(0~120km/h,包括0和120)！");
        return;
    }
    if ((document.getElementById("blocingTime").value < 0 || document.getElementById("blocingTime").value > 180)) {
        alert("红外阻断报警触发时间超过正常范围(0~180s,包括0和180)！");
        return;
    }
}

function SetButton(enable,color){
    document.getElementById("commitBT").disabled = enable; //去掉不可点击
    document.getElementById("commitBT").style.backgroundColor = color;//'#0072C6'; //设置背景色
}


//鼠标离开文本框焦点事件  
function Textblurs() {
    var channelName0 = document.getElementById("channelName0").value;
    var channelName1 = document.getElementById("channelName1").value;
    var channelName2 = document.getElementById("channelName2").value;

    var channelName0Length = channelName0.length;
    var channelName1Length = channelName1.length;
    var channelName2Length = channelName2.length;

    var patternChinese = new RegExp("[\u4E00-\u9FA5]+");
    var patternEnglish = new RegExp("[A-Za-z]+");
    var patternNumber = /^[0-9]+.?[0-9]*/;
    //判断输入中文
    if (patternChinese.test(channelName0) || patternChinese.test(channelName1) || patternChinese.test(channelName2)) {

        if (channelName0Length > 8) {
            layer.msg("最长支持8个中文字符！");
            $("#channelName0").css("border", "3px solid red");
            document.getElementById("commitBT").disabled = true;//不可点击
            document.getElementById("commitBT").style.backgroundColor = '#555555';//设置背景色
            return;
        }
        else if (channelName1Length > 8) {
            layer.msg("最长支持8个中文字符！");
            $("#channelName1").css("border", "3px solid red");
            document.getElementById("commitBT").disabled = true;//不可点击
            document.getElementById("commitBT").style.backgroundColor = '#555555';//设置背景色
            return;
        }
        else if (channelName2Length > 8) {
            layer.msg("最长支持8个中文字符！");
            $("#channelName2").css("border", "3px solid red");
            document.getElementById("commitBT").disabled = true;//不可点击
            document.getElementById("commitBT").style.backgroundColor = '#555555';//设置背景色
            return;
        }
        else {
            $("#channelName0").css("border", "");
            $("#channelName1").css("border", "");
            $("#channelName2").css("border", "");
            document.getElementById("commitBT").disabled = false; //去掉不可点击
            document.getElementById("commitBT").style.backgroundColor = '#0072C6'; //设置背景色
        }
    }
    //判断输入英文
    if (patternEnglish.test(channelName0) || patternEnglish.test(channelName1) || patternEnglish.test(channelName2)) {

        if (channelName0Length > 24) {
            layer.msg("最长支持24个英文字符！");
            $("#channelName0").css("border", "3px solid red");
            document.getElementById("commitBT").disabled = true;//不可点击
            document.getElementById("commitBT").style.backgroundColor = '#555555';//设置背景色
            return;
        } else if (channelName1Length > 24) {
            layer.msg("最长支持24个英文字符！");
            $("#channelName1").css("border", "3px solid red");
            document.getElementById("commitBT").disabled = true;//不可点击
            document.getElementById("commitBT").style.backgroundColor = '#555555';//设置背景色
            return;
        }
        else if (channelName2Length > 24) {
            layer.msg("最长支持24个英文字符！");
            $("#channelName2").css("border", "3px solid red");
            document.getElementById("commitBT").disabled = true;//不可点击
            document.getElementById("commitBT").style.backgroundColor = '#555555';//设置背景色
            return;
        }
        else {
            $("#channelName0").css("border", "");
            $("#channelName1").css("border", "");
            $("#channelName2").css("border", "");
            document.getElementById("commitBT").disabled = false; //去掉不可点击
            document.getElementById("commitBT").style.backgroundColor = '#0072C6'; //设置背景色
        }
    }
    //判断输入数字
    if (patternNumber.test(channelName0) || patternNumber.test(channelName1) || patternNumber.test(channelName2)) {
        //layer.msg("通道名称不能以数字开头！");
        if (channelName0Length > 24) {
            layer.msg("最长支持24个英文字符！");
            $("#channelName0").css("border", "3px solid red");
            document.getElementById("commitBT").disabled = true;//不可点击
            document.getElementById("commitBT").style.backgroundColor = '#555555';//设置背景色
            return;
        } else if (channelName1Length > 24) {
            layer.msg("最长支持24个英文字符！");
            $("#channelName1").css("border", "3px solid red");
            document.getElementById("commitBT").disabled = true;//不可点击
            document.getElementById("commitBT").style.backgroundColor = '#555555';//设置背景色
            return;
        }
        else if (channelName2Length > 24) {
            layer.msg("最长支持24个英文字符！");
            $("#channelName2").css("border", "3px solid red");
            document.getElementById("commitBT").disabled = true;//不可点击
            document.getElementById("commitBT").style.backgroundColor = '#555555';//设置背景色
            return;
        }
        else {
            $("#channelName0").css("border", "");
            $("#channelName1").css("border", "");
            $("#channelName2").css("border", "");
            document.getElementById("commitBT").disabled = false; //去掉不可点击
            document.getElementById("commitBT").style.backgroundColor = '#0072C6'; //设置背景色
        }
    }
}
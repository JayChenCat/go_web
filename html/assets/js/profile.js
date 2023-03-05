/*
 * Start layui2 - webtpl v1.0.0 (http://www.daimajie.com)
 * Copyright 2019-2020 daimajie
 */


/*
 * Start layui2 - webtpl v1.0.0 (http://www.daimajie.com)
 * Copyright 2019-2020 daimajie
 */
layui.define('jquery', function(exports){
	var $ = layui.jquery;


  var obj = {
  	profile : null,
  	content : null,

  	init : function(content, profile){
  		this.profile = $(profile);
  		this.content = $(content);

  		//添加导航项目
  		//this.create();
  		this.event();

  	},
  	create: function(){
  		let that = this;
  		//添加锚点导航
		  that.content.children().each(function(index, element) {
			  var tagName=$(this).get(0).tagName;
			  if(tagName.substr(0,1).toUpperCase()=="H"){  
			      
			      var contentH=$(this).html();
			      var markid="mark-"+index.toString();
			      var aid="mark-"+index.toString()+index.toString();

			      $(this).attr("id",markid);
			      that.profile.find(".profile_right_cnt_last").before("<p><a id='"+ aid +"' href='#"+markid+"'>"+contentH+"</a></p>");
			  }  
			});


  	},
    event: function(){
    	let c = this.profile.find(".profile_right_cnt>p");

			//点击文章导航 添加激活class
			c.on("click",function(){
				$(this).addClass("floatnav_a_cur").siblings().removeClass("floatnav_a_cur");
			});

			//上下滚动 锚点跟随
			$(window).scroll(function(){
				var d=$(window).scrollTop();
				c.each(function(index, el){
					let href = $(el).find('a').attr('href');

					if($(href).offset().top - $(href).height() <= d){
						$(this).addClass("floatnav_a_cur").siblings().removeClass("floatnav_a_cur");
					}
				});
			})
		}


  };
  
  //输出test接口
  exports('profile', obj);
}); 
layui.define('jquery', function(exports){
	let $ = layui.jquery;


  let obj = {
  	init: function(selector){
  		let that = this;
  		$(selector).find('.tag-item').each(function(){
  			let color = that.getColor();
		    let size = that.getSize();

		    $(this).css({
		    	'color' : color,
		    	'fontSize' : size
		    });

  		});
  	},
  	/*getTagHtml: function (tagArr){
		    let html = '';
		    if(tagArr!=null && tagArr.length>0){
		        tagArr.sort(function(){ return 0.5 - Math.random() })
		        for(let i=0;i<tagArr.length;i++){
		            let color = getColor();
		            let size = getSize();
		            html += '<a href="${ctx!}/tag/search.action?tag='+tagArr[i]+'" style="color:'+color+'; font-size:'+size+'px;">'+tagArr[i]+'</a>';
		        }
		    }
		    $("#tags").html(html);
		},*/
  	getSize: function (){
		    return Math.floor(Math.random() * (30 - 10) + 10);
		},
  	getColor: function () {
		    let colorElements = "0,1,2,3,4,5,6,7,8,9,a,b,c,d,e,f";
		    let colorArray = colorElements.split(",");
		    let color = "#";
		    for (let i = 0; i < 6; i++) {
		        color += colorArray[Math.floor(Math.random() * 16)];
		    }
		    return color;
		}
  };
  
  //输出test接口
  exports('tags', obj);
}); 
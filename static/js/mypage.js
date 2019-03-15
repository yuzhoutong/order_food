$(".paginate_button").click(function(){
	$(".paginate_button").removeClass("active")
	if ($(this).hasClass('prev')  ){
		$(this).next().addClass('active')
		return
	}
	if ($(this).hasClass("next")){
		$(this).prev().addClass('active')
		console.log($(this))
		return
	}
	$(this).addClass("active")
})

function DevPage(page,pages){
	var prev = ""
	var next = ""
	if (page >= pages){
		return
	}
	if (page < 6) {
		return
	}
	var start = page-5
	var end = page + 1

	$("#pagination").empty()
	if (page == 1){
		prev = ` <li class="paginate_button previous disabled" id="example2_previous">
                    <a href="#" aria-controls="example2" data-dt-idx="0" tabindex="0">Previous</a>
                </li>`
	}else {
		prev = ` <li class="paginate_button previous " id="example2_previous">
                	<a href="#" aria-controls="example2" data-dt-idx="0" tabindex="0">Previous</a>
            	</li>`
	}
    var next = `<li class="paginate_button next" id="example2_next"><a href="#" aria-controls="example2" data-dt-idx="7" tabindex="0">Next</a>
                </li>`
    $("#pagination").append(prev)
	for (var i = start;i<=end;i++){
		var text = "<li class='paginate_button active'><a href='#' aria-controls='example2' data-dt-idx='1' tabindex='0'>"+i+"</a></li>"
		$("#pagination").append(text)
	}
	$("#pagination").append(text)
}



$(document).on('click', '.paginate_button', function(event) {
	console.log($(this).hasClass('next'))
	var page = 0
	if ($(this).hasClass('prev') || $(this).hasClass("next")){
		page = $(this).attr("value")
		console.log(page) 	
	}else {
		page = $(this).find("a").text()
	}
	var name = ""
	$(".th").remove()
	$.ajax({
		url: 'home/userlist',
		type: 'get',
		dataType: 'json',
		data: {"page": page},
		success:function(result){
			if (result.Status){
				for (var i = 0;i<result.Object.length;i++){
					var text = "<tr class='th'>"+
                		"<td>"+result.Object[i].OrderUsersId+"</td>"+
                		"<td>"+result.Object[i].Name+"</td>"+
                		"<td>"+result.Object[i].CreateTime+"</td>"
           			 if(result.Object[i].Accountstatus==1){
           			 	text += "<td><span class='label label-success'>启用</span></td>"
           			} else if (result.Object[i].Accountstatus==2){
           				text += "<td><span class='label label-danger'>禁用</span></td>"
           			}     
            		text+="<td>"+result.Object[i].Phone+"</td>"
            		text+="<td><a class='fa  fa-trash delete' uid = "+result.Object[i].OrderUsersId+" style='cursor:pointer'></a>"
                    if(result.Object[i].Accountstatus==1){
                        text+="&nbsp; &nbsp; &nbsp; &nbsp; &nbsp; &nbsp;<a title='禁用' style='color: red;cursor:pointer' class='fa fa-minus-square state' i='+result.Object[i].Accountstatus+' uid = {{$v.OrderUsersId}}></a></td>"
                    } else{
                    	text+="&nbsp; &nbsp; &nbsp;  &nbsp; &nbsp; &nbsp;<a title='启用' style='color: #00a73e;cursor:pointer' class='fa fa-check-square state' i='+result.Object[i].Accountstatus+' uid = {{$v.OrderUsersId}}></a></td>"
                    }
            		text +="</tr>"
	            	console.log(text)
	            	$("#userbody").append(text)
				}
				DevPage(result.Page,result.Pages)
			}
		},
	})	
})


//底部扩展键
$(function() {
    $('#doc-dropdown-js').dropdown({justify: '#doc-dropdown-justify-js'});
});

$(function(){
	$(".office_text").panel({iWheelStep:32});
});

//tab for three icon	
$(document).ready(function(){
  	$(".sidestrip_icon a").click(function(){
      $(".sidestrip_icon a").eq($(this).index()).addClass("cur").siblings().removeClass('cur');
	  $(".middle").hide().eq($(this).index()).show();
    });
});

//input box focus
$(document).ready(function(){
  	$("#input_box").focus(function(){
       $('.windows_input').css('background','#fff');
       $('#input_box').css('background','#fff');
   });
    $("#input_box").blur(function(){
       $('.windows_input').css('background','');
       $('#input_box').css('background','');
    });
});

// window.onload=function b(){
// 	var text = document.getElementById('input_box');
// 	var chat = document.getElementById('chatbox');
// 	var btn = document.getElementById('send');
// 	var talk = document.getElementById('talkbox');
//     debugger
// 	btn.onclick=function(){
// 		if(text.value ==''){
//             alert('不能发送空消息');
//         }else{
// 			chat.innerHTML += '<li class="me"><img src="'+'images/own_head.jpg'+'"><span>'+text.value+'</span></li>';
// 			text.value = '';
// 			chat.scrollTop=chat.scrollHeight;
// 			talk.style.background="#fff";
// 			text.style.background="#fff";
//
//             var content = document.getElementById('input_box');
//             var user_id = document.getElementById('user_id');
//             var token = document.getElementById('token');
//             var url = document.getElementById('write');
//             console.log("url.value:", url.value)
// 			$.ajax({
//                 type:"post",
//                 data:{keyid:user_id.value, content:content.value, token:token.value},
//             	url:url,
//                 dataType:'json',
//                 success:function (data) {
//                     if (data.status == 200) {
//
//                     }else {
//                         layer.msg(data.msg)
//                     }
//                 },
//                 error:function (data) {
//                     console.log(data)
//                 }
// 			})
//         };
// 	};
// };

//三图标
window.onload = function () {
    function a() {
        var si1 = document.getElementById('si_1');
        var si2 = document.getElementById('si_2');
        var si3 = document.getElementById('si_3');
        si1.onclick = function () {
            si1.style.background = "url(/public/chat/images/icon/head_2_1.png) no-repeat"
            si2.style.background = "";
            si3.style.background = "";
        };
        si2.onclick = function () {
            si2.style.background = "url(/public/chat/images/icon/head_3_1.png) no-repeat"
            si1.style.background = "";
            si3.style.background = "";
        };
        si3.onclick = function () {
            si3.style.background = "url(/public/chat/images/icon/head_4_1.png) no-repeat"
            si1.style.background = "";
            si2.style.background = "";
        };
    };

    a();
};

$(function () {
    var ws_ip = document.getElementById('ws_ip');
    var text = document.getElementById('input_box');
    var chat = document.getElementById('chatbox');
    var btn = document.getElementById('send');
    var talk = document.getElementById('talkbox');


    var ws = new WebSocket(ws_ip.value);
    ws.onopen = function(){
        console.log("已连接上websocket")
    };
    ws.onmessage = function(e) {
        var item = JSON.parse(e.data);

        console.log("item:", item)

        if (item.uid != user_id.value) {
            var html = ' <li class="other"><img src="'+ item.url +'" title="'+ item.username+'"><span>'+ item.content +'</span></li>'
            $("#chatbox").append(html)
        }
        if (item.add.length > 0) {
            $.each(item.add, function (i, user) {
                if (user.id != user_id.value) {
                    html = '<li class= "' + user.id;
                    if (user.id == user_id.value){
                        html += ' user_active"';
                    }
                    html += '>\n' +
                        '                            <div class="user_head "><img src="' + user.Url + '" /></div>\n' +
                        '                            <div class="user_text">\n' +
                        '                                <p class="user_name">' + user.username + '</p>\n' +
                        '                                <p class="user_message">...</p>\n' +
                        '                            </div>\n' +
                        '                            <div class="user_time">下午 2：54</div>\n' +
                        '                        </li>'
                }
            });
            $(".user_list").append(html);
        }
        if (item.del.length > 0) {
            $.each(item.add, function (i, user) {
                $(".user_list li ." + user.id).remove();
            });
        }
    };
    initData();


    btn.onclick = function () {
        if (text.value == '') {
            layer.msg('不能发送空消息');
        } else {
            var img =  $('#head_img').attr("src");
            chat.innerHTML += '<li class="me"><img src="' + img + '"><span>' + text.value + '</span></li>';
            chat.scrollTop = chat.scrollHeight;
            talk.style.background = "#fff";
            text.style.background = "#fff";
            var content = text.value;
            text.value = '';

            var user_id = document.getElementById('user_id');
            var token = document.getElementById('token');
            var url = document.getElementById('write');
            $.ajax({
                type:"post",
                data:{keyid:user_id.value, content:content, token:token.value},
                url:url.value,
                dataType:'json',
                error:function (data) {
                    console.log(data)
                }
            })
        };
    };

});
function initData() {
    var content = $(".chat_content").val();
    var url = "/user/get_data";
    var user_id = document.getElementById('user_id');
    var token = document.getElementById('token');
    var html = "";
    $.ajax({
        type:"post",
        data:{keyid:user_id.value, content:content, token:token.value},
        url:url,
        dataType:'json',
        success:function (data) {
        if (data.status == 200) {
            var user = data.user;
            var _data = data.data;

            var _name = user.username + '<img src="'+ user.img +'" alt="" />';
            $('.own_name').text(user.username);

        // $('.own_head').css("background-image","url(./head/" + user._img+")")

            $('#head_img').attr("src", user.img)

            //聊天记录
            $.each(_data,function(i,item){
                html += ' <li class="';

                if (item.uid == user_id.value) {
                    html += " me"
                }else {
                    html += " other"
                }
                html += '"><img src="'+ item.url +'" title="'+ item.username+'"><span>'+ item.content +'</span></li>';
            });
            $("#chatbox").append(html);

            //用户列表
            var _users = data.users;
            $.each(_users, function (i, user) {
                html = '<li ';
                if (user.id == user_id.value){
                    html += 'class="user_active"';
                }
                html += '>\n' +
                '                            <div class="user_head"><img src="' + user.Url + '" /></div>\n' +
                '                            <div class="user_text">\n' +
                '                                <p class="user_name">' + user.username + '</p>\n' +
                '                                <p class="user_message">...</p>\n' +
                '                            </div>\n' +
                '                            <div class="user_time">下午 2：54</div>\n' +
                '                        </li>'
            });
            $(".user_list").append(html);
        }else {
            layer.msg(data.msg)
        }
    },
    error:function (data) {
        console.log(data)
    }
})
}


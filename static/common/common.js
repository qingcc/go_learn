$(function() {
//region 列表页删除选中的ID
    $('.btndel').each(function () {
        var href = $(this).attr('href');
        $(this).attr('href','javascript:void(0)');
        $(this).attr('url',href);
    });

    $('.btndel').click(function () {
        if ($(".checkall input:checked").length < 1) {
            layer.msg('对不起，请选中您要操作的记录！');
            return false;
        }
        var msg = "删除记录后不可恢复，您确定吗？";
        if (arguments.length == 2) {
            msg = objmsg;
        }
        var url = $(this).attr('url');
        var obj = this;
        layer.confirm(msg, function () {
            var data = $("form").serialize();
            $.ajax({
                type: 'POST',
                url: url,
                data: data,
                dataType: 'json',
                success: function(data){
                    if (data.status == 200) {
                        parent.layer.msg(data.msg, { icon: 1, time: 1000 }, function () {
                            location.reload(); //刷新父页面
                        });
                    }else {
                        layer.msg(data.msg,{icon:2,time:1000, skin:'layer-ext-moon'});
                    }
                },
                error:function(data) {
                    console.log(data.msg);
                },
            });
        });
    });
//endregion

    $(".ajax").click(function () {
        var url = $(this).attr("href");
        $.get(url, [],function(data){
            if (data.status == 200) {
                layer.msg(data.msg);
                setTimeout(function () {
                    window.location.href=data.url
                },2000);
            }else {
                layer.msg(data.msg);
            }
            return
        });
        return false
    });

});
function ChangeSort(obj, url, id) {
    var sort = $(obj).val();
    $.post(url, {id:id, sort:sort},function(data){
        console.log(data)
        if (data.status == 200) {
            layer.msg(data.msg);
            setTimeout(function () {
                window.location.href=data.url
            },2000);
        }else {
            layer.msg(data.msg);
        }
        return
    });
    return false
}
//region  删除-单条记录
function single_del(obj,id,url){
    layer.confirm('确认要删除吗？', function(){
        $.ajax({
            type: 'POST',
            url: url,
            data:{id:id},
            dataType: 'json',
            success: function(data){
                if (data.status == 200){
                    $(obj).parents("tr").remove();
                    layer.msg(data.msg,{icon:1,time:1000, skin:'layer-ext-moon'});
                }else{
                    layer.msg(data.msg,{icon:7,time:1000, skin:'layer-ext-moon'});
                }
            },
            error:function(data) {
                console.log(data.msg);
            },
        });
    });
}
//endregion

//region Remark: 需要关闭 悬浮窗 并 刷新 父页面 Author:Qing
function ajaxSubmit(){
    var url = $("form").attr("action");
    var data = $("form").serialize();

    $.ajax({
        type: 'POST',
        url: url,
        data: data,
        dataType: 'json',
        success: function(data){
            if (data.status == 200) {
                parent.layer.msg(data.msg, { icon: 1, time: 1000 }, function () {
                    parent.location.reload(); //刷新父页面
                    layer_close()
                });
            }else {
                layer.msg(data.msg,{icon:2,time:1000, skin:'layer-ext-moon'});
            }
        },
        error:function(data) {
            console.log(data.msg);
        },
    });
}
//endregion

//region Remark: 需要关闭 悬浮窗 并 刷新 父页面 Author:Qing
function submit(){
    debugger
    var url = $("form").attr("action");
    var data = $("form").serialize();

    $.ajax({
        type: 'POST',
        url: url,
        data: data,
        dataType: 'json',
        success: function(data){
            if (data.status == 200) {
                layer.msg(data.msg, { icon: 1, time: 1000 }, function () {
                    window.location.href = data.url; //刷新父页面
                    // layer_close()
                });
            }else {
                layer.msg(data.msg,{icon:2,time:1000, skin:'layer-ext-moon'});
            }
        },
        error:function(data) {
            console.log(data.msg);
        },
    });
}
//endregion

//region 添加 或 编辑 需要post 提交 且 需要 关闭 悬浮窗 刷新 父页面
function addOrEdit(url,data) {
    $.ajax({
        type: 'POST',
        url: url,
        data: data,
        dataType: 'json',
        success: function(data){
            if (data.status == 200) {
                parent.layer.msg(data.msg, { icon: 1, time: 1000 }, function () {
                    parent.location.reload(); //刷新父页面
                    layer_close()
                });
            }else {
                layer.msg(data.msg,{icon:2,time:1000, skin:'layer-ext-moon'});
            }
        },
        error:function(data) {
            console.log(data.msg);
        },
    });
}
//endregion

/*停用*/
function toEnable(obj,id,url){
        $.ajax({
            type: 'POST',
            url: url,
            data: {id:id},
            dataType: 'json',
            success: function(data){
                if (data.status == 200) {
                        $(obj).parents("tr").find(".td-manage").prepend('<a onClick="toAble(this,'+data.id+',\''+data.url+'\')" href="javascript:;" title="启用" style="text-decoration:none"><i class="Hui-iconfont">&#xe615;</i></a>');
                        $(obj).parents("tr").find(".td-status").html('<span class="label label-default radius">已禁用</span>');
                        $(obj).remove();
                        layer.msg('已停用!',{icon: 5,time:1000});
                }else {
                    layer.msg(data.msg,{icon:2,time:1000, skin:'layer-ext-moon'});
                }
            },
            error:function(data) {
                console.log(data.msg);
            },
        });
}

/*启用*/
function toAble(obj,id,url){
    $.ajax({
        type: 'POST',
        url: url,
        data: {id:id},
        dataType: 'json',
        success: function(data){
            if (data.status == 200) {
                $(obj).parents("tr").find(".td-manage").prepend('<a onClick="toEnable(this,'+data.id+', \''+data.url+'\')" href="javascript:;" title="停用" style="text-decoration:none"><i class="Hui-iconfont">&#xe631;</i></a>');
                $(obj).parents("tr").find(".td-status").html('<span class="label label-success radius">已启用</span>');
                $(obj).remove();
                layer.msg('已启用!', {icon: 6,time:1000});
            }else {
                layer.msg(data.msg,{icon:2,time:1000, skin:'layer-ext-moon'});
            }
        },
        error:function(data) {
            console.log(data.msg);
        },
    });
}

/*
 @file 推广位跟踪
 @author 曹宇 <caoyu#at#csdn.net>
 @version 20130902
 */
function getCookie(sName){
    var allCookie = document.cookie.split("; ");
    for (var i=0; i < allCookie.length; i++){
        var aCrumb = allCookie[i].split("=");
        if (sName == aCrumb[0])
            return aCrumb[1];
    }
    return null;
}


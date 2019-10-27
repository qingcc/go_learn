$(function () {
    var uploader = WebUploader.create({

        // 选完文件后，是否自动上传。
        auto: true,

        // swf文件路径
        swf: "/public/common/webuploader/uploader.swf",

        // 文件接收服务端。
        server: "/admin/upload/image",
        fileVal: 'FileData', //上传域的名称

        // 选择文件的按钮。可选。
        // 内部根据当前运行是创建，可能是input元素，也可能是flash.
        pick: {
            id: '#picker',
            innerHTML: '选择图片',
            multiple: true
        },

        //multiple: true,

        // 只允许选择图片文件。
        accept: {
            title: 'Images',
            extensions: 'gif,jpg,jpeg,bmp,png',
            mimeTypes: 'image/*'
        },
        thumb:{
            width:110,
            height:110,

            //图片质量,只有type为'image/jpeg'的时候才有效
            quality:70,

            //是否允许放大,如果想要生成小图的时候不失真设置为false
            allowMagnify:true,

            //是否允许裁剪
            crop:true,

            //为空的话则保留原有图片格式
            //否则强制转换成指定的类型
            type: 'image/jpeg'
        }
    });

    // 当有文件被添加进队列的时候
    /*    uploader.on( 'fileQueued', function( file ) {
            var $list = $("#thelist");
            $list.append( '<div id="' + file.id + '" class="item">' +
                '<h4 class="info">' + file.name + '</h4>' +
                '<p class="state">等待上传...</p>' +
                '</div>' );
        });*/

    uploader.on( 'uploadSuccess', function( file, response ) {
        console.log(response);
        // var html = " <li><img src='' id='imgs_" + file.id + "' alt='' />" +
        //     "<div class='mws-gallery-overlay'>" +
        //     "<div class='mws-gallery-zoom'></div>" +
        //     "<input type='hidden' name='avatar' value='" + response._raw + "'></div>" +
        //     "<a href='javascript:;' onclick='delImg(this);'>删除</a></li>";
        // $('#upload_imgs').find('.show_imgs').prepend(html);
        $('.upload-path').attr("value", response.path);
        // $('#imgs_'+file.id).attr('src', '/storage/' + response._raw);
    });

    // 文件上传失败，显示上传出错。
    uploader.on( 'uploadError', function( file ) {
        layer.msg('上传失败');
    });

    uploader.on( 'uploadComplete', function( file ) {
        $( '#'+file.id ).find('.progress').fadeOut();
    });

})
document.querySelector('.chat[data-chat=person2]').classList.add('active-chat');
document.querySelector('.person[data-chat=person2]').classList.add('active');

var friends = {
  list: document.querySelector('ul.people'),
  all: document.querySelectorAll('.left .person'),
  name: '' },

chat = {
  container: document.querySelector('.container .right'),
  current: null,
  person: null,
  name: document.querySelector('.container .right .top .name') };


friends.all.forEach(function (f) {
  f.addEventListener('mousedown', function () {
    f.classList.contains('active') || setAciveChat(f);
  });
});

function setAciveChat(f) {
  friends.list.querySelector('.active').classList.remove('active');
  f.classList.add('active');
  chat.current = chat.container.querySelector('.active-chat');
  chat.person = f.getAttribute('data-chat');
  chat.current.classList.remove('active-chat');
  chat.container.querySelector('[data-chat="' + chat.person + '"]').classList.add('active-chat');
  friends.name = f.querySelector('.name').innerText;
  chat.name.innerHTML = friends.name;
}

function post_and_open(url,args)
{
    $.post(url,args,function(response){
        var obj = window.open("about:blank","_blank");
        obj.document.write(response);
    })
}

function StandardPost (url,args)
{
    var form = $("<form method='post'></form>");
    form.attr({"action":url});
    for (arg in args)
    {
        var input = $("<input type='hidden'>");
        input.attr({"name":arg});
        input.val(args[arg]);
        form.append(input);
    }
    form.submit();
}
#如何安全接入这个聊天主页面
通过登录api `user/login` 获得id和token
返回json到前端,
前端拼接url
/chat/index.shtml?id=1&token=123
通过location.href 跳转

#如何添加/显示好友和群



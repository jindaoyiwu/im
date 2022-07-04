Git 全局设置:

git config --global user.name "jindaoyiwu"
git config --global user.email "jindaoyiwu@163.com"
创建 git 仓库:

mkdir go-im
cd go-im
git init 
touch README.md
git add README.md
git commit -m "first commit"
git remote add origin git@gitee.com:jindaoyiwu/go-im.git
git push -u origin "master"
已有仓库?

cd existing_git_repo
git remote add origin git@gitee.com:jindaoyiwu/go-im.git
git push -u origin "master"

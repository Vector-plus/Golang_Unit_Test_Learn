# Git基本使用方法

![git原理图](/Golang_Unit_Test_Learn/static/git结构图.webp)
[学习链接](https://www.jianshu.com/p/4821f3c802a6)

主要包括git的上传下载代码、代码的合并、分支的操作

## Git对代码的基本操作
**一般操作流程：**
git clone 

git status

git pull

git status

git add .

git status

git commit |-a| -m "mesg"

git status

git push

**分支操作：**

git branch -a

git branch -vv   #查看远程关联

git branch "branch_name" #新建本地分支

git checkout "branch_name" #切换分支

git branch --set-upstream-to=origin/branch_name branch_name #关联远程分支和本地分支

git checkout "master"

git merge "branch_name" #合并分支

git status
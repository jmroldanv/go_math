# v2 - Major version - going to create a new branch to a new version major (not backwards compatible) v2 
$ cd D:\workspace\code\golang\go_math
$ mkdir v2

# linux - copy the packages to new v2 directory and creates go.mod
$ cp -r calc/ geometry/ v2/
$ cd v2
$ go mod init github.com/jmroldanv/go_math/v2
$ type go.mod

# change the source code with the new updates and creates new tag in github
$ cd v2
$ git add -A
$ git commit -m "transform Add into a variadic function"
$ git tag -a v2.0.0 -m "transform Add into a variadic function"
$ git push origin master --tags

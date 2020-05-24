- my modules: https://github.com/jmroldanv/go_math

# init modules (enable using modules) and creates go.mod
- $ cd D:\workspace\code\golang\go_math
- $ go mod init github.com/jmroldanv/go_math

# publish to github
- $ cd D:\workspace\code\golang\go_math
- $ git init
- $ dir /a       - to see hidden files/directories as .git directory 
- $ git remote add origin https://github.com/jmroldanv/go_math.git
- $ git remote -v
- $ git add -A
- $ git commit -m "first commit"
- $ git push -u -f origin master

# tagging the version ( vMajor.Minor.Patch )
- $ git tag -a v1.0.0 -m "initial release"
- $ git push origin master --tags
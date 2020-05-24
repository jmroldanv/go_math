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

# notes - to delete a tag 
- $ git tag -l
- $ git tag -d v1.0.0                   // delete locally
- $ git push --delete origin v1.0.0     // delelte remotely

# updating the dependencies to a new version
# we could do it manually or automatically

# manually
- $ del go.sum
- edit go.mod exchange the version from v1.0.0 to v1.0.1
- $ go run main.go

# automatically (update .mod and .sum files)
$ go get -u
$ go run main.go
#!/bin/bash
set -x
#
# リストのメンバー取得関数は follow_follower に追加されていた。(Github)
# ~/go/pkg/mod/github.com/!chimera!coder/anaconda@v2.0.0+incompatible
# を使わずに
# ~/go/src/github.com/ChimeraCoder/anaconda  をリンクしたいために、
# GO111MODULE=off が必要だった。


GO111MODULE=off go build twlist.go

#go env -w GO111MODULE=auto

## see env
#go env GOENV
#cat "$(go env GOENV)

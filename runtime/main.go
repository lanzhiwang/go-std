package main

import (
	"fmt"
	"os"
	"runtime"
)

func main() {
	fmt.Println(runtime.Compiler)
	fmt.Println(runtime.GOARCH)
	fmt.Println(runtime.GOOS)
	fmt.Println(os.Environ())
}

// $ ./main
// gc
// amd64
// linux
// [HOSTNAME=ff3dbb77cc10 PWD=/go-std HOME=/root GOLANG_VERSION=1.22.4 TERM=xterm SHLVL=1 GOTOOLCHAIN=local PATH=/go/bin:/usr/local/go/bin:/usr/local/sbin:/usr/local/bin:/usr/sbin:/usr/bin:/sbin:/bin GOPATH=/go _=./main]
// $

// $ ./main
// gc
// amd64
// darwin
// [TERM_SESSION_ID=w0t0p0:C5D0C579-C35A-4165-BD4A-37866E81906B SSH_AUTH_SOCK=/private/tmp/com.apple.launchd.OkTKKV11d5/Listeners LC_TERMINAL_VERSION=3.3.7 COLORFGBG=7;0 ITERM_PROFILE=huzhi-mac XPC_FLAGS=0x0 LANG=zh_CN.UTF-8 PWD=/Users/huzhi/work/code/go_code/go-std SHELL=/bin/zsh __CFBundleIdentifier=com.googlecode.iterm2 SECURITYSESSIONID=186a6 TERM_PROGRAM_VERSION=3.3.7 TERM_PROGRAM=iTerm.app PATH=/usr/local/opt/scala/bin:/usr/local/opt/git/bin:/Users/huzhi/go/bin:/usr/local/opt/go@1.20/bin:/usr/local/opt/openjdk@8/bin:/usr/local/bin:/usr/bin:/bin:/usr/sbin:/sbin LC_TERMINAL=iTerm2 COLORTERM=truecolor COMMAND_MODE=unix2003 TERM=xterm-256color HOME=/Users/huzhi TMPDIR=/var/folders/4f/gy92m6hd2xj3c6_yz0dht6880000gn/T/ USER=huzhi XPC_SERVICE_NAME=0 LOGNAME=huzhi LaunchInstanceID=F00F4799-21B6-4F63-A45F-C01994AF8D3A __CF_USER_TEXT_ENCODING=0x0:0x19:0x34 ITERM_SESSION_ID=w0t0p0:C5D0C579-C35A-4165-BD4A-37866E81906B SHLVL=1 OLDPWD=/Users/huzhi ZSH=/Users/huzhi/.oh-my-zsh PAGER=less LESS=-R LSCOLORS=Gxfxcxdxbxegedabagacad LS_COLORS=di=1;36:ln=35:so=32:pi=33:ex=31:bd=34;46:cd=34;43:su=30;41:sg=30;46:tw=30;42:ow=30;43 _=/Users/huzhi/work/code/go_code/go-std/./main]
// $

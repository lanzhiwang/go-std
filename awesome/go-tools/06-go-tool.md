```bash
$ go help tool

usage: go tool [-n] command [args...]

Tool runs the go tool command identified by the arguments.
With no arguments it prints the list of known tools.
工具运行由参数标识的 go tool 命令。如果没有参数，它会打印已知工具的列表。

The -n flag causes tool to print the command that would be
executed but not execute it.
-n 标志使工具打印将要执行的命令，但不执行它。

For more about each tool command, see 'go doc cmd/<command>'.

$ go tool
asm
buildid
cgo
compile
covdata
cover
dist #
distpack #
doc
fix
link
nm
objdump
pack
pprof
test2json
trace
vet

# 预处理
# 编译
# 汇编
# 链接

/usr/local/go/pkg/tool/linux_amd64/compile # 编译
/usr/local/go/pkg/tool/linux_amd64/asm # 汇编
/usr/local/go/pkg/tool/linux_amd64/link # 链接
/usr/local/go/pkg/tool/linux_amd64/buildid # Buildid displays or updates the build ID stored in a Go package or binary.
/usr/local/go/pkg/tool/linux_amd64/pack # Pack 是传统 Unix ar 工具的简化版本。它仅实现了 Go 所需的操作。

/usr/local/go/pkg/tool/linux_amd64/cgo

$ ls -al /usr/local/go/pkg/tool/linux_amd64/
total 86468
drwxr-xr-x 2 root root     4096 May 30 19:26 .
drwxr-xr-x 3 root root     4096 May 30 19:26 ..
-rwxr-xr-x 1 root root  2438886 May 30 19:26 addr2line
-rwxr-xr-x 1 root root  3831435 May 30 19:26 asm
-rwxr-xr-x 1 root root  1912930 May 30 19:26 buildid
-rwxr-xr-x 1 root root  3740784 May 30 19:26 cgo
-rwxr-xr-x 1 root root 19344631 May 30 19:26 compile
-rwxr-xr-x 1 root root  2469664 May 30 19:26 covdata
-rwxr-xr-x 1 root root  4262861 May 30 19:26 cover
-rwxr-xr-x 1 root root  3092557 May 30 19:26 doc
-rwxr-xr-x 1 root root  2587022 May 30 19:26 fix
-rwxr-xr-x 1 root root  5294592 May 30 19:26 link
-rwxr-xr-x 1 root root  2383139 May 30 19:26 nm
-rwxr-xr-x 1 root root  3474997 May 30 19:26 objdump
-rwxr-xr-x 1 root root  1648605 May 30 19:26 pack
-rwxr-xr-x 1 root root 11756474 May 30 19:26 pprof
-rwxr-xr-x 1 root root  2034509 May 30 19:26 test2json
-rwxr-xr-x 1 root root 12107504 May 30 19:26 trace
-rwxr-xr-x 1 root root  6119508 May 30 19:26 vet
$

$  GO111MODULE='on' GOPROXY='https://goproxy.cn,direct' go build -o main -v -work -x -a -p 1 context/v1-example-waitgroup/main.go

WORK=/tmp/go-build2485170101

##############################################################################

internal/goarch
mkdir -p $WORK/b006/
echo '# import config' > $WORK/b006/importcfg # internal
cd /go-std
/usr/local/go/pkg/tool/linux_amd64/compile \
-o $WORK/b006/_pkg_.a \
-trimpath "$WORK/b006=>" \
-p internal/goarch \
-std \
-complete \
-buildid W8XbrYSTrZnxJ0wap0fC/W8XbrYSTrZnxJ0wap0fC \
-goversion go1.22.4 \
-c=16 \
-nolocalimports \
-importcfg $WORK/b006/importcfg \
-pack /usr/local/go/src/internal/goarch/goarch.go \
      /usr/local/go/src/internal/goarch/goarch_amd64.go \
      /usr/local/go/src/internal/goarch/zgoarch_amd64.go

/usr/local/go/pkg/tool/linux_amd64/buildid -w $WORK/b006/_pkg_.a # internal

cp $WORK/b006/_pkg_.a /root/.cache/go-build/95/95d75a340f1de5a36d487bc25d45647404c4ce3fabfeb3e56dcada4a188f2f88-d # internal

##############################################################################

internal/abi
mkdir -p $WORK/b005/
echo -n > $WORK/b005/go_asm.h # internal
cd /usr/local/go/src/internal/abi

/usr/local/go/pkg/tool/linux_amd64/asm \
-p internal/abi \
-trimpath "$WORK/b005=>" \
-I $WORK/b005/ \
-I /usr/local/go/pkg/include \
-D GOOS_linux \
-D GOARCH_amd64 \
-D GOAMD64_v1 \
-gensymabis \
-o $WORK/b005/symabis ./abi_test.s ./stub.s

cat >/tmp/go-build2485170101/b005/importcfg << 'EOF' # internal
# import config
packagefile internal/goarch=/tmp/go-build2485170101/b006/_pkg_.a
EOF

cd /go-std

/usr/local/go/pkg/tool/linux_amd64/compile \
-o $WORK/b005/_pkg_.a \
-trimpath "$WORK/b005=>" \
-p internal/abi \
-std -buildid sxLkIWkON61qOkHd3yTl/sxLkIWkON61qOkHd3yTl \
-goversion go1.22.4 \
-symabis $WORK/b005/symabis \
-c=16 \
-nolocalimports \
-importcfg $WORK/b005/importcfg \
-pack -asmhdr \
$WORK/b005/go_asm.h \
/usr/local/go/src/internal/abi/abi.go \
/usr/local/go/src/internal/abi/abi_amd64.go \
/usr/local/go/src/internal/abi/compiletype.go \
/usr/local/go/src/internal/abi/funcpc.go \
/usr/local/go/src/internal/abi/map.go \
/usr/local/go/src/internal/abi/stack.go \
/usr/local/go/src/internal/abi/switch.go \
/usr/local/go/src/internal/abi/symtab.go \
/usr/local/go/src/internal/abi/type.go

cd /usr/local/go/src/internal/abi

/usr/local/go/pkg/tool/linux_amd64/asm -p internal/abi -trimpath "$WORK/b005=>" -I $WORK/b005/ -I /usr/local/go/pkg/include -D GOOS_linux -D GOARCH_amd64 -D GOAMD64_v1 -o $WORK/b005/abi_test.o ./abi_test.s

/usr/local/go/pkg/tool/linux_amd64/asm -p internal/abi -trimpath "$WORK/b005=>" -I $WORK/b005/ -I /usr/local/go/pkg/include -D GOOS_linux -D GOARCH_amd64 -D GOAMD64_v1 -o $WORK/b005/stub.o ./stub.s

/usr/local/go/pkg/tool/linux_amd64/pack r $WORK/b005/_pkg_.a $WORK/b005/abi_test.o $WORK/b005/stub.o # internal

/usr/local/go/pkg/tool/linux_amd64/buildid -w $WORK/b005/_pkg_.a # internal

cp $WORK/b005/_pkg_.a /root/.cache/go-build/f4/f42b2333ad4fe9f65cc13968a64c2f3f67fdf103be62ec3f3400f556e5b394dd-d # internal

##############################################################################

internal/unsafeheader
mkdir -p $WORK/b008/
echo '# import config' > $WORK/b008/importcfg # internal
cd /go-std

/usr/local/go/pkg/tool/linux_amd64/compile -o $WORK/b008/_pkg_.a -trimpath "$WORK/b008=>" -p internal/unsafeheader -std -complete -buildid cwzjX3XBYecsYqxkVo2z/cwzjX3XBYecsYqxkVo2z -goversion go1.22.4 -c=16 -nolocalimports -importcfg $WORK/b008/importcfg -pack /usr/local/go/src/internal/unsafeheader/unsafeheader.go

/usr/local/go/pkg/tool/linux_amd64/buildid -w $WORK/b008/_pkg_.a # internal
cp $WORK/b008/_pkg_.a /root/.cache/go-build/f3/f3d39875f4f2f30755d19575d741056bc2351202a6b0f3daf2419def84ebc9c6-d # internal

##############################################################################

internal/cpu
mkdir -p $WORK/b011/
echo -n > $WORK/b011/go_asm.h # internal
cd /usr/local/go/src/internal/cpu

/usr/local/go/pkg/tool/linux_amd64/asm -p internal/cpu -trimpath "$WORK/b011=>" -I $WORK/b011/ -I /usr/local/go/pkg/include -D GOOS_linux -D GOARCH_amd64 -D GOAMD64_v1 -gensymabis -o $WORK/b011/symabis ./cpu.s ./cpu_x86.s

echo '# import config' > $WORK/b011/importcfg # internal

cd /go-std

/usr/local/go/pkg/tool/linux_amd64/compile -o $WORK/b011/_pkg_.a -trimpath "$WORK/b011=>" -p internal/cpu -std -buildid xSp7WgDXlt4PpPkrfuIs/xSp7WgDXlt4PpPkrfuIs -goversion go1.22.4 -symabis $WORK/b011/symabis -c=16 -nolocalimports -importcfg $WORK/b011/importcfg -pack -asmhdr $WORK/b011/go_asm.h /usr/local/go/src/internal/cpu/cpu.go /usr/local/go/src/internal/cpu/cpu_x86.go

cd /usr/local/go/src/internal/cpu

/usr/local/go/pkg/tool/linux_amd64/asm -p internal/cpu -trimpath "$WORK/b011=>" -I $WORK/b011/ -I /usr/local/go/pkg/include -D GOOS_linux -D GOARCH_amd64 -D GOAMD64_v1 -o $WORK/b011/cpu.o ./cpu.s

/usr/local/go/pkg/tool/linux_amd64/asm -p internal/cpu -trimpath "$WORK/b011=>" -I $WORK/b011/ -I /usr/local/go/pkg/include -D GOOS_linux -D GOARCH_amd64 -D GOAMD64_v1 -o $WORK/b011/cpu_x86.o ./cpu_x86.s

/usr/local/go/pkg/tool/linux_amd64/pack r $WORK/b011/_pkg_.a $WORK/b011/cpu.o $WORK/b011/cpu_x86.o # internal

/usr/local/go/pkg/tool/linux_amd64/buildid -w $WORK/b011/_pkg_.a # internal
cp $WORK/b011/_pkg_.a /root/.cache/go-build/2a/2a1f2962348d30d72f91bf592d1e6da93fac231d6cc475e337bdeff16e391bf7-d # internal

##############################################################################

internal/bytealg
mkdir -p $WORK/b010/
echo -n > $WORK/b010/go_asm.h # internal
cd /usr/local/go/src/internal/bytealg

/usr/local/go/pkg/tool/linux_amd64/asm -p internal/bytealg -trimpath "$WORK/b010=>" -I $WORK/b010/ -I /usr/local/go/pkg/include -D GOOS_linux -D GOARCH_amd64 -D GOAMD64_v1 -gensymabis -o $WORK/b010/symabis ./compare_amd64.s ./count_amd64.s ./equal_amd64.s ./index_amd64.s ./indexbyte_amd64.s

cat >/tmp/go-build2485170101/b010/importcfg << 'EOF' # internal
# import config
packagefile internal/cpu=/tmp/go-build2485170101/b011/_pkg_.a
EOF

cd /go-std

/usr/local/go/pkg/tool/linux_amd64/compile -o $WORK/b010/_pkg_.a -trimpath "$WORK/b010=>" -p internal/bytealg -std -buildid OGz252pe-DxjNWnXS_0e/OGz252pe-DxjNWnXS_0e -goversion go1.22.4 -symabis $WORK/b010/symabis -c=16 -nolocalimports -importcfg $WORK/b010/importcfg -pack -asmhdr $WORK/b010/go_asm.h /usr/local/go/src/internal/bytealg/bytealg.go /usr/local/go/src/internal/bytealg/compare_native.go /usr/local/go/src/internal/bytealg/count_native.go /usr/local/go/src/internal/bytealg/equal_generic.go /usr/local/go/src/internal/bytealg/equal_native.go /usr/local/go/src/internal/bytealg/index_amd64.go /usr/local/go/src/internal/bytealg/index_native.go /usr/local/go/src/internal/bytealg/indexbyte_native.go /usr/local/go/src/internal/bytealg/lastindexbyte_generic.go

cd /usr/local/go/src/internal/bytealg

/usr/local/go/pkg/tool/linux_amd64/asm -p internal/bytealg -trimpath "$WORK/b010=>" -I $WORK/b010/ -I /usr/local/go/pkg/include -D GOOS_linux -D GOARCH_amd64 -D GOAMD64_v1 -o $WORK/b010/compare_amd64.o ./compare_amd64.s

/usr/local/go/pkg/tool/linux_amd64/asm -p internal/bytealg -trimpath "$WORK/b010=>" -I $WORK/b010/ -I /usr/local/go/pkg/include -D GOOS_linux -D GOARCH_amd64 -D GOAMD64_v1 -o $WORK/b010/count_amd64.o ./count_amd64.s

/usr/local/go/pkg/tool/linux_amd64/asm -p internal/bytealg -trimpath "$WORK/b010=>" -I $WORK/b010/ -I /usr/local/go/pkg/include -D GOOS_linux -D GOARCH_amd64 -D GOAMD64_v1 -o $WORK/b010/equal_amd64.o ./equal_amd64.s

/usr/local/go/pkg/tool/linux_amd64/asm -p internal/bytealg -trimpath "$WORK/b010=>" -I $WORK/b010/ -I /usr/local/go/pkg/include -D GOOS_linux -D GOARCH_amd64 -D GOAMD64_v1 -o $WORK/b010/index_amd64.o ./index_amd64.s

/usr/local/go/pkg/tool/linux_amd64/asm -p internal/bytealg -trimpath "$WORK/b010=>" -I $WORK/b010/ -I /usr/local/go/pkg/include -D GOOS_linux -D GOARCH_amd64 -D GOAMD64_v1 -o $WORK/b010/indexbyte_amd64.o ./indexbyte_amd64.s

/usr/local/go/pkg/tool/linux_amd64/pack r $WORK/b010/_pkg_.a $WORK/b010/compare_amd64.o $WORK/b010/count_amd64.o $WORK/b010/equal_amd64.o $WORK/b010/index_amd64.o $WORK/b010/indexbyte_amd64.o # internal

/usr/local/go/pkg/tool/linux_amd64/buildid -w $WORK/b010/_pkg_.a # internal

cp $WORK/b010/_pkg_.a /root/.cache/go-build/e2/e290284795322c6f2cc123c16f2d0dd0cc2d86543272cd59079a16ea60a3eafd-d # internal

##############################################################################

internal/chacha8rand
mkdir -p $WORK/b012/
echo -n > $WORK/b012/go_asm.h # internal
cd /usr/local/go/src/internal/chacha8rand

/usr/local/go/pkg/tool/linux_amd64/asm -p internal/chacha8rand -trimpath "$WORK/b012=>" -I $WORK/b012/ -I /usr/local/go/pkg/include -D GOOS_linux -D GOARCH_amd64 -D GOAMD64_v1 -gensymabis -o $WORK/b012/symabis ./chacha8_amd64.s

cat >/tmp/go-build2485170101/b012/importcfg << 'EOF' # internal
# import config
packagefile internal/goarch=/tmp/go-build2485170101/b006/_pkg_.a
EOF

cd /go-std
/usr/local/go/pkg/tool/linux_amd64/compile -o $WORK/b012/_pkg_.a -trimpath "$WORK/b012=>" -p internal/chacha8rand -std -buildid 7Ob2u1xNCcmLrDJt0Ooq/7Ob2u1xNCcmLrDJt0Ooq -goversion go1.22.4 -symabis $WORK/b012/symabis -c=16 -nolocalimports -importcfg $WORK/b012/importcfg -pack -asmhdr $WORK/b012/go_asm.h /usr/local/go/src/internal/chacha8rand/chacha8.go /usr/local/go/src/internal/chacha8rand/chacha8_generic.go

cd /usr/local/go/src/internal/chacha8rand

/usr/local/go/pkg/tool/linux_amd64/asm -p internal/chacha8rand -trimpath "$WORK/b012=>" -I $WORK/b012/ -I /usr/local/go/pkg/include -D GOOS_linux -D GOARCH_amd64 -D GOAMD64_v1 -o $WORK/b012/chacha8_amd64.o ./chacha8_amd64.s

/usr/local/go/pkg/tool/linux_amd64/pack r $WORK/b012/_pkg_.a $WORK/b012/chacha8_amd64.o # internal

/usr/local/go/pkg/tool/linux_amd64/buildid -w $WORK/b012/_pkg_.a # internal

cp $WORK/b012/_pkg_.a /root/.cache/go-build/10/1058d8a4a7de9ab1bf6215037a50692ca53209f80b7e5ff09139d0db5a097e98-d # internal

##############################################################################

internal/coverage/rtcov
mkdir -p $WORK/b013/
echo '# import config' > $WORK/b013/importcfg # internal
cd /go-std

/usr/local/go/pkg/tool/linux_amd64/compile -o $WORK/b013/_pkg_.a -trimpath "$WORK/b013=>" -p internal/coverage/rtcov -std -complete -buildid dtuSBwdYW-OGnA1MSjA9/dtuSBwdYW-OGnA1MSjA9 -goversion go1.22.4 -c=16 -nolocalimports -importcfg $WORK/b013/importcfg -pack /usr/local/go/src/internal/coverage/rtcov/rtcov.go

/usr/local/go/pkg/tool/linux_amd64/buildid -w $WORK/b013/_pkg_.a # internal

cp $WORK/b013/_pkg_.a /root/.cache/go-build/56/56e70fe7de0cdf8674057760e2f203ac4f8563f53ee0c345a4c24960b7a523dc-d # internal

##############################################################################

internal/godebugs
mkdir -p $WORK/b014/
echo '# import config' > $WORK/b014/importcfg # internal

/usr/local/go/pkg/tool/linux_amd64/compile -o $WORK/b014/_pkg_.a -trimpath "$WORK/b014=>" -p internal/godebugs -std -complete -buildid SaVkc0586JZwgoMO7juS/SaVkc0586JZwgoMO7juS -goversion go1.22.4 -c=16 -nolocalimports -importcfg $WORK/b014/importcfg -pack /usr/local/go/src/internal/godebugs/table.go

/usr/local/go/pkg/tool/linux_amd64/buildid -w $WORK/b014/_pkg_.a # internal

cp $WORK/b014/_pkg_.a /root/.cache/go-build/bf/bf7cda828af92519f6807d7d6a6828ca4ab1054a826c738713b37e81f98367ee-d # internal

##############################################################################

internal/goexperiment
mkdir -p $WORK/b015/
echo '# import config' > $WORK/b015/importcfg # internal

/usr/local/go/pkg/tool/linux_amd64/compile -o $WORK/b015/_pkg_.a -trimpath "$WORK/b015=>" -p internal/goexperiment -std -complete -buildid E8DjjOlAxJlu4t9GdB9f/E8DjjOlAxJlu4t9GdB9f -goversion go1.22.4 -c=16 -nolocalimports -importcfg $WORK/b015/importcfg -pack /usr/local/go/src/internal/goexperiment/exp_allocheaders_on.go /usr/local/go/src/internal/goexperiment/exp_arenas_off.go /usr/local/go/src/internal/goexperiment/exp_boringcrypto_off.go /usr/local/go/src/internal/goexperiment/exp_cacheprog_off.go /usr/local/go/src/internal/goexperiment/exp_cgocheck2_off.go /usr/local/go/src/internal/goexperiment/exp_coverageredesign_on.go /usr/local/go/src/internal/goexperiment/exp_exectracer2_on.go /usr/local/go/src/internal/goexperiment/exp_fieldtrack_off.go /usr/local/go/src/internal/goexperiment/exp_heapminimum512kib_off.go /usr/local/go/src/internal/goexperiment/exp_loopvar_off.go /usr/local/go/src/internal/goexperiment/exp_newinliner_off.go /usr/local/go/src/internal/goexperiment/exp_pagetrace_off.go /usr/local/go/src/internal/goexperiment/exp_preemptibleloops_off.go /usr/local/go/src/internal/goexperiment/exp_rangefunc_off.go /usr/local/go/src/internal/goexperiment/exp_regabiargs_on.go /usr/local/go/src/internal/goexperiment/exp_regabiwrappers_on.go /usr/local/go/src/internal/goexperiment/exp_staticlockranking_off.go /usr/local/go/src/internal/goexperiment/flags.go

/usr/local/go/pkg/tool/linux_amd64/buildid -w $WORK/b015/_pkg_.a # internal

cp $WORK/b015/_pkg_.a /root/.cache/go-build/9b/9be16b5a15c10e4bd43dcf2bbca52b58939862e4dbad148f0af7186ea3076f88-d # internal

##############################################################################

internal/goos
mkdir -p $WORK/b016/
echo '# import config' > $WORK/b016/importcfg # internal

/usr/local/go/pkg/tool/linux_amd64/compile -o $WORK/b016/_pkg_.a -trimpath "$WORK/b016=>" -p internal/goos -std -complete -buildid v3LtoeBSJC6VOSIpXGRT/v3LtoeBSJC6VOSIpXGRT -goversion go1.22.4 -c=16 -nolocalimports -importcfg $WORK/b016/importcfg -pack /usr/local/go/src/internal/goos/goos.go /usr/local/go/src/internal/goos/unix.go /usr/local/go/src/internal/goos/zgoos_linux.go

/usr/local/go/pkg/tool/linux_amd64/buildid -w $WORK/b016/_pkg_.a # internal
cp $WORK/b016/_pkg_.a /root/.cache/go-build/37/37f751ab5d7e7b19ddad9eda105e8e849a2bb518b169ef787a5a55acc33be8bf-d # internal

##############################################################################

runtime/internal/atomic
mkdir -p $WORK/b017/
echo -n > $WORK/b017/go_asm.h # internal
cd /usr/local/go/src/runtime/internal/atomic
/usr/local/go/pkg/tool/linux_amd64/asm -p runtime/internal/atomic -trimpath "$WORK/b017=>" -I $WORK/b017/ -I /usr/local/go/pkg/include -D GOOS_linux -D GOARCH_amd64 -D GOAMD64_v1 -gensymabis -o $WORK/b017/symabis ./atomic_amd64.s

echo '# import config' > $WORK/b017/importcfg # internal

cd /go-std

/usr/local/go/pkg/tool/linux_amd64/compile -o $WORK/b017/_pkg_.a -trimpath "$WORK/b017=>" -p runtime/internal/atomic -std -buildid 6xOg5AdpEFCKQQG96m4S/6xOg5AdpEFCKQQG96m4S -goversion go1.22.4 -symabis $WORK/b017/symabis -c=16 -nolocalimports -importcfg $WORK/b017/importcfg -pack -asmhdr $WORK/b017/go_asm.h /usr/local/go/src/runtime/internal/atomic/atomic_amd64.go /usr/local/go/src/runtime/internal/atomic/doc.go /usr/local/go/src/runtime/internal/atomic/stubs.go /usr/local/go/src/runtime/internal/atomic/types.go /usr/local/go/src/runtime/internal/atomic/types_64bit.go /usr/local/go/src/runtime/internal/atomic/unaligned.go

cd /usr/local/go/src/runtime/internal/atomic

/usr/local/go/pkg/tool/linux_amd64/asm -p runtime/internal/atomic -trimpath "$WORK/b017=>" -I $WORK/b017/ -I /usr/local/go/pkg/include -D GOOS_linux -D GOARCH_amd64 -D GOAMD64_v1 -o $WORK/b017/atomic_amd64.o ./atomic_amd64.s

/usr/local/go/pkg/tool/linux_amd64/pack r $WORK/b017/_pkg_.a $WORK/b017/atomic_amd64.o # internal

/usr/local/go/pkg/tool/linux_amd64/buildid -w $WORK/b017/_pkg_.a # internal

cp $WORK/b017/_pkg_.a /root/.cache/go-build/fa/fa0ebb893b924a075ac695fb2141b4a752f82073033ef868e71af3aa29401d84-d # internal

##############################################################################

runtime/internal/math
mkdir -p $WORK/b018/

cat >/tmp/go-build2485170101/b018/importcfg << 'EOF' # internal
# import config
packagefile internal/goarch=/tmp/go-build2485170101/b006/_pkg_.a
EOF

cd /go-std

/usr/local/go/pkg/tool/linux_amd64/compile -o $WORK/b018/_pkg_.a -trimpath "$WORK/b018=>" -p runtime/internal/math -std -complete -buildid Ua58p01xnrxrWgsDh9BS/Ua58p01xnrxrWgsDh9BS -goversion go1.22.4 -c=16 -nolocalimports -importcfg $WORK/b018/importcfg -pack /usr/local/go/src/runtime/internal/math/math.go

/usr/local/go/pkg/tool/linux_amd64/buildid -w $WORK/b018/_pkg_.a # internal

cp $WORK/b018/_pkg_.a /root/.cache/go-build/d7/d7a90d2a206876d53a91125d41f2f54f15dc6b9624bb3d677f4e05c933feaad9-d # internal

##############################################################################

runtime/internal/sys
mkdir -p $WORK/b019/
cat >/tmp/go-build2485170101/b019/importcfg << 'EOF' # internal
# import config
packagefile internal/goarch=/tmp/go-build2485170101/b006/_pkg_.a
packagefile internal/goos=/tmp/go-build2485170101/b016/_pkg_.a
EOF

/usr/local/go/pkg/tool/linux_amd64/compile -o $WORK/b019/_pkg_.a -trimpath "$WORK/b019=>" -p runtime/internal/sys -std -complete -buildid UOIfLKZiTQdoJLbOVF8f/UOIfLKZiTQdoJLbOVF8f -goversion go1.22.4 -c=16 -nolocalimports -importcfg $WORK/b019/importcfg -pack /usr/local/go/src/runtime/internal/sys/consts.go /usr/local/go/src/runtime/internal/sys/consts_norace.go /usr/local/go/src/runtime/internal/sys/intrinsics.go /usr/local/go/src/runtime/internal/sys/nih.go /usr/local/go/src/runtime/internal/sys/sys.go /usr/local/go/src/runtime/internal/sys/zversion.go

/usr/local/go/pkg/tool/linux_amd64/buildid -w $WORK/b019/_pkg_.a # internal

cp $WORK/b019/_pkg_.a /root/.cache/go-build/c6/c69e15df4d5d6b6afc87051da30e5cf7c646814c2cfabbfe70a194e6b600bd87-d # internal

##############################################################################

runtime/internal/syscall
mkdir -p $WORK/b020/
echo -n > $WORK/b020/go_asm.h # internal
cd /usr/local/go/src/runtime/internal/syscall

/usr/local/go/pkg/tool/linux_amd64/asm -p runtime/internal/syscall -trimpath "$WORK/b020=>" -I $WORK/b020/ -I /usr/local/go/pkg/include -D GOOS_linux -D GOARCH_amd64 -D GOAMD64_v1 -gensymabis -o $WORK/b020/symabis ./asm_linux_amd64.s

echo '# import config' > $WORK/b020/importcfg # internal

cd /go-std

/usr/local/go/pkg/tool/linux_amd64/compile -o $WORK/b020/_pkg_.a -trimpath "$WORK/b020=>" -p runtime/internal/syscall -std -buildid ruUY2seTm9ocdknoSUjx/ruUY2seTm9ocdknoSUjx -goversion go1.22.4 -symabis $WORK/b020/symabis -c=16 -nolocalimports -importcfg $WORK/b020/importcfg -pack -asmhdr $WORK/b020/go_asm.h /usr/local/go/src/runtime/internal/syscall/defs_linux_amd64.go /usr/local/go/src/runtime/internal/syscall/syscall_linux.go

cd /usr/local/go/src/runtime/internal/syscall

/usr/local/go/pkg/tool/linux_amd64/asm -p runtime/internal/syscall -trimpath "$WORK/b020=>" -I $WORK/b020/ -I /usr/local/go/pkg/include -D GOOS_linux -D GOARCH_amd64 -D GOAMD64_v1 -o $WORK/b020/asm_linux_amd64.o ./asm_linux_amd64.s

/usr/local/go/pkg/tool/linux_amd64/pack r $WORK/b020/_pkg_.a $WORK/b020/asm_linux_amd64.o # internal

/usr/local/go/pkg/tool/linux_amd64/buildid -w $WORK/b020/_pkg_.a # internal

cp $WORK/b020/_pkg_.a /root/.cache/go-build/48/4803738011c724c5df9a469dac934f20a3a6741535265181fb5025a589713f08-d # internal

##############################################################################

runtime
mkdir -p $WORK/b009/
echo -n > $WORK/b009/go_asm.h # internal
cd /usr/local/go/src/runtime

/usr/local/go/pkg/tool/linux_amd64/asm -p runtime -trimpath "$WORK/b009=>" -I $WORK/b009/ -I /usr/local/go/pkg/include -D GOOS_linux -D GOARCH_amd64 -D GOAMD64_v1 -gensymabis -o $WORK/b009/symabis ./asm.s ./asm_amd64.s ./duff_amd64.s ./memclr_amd64.s ./memmove_amd64.s ./preempt_amd64.s ./rt0_linux_amd64.s ./sys_linux_amd64.s ./test_amd64.s ./time_linux_amd64.s

cat >/tmp/go-build2485170101/b009/importcfg << 'EOF' # internal
# import config
packagefile internal/abi=/tmp/go-build2485170101/b005/_pkg_.a
packagefile internal/bytealg=/tmp/go-build2485170101/b010/_pkg_.a
packagefile internal/chacha8rand=/tmp/go-build2485170101/b012/_pkg_.a
packagefile internal/coverage/rtcov=/tmp/go-build2485170101/b013/_pkg_.a
packagefile internal/cpu=/tmp/go-build2485170101/b011/_pkg_.a
packagefile internal/goarch=/tmp/go-build2485170101/b006/_pkg_.a
packagefile internal/godebugs=/tmp/go-build2485170101/b014/_pkg_.a
packagefile internal/goexperiment=/tmp/go-build2485170101/b015/_pkg_.a
packagefile internal/goos=/tmp/go-build2485170101/b016/_pkg_.a
packagefile runtime/internal/atomic=/tmp/go-build2485170101/b017/_pkg_.a
packagefile runtime/internal/math=/tmp/go-build2485170101/b018/_pkg_.a
packagefile runtime/internal/sys=/tmp/go-build2485170101/b019/_pkg_.a
packagefile runtime/internal/syscall=/tmp/go-build2485170101/b020/_pkg_.a
EOF

cd /go-std

/usr/local/go/pkg/tool/linux_amd64/compile \
-o $WORK/b009/_pkg_.a \
-trimpath "$WORK/b009=>" \
-p runtime \
-std \
-buildid EPdUEQ9F-ABjjMT6xhlM/EPdUEQ9F-ABjjMT6xhlM \
-goversion go1.22.4 \
-symabis $WORK/b009/symabis \
-c=16 \
-nolocalimports \
-importcfg $WORK/b009/importcfg \
-pack -asmhdr \
$WORK/b009/go_asm.h \
/usr/local/go/src/runtime/alg.go /usr/local/go/src/runtime/arena.go /usr/local/go/src/runtime/asan0.go /usr/local/go/src/runtime/atomic_pointer.go /usr/local/go/src/runtime/cgo.go /usr/local/go/src/runtime/cgo_mmap.go /usr/local/go/src/runtime/cgo_sigaction.go /usr/local/go/src/runtime/cgocall.go /usr/local/go/src/runtime/cgocallback.go /usr/local/go/src/runtime/cgocheck.go /usr/local/go/src/runtime/chan.go /usr/local/go/src/runtime/checkptr.go /usr/local/go/src/runtime/compiler.go /usr/local/go/src/runtime/complex.go /usr/local/go/src/runtime/coro.go /usr/local/go/src/runtime/covercounter.go /usr/local/go/src/runtime/covermeta.go /usr/local/go/src/runtime/cpuflags.go /usr/local/go/src/runtime/cpuflags_amd64.go /usr/local/go/src/runtime/cpuprof.go /usr/local/go/src/runtime/cputicks.go /usr/local/go/src/runtime/create_file_unix.go /usr/local/go/src/runtime/debug.go /usr/local/go/src/runtime/debugcall.go /usr/local/go/src/runtime/debuglog.go /usr/local/go/src/runtime/debuglog_off.go /usr/local/go/src/runtime/defs_linux_amd64.go /usr/local/go/src/runtime/env_posix.go /usr/local/go/src/runtime/error.go /usr/local/go/src/runtime/exithook.go /usr/local/go/src/runtime/extern.go /usr/local/go/src/runtime/fastlog2.go /usr/local/go/src/runtime/fastlog2table.go /usr/local/go/src/runtime/fds_unix.go /usr/local/go/src/runtime/float.go /usr/local/go/src/runtime/hash64.go /usr/local/go/src/runtime/heapdump.go /usr/local/go/src/runtime/histogram.go /usr/local/go/src/runtime/iface.go /usr/local/go/src/runtime/lfstack.go /usr/local/go/src/runtime/lock_futex.go /usr/local/go/src/runtime/lockrank.go /usr/local/go/src/runtime/lockrank_off.go /usr/local/go/src/runtime/malloc.go /usr/local/go/src/runtime/map.go /usr/local/go/src/runtime/map_fast32.go /usr/local/go/src/runtime/map_fast64.go /usr/local/go/src/runtime/map_faststr.go /usr/local/go/src/runtime/mbarrier.go /usr/local/go/src/runtime/mbitmap.go /usr/local/go/src/runtime/mbitmap_allocheaders.go /usr/local/go/src/runtime/mcache.go /usr/local/go/src/runtime/mcentral.go /usr/local/go/src/runtime/mcheckmark.go /usr/local/go/src/runtime/mem.go /usr/local/go/src/runtime/mem_linux.go /usr/local/go/src/runtime/metrics.go /usr/local/go/src/runtime/mfinal.go /usr/local/go/src/runtime/mfixalloc.go /usr/local/go/src/runtime/mgc.go /usr/local/go/src/runtime/mgclimit.go /usr/local/go/src/runtime/mgcmark.go /usr/local/go/src/runtime/mgcpacer.go /usr/local/go/src/runtime/mgcscavenge.go /usr/local/go/src/runtime/mgcstack.go /usr/local/go/src/runtime/mgcsweep.go /usr/local/go/src/runtime/mgcwork.go /usr/local/go/src/runtime/mheap.go /usr/local/go/src/runtime/minmax.go /usr/local/go/src/runtime/mpagealloc.go /usr/local/go/src/runtime/mpagealloc_64bit.go /usr/local/go/src/runtime/mpagecache.go /usr/local/go/src/runtime/mpallocbits.go /usr/local/go/src/runtime/mprof.go /usr/local/go/src/runtime/mranges.go /usr/local/go/src/runtime/msan0.go /usr/local/go/src/runtime/msize_allocheaders.go /usr/local/go/src/runtime/mspanset.go /usr/local/go/src/runtime/mstats.go /usr/local/go/src/runtime/mwbbuf.go /usr/local/go/src/runtime/nbpipe_pipe2.go /usr/local/go/src/runtime/netpoll.go /usr/local/go/src/runtime/netpoll_epoll.go /usr/local/go/src/runtime/nonwindows_stub.go /usr/local/go/src/runtime/os_linux.go /usr/local/go/src/runtime/os_linux_generic.go /usr/local/go/src/runtime/os_linux_noauxv.go /usr/local/go/src/runtime/os_linux_x86.go /usr/local/go/src/runtime/os_nonopenbsd.go /usr/local/go/src/runtime/os_unix.go /usr/local/go/src/runtime/pagetrace_off.go /usr/local/go/src/runtime/panic.go /usr/local/go/src/runtime/pinner.go /usr/local/go/src/runtime/plugin.go /usr/local/go/src/runtime/preempt.go /usr/local/go/src/runtime/preempt_nonwindows.go /usr/local/go/src/runtime/print.go /usr/local/go/src/runtime/proc.go /usr/local/go/src/runtime/profbuf.go /usr/local/go/src/runtime/proflabel.go /usr/local/go/src/runtime/race0.go /usr/local/go/src/runtime/rand.go /usr/local/go/src/runtime/rdebug.go /usr/local/go/src/runtime/retry.go /usr/local/go/src/runtime/runtime.go /usr/local/go/src/runtime/runtime1.go /usr/local/go/src/runtime/runtime2.go /usr/local/go/src/runtime/runtime_boring.go /usr/local/go/src/runtime/rwmutex.go /usr/local/go/src/runtime/security_linux.go /usr/local/go/src/runtime/security_unix.go /usr/local/go/src/runtime/select.go /usr/local/go/src/runtime/sema.go /usr/local/go/src/runtime/signal_amd64.go /usr/local/go/src/runtime/signal_linux_amd64.go /usr/local/go/src/runtime/signal_unix.go /usr/local/go/src/runtime/sigqueue.go /usr/local/go/src/runtime/sigqueue_note.go /usr/local/go/src/runtime/sigtab_linux_generic.go /usr/local/go/src/runtime/sizeclasses.go /usr/local/go/src/runtime/slice.go /usr/local/go/src/runtime/softfloat64.go /usr/local/go/src/runtime/stack.go /usr/local/go/src/runtime/stkframe.go /usr/local/go/src/runtime/string.go /usr/local/go/src/runtime/stubs.go /usr/local/go/src/runtime/stubs2.go /usr/local/go/src/runtime/stubs3.go /usr/local/go/src/runtime/stubs_amd64.go /usr/local/go/src/runtime/stubs_linux.go /usr/local/go/src/runtime/symtab.go /usr/local/go/src/runtime/symtabinl.go /usr/local/go/src/runtime/sys_nonppc64x.go /usr/local/go/src/runtime/sys_x86.go /usr/local/go/src/runtime/tagptr.go /usr/local/go/src/runtime/tagptr_64bit.go /usr/local/go/src/runtime/test_amd64.go /usr/local/go/src/runtime/time.go /usr/local/go/src/runtime/time_nofake.go /usr/local/go/src/runtime/timeasm.go /usr/local/go/src/runtime/tls_stub.go /usr/local/go/src/runtime/trace2.go /usr/local/go/src/runtime/trace2buf.go /usr/local/go/src/runtime/trace2cpu.go /usr/local/go/src/runtime/trace2event.go /usr/local/go/src/runtime/trace2map.go /usr/local/go/src/runtime/trace2region.go /usr/local/go/src/runtime/trace2runtime.go /usr/local/go/src/runtime/trace2stack.go /usr/local/go/src/runtime/trace2status.go /usr/local/go/src/runtime/trace2string.go /usr/local/go/src/runtime/trace2time.go /usr/local/go/src/runtime/traceback.go /usr/local/go/src/runtime/type.go /usr/local/go/src/runtime/typekind.go /usr/local/go/src/runtime/unsafe.go /usr/local/go/src/runtime/utf8.go /usr/local/go/src/runtime/vdso_elf64.go /usr/local/go/src/runtime/vdso_linux.go /usr/local/go/src/runtime/vdso_linux_amd64.go /usr/local/go/src/runtime/write_err.go

cp /usr/local/go/src/runtime/asm_amd64.h $WORK/b009/asm_GOARCH.h

cd /usr/local/go/src/runtime

/usr/local/go/pkg/tool/linux_amd64/asm -p runtime -trimpath "$WORK/b009=>" -I $WORK/b009/ -I /usr/local/go/pkg/include -D GOOS_linux -D GOARCH_amd64 -D GOAMD64_v1 -o $WORK/b009/asm.o ./asm.s

/usr/local/go/pkg/tool/linux_amd64/asm -p runtime -trimpath "$WORK/b009=>" -I $WORK/b009/ -I /usr/local/go/pkg/include -D GOOS_linux -D GOARCH_amd64 -D GOAMD64_v1 -o $WORK/b009/asm_amd64.o ./asm_amd64.s

/usr/local/go/pkg/tool/linux_amd64/asm -p runtime -trimpath "$WORK/b009=>" -I $WORK/b009/ -I /usr/local/go/pkg/include -D GOOS_linux -D GOARCH_amd64 -D GOAMD64_v1 -o $WORK/b009/duff_amd64.o ./duff_amd64.s

/usr/local/go/pkg/tool/linux_amd64/asm -p runtime -trimpath "$WORK/b009=>" -I $WORK/b009/ -I /usr/local/go/pkg/include -D GOOS_linux -D GOARCH_amd64 -D GOAMD64_v1 -o $WORK/b009/memclr_amd64.o ./memclr_amd64.s

/usr/local/go/pkg/tool/linux_amd64/asm -p runtime -trimpath "$WORK/b009=>" -I $WORK/b009/ -I /usr/local/go/pkg/include -D GOOS_linux -D GOARCH_amd64 -D GOAMD64_v1 -o $WORK/b009/memmove_amd64.o ./memmove_amd64.s

/usr/local/go/pkg/tool/linux_amd64/asm -p runtime -trimpath "$WORK/b009=>" -I $WORK/b009/ -I /usr/local/go/pkg/include -D GOOS_linux -D GOARCH_amd64 -D GOAMD64_v1 -o $WORK/b009/preempt_amd64.o ./preempt_amd64.s

/usr/local/go/pkg/tool/linux_amd64/asm -p runtime -trimpath "$WORK/b009=>" -I $WORK/b009/ -I /usr/local/go/pkg/include -D GOOS_linux -D GOARCH_amd64 -D GOAMD64_v1 -o $WORK/b009/rt0_linux_amd64.o ./rt0_linux_amd64.s

/usr/local/go/pkg/tool/linux_amd64/asm -p runtime -trimpath "$WORK/b009=>" -I $WORK/b009/ -I /usr/local/go/pkg/include -D GOOS_linux -D GOARCH_amd64 -D GOAMD64_v1 -o $WORK/b009/sys_linux_amd64.o ./sys_linux_amd64.s

/usr/local/go/pkg/tool/linux_amd64/asm -p runtime -trimpath "$WORK/b009=>" -I $WORK/b009/ -I /usr/local/go/pkg/include -D GOOS_linux -D GOARCH_amd64 -D GOAMD64_v1 -o $WORK/b009/test_amd64.o ./test_amd64.s

/usr/local/go/pkg/tool/linux_amd64/asm -p runtime -trimpath "$WORK/b009=>" -I $WORK/b009/ -I /usr/local/go/pkg/include -D GOOS_linux -D GOARCH_amd64 -D GOAMD64_v1 -o $WORK/b009/time_linux_amd64.o ./time_linux_amd64.s

/usr/local/go/pkg/tool/linux_amd64/pack r $WORK/b009/_pkg_.a $WORK/b009/asm.o $WORK/b009/asm_amd64.o $WORK/b009/duff_amd64.o $WORK/b009/memclr_amd64.o $WORK/b009/memmove_amd64.o $WORK/b009/preempt_amd64.o $WORK/b009/rt0_linux_amd64.o $WORK/b009/sys_linux_amd64.o $WORK/b009/test_amd64.o $WORK/b009/time_linux_amd64.o # internal

/usr/local/go/pkg/tool/linux_amd64/buildid -w $WORK/b009/_pkg_.a # internal

cp $WORK/b009/_pkg_.a /root/.cache/go-build/b6/b6038b800a3440102bcc90ba50c52f533b73d74a719e30fe432a91ed350f227b-d # internal

##############################################################################

internal/reflectlite
mkdir -p $WORK/b004/
echo -n > $WORK/b004/go_asm.h # internal
cd /usr/local/go/src/internal/reflectlite

/usr/local/go/pkg/tool/linux_amd64/asm -p internal/reflectlite -trimpath "$WORK/b004=>" -I $WORK/b004/ -I /usr/local/go/pkg/include -D GOOS_linux -D GOARCH_amd64 -D GOAMD64_v1 -gensymabis -o $WORK/b004/symabis ./asm.s

cat >/tmp/go-build2485170101/b004/importcfg << 'EOF' # internal
# import config
packagefile internal/abi=/tmp/go-build2485170101/b005/_pkg_.a
packagefile internal/goarch=/tmp/go-build2485170101/b006/_pkg_.a
packagefile internal/unsafeheader=/tmp/go-build2485170101/b008/_pkg_.a
packagefile runtime=/tmp/go-build2485170101/b009/_pkg_.a
EOF

cd /go-std

/usr/local/go/pkg/tool/linux_amd64/compile -o $WORK/b004/_pkg_.a -trimpath "$WORK/b004=>" -p internal/reflectlite -std -buildid h0S7vyeKPdEFf8SOEgEI/h0S7vyeKPdEFf8SOEgEI -goversion go1.22.4 -symabis $WORK/b004/symabis -c=16 -nolocalimports -importcfg $WORK/b004/importcfg -pack -asmhdr $WORK/b004/go_asm.h /usr/local/go/src/internal/reflectlite/swapper.go /usr/local/go/src/internal/reflectlite/type.go /usr/local/go/src/internal/reflectlite/value.go

cd /usr/local/go/src/internal/reflectlite

/usr/local/go/pkg/tool/linux_amd64/asm -p internal/reflectlite -trimpath "$WORK/b004=>" -I $WORK/b004/ -I /usr/local/go/pkg/include -D GOOS_linux -D GOARCH_amd64 -D GOAMD64_v1 -o $WORK/b004/asm.o ./asm.s

/usr/local/go/pkg/tool/linux_amd64/pack r $WORK/b004/_pkg_.a $WORK/b004/asm.o # internal

/usr/local/go/pkg/tool/linux_amd64/buildid -w $WORK/b004/_pkg_.a # internal

cp $WORK/b004/_pkg_.a /root/.cache/go-build/6e/6ed359e02323b2f4016ec246547b425bb60348a7aa3e2fab2884a75593b49e32-d # internal

##############################################################################

errors
mkdir -p $WORK/b003/

cat >/tmp/go-build2485170101/b003/importcfg << 'EOF' # internal
# import config
packagefile internal/reflectlite=/tmp/go-build2485170101/b004/_pkg_.a
EOF

cd /go-std

/usr/local/go/pkg/tool/linux_amd64/compile -o $WORK/b003/_pkg_.a -trimpath "$WORK/b003=>" -p errors -std -complete -buildid Q1ZeLZdAnYTGKnkw6EEM/Q1ZeLZdAnYTGKnkw6EEM -goversion go1.22.4 -c=16 -nolocalimports -importcfg $WORK/b003/importcfg -pack /usr/local/go/src/errors/errors.go /usr/local/go/src/errors/join.go /usr/local/go/src/errors/wrap.go

/usr/local/go/pkg/tool/linux_amd64/buildid -w $WORK/b003/_pkg_.a # internal

cp $WORK/b003/_pkg_.a /root/.cache/go-build/74/745cf7d8fc8d63547629a9189fc602c9509900cd8569c987b27a69c57a9fd5ad-d # internal

##############################################################################

internal/itoa
mkdir -p $WORK/b023/
echo '# import config' > $WORK/b023/importcfg # internal

/usr/local/go/pkg/tool/linux_amd64/compile -o $WORK/b023/_pkg_.a -trimpath "$WORK/b023=>" -p internal/itoa -std -complete -buildid o5aD7UZtPDa3umw05j6L/o5aD7UZtPDa3umw05j6L -goversion go1.22.4 -c=16 -nolocalimports -importcfg $WORK/b023/importcfg -pack /usr/local/go/src/internal/itoa/itoa.go

/usr/local/go/pkg/tool/linux_amd64/buildid -w $WORK/b023/_pkg_.a # internal

cp $WORK/b023/_pkg_.a /root/.cache/go-build/72/7289ff1401de0607f9276b107de732f725f3514ba316106b66af7363d8415ab0-d # internal

##############################################################################

math/bits
mkdir -p $WORK/b025/
echo '# import config' > $WORK/b025/importcfg # internal

/usr/local/go/pkg/tool/linux_amd64/compile -o $WORK/b025/_pkg_.a -trimpath "$WORK/b025=>" -p math/bits -std -complete -buildid 9JLPifiVRAh3D2u7R8oC/9JLPifiVRAh3D2u7R8oC -goversion go1.22.4 -c=16 -nolocalimports -importcfg $WORK/b025/importcfg -pack /usr/local/go/src/math/bits/bits.go /usr/local/go/src/math/bits/bits_errors.go /usr/local/go/src/math/bits/bits_tables.go

/usr/local/go/pkg/tool/linux_amd64/buildid -w $WORK/b025/_pkg_.a # internal

cp $WORK/b025/_pkg_.a /root/.cache/go-build/a8/a85d3cffdc99e2d02b0c6a68ad0d018a524386fe4313c30bb09c87d526c2f150-d # internal

##############################################################################

math
mkdir -p $WORK/b024/
echo -n > $WORK/b024/go_asm.h # internal
cd /usr/local/go/src/math

/usr/local/go/pkg/tool/linux_amd64/asm -p math -trimpath "$WORK/b024=>" -I $WORK/b024/ -I /usr/local/go/pkg/include -D GOOS_linux -D GOARCH_amd64 -D GOAMD64_v1 -gensymabis -o $WORK/b024/symabis ./dim_amd64.s ./exp_amd64.s ./floor_amd64.s ./hypot_amd64.s ./log_amd64.s

cat >/tmp/go-build2485170101/b024/importcfg << 'EOF' # internal
# import config
packagefile internal/cpu=/tmp/go-build2485170101/b011/_pkg_.a
packagefile math/bits=/tmp/go-build2485170101/b025/_pkg_.a
EOF

cd /go-std

/usr/local/go/pkg/tool/linux_amd64/compile -o $WORK/b024/_pkg_.a -trimpath "$WORK/b024=>" -p math -std -buildid R9rP9BACOM6e_UD4c6wA/R9rP9BACOM6e_UD4c6wA -goversion go1.22.4 -symabis $WORK/b024/symabis -c=16 -nolocalimports -importcfg $WORK/b024/importcfg -pack -asmhdr $WORK/b024/go_asm.h /usr/local/go/src/math/abs.go /usr/local/go/src/math/acosh.go /usr/local/go/src/math/asin.go /usr/local/go/src/math/asinh.go /usr/local/go/src/math/atan.go /usr/local/go/src/math/atan2.go /usr/local/go/src/math/atanh.go /usr/local/go/src/math/bits.go /usr/local/go/src/math/cbrt.go /usr/local/go/src/math/const.go /usr/local/go/src/math/copysign.go /usr/local/go/src/math/dim.go /usr/local/go/src/math/dim_asm.go /usr/local/go/src/math/erf.go /usr/local/go/src/math/erfinv.go /usr/local/go/src/math/exp.go /usr/local/go/src/math/exp2_noasm.go /usr/local/go/src/math/exp_amd64.go /usr/local/go/src/math/exp_asm.go /usr/local/go/src/math/expm1.go /usr/local/go/src/math/floor.go /usr/local/go/src/math/floor_asm.go /usr/local/go/src/math/fma.go /usr/local/go/src/math/frexp.go /usr/local/go/src/math/gamma.go /usr/local/go/src/math/hypot.go /usr/local/go/src/math/hypot_asm.go /usr/local/go/src/math/j0.go /usr/local/go/src/math/j1.go /usr/local/go/src/math/jn.go /usr/local/go/src/math/ldexp.go /usr/local/go/src/math/lgamma.go /usr/local/go/src/math/log.go /usr/local/go/src/math/log10.go /usr/local/go/src/math/log1p.go /usr/local/go/src/math/log_asm.go /usr/local/go/src/math/logb.go /usr/local/go/src/math/mod.go /usr/local/go/src/math/modf.go /usr/local/go/src/math/modf_noasm.go /usr/local/go/src/math/nextafter.go /usr/local/go/src/math/pow.go /usr/local/go/src/math/pow10.go /usr/local/go/src/math/remainder.go /usr/local/go/src/math/signbit.go /usr/local/go/src/math/sin.go /usr/local/go/src/math/sincos.go /usr/local/go/src/math/sinh.go /usr/local/go/src/math/sqrt.go /usr/local/go/src/math/stubs.go /usr/local/go/src/math/tan.go /usr/local/go/src/math/tanh.go /usr/local/go/src/math/trig_reduce.go /usr/local/go/src/math/unsafe.go

cd /usr/local/go/src/math

/usr/local/go/pkg/tool/linux_amd64/asm -p math -trimpath "$WORK/b024=>" -I $WORK/b024/ -I /usr/local/go/pkg/include -D GOOS_linux -D GOARCH_amd64 -D GOAMD64_v1 -o $WORK/b024/dim_amd64.o ./dim_amd64.s

/usr/local/go/pkg/tool/linux_amd64/asm -p math -trimpath "$WORK/b024=>" -I $WORK/b024/ -I /usr/local/go/pkg/include -D GOOS_linux -D GOARCH_amd64 -D GOAMD64_v1 -o $WORK/b024/exp_amd64.o ./exp_amd64.s

/usr/local/go/pkg/tool/linux_amd64/asm -p math -trimpath "$WORK/b024=>" -I $WORK/b024/ -I /usr/local/go/pkg/include -D GOOS_linux -D GOARCH_amd64 -D GOAMD64_v1 -o $WORK/b024/floor_amd64.o ./floor_amd64.s

/usr/local/go/pkg/tool/linux_amd64/asm -p math -trimpath "$WORK/b024=>" -I $WORK/b024/ -I /usr/local/go/pkg/include -D GOOS_linux -D GOARCH_amd64 -D GOAMD64_v1 -o $WORK/b024/hypot_amd64.o ./hypot_amd64.s

/usr/local/go/pkg/tool/linux_amd64/asm -p math -trimpath "$WORK/b024=>" -I $WORK/b024/ -I /usr/local/go/pkg/include -D GOOS_linux -D GOARCH_amd64 -D GOAMD64_v1 -o $WORK/b024/log_amd64.o ./log_amd64.s

/usr/local/go/pkg/tool/linux_amd64/pack r $WORK/b024/_pkg_.a $WORK/b024/dim_amd64.o $WORK/b024/exp_amd64.o $WORK/b024/floor_amd64.o $WORK/b024/hypot_amd64.o $WORK/b024/log_amd64.o # internal

/usr/local/go/pkg/tool/linux_amd64/buildid -w $WORK/b024/_pkg_.a # internal

cp $WORK/b024/_pkg_.a /root/.cache/go-build/9d/9d848e24b1bacd79b1be6535f07619bef71a6b871c124ec1e1454a76518beb4b-d # internal

##############################################################################

unicode/utf8
mkdir -p $WORK/b027/
echo '# import config' > $WORK/b027/importcfg # internal
cd /go-std

/usr/local/go/pkg/tool/linux_amd64/compile -o $WORK/b027/_pkg_.a -trimpath "$WORK/b027=>" -p unicode/utf8 -std -complete -buildid txDdXjeJM8CGv03nM1gV/txDdXjeJM8CGv03nM1gV -goversion go1.22.4 -c=16 -nolocalimports -importcfg $WORK/b027/importcfg -pack /usr/local/go/src/unicode/utf8/utf8.go

/usr/local/go/pkg/tool/linux_amd64/buildid -w $WORK/b027/_pkg_.a # internal

cp $WORK/b027/_pkg_.a /root/.cache/go-build/be/be08f8e664237e938d0bbf6b9a75e0154b109b41a0b033951dbc366f3e2ee6ab-d # internal

##############################################################################

strconv
mkdir -p $WORK/b026/

cat >/tmp/go-build2485170101/b026/importcfg << 'EOF' # internal
# import config
packagefile errors=/tmp/go-build2485170101/b003/_pkg_.a
packagefile internal/bytealg=/tmp/go-build2485170101/b010/_pkg_.a
packagefile math=/tmp/go-build2485170101/b024/_pkg_.a
packagefile math/bits=/tmp/go-build2485170101/b025/_pkg_.a
packagefile unicode/utf8=/tmp/go-build2485170101/b027/_pkg_.a
EOF

/usr/local/go/pkg/tool/linux_amd64/compile -o $WORK/b026/_pkg_.a -trimpath "$WORK/b026=>" -p strconv -std -complete -buildid kj-QcpbSm7YrjyecDRa7/kj-QcpbSm7YrjyecDRa7 -goversion go1.22.4 -c=16 -nolocalimports -importcfg $WORK/b026/importcfg -pack /usr/local/go/src/strconv/atob.go /usr/local/go/src/strconv/atoc.go /usr/local/go/src/strconv/atof.go /usr/local/go/src/strconv/atoi.go /usr/local/go/src/strconv/bytealg.go /usr/local/go/src/strconv/ctoa.go /usr/local/go/src/strconv/decimal.go /usr/local/go/src/strconv/doc.go /usr/local/go/src/strconv/eisel_lemire.go /usr/local/go/src/strconv/ftoa.go /usr/local/go/src/strconv/ftoaryu.go /usr/local/go/src/strconv/isprint.go /usr/local/go/src/strconv/itoa.go /usr/local/go/src/strconv/quote.go

/usr/local/go/pkg/tool/linux_amd64/buildid -w $WORK/b026/_pkg_.a # internal

cp $WORK/b026/_pkg_.a /root/.cache/go-build/56/56a03f66beb4da0d55134646ba564cad6ac19d76af6600e5a485d70dbcc8d073-d # internal

##############################################################################

internal/race
mkdir -p $WORK/b029/
echo '# import config' > $WORK/b029/importcfg # internal

/usr/local/go/pkg/tool/linux_amd64/compile -o $WORK/b029/_pkg_.a -trimpath "$WORK/b029=>" -p internal/race -std -complete -buildid XJIBhGQlxSjpLwMM2dnQ/XJIBhGQlxSjpLwMM2dnQ -goversion go1.22.4 -c=16 -nolocalimports -importcfg $WORK/b029/importcfg -pack /usr/local/go/src/internal/race/doc.go /usr/local/go/src/internal/race/norace.go

/usr/local/go/pkg/tool/linux_amd64/buildid -w $WORK/b029/_pkg_.a # internal

cp $WORK/b029/_pkg_.a /root/.cache/go-build/86/867e3a3c64ef622e90ed5ca0e98c513033f04de1ae0b85ea58b5bf23da5fa5e0-d # internal

##############################################################################

sync/atomic
mkdir -p $WORK/b030/
echo -n > $WORK/b030/go_asm.h # internal
cd /usr/local/go/src/sync/atomic

/usr/local/go/pkg/tool/linux_amd64/asm -p sync/atomic -trimpath "$WORK/b030=>" -I $WORK/b030/ -I /usr/local/go/pkg/include -D GOOS_linux -D GOARCH_amd64 -D GOAMD64_v1 -gensymabis -o $WORK/b030/symabis ./asm.s

echo '# import config' > $WORK/b030/importcfg # internal

cd /go-std

/usr/local/go/pkg/tool/linux_amd64/compile -o $WORK/b030/_pkg_.a -trimpath "$WORK/b030=>" -p sync/atomic -std -buildid 8VEN57uVPKpavVjgloAE/8VEN57uVPKpavVjgloAE -goversion go1.22.4 -symabis $WORK/b030/symabis -c=16 -nolocalimports -importcfg $WORK/b030/importcfg -pack -asmhdr $WORK/b030/go_asm.h /usr/local/go/src/sync/atomic/doc.go /usr/local/go/src/sync/atomic/type.go /usr/local/go/src/sync/atomic/value.go

cd /usr/local/go/src/sync/atomic

/usr/local/go/pkg/tool/linux_amd64/asm -p sync/atomic -trimpath "$WORK/b030=>" -I $WORK/b030/ -I /usr/local/go/pkg/include -D GOOS_linux -D GOARCH_amd64 -D GOAMD64_v1 -o $WORK/b030/asm.o ./asm.s

/usr/local/go/pkg/tool/linux_amd64/pack r $WORK/b030/_pkg_.a $WORK/b030/asm.o # internal

/usr/local/go/pkg/tool/linux_amd64/buildid -w $WORK/b030/_pkg_.a # internal

cp $WORK/b030/_pkg_.a /root/.cache/go-build/d1/d1b3c91c3569ae042bb991aa23c4f19a3442882608ec88856eb5d134dc247cfb-d # internal

##############################################################################

sync
mkdir -p $WORK/b028/

cat >/tmp/go-build2485170101/b028/importcfg << 'EOF' # internal
# import config
packagefile internal/race=/tmp/go-build2485170101/b029/_pkg_.a
packagefile runtime=/tmp/go-build2485170101/b009/_pkg_.a
packagefile sync/atomic=/tmp/go-build2485170101/b030/_pkg_.a
EOF

cd /go-std

/usr/local/go/pkg/tool/linux_amd64/compile -o $WORK/b028/_pkg_.a -trimpath "$WORK/b028=>" -p sync -std -buildid C_Mbea2uwyX86sWwE1sA/C_Mbea2uwyX86sWwE1sA -goversion go1.22.4 -c=16 -nolocalimports -importcfg $WORK/b028/importcfg -pack /usr/local/go/src/sync/cond.go /usr/local/go/src/sync/map.go /usr/local/go/src/sync/mutex.go /usr/local/go/src/sync/once.go /usr/local/go/src/sync/oncefunc.go /usr/local/go/src/sync/pool.go /usr/local/go/src/sync/poolqueue.go /usr/local/go/src/sync/runtime.go /usr/local/go/src/sync/runtime2.go /usr/local/go/src/sync/rwmutex.go /usr/local/go/src/sync/waitgroup.go

/usr/local/go/pkg/tool/linux_amd64/buildid -w $WORK/b028/_pkg_.a # internal

cp $WORK/b028/_pkg_.a /root/.cache/go-build/d4/d498155d59c1f5b3b759ba5c8fd9347474b0611cfb58b57567131ea23ee2c52a-d # internal

##############################################################################

unicode
mkdir -p $WORK/b031/
echo '# import config' > $WORK/b031/importcfg # internal

/usr/local/go/pkg/tool/linux_amd64/compile -o $WORK/b031/_pkg_.a -trimpath "$WORK/b031=>" -p unicode -std -complete -buildid ncgJdBBQrZea3puCN1uL/ncgJdBBQrZea3puCN1uL -goversion go1.22.4 -c=16 -nolocalimports -importcfg $WORK/b031/importcfg -pack /usr/local/go/src/unicode/casetables.go /usr/local/go/src/unicode/digit.go /usr/local/go/src/unicode/graphic.go /usr/local/go/src/unicode/letter.go /usr/local/go/src/unicode/tables.go

/usr/local/go/pkg/tool/linux_amd64/buildid -w $WORK/b031/_pkg_.a # internal

cp $WORK/b031/_pkg_.a /root/.cache/go-build/9a/9ac3ed0ba8260d5cec866c63f9a796c2ab7cf0f7ce682a19a676cfdc7d4cc587-d # internal

##############################################################################

reflect
mkdir -p $WORK/b022/
echo -n > $WORK/b022/go_asm.h # internal
cd /usr/local/go/src/reflect

/usr/local/go/pkg/tool/linux_amd64/asm -p reflect -trimpath "$WORK/b022=>" -I $WORK/b022/ -I /usr/local/go/pkg/include -D GOOS_linux -D GOARCH_amd64 -D GOAMD64_v1 -gensymabis -o $WORK/b022/symabis ./asm_amd64.s

cat >/tmp/go-build2485170101/b022/importcfg << 'EOF' # internal
# import config
packagefile errors=/tmp/go-build2485170101/b003/_pkg_.a
packagefile internal/abi=/tmp/go-build2485170101/b005/_pkg_.a
packagefile internal/bytealg=/tmp/go-build2485170101/b010/_pkg_.a
packagefile internal/goarch=/tmp/go-build2485170101/b006/_pkg_.a
packagefile internal/itoa=/tmp/go-build2485170101/b023/_pkg_.a
packagefile internal/unsafeheader=/tmp/go-build2485170101/b008/_pkg_.a
packagefile math=/tmp/go-build2485170101/b024/_pkg_.a
packagefile runtime=/tmp/go-build2485170101/b009/_pkg_.a
packagefile strconv=/tmp/go-build2485170101/b026/_pkg_.a
packagefile sync=/tmp/go-build2485170101/b028/_pkg_.a
packagefile unicode=/tmp/go-build2485170101/b031/_pkg_.a
packagefile unicode/utf8=/tmp/go-build2485170101/b027/_pkg_.a
EOF

cd /go-std

/usr/local/go/pkg/tool/linux_amd64/compile -o $WORK/b022/_pkg_.a -trimpath "$WORK/b022=>" -p reflect -std -buildid O2DAf7C9ilEpaaI0nEs7/O2DAf7C9ilEpaaI0nEs7 -goversion go1.22.4 -symabis $WORK/b022/symabis -c=16 -nolocalimports -importcfg $WORK/b022/importcfg -pack -asmhdr $WORK/b022/go_asm.h /usr/local/go/src/reflect/abi.go /usr/local/go/src/reflect/deepequal.go /usr/local/go/src/reflect/float32reg_generic.go /usr/local/go/src/reflect/makefunc.go /usr/local/go/src/reflect/swapper.go /usr/local/go/src/reflect/type.go /usr/local/go/src/reflect/value.go /usr/local/go/src/reflect/visiblefields.go

cd /usr/local/go/src/reflect

/usr/local/go/pkg/tool/linux_amd64/asm -p reflect -trimpath "$WORK/b022=>" -I $WORK/b022/ -I /usr/local/go/pkg/include -D GOOS_linux -D GOARCH_amd64 -D GOAMD64_v1 -o $WORK/b022/asm_amd64.o ./asm_amd64.s

/usr/local/go/pkg/tool/linux_amd64/pack r $WORK/b022/_pkg_.a $WORK/b022/asm_amd64.o # internal

/usr/local/go/pkg/tool/linux_amd64/buildid -w $WORK/b022/_pkg_.a # internal

cp $WORK/b022/_pkg_.a /root/.cache/go-build/4b/4bd691da8f5f0b06b99083f0678ac1013e51f5076427f6de0d557e5dca1ce5b7-d # internal

##############################################################################

cmp
mkdir -p $WORK/b034/
echo '# import config' > $WORK/b034/importcfg # internal
cd /go-std

/usr/local/go/pkg/tool/linux_amd64/compile -o $WORK/b034/_pkg_.a -trimpath "$WORK/b034=>" -p cmp -std -complete -buildid Y5f64ds_o1w1kVn-AJEu/Y5f64ds_o1w1kVn-AJEu -goversion go1.22.4 -c=16 -nolocalimports -importcfg $WORK/b034/importcfg -pack /usr/local/go/src/cmp/cmp.go

/usr/local/go/pkg/tool/linux_amd64/buildid -w $WORK/b034/_pkg_.a # internal

cp $WORK/b034/_pkg_.a /root/.cache/go-build/16/1626237c8e97ddc513281468e43264dc4404ac14f2f776ea17e23d158e206f4d-d # internal

##############################################################################

slices
mkdir -p $WORK/b033/

cat >/tmp/go-build2485170101/b033/importcfg << 'EOF' # internal
# import config
packagefile cmp=/tmp/go-build2485170101/b034/_pkg_.a
packagefile math/bits=/tmp/go-build2485170101/b025/_pkg_.a
EOF

/usr/local/go/pkg/tool/linux_amd64/compile -o $WORK/b033/_pkg_.a -trimpath "$WORK/b033=>" -p slices -std -complete -buildid XQ8WGCsae0_Y5AmxKXF6/XQ8WGCsae0_Y5AmxKXF6 -goversion go1.22.4 -c=16 -nolocalimports -importcfg $WORK/b033/importcfg -pack /usr/local/go/src/slices/slices.go /usr/local/go/src/slices/sort.go /usr/local/go/src/slices/zsortanyfunc.go /usr/local/go/src/slices/zsortordered.go

/usr/local/go/pkg/tool/linux_amd64/buildid -w $WORK/b033/_pkg_.a # internal

cp $WORK/b033/_pkg_.a /root/.cache/go-build/dc/dcb5d544ccdae464b7568bf8972ce0fce2c483412aab8e163901f82f0657fc89-d # internal

##############################################################################

sort
mkdir -p $WORK/b032/

cat >/tmp/go-build2485170101/b032/importcfg << 'EOF' # internal
# import config
packagefile internal/reflectlite=/tmp/go-build2485170101/b004/_pkg_.a
packagefile math/bits=/tmp/go-build2485170101/b025/_pkg_.a
packagefile slices=/tmp/go-build2485170101/b033/_pkg_.a
EOF

/usr/local/go/pkg/tool/linux_amd64/compile -o $WORK/b032/_pkg_.a -trimpath "$WORK/b032=>" -p sort -std -complete -buildid wWnwpXZbUGL9NEF1SJG0/wWnwpXZbUGL9NEF1SJG0 -goversion go1.22.4 -c=16 -nolocalimports -importcfg $WORK/b032/importcfg -pack /usr/local/go/src/sort/search.go /usr/local/go/src/sort/slice.go /usr/local/go/src/sort/sort.go /usr/local/go/src/sort/sort_impl_go121.go /usr/local/go/src/sort/zsortfunc.go /usr/local/go/src/sort/zsortinterface.go

/usr/local/go/pkg/tool/linux_amd64/buildid -w $WORK/b032/_pkg_.a # internal

cp $WORK/b032/_pkg_.a /root/.cache/go-build/1f/1fbe7109b7544c57a293b89023be942937f501d933b18cdcde22b04b1e2443d5-d # internal

##############################################################################

internal/fmtsort
mkdir -p $WORK/b021/

cat >/tmp/go-build2485170101/b021/importcfg << 'EOF' # internal
# import config
packagefile reflect=/tmp/go-build2485170101/b022/_pkg_.a
packagefile sort=/tmp/go-build2485170101/b032/_pkg_.a
EOF

/usr/local/go/pkg/tool/linux_amd64/compile -o $WORK/b021/_pkg_.a -trimpath "$WORK/b021=>" -p internal/fmtsort -std -complete -buildid 6zSBp3hHs1RTFI_urJdP/6zSBp3hHs1RTFI_urJdP -goversion go1.22.4 -c=16 -nolocalimports -importcfg $WORK/b021/importcfg -pack /usr/local/go/src/internal/fmtsort/sort.go

/usr/local/go/pkg/tool/linux_amd64/buildid -w $WORK/b021/_pkg_.a # internal

cp $WORK/b021/_pkg_.a /root/.cache/go-build/24/24a0086b902632014addc7c1006562a1ea35667c9924a256fab163e62fbe0669-d # internal

##############################################################################

io
mkdir -p $WORK/b035/

cat >/tmp/go-build2485170101/b035/importcfg << 'EOF' # internal
# import config
packagefile errors=/tmp/go-build2485170101/b003/_pkg_.a
packagefile sync=/tmp/go-build2485170101/b028/_pkg_.a
EOF

/usr/local/go/pkg/tool/linux_amd64/compile -o $WORK/b035/_pkg_.a -trimpath "$WORK/b035=>" -p io -std -complete -buildid YyJagy-FPkuo8Gq4tGzd/YyJagy-FPkuo8Gq4tGzd -goversion go1.22.4 -c=16 -nolocalimports -importcfg $WORK/b035/importcfg -pack /usr/local/go/src/io/io.go /usr/local/go/src/io/multi.go /usr/local/go/src/io/pipe.go

/usr/local/go/pkg/tool/linux_amd64/buildid -w $WORK/b035/_pkg_.a # internal

cp $WORK/b035/_pkg_.a /root/.cache/go-build/b7/b7c57d2fdd67685beec674b1a83dec003bcf2c51ed1690e419406082f582c099-d # internal

##############################################################################

internal/oserror
mkdir -p $WORK/b040/

cat >/tmp/go-build2485170101/b040/importcfg << 'EOF' # internal
# import config
packagefile errors=/tmp/go-build2485170101/b003/_pkg_.a
EOF

/usr/local/go/pkg/tool/linux_amd64/compile -o $WORK/b040/_pkg_.a -trimpath "$WORK/b040=>" -p internal/oserror -std -complete -buildid gsvIsczcKRoPZUNVIjfZ/gsvIsczcKRoPZUNVIjfZ -goversion go1.22.4 -c=16 -nolocalimports -importcfg $WORK/b040/importcfg -pack /usr/local/go/src/internal/oserror/errors.go

/usr/local/go/pkg/tool/linux_amd64/buildid -w $WORK/b040/_pkg_.a # internal

cp $WORK/b040/_pkg_.a /root/.cache/go-build/71/7145e08320bea2a6e58ede0fd11caac754a16164795146505dc51a1d331bf347-d # internal

##############################################################################

syscall
mkdir -p $WORK/b039/
echo -n > $WORK/b039/go_asm.h # internal
cd /usr/local/go/src/syscall

/usr/local/go/pkg/tool/linux_amd64/asm -p syscall -trimpath "$WORK/b039=>" -I $WORK/b039/ -I /usr/local/go/pkg/include -D GOOS_linux -D GOARCH_amd64 -D GOAMD64_v1 -gensymabis -o $WORK/b039/symabis ./asm_linux_amd64.s

cat >/tmp/go-build2485170101/b039/importcfg << 'EOF' # internal
# import config
packagefile errors=/tmp/go-build2485170101/b003/_pkg_.a
packagefile internal/bytealg=/tmp/go-build2485170101/b010/_pkg_.a
packagefile internal/itoa=/tmp/go-build2485170101/b023/_pkg_.a
packagefile internal/oserror=/tmp/go-build2485170101/b040/_pkg_.a
packagefile internal/race=/tmp/go-build2485170101/b029/_pkg_.a
packagefile runtime=/tmp/go-build2485170101/b009/_pkg_.a
packagefile sync=/tmp/go-build2485170101/b028/_pkg_.a
packagefile sync/atomic=/tmp/go-build2485170101/b030/_pkg_.a
EOF

cd /go-std

/usr/local/go/pkg/tool/linux_amd64/compile -o $WORK/b039/_pkg_.a -trimpath "$WORK/b039=>" -p syscall -std -buildid K9IiJgkCJdXAxrTBZp12/K9IiJgkCJdXAxrTBZp12 -goversion go1.22.4 -symabis $WORK/b039/symabis -c=16 -nolocalimports -importcfg $WORK/b039/importcfg -pack -asmhdr $WORK/b039/go_asm.h /usr/local/go/src/syscall/asan0.go /usr/local/go/src/syscall/dirent.go /usr/local/go/src/syscall/endian_little.go /usr/local/go/src/syscall/env_unix.go /usr/local/go/src/syscall/exec_linux.go /usr/local/go/src/syscall/exec_unix.go /usr/local/go/src/syscall/flock_linux.go /usr/local/go/src/syscall/forkpipe2.go /usr/local/go/src/syscall/lsf_linux.go /usr/local/go/src/syscall/msan0.go /usr/local/go/src/syscall/net.go /usr/local/go/src/syscall/netlink_linux.go /usr/local/go/src/syscall/rlimit.go /usr/local/go/src/syscall/rlimit_stub.go /usr/local/go/src/syscall/setuidgid_linux.go /usr/local/go/src/syscall/sockcmsg_linux.go /usr/local/go/src/syscall/sockcmsg_unix.go /usr/local/go/src/syscall/sockcmsg_unix_other.go /usr/local/go/src/syscall/syscall.go /usr/local/go/src/syscall/syscall_linux.go /usr/local/go/src/syscall/syscall_linux_accept4.go /usr/local/go/src/syscall/syscall_linux_amd64.go /usr/local/go/src/syscall/syscall_unix.go /usr/local/go/src/syscall/time_nofake.go /usr/local/go/src/syscall/timestruct.go /usr/local/go/src/syscall/zerrors_linux_amd64.go /usr/local/go/src/syscall/zsyscall_linux_amd64.go /usr/local/go/src/syscall/zsysnum_linux_amd64.go /usr/local/go/src/syscall/ztypes_linux_amd64.go

cd /usr/local/go/src/syscall

/usr/local/go/pkg/tool/linux_amd64/asm -p syscall -trimpath "$WORK/b039=>" -I $WORK/b039/ -I /usr/local/go/pkg/include -D GOOS_linux -D GOARCH_amd64 -D GOAMD64_v1 -o $WORK/b039/asm_linux_amd64.o ./asm_linux_amd64.s

/usr/local/go/pkg/tool/linux_amd64/pack r $WORK/b039/_pkg_.a $WORK/b039/asm_linux_amd64.o # internal

/usr/local/go/pkg/tool/linux_amd64/buildid -w $WORK/b039/_pkg_.a # internal

cp $WORK/b039/_pkg_.a /root/.cache/go-build/8f/8f8700d5d961d718530eeabf61a99c0d55482887eeb21ee5ff5855fd5ddf0b1a-d # internal

##############################################################################

internal/syscall/unix
mkdir -p $WORK/b038/

cat >/tmp/go-build2485170101/b038/importcfg << 'EOF' # internal
# import config
packagefile sync/atomic=/tmp/go-build2485170101/b030/_pkg_.a
packagefile syscall=/tmp/go-build2485170101/b039/_pkg_.a
EOF

cd /go-std

/usr/local/go/pkg/tool/linux_amd64/compile -o $WORK/b038/_pkg_.a -trimpath "$WORK/b038=>" -p internal/syscall/unix -std -complete -buildid sU_TG0ZgDA70TbqgadHA/sU_TG0ZgDA70TbqgadHA -goversion go1.22.4 -c=16 -nolocalimports -importcfg $WORK/b038/importcfg -pack /usr/local/go/src/internal/syscall/unix/at.go /usr/local/go/src/internal/syscall/unix/at_fstatat.go /usr/local/go/src/internal/syscall/unix/at_sysnum_linux.go /usr/local/go/src/internal/syscall/unix/at_sysnum_newfstatat_linux.go /usr/local/go/src/internal/syscall/unix/constants.go /usr/local/go/src/internal/syscall/unix/copy_file_range_linux.go /usr/local/go/src/internal/syscall/unix/eaccess_linux.go /usr/local/go/src/internal/syscall/unix/fcntl_unix.go /usr/local/go/src/internal/syscall/unix/getrandom.go /usr/local/go/src/internal/syscall/unix/getrandom_linux.go /usr/local/go/src/internal/syscall/unix/kernel_version_linux.go /usr/local/go/src/internal/syscall/unix/net.go /usr/local/go/src/internal/syscall/unix/nonblocking_unix.go /usr/local/go/src/internal/syscall/unix/pidfd_linux.go /usr/local/go/src/internal/syscall/unix/sysnum_linux_amd64.go

/usr/local/go/pkg/tool/linux_amd64/buildid -w $WORK/b038/_pkg_.a # internal

cp $WORK/b038/_pkg_.a /root/.cache/go-build/2f/2f63a3c6c158e0f4e033dedf8ad87374672c0d04d9067b38432b84c69e71d2a5-d # internal

##############################################################################

time
mkdir -p $WORK/b041/

cat >/tmp/go-build2485170101/b041/importcfg << 'EOF' # internal
# import config
packagefile errors=/tmp/go-build2485170101/b003/_pkg_.a
packagefile runtime=/tmp/go-build2485170101/b009/_pkg_.a
packagefile sync=/tmp/go-build2485170101/b028/_pkg_.a
packagefile syscall=/tmp/go-build2485170101/b039/_pkg_.a
EOF

/usr/local/go/pkg/tool/linux_amd64/compile -o $WORK/b041/_pkg_.a -trimpath "$WORK/b041=>" -p time -std -buildid ZwIL7_SOSCwk7r-Yf8Mm/ZwIL7_SOSCwk7r-Yf8Mm -goversion go1.22.4 -c=16 -nolocalimports -importcfg $WORK/b041/importcfg -pack /usr/local/go/src/time/format.go /usr/local/go/src/time/format_rfc3339.go /usr/local/go/src/time/sleep.go /usr/local/go/src/time/sys_unix.go /usr/local/go/src/time/tick.go /usr/local/go/src/time/time.go /usr/local/go/src/time/zoneinfo.go /usr/local/go/src/time/zoneinfo_goroot.go /usr/local/go/src/time/zoneinfo_read.go /usr/local/go/src/time/zoneinfo_unix.go

/usr/local/go/pkg/tool/linux_amd64/buildid -w $WORK/b041/_pkg_.a # internal

cp $WORK/b041/_pkg_.a /root/.cache/go-build/e8/e8a55a4284b6b145bf9ff67622d89a597671883ca15710c25057a9530290ddd7-d # internal

##############################################################################

internal/poll
mkdir -p $WORK/b037/

cat >/tmp/go-build2485170101/b037/importcfg << 'EOF' # internal
# import config
packagefile errors=/tmp/go-build2485170101/b003/_pkg_.a
packagefile internal/syscall/unix=/tmp/go-build2485170101/b038/_pkg_.a
packagefile io=/tmp/go-build2485170101/b035/_pkg_.a
packagefile runtime=/tmp/go-build2485170101/b009/_pkg_.a
packagefile sync=/tmp/go-build2485170101/b028/_pkg_.a
packagefile sync/atomic=/tmp/go-build2485170101/b030/_pkg_.a
packagefile syscall=/tmp/go-build2485170101/b039/_pkg_.a
packagefile time=/tmp/go-build2485170101/b041/_pkg_.a
EOF

/usr/local/go/pkg/tool/linux_amd64/compile -o $WORK/b037/_pkg_.a -trimpath "$WORK/b037=>" -p internal/poll -std -buildid TU0brrBWZrQR4_0AIqDf/TU0brrBWZrQR4_0AIqDf -goversion go1.22.4 -c=16 -nolocalimports -importcfg $WORK/b037/importcfg -pack /usr/local/go/src/internal/poll/copy_file_range_linux.go /usr/local/go/src/internal/poll/errno_unix.go /usr/local/go/src/internal/poll/fd.go /usr/local/go/src/internal/poll/fd_fsync_posix.go /usr/local/go/src/internal/poll/fd_mutex.go /usr/local/go/src/internal/poll/fd_poll_runtime.go /usr/local/go/src/internal/poll/fd_posix.go /usr/local/go/src/internal/poll/fd_unix.go /usr/local/go/src/internal/poll/fd_unixjs.go /usr/local/go/src/internal/poll/fd_writev_unix.go /usr/local/go/src/internal/poll/hook_cloexec.go /usr/local/go/src/internal/poll/hook_unix.go /usr/local/go/src/internal/poll/iovec_unix.go /usr/local/go/src/internal/poll/sendfile_linux.go /usr/local/go/src/internal/poll/sock_cloexec.go /usr/local/go/src/internal/poll/sockopt.go /usr/local/go/src/internal/poll/sockopt_linux.go /usr/local/go/src/internal/poll/sockopt_unix.go /usr/local/go/src/internal/poll/sockoptip.go /usr/local/go/src/internal/poll/splice_linux.go /usr/local/go/src/internal/poll/writev.go

/usr/local/go/pkg/tool/linux_amd64/buildid -w $WORK/b037/_pkg_.a # internal

cp $WORK/b037/_pkg_.a /root/.cache/go-build/7a/7a1ab9f191b86cfc6784c180fe9a711625a5d5a277775abfb63b4af4fe29799b-d # internal

##############################################################################

internal/safefilepath
mkdir -p $WORK/b042/

cat >/tmp/go-build2485170101/b042/importcfg << 'EOF' # internal
# import config
packagefile errors=/tmp/go-build2485170101/b003/_pkg_.a
packagefile runtime=/tmp/go-build2485170101/b009/_pkg_.a
EOF

/usr/local/go/pkg/tool/linux_amd64/compile -o $WORK/b042/_pkg_.a -trimpath "$WORK/b042=>" -p internal/safefilepath -std -complete -buildid LGFSl5cyK0fJnNBap_Uj/LGFSl5cyK0fJnNBap_Uj -goversion go1.22.4 -c=16 -nolocalimports -importcfg $WORK/b042/importcfg -pack /usr/local/go/src/internal/safefilepath/path.go /usr/local/go/src/internal/safefilepath/path_other.go

/usr/local/go/pkg/tool/linux_amd64/buildid -w $WORK/b042/_pkg_.a # internal

cp $WORK/b042/_pkg_.a /root/.cache/go-build/f1/f110717f3572649e91407da9dd4acb4c4d55e0c2c75e3db6ca31ee5375351d56-d # internal

##############################################################################

internal/syscall/execenv
mkdir -p $WORK/b043/

cat >/tmp/go-build2485170101/b043/importcfg << 'EOF' # internal
# import config
packagefile syscall=/tmp/go-build2485170101/b039/_pkg_.a
EOF

/usr/local/go/pkg/tool/linux_amd64/compile -o $WORK/b043/_pkg_.a -trimpath "$WORK/b043=>" -p internal/syscall/execenv -std -complete -buildid 4oeQwe6JF9uSGqlg3t55/4oeQwe6JF9uSGqlg3t55 -goversion go1.22.4 -c=16 -nolocalimports -importcfg $WORK/b043/importcfg -pack /usr/local/go/src/internal/syscall/execenv/execenv_default.go

/usr/local/go/pkg/tool/linux_amd64/buildid -w $WORK/b043/_pkg_.a # internal

cp $WORK/b043/_pkg_.a /root/.cache/go-build/bb/bbc06c82d6f1a4c5518180ee2e251f30e5217e7c7f340b8621240a2d61cb1178-d # internal

##############################################################################

internal/testlog
mkdir -p $WORK/b044/

cat >/tmp/go-build2485170101/b044/importcfg << 'EOF' # internal
# import config
packagefile sync=/tmp/go-build2485170101/b028/_pkg_.a
packagefile sync/atomic=/tmp/go-build2485170101/b030/_pkg_.a
EOF

/usr/local/go/pkg/tool/linux_amd64/compile -o $WORK/b044/_pkg_.a -trimpath "$WORK/b044=>" -p internal/testlog -std -complete -buildid -yo5Tx3qV0vf_KfkwApo/-yo5Tx3qV0vf_KfkwApo -goversion go1.22.4 -c=16 -nolocalimports -importcfg $WORK/b044/importcfg -pack /usr/local/go/src/internal/testlog/exit.go /usr/local/go/src/internal/testlog/log.go

/usr/local/go/pkg/tool/linux_amd64/buildid -w $WORK/b044/_pkg_.a # internal

cp $WORK/b044/_pkg_.a /root/.cache/go-build/12/12758729bcc51d140a1cee29eaa5d9fedb13e3f4cd3767665f2c4c88aece8cc9-d # internal

##############################################################################

path
mkdir -p $WORK/b046/

cat >/tmp/go-build2485170101/b046/importcfg << 'EOF' # internal
# import config
packagefile errors=/tmp/go-build2485170101/b003/_pkg_.a
packagefile internal/bytealg=/tmp/go-build2485170101/b010/_pkg_.a
packagefile unicode/utf8=/tmp/go-build2485170101/b027/_pkg_.a
EOF

/usr/local/go/pkg/tool/linux_amd64/compile -o $WORK/b046/_pkg_.a -trimpath "$WORK/b046=>" -p path -std -complete -buildid -LmcYDWNcA34E8LTSUwJ/-LmcYDWNcA34E8LTSUwJ -goversion go1.22.4 -c=16 -nolocalimports -importcfg $WORK/b046/importcfg -pack /usr/local/go/src/path/match.go /usr/local/go/src/path/path.go

/usr/local/go/pkg/tool/linux_amd64/buildid -w $WORK/b046/_pkg_.a # internal

cp $WORK/b046/_pkg_.a /root/.cache/go-build/f5/f57eef3f5cabeb1bf9583d89e93598f76f088d54f23a02332bcd711ffdbbdad9-d # internal

##############################################################################

io/fs
mkdir -p $WORK/b045/

cat >/tmp/go-build2485170101/b045/importcfg << 'EOF' # internal
# import config
packagefile errors=/tmp/go-build2485170101/b003/_pkg_.a
packagefile internal/oserror=/tmp/go-build2485170101/b040/_pkg_.a
packagefile io=/tmp/go-build2485170101/b035/_pkg_.a
packagefile path=/tmp/go-build2485170101/b046/_pkg_.a
packagefile sort=/tmp/go-build2485170101/b032/_pkg_.a
packagefile time=/tmp/go-build2485170101/b041/_pkg_.a
packagefile unicode/utf8=/tmp/go-build2485170101/b027/_pkg_.a
EOF

/usr/local/go/pkg/tool/linux_amd64/compile -o $WORK/b045/_pkg_.a -trimpath "$WORK/b045=>" -p io/fs -std -complete -buildid zDRVb-8M60hExKu10KPA/zDRVb-8M60hExKu10KPA -goversion go1.22.4 -c=16 -nolocalimports -importcfg $WORK/b045/importcfg -pack /usr/local/go/src/io/fs/format.go /usr/local/go/src/io/fs/fs.go /usr/local/go/src/io/fs/glob.go /usr/local/go/src/io/fs/readdir.go /usr/local/go/src/io/fs/readfile.go /usr/local/go/src/io/fs/stat.go /usr/local/go/src/io/fs/sub.go /usr/local/go/src/io/fs/walk.go

/usr/local/go/pkg/tool/linux_amd64/buildid -w $WORK/b045/_pkg_.a # internal

cp $WORK/b045/_pkg_.a /root/.cache/go-build/88/881822189da94d39688967b48ba9accee04958e8a7adb194fbcaed7dba054396-d # internal

##############################################################################

os
mkdir -p $WORK/b036/

cat >/tmp/go-build2485170101/b036/importcfg << 'EOF' # internal
# import config
packagefile errors=/tmp/go-build2485170101/b003/_pkg_.a
packagefile internal/bytealg=/tmp/go-build2485170101/b010/_pkg_.a
packagefile internal/itoa=/tmp/go-build2485170101/b023/_pkg_.a
packagefile internal/poll=/tmp/go-build2485170101/b037/_pkg_.a
packagefile internal/safefilepath=/tmp/go-build2485170101/b042/_pkg_.a
packagefile internal/syscall/execenv=/tmp/go-build2485170101/b043/_pkg_.a
packagefile internal/syscall/unix=/tmp/go-build2485170101/b038/_pkg_.a
packagefile internal/testlog=/tmp/go-build2485170101/b044/_pkg_.a
packagefile io=/tmp/go-build2485170101/b035/_pkg_.a
packagefile io/fs=/tmp/go-build2485170101/b045/_pkg_.a
packagefile runtime=/tmp/go-build2485170101/b009/_pkg_.a
packagefile sort=/tmp/go-build2485170101/b032/_pkg_.a
packagefile sync=/tmp/go-build2485170101/b028/_pkg_.a
packagefile sync/atomic=/tmp/go-build2485170101/b030/_pkg_.a
packagefile syscall=/tmp/go-build2485170101/b039/_pkg_.a
packagefile time=/tmp/go-build2485170101/b041/_pkg_.a
EOF

/usr/local/go/pkg/tool/linux_amd64/compile -o $WORK/b036/_pkg_.a -trimpath "$WORK/b036=>" -p os -std -buildid I4gSFGoAGo3qAq6NNxml/I4gSFGoAGo3qAq6NNxml -goversion go1.22.4 -c=16 -nolocalimports -importcfg $WORK/b036/importcfg -pack /usr/local/go/src/os/dir.go /usr/local/go/src/os/dir_unix.go /usr/local/go/src/os/dirent_linux.go /usr/local/go/src/os/endian_little.go /usr/local/go/src/os/env.go /usr/local/go/src/os/error.go /usr/local/go/src/os/error_errno.go /usr/local/go/src/os/error_posix.go /usr/local/go/src/os/exec.go /usr/local/go/src/os/exec_posix.go /usr/local/go/src/os/exec_unix.go /usr/local/go/src/os/executable.go /usr/local/go/src/os/executable_procfs.go /usr/local/go/src/os/file.go /usr/local/go/src/os/file_open_unix.go /usr/local/go/src/os/file_posix.go /usr/local/go/src/os/file_unix.go /usr/local/go/src/os/getwd.go /usr/local/go/src/os/path.go /usr/local/go/src/os/path_unix.go /usr/local/go/src/os/pipe2_unix.go /usr/local/go/src/os/proc.go /usr/local/go/src/os/rawconn.go /usr/local/go/src/os/removeall_at.go /usr/local/go/src/os/stat.go /usr/local/go/src/os/stat_linux.go /usr/local/go/src/os/stat_unix.go /usr/local/go/src/os/sticky_notbsd.go /usr/local/go/src/os/sys.go /usr/local/go/src/os/sys_linux.go /usr/local/go/src/os/sys_unix.go /usr/local/go/src/os/tempfile.go /usr/local/go/src/os/types.go /usr/local/go/src/os/types_unix.go /usr/local/go/src/os/wait_waitid.go /usr/local/go/src/os/zero_copy_linux.go

/usr/local/go/pkg/tool/linux_amd64/buildid -w $WORK/b036/_pkg_.a # internal

cp $WORK/b036/_pkg_.a /root/.cache/go-build/68/685fc8fa062b6a6c984811f0c557ed952a69e12f0af051e2ae0b3b42d6efaa2d-d # internal

##############################################################################

fmt
mkdir -p $WORK/b002/

cat >/tmp/go-build2485170101/b002/importcfg << 'EOF' # internal
# import config
packagefile errors=/tmp/go-build2485170101/b003/_pkg_.a
packagefile internal/fmtsort=/tmp/go-build2485170101/b021/_pkg_.a
packagefile io=/tmp/go-build2485170101/b035/_pkg_.a
packagefile math=/tmp/go-build2485170101/b024/_pkg_.a
packagefile os=/tmp/go-build2485170101/b036/_pkg_.a
packagefile reflect=/tmp/go-build2485170101/b022/_pkg_.a
packagefile sort=/tmp/go-build2485170101/b032/_pkg_.a
packagefile strconv=/tmp/go-build2485170101/b026/_pkg_.a
packagefile sync=/tmp/go-build2485170101/b028/_pkg_.a
packagefile unicode/utf8=/tmp/go-build2485170101/b027/_pkg_.a
EOF

/usr/local/go/pkg/tool/linux_amd64/compile -o $WORK/b002/_pkg_.a -trimpath "$WORK/b002=>" -p fmt -std -complete -buildid _MevO0BbyQUFe2sBiyq9/_MevO0BbyQUFe2sBiyq9 -goversion go1.22.4 -c=16 -nolocalimports -importcfg $WORK/b002/importcfg -pack /usr/local/go/src/fmt/doc.go /usr/local/go/src/fmt/errors.go /usr/local/go/src/fmt/format.go /usr/local/go/src/fmt/print.go /usr/local/go/src/fmt/scan.go

/usr/local/go/pkg/tool/linux_amd64/buildid -w $WORK/b002/_pkg_.a # internal

cp $WORK/b002/_pkg_.a /root/.cache/go-build/57/57257950e0d3af3b43df5adac17bc79e21c9333e64baae1923915e0f922deb6f-d # internal

##############################################################################

command-line-arguments
mkdir -p $WORK/b001/

cat >/tmp/go-build2485170101/b001/importcfg << 'EOF' # internal
# import config
packagefile fmt=/tmp/go-build2485170101/b002/_pkg_.a
packagefile sync=/tmp/go-build2485170101/b028/_pkg_.a
packagefile time=/tmp/go-build2485170101/b041/_pkg_.a
packagefile runtime=/tmp/go-build2485170101/b009/_pkg_.a
EOF

/usr/local/go/pkg/tool/linux_amd64/compile \
-o $WORK/b001/_pkg_.a \
-trimpath "$WORK/b001=>" \
-p main -complete \
-buildid 6DQl8Z0_8qY-JAGGHfrP/6DQl8Z0_8qY-JAGGHfrP \
-goversion go1.22.4 \
-c=16 -nolocalimports \
-importcfg $WORK/b001/importcfg \
-pack ./context/v1-example-waitgroup/main.go

/usr/local/go/pkg/tool/linux_amd64/buildid -w $WORK/b001/_pkg_.a # internal

cp $WORK/b001/_pkg_.a /root/.cache/go-build/05/0524193eb9c39f31d798aa08a24a49f716507e77f76a7f713c798d104885a893-d # internal

##############################################################################

cat >/tmp/go-build2485170101/b001/importcfg.link << 'EOF' # internal
packagefile command-line-arguments=/tmp/go-build2485170101/b001/_pkg_.a
packagefile fmt=/tmp/go-build2485170101/b002/_pkg_.a
packagefile sync=/tmp/go-build2485170101/b028/_pkg_.a
packagefile time=/tmp/go-build2485170101/b041/_pkg_.a
packagefile runtime=/tmp/go-build2485170101/b009/_pkg_.a
packagefile errors=/tmp/go-build2485170101/b003/_pkg_.a
packagefile internal/fmtsort=/tmp/go-build2485170101/b021/_pkg_.a
packagefile io=/tmp/go-build2485170101/b035/_pkg_.a
packagefile math=/tmp/go-build2485170101/b024/_pkg_.a
packagefile os=/tmp/go-build2485170101/b036/_pkg_.a
packagefile reflect=/tmp/go-build2485170101/b022/_pkg_.a
packagefile sort=/tmp/go-build2485170101/b032/_pkg_.a
packagefile strconv=/tmp/go-build2485170101/b026/_pkg_.a
packagefile unicode/utf8=/tmp/go-build2485170101/b027/_pkg_.a
packagefile internal/race=/tmp/go-build2485170101/b029/_pkg_.a
packagefile sync/atomic=/tmp/go-build2485170101/b030/_pkg_.a
packagefile syscall=/tmp/go-build2485170101/b039/_pkg_.a
packagefile internal/abi=/tmp/go-build2485170101/b005/_pkg_.a
packagefile internal/bytealg=/tmp/go-build2485170101/b010/_pkg_.a
packagefile internal/chacha8rand=/tmp/go-build2485170101/b012/_pkg_.a
packagefile internal/coverage/rtcov=/tmp/go-build2485170101/b013/_pkg_.a
packagefile internal/cpu=/tmp/go-build2485170101/b011/_pkg_.a
packagefile internal/goarch=/tmp/go-build2485170101/b006/_pkg_.a
packagefile internal/godebugs=/tmp/go-build2485170101/b014/_pkg_.a
packagefile internal/goexperiment=/tmp/go-build2485170101/b015/_pkg_.a
packagefile internal/goos=/tmp/go-build2485170101/b016/_pkg_.a
packagefile runtime/internal/atomic=/tmp/go-build2485170101/b017/_pkg_.a
packagefile runtime/internal/math=/tmp/go-build2485170101/b018/_pkg_.a
packagefile runtime/internal/sys=/tmp/go-build2485170101/b019/_pkg_.a
packagefile runtime/internal/syscall=/tmp/go-build2485170101/b020/_pkg_.a
packagefile internal/reflectlite=/tmp/go-build2485170101/b004/_pkg_.a
packagefile math/bits=/tmp/go-build2485170101/b025/_pkg_.a
packagefile internal/itoa=/tmp/go-build2485170101/b023/_pkg_.a
packagefile internal/poll=/tmp/go-build2485170101/b037/_pkg_.a
packagefile internal/safefilepath=/tmp/go-build2485170101/b042/_pkg_.a
packagefile internal/syscall/execenv=/tmp/go-build2485170101/b043/_pkg_.a
packagefile internal/syscall/unix=/tmp/go-build2485170101/b038/_pkg_.a
packagefile internal/testlog=/tmp/go-build2485170101/b044/_pkg_.a
packagefile io/fs=/tmp/go-build2485170101/b045/_pkg_.a
packagefile internal/unsafeheader=/tmp/go-build2485170101/b008/_pkg_.a
packagefile unicode=/tmp/go-build2485170101/b031/_pkg_.a
packagefile slices=/tmp/go-build2485170101/b033/_pkg_.a
packagefile internal/oserror=/tmp/go-build2485170101/b040/_pkg_.a
packagefile path=/tmp/go-build2485170101/b046/_pkg_.a
packagefile cmp=/tmp/go-build2485170101/b034/_pkg_.a
modinfo "0w\xaf\f\x92t\b\x02A\xe1\xc1\a\xe6\xd6\x18\xe6path\tcommand-line-arguments\nbuild\t-buildmode=exe\nbuild\t-compiler=gc\nbuild\tCGO_ENABLED=1\nbuild\tCGO_CFLAGS=\nbuild\tCGO_CPPFLAGS=\nbuild\tCGO_CXXFLAGS=\nbuild\tCGO_LDFLAGS=\nbuild\tGOARCH=amd64\nbuild\tGOOS=linux\nbuild\tGOAMD64=v1\n\xf92C1\x86\x18 r\x00\x82B\x10A\x16\xd8\xf2"
EOF

mkdir -p $WORK/b001/exe/

cd .

/usr/local/go/pkg/tool/linux_amd64/link \
-o $WORK/b001/exe/a.out \
-importcfg $WORK/b001/importcfg.link \
-buildmode=exe \
-buildid=tVrWC8HJQKygn2OGXbhA/6DQl8Z0_8qY-JAGGHfrP/mu2RaLmPj5PDGD_ZnxlK/tVrWC8HJQKygn2OGXbhA \
-extld=gcc \
$WORK/b001/_pkg_.a

/usr/local/go/pkg/tool/linux_amd64/buildid -w $WORK/b001/exe/a.out # internal

cp $WORK/b001/exe/a.out main

$

```


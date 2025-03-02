ANDROID_NDK_BASE="$ANDROID_HOME/ndk/27.1.12297006/toolchains/llvm/prebuilt/linux-x86_64/bin"
ANDROID_SDK_TARGET="21"

echo "Building"
CC="$ANDROID_NDK_BASE/aarch64-linux-android$ANDROID_SDK_TARGET-clang" CXX="$ANDROID_NDK_BASE/aarch64-linux-android$ANDROID_SDK_TARGET-clang++" CFLAGS="-static -fPIC" CXXFLAGS="-static" LDFLAGS="-pie -fuse-ld=bfd" GOOS=android GOARCH=arm64 CGO_ENABLED=1 go build cmd/main.go

if [ $? -eq 0 ]; then
    echo "Built"
    python tools/adb.py push ./main /data/data/com.veluga.fm.echo/main -m 777
else
    echo "🟥🟥🟥 Error"
fi



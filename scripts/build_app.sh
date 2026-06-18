#!/bin/bash
set -e

APP_NAME="InterviewPro"
BUNDLE_ID="com.chenyongzhi.interviewpro"
VERSION="1.0.0"
BINARY="$APP_NAME"
APP_DIR="$APP_NAME.app"

cd "$(dirname "$0")/.."

echo "▶ 编译 Go 二进制..."
CGO_ENABLED=1 go build \
  -ldflags="-s -w -X main.version=$VERSION" \
  -o "$BINARY" .

echo "▶ 创建 .app 目录结构..."
rm -rf "$APP_DIR"
mkdir -p "$APP_DIR/Contents/MacOS"
mkdir -p "$APP_DIR/Contents/Resources"

echo "▶ 复制二进制文件..."
cp "$BINARY" "$APP_DIR/Contents/MacOS/$APP_NAME"
chmod +x "$APP_DIR/Contents/MacOS/$APP_NAME"

echo "▶ 生成 Info.plist..."
cat > "$APP_DIR/Contents/Info.plist" << PLIST
<?xml version="1.0" encoding="UTF-8"?>
<!DOCTYPE plist PUBLIC "-//Apple//DTD PLIST 1.0//EN" "http://www.apple.com/DTDs/PropertyList-1.0.dtd">
<plist version="1.0">
<dict>
    <key>CFBundleExecutable</key>
    <string>$APP_NAME</string>
    <key>CFBundleIdentifier</key>
    <string>$BUNDLE_ID</string>
    <key>CFBundleName</key>
    <string>$APP_NAME</string>
    <key>CFBundleDisplayName</key>
    <string>大厂面试官 Agent</string>
    <key>CFBundleVersion</key>
    <string>$VERSION</string>
    <key>CFBundleShortVersionString</key>
    <string>$VERSION</string>
    <key>CFBundlePackageType</key>
    <string>APPL</string>
    <key>CFBundleIconFile</key>
    <string>AppIcon</string>
    <key>LSUIElement</key>
    <true/>
    <key>NSHighResolutionCapable</key>
    <true/>
    <key>NSHumanReadableCopyright</key>
    <string>Copyright © 2025 chenyongzhi1119. MIT License.</string>
</dict>
</plist>
PLIST

echo "▶ 生成应用图标..."
if [ -f "assets/app_icon_512.png" ]; then
  ICONSET_DIR="$APP_DIR/Contents/Resources/AppIcon.iconset"
  mkdir -p "$ICONSET_DIR"
  # 生成各尺寸图标（macOS 要求）
  for size in 16 32 64 128 256 512; do
    sips -z $size $size assets/app_icon_512.png \
      --out "$ICONSET_DIR/icon_${size}x${size}.png" > /dev/null 2>&1
    double=$((size * 2))
    if [ $double -le 512 ]; then
      sips -z $double $double assets/app_icon_512.png \
        --out "$ICONSET_DIR/icon_${size}x${size}@2x.png" > /dev/null 2>&1
    fi
  done
  iconutil -c icns "$ICONSET_DIR" -o "$APP_DIR/Contents/Resources/AppIcon.icns"
  rm -rf "$ICONSET_DIR"
  echo "   图标生成完成"
else
  echo "   ⚠ 未找到 assets/app_icon_512.png，跳过图标"
fi

echo "▶ 清理临时文件..."
rm -f "$BINARY"

echo ""
echo "✅ 打包完成：$APP_DIR"
echo ""
echo "使用方式："
echo "  双击运行：open $APP_DIR"
echo "  设置 API Key 后使用（或在页面「设置」中填写）"
echo ""
echo "设置 API Key 并启动（可选）："
echo "  DEEPSEEK_API_KEY=sk-xxx open $APP_DIR"

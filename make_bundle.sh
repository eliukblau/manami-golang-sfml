#!/bin/sh

echo "\nLimpiando Directorios..."
find . -name ".DS_Store" -depth -exec rm -f {} \;
rm -rf build
rm -rf dist

echo "\nGenerando el binario..."
go clean
go build -o manami

echo "\nCreando App Bundle (py2app)..."
touch dummy.py
python3 setup.py py2app --semi-standalone
rm -f dummy.py

echo "\nAdaptando el App Bundle..."
rm -rf dist/Manami.app/Contents/MacOS/**
rm -rf dist/Manami.app/Contents/Resources/**

echo "\nNormalizando los Frameworks y DyLibs..."
rm -rf dist/Manami.app/Contents/Frameworks/Python.framework
rm -rf dist/Manami.app/Contents/Frameworks/libcrypto.1.0.0.dylib
rm -rf dist/Manami.app/Contents/Frameworks/libssl.1.0.0.dylib
rm -rf dist/Manami.app/Contents/Frameworks/liblzma.5.dylib

echo "\nCorrigiendo estructura..."
cp -v ./manami dist/Manami.app/Contents/MacOS/Manami
cp -v bundle_res/Info.plist dist/Manami.app/Contents/
cp -v bundle_res/app.icns dist/Manami.app/Contents/Resources/

echo "\nCorrigiendo enlaces del binario..."
install_name_tool -change \
  @rpath/SDL2.framework/Versions/A/SDL2 \
  @executable_path/../Frameworks/SDL2.framework/Versions/A/SDL2 \
  dist/Manami.app/Contents/MacOS/Manami

install_name_tool -change \
  @rpath/SDL2_image.framework/Versions/A/SDL2_image \
  @executable_path/../Frameworks/SDL2_image.framework/Versions/A/SDL2_image \
  dist/Manami.app/Contents/MacOS/Manami

install_name_tool -change \
  @rpath/SDL2_mixer.framework/Versions/A/SDL2_mixer \
  @executable_path/../Frameworks/SDL2_mixer.framework/Versions/A/SDL2_mixer \
  dist/Manami.app/Contents/MacOS/Manami

install_name_tool -change \
  @rpath/SDL2_ttf.framework/Versions/A/SDL2_ttf \
  @executable_path/../Frameworks/SDL2_ttf.framework/Versions/A/SDL2_ttf \
  dist/Manami.app/Contents/MacOS/Manami

echo "\nCopiando recursos..."
cp -rv gfx dist/Manami.app/Contents/Resources/
cp -rv sfx dist/Manami.app/Contents/Resources/

echo "\nLimpiando resultados..."
rm -rf build

echo "\n*** App Bundle creado con exito! :) ***\n"

#!/bin/sh

echo "\nLimpiando Directorios..."
find . -name ".DS_Store" -depth -exec rm -f {} \;
rm -rf build
rm -rf dist

echo "\nGenerando el binario..."
cd src
go clean
go build -tags appbundle -o Manami
cd ..

echo "\nCreando App Bundle (py2app)..."
touch dummy.py
python3 setup.py py2app --semi-standalone &> /dev/null
rm -f dummy.py

echo "\nAdaptando el App Bundle..."
rm -rf dist/Manami.app/Contents/MacOS/**
rm -rf dist/Manami.app/Contents/Resources/**

echo "\nNormalizando los Frameworks y dylibs..."
rm -rf dist/Manami.app/Contents/Frameworks/Python.framework
rm -rf dist/Manami.app/Contents/Frameworks/libcrypto.1.0.0.dylib
rm -rf dist/Manami.app/Contents/Frameworks/libssl.1.0.0.dylib
rm -rf dist/Manami.app/Contents/Frameworks/liblzma.5.dylib

echo "\nCorrigiendo estructura..."
mv -v src/Manami dist/Manami.app/Contents/MacOS/Manami
cp -v bundle_res/Info.plist dist/Manami.app/Contents/
cp -v bundle_res/App.icns dist/Manami.app/Contents/Resources/

echo "\nCorrigiendo Frameworks y dylibs del binario..."
install_name_tool -change \
  /usr/local/opt/csfml/lib/libcsfml-window.2.3.dylib \
  @executable_path/../Frameworks/libcsfml-window.2.3.dylib \
  dist/Manami.app/Contents/MacOS/Manami

install_name_tool -change \
  /usr/local/opt/csfml/lib/libcsfml-graphics.2.3.dylib \
  @executable_path/../Frameworks/libcsfml-graphics.2.3.dylib \
  dist/Manami.app/Contents/MacOS/Manami

install_name_tool -change \
  /usr/local/opt/csfml/lib/libcsfml-audio.2.3.dylib \
  @executable_path/../Frameworks/libcsfml-audio.2.3.dylib \
  dist/Manami.app/Contents/MacOS/Manami

install_name_tool -change \
  /usr/local/opt/csfml/lib/libcsfml-system.2.3.dylib \
  @executable_path/../Frameworks/libcsfml-system.2.3.dylib \
  dist/Manami.app/Contents/MacOS/Manami

echo "\nCopiando recursos..."
cp -rv src/gfx dist/Manami.app/Contents/Resources/
cp -rv src/sfx dist/Manami.app/Contents/Resources/

echo "\nLimpiando resultados..."
rm -rf build

echo "\n*** App Bundle creado con exito! :) ***\n"

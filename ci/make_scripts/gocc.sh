#!/bin/bash

SCRIPT_DIR=`cd -- "$( dirname -- "${BASH_SOURCE[0]}" )" &> /dev/null && pwd`
CI_DIR=`dirname $SCRIPT_DIR`
PROJECT_DIR=`dirname $CI_DIR`
PKG_DIR=$PROJECT_DIR/server/interpreter

rm -rf $PKG_DIR/gocc
gocc -p $PKG_DIR/gocc -o $PKG_DIR/gocc $PKG_DIR/lang.bnf
gofumpt -w $PKG_DIR/gocc/parser/

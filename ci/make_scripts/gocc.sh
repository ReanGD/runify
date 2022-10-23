#!/bin/bash

SCRIPT_DIR=`cd -- "$( dirname -- "${BASH_SOURCE[0]}" )" &> /dev/null && pwd`
CI_DIR=`dirname $SCRIPT_DIR`
PROJECT_DIR=`dirname $CI_DIR`
CALC_PKG_DIR=$PROJECT_DIR/server/provider/calculator

rm -rf $CALC_PKG_DIR/gocc
gocc -p github.com/ReanGD/runify/server/calc/gocc -o $CALC_PKG_DIR/gocc $CALC_PKG_DIR/calc.bnf

#!/bin/bash
set -e

OLDGOPATH=$GOPATH

basepath=$(cd `dirname $0`; pwd)
cd $basepath

export GOPATH=`pwd`
echo 'GOPATH IS:' $GOPATH

if [ $1 == "install" ]; then
    echo 'go install ' $2
    go install servers/$2
elif [ $1 == "run" ]; then
    echo 'go run'
    go install servers/$2
    bin/$2
fi

export GOPATH=$OLDGOPATH
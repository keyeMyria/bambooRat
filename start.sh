#!/bin/bash
set -e

MYLODER=$(cd `dirname ${0}`; pwd)
# 先设置环境变量
source ${MYLODER}/tools/goenv.sh

GOFOLDER="${MYLODER}/go"
GOPROJECTS="${MYLODER}/goprojects"

if [ -d ${GOFOLDER} ];
then
    echo "locate the same GOROOT in ${GOROOT}"
    if [ -d ${GOPROJECTS} ];
    then
        echo "locate the same GOPATH in ${GOPATH}"
        exit 1
    else
        rm -rf ${GOPATH}
        echo "Clean old GOPATH ...Done"
    fi
else
    rm -rf ${GOROOT}
    echo "Clean old GOROOT ...Done"
fi
echo "configure go env done !"

set -e
# 安装go pkg
source ${MYLODER}/tools/goinstall.sh

# 设置环境变量
echo "update ~/.bash_profile ..."
PFILE="$HOME/.bash_profile"
touch "${PFILE}"

set +e
sed -i ".back" "/.*GOHOME.*/d" ${PFILE}
sed -i ".back" "/.*GOROOT.*/d" ${PFILE}
sed -i ".back" "/.*GOPATH.*/d" ${PFILE}
sed -i ".back" "/.*GOARCH.*/d" ${PFILE}
sed -i ".back" "/.*GOOS.*/d"   ${PFILE}
set -e

echo "export GOHOME=${GOHOME}"    >> ${PFILE}
echo "export GOROOT=${GOROOT}"    >> ${PFILE}
echo "export GOPATH=${GOPATH}"    >> ${PFILE}
echo "export PATH=\${GOROOT}/bin:\${GOPATH}/bin:\${PATH}" >> ${PFILE}
source $PFILE
echo "环境变量发生变化，你需要重新登录才可能生效"
echo "configure go env done !"


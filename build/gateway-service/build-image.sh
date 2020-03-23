#!/bin/bash
#****************************************************************#
# Create Date: 2019-02-02 22:16
#********************************* ******************************#
ROOTDIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"

NAME="${1}"
[ -z "${1}" ] && echo "component is empty" && exit 1

NS="knative-sample"
GIT_COMMIT="$(git rev-parse --verify HEAD)"
GIT_BRANCH=`git branch | grep \* | cut -d ' ' -f2`
TAG="$(date +%Y''%m''%d''%H''%M''%S)"
if [ ! -z "${GIT_COMMIT}" -a "${GIT_BRANCH}" != " " ]; then
  TAG="${GIT_BRANCH}_${GIT_COMMIT:0:8}-$(date +%Y''%m''%d''%H''%M''%S)"
fi

docker build -t "${NAME}:${TAG}" -f ${ROOTDIR}/Dockerfile ${ROOTDIR}/

array=( registry.cn-hangzhou.aliyuncs.com )
for registry in "${array[@]}"
do
    echo "push images to ${registry}/${NS}/${NAME}:${TAG}"
    docker tag "${NAME}:${TAG}" "${registry}/${NS}/${NAME}:${TAG}"
    docker push "${registry}/${NS}/${NAME}:${TAG}"

    docker tag "${NAME}:${TAG}" "${registry}/${NS}/${NAME}:latest"
    docker push "${registry}/${NS}/${NAME}:latest"
done

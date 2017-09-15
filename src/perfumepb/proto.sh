#!/usr/bin/env bash

if [ "$ROOT_DIR" = "" ];then
    ROOT_DIR=$PWD
fi

CWD=${ROOT_DIR}/src/perfumepb
GOGO_PATH=${ROOT_DIR}/vendor/src/github.com/gogo/protobuf/gogoproto
OTSIMOPB_PATH=${ROOT_DIR}/vendor/src/github.com/otsimo/otsimopb
COMMON_PATH=${ROOT_DIR}/src

export IMPORT_PATH=${CWD}:${ROOT_DIR}/vendor/src:${GOOGLE_APIPATH}:${GOGO_PATH}:${COMMON_PATH}
export GENERATOR="gofast_out"
export OUTPUT_DIR=${CWD}
export PROTO_FILES="$CWD/*.proto"

export OPTIONS_API="Mgoogle/api/annotations.proto=google.golang.org/genproto/googleapis/api/annotations,Mgoogle/protobuf/field_mask.proto=github.com/gogo/protobuf/types"
export OPTIONS_PROTO="${OPTIONS_API}"

protoc --proto_path=${IMPORT_PATH} \
       --${GENERATOR}=${OPTIONS_PROTO},plugins=grpc:${OUTPUT_DIR} \
       ${PROTO_FILES}

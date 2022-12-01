#!/usr/bin/env bash

set -eo pipefail

# go get ./...
# go install github.com/cosmos/gogoproto/protoc-gen-gocosmos
# apk add nodejs-current 
# npm install -g swagger-combine

BASEDIR=$(dirname $0)
PROJECTDIR=$BASEDIR/..
# default is tmp folder
SOURCEDIR=$(realpath ${1:-$PROJECTDIR/tmp})
DOC_DIR=$(realpath $PROJECTDIR/doc)

COSMOS_SDK_DIR=${COSMOS_SDK_DIR:-$(go list -f "{{ .Dir }}" -m github.com/cosmos/cosmos-sdk)}
COSMOS_WASM_DIR=${COSMOS_WASM_DIR:-$(go list -f "{{ .Dir }}" -m github.com/CosmWasm/wasmd)}
IBC_DIR=${IBC_DIR:-$(go list -f "{{ .Dir }}" -m github.com/cosmos/ibc-go/v3)}

# scan all folders that contain proto file
proto_dirs=$(find $COSMOS_SDK_DIR/proto $COSMOS_SDK_DIR/third_party/proto $IBC_DIR/proto $COSMOS_WASM_DIR -path -prune -o -name '*.proto' -print0 | xargs -0 -n1 dirname | sort | uniq)

GEN_DIR=$SOURCEDIR/swagger-gen
# clean swagger files
rm -rf $GEN_DIR
mkdir -p $GEN_DIR

for dir in $proto_dirs; do
  # generate swagger files (filter query files)
  query_file=$(find "${dir}" -maxdepth 1 \( -name 'query.proto' -o -name 'service.proto' \))
  if [[ ! -z "$query_file" ]]; then   
    buf alpha protoc  \
    -I=. \
    -I="$COSMOS_WASM_DIR/proto" \
    -I="$COSMOS_SDK_DIR/third_party/proto" \
    -I="$COSMOS_SDK_DIR/proto" \
    --gocosmos_out=Mgoogle/protobuf/any.proto=github.com/cosmos/cosmos-sdk/codec/types,Mgoogle/protobuf/empty.proto=github.com/gogo/protobuf/types,plugins=interfacetype+grpc,paths=source_relative:$COSMOS_SDK_DIR \
    --swagger_out=$GEN_DIR \
    --swagger_opt=logtostderr=true,fqn_for_swagger_name=true,simple_operation_ids=true \
    $query_file
        
  fi
done

node -e "var fs = require('fs'),file='$COSMOS_SDK_DIR/client/docs/config.json',result = fs.readFileSync(file).toString().replace('./client','$COSMOS_SDK_DIR/client').replace(/.\/tmp-swagger-gen/g, '$GEN_DIR');
var baseModuleDir = '$GEN_DIR/x', obj = JSON.parse(result);

  obj.apis.push({
      url: '$GEN_DIR/cosmwasm/wasm/v1/query.swagger.json',
      operationIds: {
          rename: {
              Params: 'CosmWasmParams'
          }    
      }
  });


fs.writeFileSync('$GEN_DIR/config.json', JSON.stringify(obj, null, 2));
"


# combine swagger files
# all the individual swagger files need to be configured in `config.json` for merging
swagger-combine $GEN_DIR/config.json -o $DOC_DIR/swagger-ui/swagger.yaml -f yaml --continueOnConflictingPaths true --includeDefinitions true

#!/bin/bash
# present working dir you are running the command from
PWD=`pwd`
# script dir
if [ `uname` == "Linux" ]; then
    SCRIPT_DIR=$(dirname $(readlink -f $0))
else
    SCRIPT_DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"
fi
# python dir
PYTHON_DIR=$SCRIPT_DIR/../python
# destination dir
DST_DIR=$PYTHON_DIR/
# source directory for proto
PROTO_DIR=$SCRIPT_DIR/../proto
# pkg directory
PKG_DIR=$SCRIPT_DIR/../proto/schema/
# switch to the script dir in case of running the script ouside of this dir
cd $SCRIPT_DIR
# convert the actions.json to a actions.p (pickle), store the version from
# package.json
JSON2DICT=$(python <<EOF
try:
    import _pickle as pickle
except ImportError:
    import cPickle as pickle
from json import loads
for j in ['actions', 'events']:
    with open('../{0}.json'.format(j), 'r') as _json:
        _str = _json.read()
        _dict = loads(_str)
        with open('../script/version.json', 'r') as version_file:
            json_str = version_file.read()
            version = loads(json_str)['version']
            _dict['_version'] = version
    with open('../python/schema/{0}.p'.format(j), 'wb') as fp:
        pickle.dump(_dict, fp)
EOF
)
# build protobufs for grpc
protoc -I=$PROTO_DIR --python_out=plugins=grpc:$DST_DIR $PKG_DIR/worm.proto
# build protobufs without grpc
protoc -I $PROTO_DIR --python_out=$DST_DIR $PKG_DIR/Metadata.proto
protoc -I $PROTO_DIR --python_out=$DST_DIR $PKG_DIR/Command.proto
protoc -I=$PROTO_DIR --python_out=$DST_DIR $PKG_DIR/Event.proto
protoc -I=$PROTO_DIR --python_out=$DST_DIR $PKG_DIR/Msg.proto
protoc -I=$PROTO_DIR --python_out=$DST_DIR $PKG_DIR/Feature.proto
# switch to the python dir
cd $PYTHON_DIR
# build sdist package
python setup.py sdist
# switch back to the dir you ran the script from
cd $PWD

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
DST_DIR=$PYTHON_DIR/boundlessgeo_schema
# source directory for proto
SRC_DIR=$SCRIPT_DIR/../proto
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
with open('../actions.json', 'r') as actions_json:
    actions_str = actions_json.read()
    actions_dict = loads(actions_str)
    with open('../script/package.json', 'r') as version_file:
        json_str = version_file.read()
        version = loads(json_str)['version']
        actions_dict['_version'] = version
with open('../python/boundlessgeo_schema/actions.p', 'wb') as fp:
    pickle.dump(actions_dict, fp)
EOF
)
protoc -I=$SRC_DIR --python_out=$DST_DIR $SRC_DIR/Msg.proto
# switch to the python dir
cd $PYTHON_DIR
# build sdist package
python setup.py sdist
# switch back to the dir you ran the script from
cd $PWD

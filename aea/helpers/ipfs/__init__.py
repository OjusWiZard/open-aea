# -*- coding: utf-8 -*-
# ------------------------------------------------------------------------------
#
#   Copyright 2022-2023 Valory AG
#   Copyright 2018-2021 Fetch.AI Limited
#
#   Licensed under the Apache License, Version 2.0 (the "License");
#   you may not use this file except in compliance with the License.
#   You may obtain a copy of the License at
#
#       http://www.apache.org/licenses/LICENSE-2.0
#
#   Unless required by applicable law or agreed to in writing, software
#   distributed under the License is distributed on an "AS IS" BASIS,
#   WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
#   See the License for the specific language governing permissions and
#   limitations under the License.
#
# ------------------------------------------------------------------------------
"""This module contains helper methods and classes for the 'aea' package."""
from aea.helpers.ipfs.utils import _protobuf_python_implementation


# fix for ipfs hashes, preload protobuf classes with protobuf python implementation
with _protobuf_python_implementation():
    from aea.helpers.ipfs.pb import (  # type: ignore  # noqa: F401   # pylint: disable=import-outside-toplevel,unused-import
        merkledag_pb2,
        unixfs_pb2,
    )
    from aea.helpers.ipfs.pb.merkledag_pb2 import (  # type: ignore  # noqa: F401   # pylint: disable=import-outside-toplevel,unused-import
        PBNode,
    )

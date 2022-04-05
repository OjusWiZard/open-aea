# -*- coding: utf-8 -*-
# ------------------------------------------------------------------------------
#
#   Copyright 2021 Valory AG
#   Copyright 2018-2019 Fetch.AI Limited
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

"""Module with methods for ipfs registry."""

import json
import logging
import os
from pathlib import Path
from typing import Dict, List, Optional, Union

import jsonschema
from aea_cli_ipfs.ipfs_utils import DownloadError, IPFSTool, NodeError

from aea.cli.registry.settings import (
    DEFAULT_IPFS_URL,
    REGISTRY_CONFIG_KEY,
    REGISTRY_IPFS,
)
from aea.cli.utils.config import get_or_create_cli_config
from aea.configurations.base import ExtendedPublicId


_default_logger = logging.getLogger(__name__)

LocalRegistry = Dict[str, Dict[str, str]]

LOCAL_REGISTRY_PATH = os.path.join(
    os.path.expanduser("~"), ".aea", "local_registry.json"
)
LOCAL_REGISTRY_DEFAULT: LocalRegistry = {
    "protocols": {},
    "skills": {},
    "connections": {},
    "contracts": {},
    "agents": {},
}

LOCAL_REGISTRY_SCHEMA = {
    "type": "object",
    "properties": {
        "protocols": {
            "type": "object",
            "propertyNames": {"pattern": r"^[a-z][a-z0-9_]+\/[a-z_0-9]+:\d\.\d\.\d$"},
        },
        "skills": {"type": "object"},
        "connections": {"type": "object"},
        "contracts": {"type": "object"},
        "agents": {"type": "object"},
    },
    "required": ["protocols", "skills", "connections", "contracts", "agents"],
}


def validate_registry(registry_data: LocalRegistry) -> None:
    """
    Validate local registry data.

    :param registry_data: json like object containing registry data.
    """
    try:
        jsonschema.validate(registry_data, schema=LOCAL_REGISTRY_SCHEMA)
    except jsonschema.ValidationError as e:
        _default_logger.debug("Registry Not Valid")
        raise ValueError(str(e))


def write_local_registry(
    registry_data: LocalRegistry, registry_path: str = LOCAL_REGISTRY_PATH
) -> None:
    """
    Write registry data to file.

    :param registry_data: json like object containing registry data.
    :param registry_path: local registry path.
    """
    validate_registry(registry_data)
    with open(registry_path, mode="w+", encoding="utf-8") as fp:
        json.dump(registry_data, fp)


def load_local_registry(registry_path: str = LOCAL_REGISTRY_PATH) -> LocalRegistry:
    """Returns local registry data."""

    local_registry_path = Path(registry_path)
    if not local_registry_path.is_file():
        write_local_registry(LOCAL_REGISTRY_DEFAULT)
        return LOCAL_REGISTRY_DEFAULT

    with open(local_registry_path, mode="r", encoding="utf-8") as fp:
        registry_data = json.load(fp)
        validate_registry(registry_data)
        return registry_data


def get_ipfs_hash_from_public_id(
    item_type: str,
    public_id: ExtendedPublicId,
    registry_path: str = LOCAL_REGISTRY_PATH,
) -> Optional[str]:
    """Get IPFS hash from local registry."""

    registry_data = load_local_registry(registry_path=registry_path)
    if public_id.package_version.is_latest:
        package_versions: List[ExtendedPublicId] = [
            ExtendedPublicId.from_str(_public_id)
            for _public_id in registry_data.get(f"{item_type}s", {}).keys()
            if public_id.same_prefix(ExtendedPublicId.from_str(_public_id))
        ]
        package_versions = list(
            reversed(sorted(package_versions, key=lambda x: x.package_version))
        )
        if len(package_versions) == 0:
            return None
        public_id, *_ = package_versions

    return registry_data.get(f"{item_type}s", {}).get(str(public_id), None)


def register_item_to_local_registry(
    item_type: str,
    public_id: Union[str, ExtendedPublicId],
    package_hash: str,
    registry_path: str = LOCAL_REGISTRY_PATH,
) -> None:
    """
    Add ExtendedPublicId to hash mapping in the local registry.

    :param item_type: item type.
    :param public_id: public id of package.
    :param package_hash: hash of package.
    :param registry_path: local registry path.
    """

    registry_data = load_local_registry(registry_path=registry_path)
    registry_data[f"{item_type}s"][str(public_id)] = str(package_hash)
    write_local_registry(registry_data, registry_path)


def fetch_ipfs(
    item_type: str, public_id: ExtendedPublicId, dest: str, remote: bool = True,
) -> Optional[Path]:
    """Fetch a package from IPFS node."""
    if remote:
        multiaddr = (
            get_or_create_cli_config()
            .get(REGISTRY_CONFIG_KEY, {})
            .get("settings", {})
            .get(REGISTRY_IPFS, {})
            .get("ipfs_node")
        )
        ipfs_tool = IPFSTool(multiaddr)
    else:

        ipfs_tool = IPFSTool(addr=DEFAULT_IPFS_URL)

    try:
        package_hash = public_id.hash
    except ValueError:
        package_hash = (
            None if remote else get_ipfs_hash_from_public_id(item_type, public_id)
        )

    if package_hash is None:
        raise Exception("Please provide hash.")

    try:
        ipfs_tool.check_ipfs_node_running()
    except NodeError:  # pragma: nocover
        if not remote:
            ipfs_tool.daemon.start()
        else:
            raise Exception(f"Cannot connect to node with addr: {ipfs_tool.addr}")

    try:
        *_download_dir, _ = os.path.split(dest)
        download_dir = os.path.sep.join(_download_dir)
        ipfs_tool.download(package_hash, download_dir)
        package_path = Path(dest).absolute()
        ipfs_tool.daemon.stop()
        return package_path

    except DownloadError as e:  # pragma: nocover
        ipfs_tool.daemon.stop()
        raise Exception(str(e)) from e

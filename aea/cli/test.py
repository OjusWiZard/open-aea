# -*- coding: utf-8 -*-
# ------------------------------------------------------------------------------
#
#   Copyright 2021-2022 Valory AG
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

"""Implementation of the 'aea test' command."""
import os
import sys
import tempfile
import time
from pathlib import Path
from typing import Callable, Optional, Sequence, Set, cast

import click
import pytest
from coverage.cmdline import main as coverage

from aea.cli.utils.click_utils import (
    PublicIdParameter,
    determine_package_type_for_directory,
)
from aea.cli.utils.context import Context
from aea.cli.utils.decorators import check_aea_project, pass_ctx, pytest_args
from aea.cli.utils.package_utils import get_package_path
from aea.components.base import load_aea_package
from aea.configurations.base import ComponentConfiguration
from aea.configurations.constants import (
    AEA_TEST_DIRNAME,
    CONNECTION,
    CONTRACT,
    PROTOCOL,
    SKILL,
)
from aea.configurations.data_types import (
    ComponentId,
    ComponentType,
    PackageType,
    PublicId,
)
from aea.configurations.loader import load_component_configuration
from aea.configurations.manager import find_component_directory_from_component_id
from aea.exceptions import enforce
from aea.helpers.base import cd
from aea.helpers.dependency_tree import DependencyTree


COVERAGERC_FILE = ".coveragerc"
COVERAGERC_CONFIG = """[run]
omit =
    */tests/*

[html]
directory = {root_dir}/htmlcov

[xml]
output = {root_dir}/coverage.xml
"""


@click.group(invoke_without_command=True)
@click.pass_context
@click.option(
    "--cov",
    is_flag=True,
    default=False,
    help="Use this flag to enable code coverage checks.",
)
@pytest_args
def test(click_context: click.Context, cov: bool, args: Sequence[str]) -> None:
    """Run tests of an AEA project."""
    ctx = cast(Context, click_context.obj)
    ctx.config["cov"] = cov
    if click_context.invoked_subcommand is None:
        test_aea_project(click_context, Path(ctx.cwd), args)


@test.command()
@click.argument("connection_public_id", type=PublicIdParameter(), required=True)
@pytest_args
@pass_ctx
def connection(
    ctx: Context, connection_public_id: PublicId, args: Sequence[str]
) -> None:
    """Executes a test suite of a connection package dependency."""
    test_item(ctx, CONNECTION, connection_public_id, args)


@test.command()
@click.argument("contract_public_id", type=PublicIdParameter(), required=True)
@pytest_args
@pass_ctx
def contract(ctx: Context, contract_public_id: PublicId, args: Sequence[str]) -> None:
    """Executes a test suite of a contract package dependency."""
    test_item(ctx, CONTRACT, contract_public_id, args)


@test.command()
@click.argument("protocol_public_id", type=PublicIdParameter(), required=True)
@pytest_args
@pass_ctx
def protocol(ctx: Context, protocol_public_id: PublicId, args: Sequence[str]) -> None:
    """Executes a test suite of a protocol package dependency."""
    test_item(ctx, PROTOCOL, protocol_public_id, args)


@test.command()
@click.argument("skill_public_id", type=PublicIdParameter(), required=True)
@pytest_args
@pass_ctx
def skill(ctx: Context, skill_public_id: PublicId, args: Sequence[str]) -> None:
    """Executes a test suite of a skill package dependency."""
    test_item(ctx, SKILL, skill_public_id, args)


@test.command()
@click.argument(
    "path", type=click.Path(exists=True, file_okay=False, dir_okay=True), required=True
)
@pytest_args
@pass_ctx
def by_path(
    ctx: Context,
    path: str,
    args: Sequence[str],
) -> None:
    """Executes a test suite of a package specified by a path."""
    click.echo(f"Executing tests of package at {path}'...")
    full_path = Path(ctx.cwd) / Path(path)
    test_package_by_path(
        full_path,
        args,
        packages_dir=Path(ctx.registry_path),
        cov=ctx.config.get("cov", False),
    )


@test.command()
@pytest_args
@pass_ctx
def packages(
    ctx: Context,
    args: Sequence[str],
) -> None:
    """Executes a test suite of a package specified by a path."""
    packages_dir = Path(ctx.registry_path)
    available_packages = DependencyTree.find_packages_in_a_local_repository(
        packages_dir
    )

    coverage_data = []
    failures = []
    start_time = time.perf_counter()

    with tempfile.TemporaryDirectory() as temp_dir:
        covrc_file = Path(temp_dir, COVERAGERC_FILE)
        covrc_file.write_text(
            COVERAGERC_CONFIG.format(root_dir=str(packages_dir.parent))
        )
        for package_type, package_dir in available_packages[:20]:
            if package_type == PackageType.AGENT.value:
                continue

            test_dir = package_dir / AEA_TEST_DIRNAME
            if not test_dir.exists():
                continue

            load_package(package_dir, packages_dir=packages_dir)
            with cd(package_dir):
                click.echo(
                    f"Running tests for {package_dir.name} of type {package_type}"
                )
                exit_code = pytest.main(
                    [
                        package_dir / AEA_TEST_DIRNAME,
                        f"--cov={package_dir.absolute()}",
                        f"--doctest-modules",
                        str(package_dir.absolute()),
                        f"--cov-config={covrc_file}",
                        "--cov-report=term",
                    ]
                )
                if exit_code:
                    click.echo(
                        f"Running tests for for {package_dir.name} of type {package_type} failed"
                    )
                    failures.append(
                        (exit_code, os.path.sep.join(package_dir.parts[-3:]))
                    )
            coverage_data.append(str(package_dir / ".coverage"))
        total_time = (time.perf_counter() - start_time) // 60

        click.echo(f"Time : {total_time}")
        click.echo("Generating coverage reports.")

        coverage(argv=["combine", f"--rcfile={covrc_file}", *coverage_data])
        coverage(argv=["html", f"--rcfile={covrc_file}"])
        coverage(argv=["xml", f"--rcfile={covrc_file}"])

    if len(failures):
        click.echo("Failed tests")
        click.echo("Exit Code\tPackage")
        for exit_code, package in failures:
            click.echo(f"{exit_code}       \t{package}")
        sys.exit(1)


def test_item(
    ctx: Context,
    item_type: str,
    item_public_id: PublicId,
    pytest_arguments: Sequence[str],
) -> None:
    """
    Run tests of a package dependency.

    :param ctx: the context.
    :param item_type: the item type.
    :param item_public_id: the item public id.
    :param pytest_arguments: arguments to forward to Pytest
    """
    click.echo(
        "Executing tests of component of type {}, {}' ...".format(
            item_type, item_public_id
        )
    )
    package_dirpath = Path(
        get_package_path(ctx.cwd, item_type, item_public_id, is_vendor=False)
    )
    if not package_dirpath.exists():
        # check if it is a vendor package
        package_dirpath = Path(
            get_package_path(ctx.cwd, item_type, item_public_id, is_vendor=True)
        )
        enforce(
            package_dirpath.exists(),
            exception_text=f"package {item_public_id} of type {item_type} not found",
            exception_class=click.ClickException,
        )
    # for a package in an AEA project, the "packages" dir is the AEA project dir
    aea_project_path = Path(ctx.cwd)
    test_package_by_path(
        package_dirpath,
        pytest_arguments,
        aea_project_path=aea_project_path,
        cov=ctx.config.get("cov", False),
    )


def load_package(
    package_dir: Path,
    aea_project_path: Optional[Path] = None,
    packages_dir: Optional[Path] = None,
) -> None:
    """Load packages into cache."""
    enforce(
        (aea_project_path is None) != (packages_dir is None),
        "one of either aea_project_path or packages_dir must be specified",
    )
    root_packages = aea_project_path if aea_project_path else packages_dir

    package_path_finder = (
        find_component_directory_from_component_id
        if aea_project_path
        else find_component_directory_from_component_id_in_registry
    )

    # check the path points to a valid AEA package
    package_type = determine_package_type_for_directory(package_dir)

    if package_type != PackageType.AGENT:
        component_type = ComponentType(package_type.value)
        configuration = load_component_configuration(component_type, package_dir)
        configuration.directory = package_dir
        load_aea_packages_recursively(configuration, package_path_finder, root_packages)

    test_package_dir = package_dir / AEA_TEST_DIRNAME
    enforce(
        test_package_dir.exists(),
        f"tests directory in {package_dir} not found",
        click.ClickException,
    )


def test_package_by_path(
    package_dir: Path,
    pytest_arguments: Sequence[str],
    aea_project_path: Optional[Path] = None,
    packages_dir: Optional[Path] = None,
    cov: bool = False,
) -> None:
    """
    Fingerprint package placed in package_dir.

    :param package_dir: directory of the package
    :param pytest_arguments: arguments to forward to Pytest
    :param aea_project_path: directory to the AEA project
    :param packages_dir: directory of the packages to import from
    """

    load_package(package_dir, aea_project_path, packages_dir)
    exit_code = run_pytest(package_dir, pytest_arguments=pytest_arguments, cov=cov)
    sys.exit(exit_code)


def run_pytest(
    package_dir: Path,
    pytest_arguments: Sequence[str],
    cov: bool = False,
) -> int:
    """Run pytest."""

    runtime_args = [
        AEA_TEST_DIRNAME,
        *pytest_arguments,
    ]

    if cov:
        runtime_args.extend(
            [
                f"--cov={package_dir.absolute()}",
                f"--doctest-modules",
                str(package_dir.absolute()),
                "--cov-report=term",
                "--cov-report=term-missing",
            ]
        )

    with cd(package_dir):
        return pytest.main(runtime_args)


@check_aea_project
def test_aea_project(
    click_context: click.Context, aea_project_dirpath: Path, args: Sequence[str]
) -> None:
    """Run tests of an AEA project."""
    click.echo("Executing tests of the AEA project...")
    ctx = cast(Context, click_context.obj)
    # in case of an AEA project, the 'packages' directory is the AEA project path itself
    test_package_by_path(aea_project_dirpath, args, aea_project_path=Path(ctx.cwd))


def load_aea_packages_recursively(
    config: ComponentConfiguration,
    package_path_finder: Callable[[Path, ComponentId], Path],
    root_packages: Path,
    already_loaded: Optional[Set[ComponentId]] = None,
) -> None:
    """
    Load all AEA packages recursively.

    It works like 'load_aea_package', but recursively imports all dependencies.

    :param config: the component configuration
    :param package_path_finder: a function that find packages from the packages dir
    :param root_packages: the path to the root of packages dir
    :param already_loaded: the already loaded component ids
    """
    already_loaded = already_loaded if already_loaded else set()
    for dependency_id in config.package_dependencies:
        # TODO: load packages in topological order? Should not matter as at the moment we are not
        #       actually running the modules, just populating sys.modules
        dependency_path = package_path_finder(root_packages, dependency_id)
        dependency_configuration = load_component_configuration(
            dependency_id.component_type, dependency_path
        )
        dependency_configuration.directory = dependency_path
        load_aea_packages_recursively(
            dependency_configuration, package_path_finder, root_packages, already_loaded
        )
    load_aea_package(config)
    already_loaded.add(config.component_id)


def find_component_directory_from_component_id_in_registry(
    registry_path: Path, component_id: ComponentId
) -> Path:
    """Find a component directory from component id in a registry."""
    package_path = (
        registry_path
        / component_id.author
        / component_id.component_type.to_plural()
        / component_id.public_id.name
    )
    if package_path.exists() and package_path.is_dir():
        return package_path

    raise ValueError("Package {} not found.".format(component_id))

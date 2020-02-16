#!/usr/bin/env python3
# @generated AUTOGENERATED file. Do not Change!

from dataclasses import dataclass, field
from datetime import datetime
from enum import Enum
from functools import partial
from numbers import Number
from typing import Any, Callable, List, Mapping, Optional

from dataclasses_json import dataclass_json
from marshmallow import fields as marshmallow_fields

from .datetime_utils import fromisoformat


DATETIME_FIELD = field(
    metadata={
        "dataclasses_json": {
            "encoder": datetime.isoformat,
            "decoder": fromisoformat,
            "mm_field": marshmallow_fields.DateTime(format="iso"),
        }
    }
)


@dataclass_json
@dataclass
class LatestPythonPackageQuery:
    __QUERY__ = """
    query LatestPythonPackageQuery {
  latestPythonPackage {
    lastPythonPackage {
      version
    }
    lastBreakingPythonPackage {
      version
    }
  }
}

    """

    @dataclass_json
    @dataclass
    class LatestPythonPackageQueryData:
        @dataclass_json
        @dataclass
        class LatestPythonPackageResult:
            @dataclass_json
            @dataclass
            class PythonPackage:
                version: str

            @dataclass_json
            @dataclass
            class PythonPackage:
                version: str

            lastPythonPackage: Optional[PythonPackage] = None
            lastBreakingPythonPackage: Optional[PythonPackage] = None

        latestPythonPackage: Optional[LatestPythonPackageResult] = None

    data: Optional[LatestPythonPackageQueryData] = None
    errors: Optional[Any] = None

    @classmethod
    # fmt: off
    def execute(cls, client):
        # fmt: off
        variables = None
        response_text = client.call(cls.__QUERY__, variables=variables)
        return cls.from_json(response_text).data
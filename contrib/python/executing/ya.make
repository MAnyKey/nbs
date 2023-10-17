# Generated by devtools/yamaker (pypi).

PY3_LIBRARY()

VERSION(2.0.0)

LICENSE(MIT)

NO_LINT()

PY_SRCS(
    TOP_LEVEL
    executing/__init__.py
    executing/_exceptions.py
    executing/_position_node_finder.py
    executing/executing.py
    executing/version.py
)

RESOURCE_FILES(
    PREFIX contrib/python/executing/
    .dist-info/METADATA
    .dist-info/top_level.txt
    executing/py.typed
)

END()

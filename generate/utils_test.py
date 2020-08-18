from generate.utils import index_or_none, remove_upper_next, trim_by_index, format_to_go_struct, format_golang_comment, last_index, rename_reserved


def test_rename_reserver():
    assert "import_" == rename_reserved("import")
    assert "package_" == rename_reserved("package")
    assert "type_" == rename_reserved("type")
    assert "good" == rename_reserved("good")


def test_format_to_go_struct():
    _c = "aasdasd-dcas_asdad-a"
    assert "AasdasdDcasAsdadA" == format_to_go_struct(_c)


def test_last_index():
    _c = "asdasd-zxczasd-asd-43-g-f"
    assert 23 == last_index(_c, "-")


def test_format_golang_comment():
    _c = "asdasd\nasdsads\n"
    assert "//asdasd\n//asdsads\n//" == format_golang_comment(_c)


def test_remove_upper_next():
    _c = "bad-command-name"
    assert "BadCommandName" == remove_upper_next(_c, "-")


def test_index_or_none():
    _c = "char-c-a"
    while _i := index_or_none(_c, "-"):
        _c = _c.replace("-", "")
    assert _c == "charca"


def test_trim_by_index():
    _c = "char-ad-sd-c-cc"
    assert "char-ad-sd" == trim_by_index(_c, "-c-cc")

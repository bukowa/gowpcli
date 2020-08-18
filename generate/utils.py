import random
import string
from typing import Union


def make_tabs(s: str) -> str:
    """
    if there is no space after \n add a tab
    """
    s.rsplit()

def rename_reserved(s: str) -> str:
    """
    some names are reserved in go, handle them here
    """
    if any([
        "import" == s.lower(),
        "type" == s.lower(),
        "package" == s.lower(),
    ]):
        return s + "_"
    return s


def random_string(n: int) -> str:
    """
    return random string of len :param n
    """
    return "".join([random.choice(string.ascii_letters) for _ in range(n)])


def last_index(s: str, char: str) -> int:
    """
    return last index of :param char in s
    """
    return s.rindex(char)


def format_golang_comment(s: str) -> str:
    """
    format string to be displayed as golang comment.
    """
    if not s.startswith("//"):
        s = "//" + s
    return s.replace("\n", "\n//")


def remove_upper_next(s: str, char: str) -> str:
    """
    removes all chars :param char from :param s
    and capitalizes char after :param char
    """
    while _i := index_or_none(s, char):
        s = s.replace(char, "", 1)
        s = s[:_i] + s[_i].upper() + s[_i+1:]
    s = s[0].upper() + s[1:]
    return s


def format_to_go_struct(s: str):
    """
    format to what will will appear in golang struct
    """
    s = remove_upper_next(s, "-")
    s = remove_upper_next(s, "_")
    s = s[0].upper() + s[1:]
    return s


def index_or_none(s: str, char: str) -> Union[int, None]:
    """
    returns index of :param char in :param s or None
    """
    try: return s.index(char)
    except ValueError: return None


def trim_by_index(s: str, char: str) -> str:
    """
    remove all chars behind s.index(char)
    """
    try:
        _i = s.index(char)
    except ValueError: pass
    else:
        return s[:_i]

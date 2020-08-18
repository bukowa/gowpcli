
class Synopsis:
    struct: str = ""

    def __init__(self, text: str):
        self.text = text
        self.params = text.split(" ")

    @classmethod
    def field(cls, param: str) -> str:
        """
        Create field for Go struct.
        """
        def _replacer(_param: str) -> str:
            for _x in ["<", ">", "[", "]", "--", "..."]:
                _param = _param.replace(_x, "")
            return _param

        if "=" in param:
            try:
                _i = param.index("=")
                param = param[:_i]
                value = param[_i+1:]
            except ValueError:
                pass

        return _replacer(param)

    @classmethod
    def named_map(cls, param: str) -> bool:
        """
        Param can be key=value '[--<field>=<value>]'.
        This is a special treatment.
        Additional field can be anything,
        this should me a Go map then.
        """
        if "[--<field>=<value>]" in param:
            return True
        return False

    @classmethod
    def option(cls, param: str) -> bool:
        """
        Param that can be set or not.
        Its a boolean without value.
        It can be set or not.
        [--dir] | [--all]
        """
        return all([
            "--" in param,
            "=" not in param,
        ])

    @classmethod
    def optional(cls, param: str) -> bool:
        return not cls.required(param)

    @classmethod
    def multiple(cls, param: str) -> bool:
        """
        Param that accepts multiple values.
        """
        return any([
            param.endswith("..."),
            param.endswith("...]")
        ])

    @classmethod
    def required(cls, param: str) -> bool:
        return not param.startswith("[")

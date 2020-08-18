import json
import os
import shutil
from dataclasses import dataclass, field
from typing import List, Union

from jinja2 import Template
from generate.synopsis import Synopsis
from generate.utils import trim_by_index, format_to_go_struct, format_golang_comment, random_string, rename_reserved

COMMANDS: List['Command'] = []
MAIN_IMPORTS: List['GoCommand'] = []

TEMPLATE_STRUCT = Template(open("template_struct.html", "r", encoding="utf-8").read())
TEMPLATE_MAIN = Template(open("template_main.html", "r", encoding="utf-8").read())


@dataclass
class GoField:
    param: str                                          # raw param like '[--field=<field>]'
    name: str = field(init=False, default="")           # name of the field in the struct 'Postid'
    type_: str = field(init=False, default="")          # type of the field in the struct '[]string'
    required: bool = field(init=False, default=False)   # some params are required

    def __post_init__(self):
        self.name = Synopsis.field(self.param).replace("|", "").capitalize()
        if Synopsis.required(self.param):
            self.required = True
        if Synopsis.multiple(self.param):
            self.type_ = "[]"
        if Synopsis.option(self.param):
            self.type_ += "bool"
        else:
            self.type_ += "string"
        if Synopsis.named_map(self.param):
            self.type_ = "map[string]string"
            self.name += "Map"
        self.name = format_to_go_struct(self.name)


@dataclass
class GoCommand:
    command: 'Command'
    package: str = None
    struct_name: str = None
    fields: List[GoField] = field(default_factory=list)
    folder: str = field(init=False, default=None)
    path: str = field(init=False, default=None)
    import_: str = field(init=False, default=None)

    def __post_init__(self):
        # format data to golang
        self.struct_name = format_to_go_struct(self.command.name)
        self.struct_name = rename_reserved(self.struct_name)
        self.package = self.struct_name.replace("-", "_")
        self.folder = self.package.lower()
        self.command.longdesc = trim_by_index(self.command.longdesc, "## GLOBAL PARAMETERS")
        self.command.longdesc = format_golang_comment(self.command.longdesc)
        self.command.description = format_golang_comment(self.command.description)
        self.import_ = random_string(30)

        # build golang fields
        if self.command.synopsis:
            synopsis = Synopsis(text=self.command.synopsis)
            self.fields = [GoField(_x) for _x in synopsis.params]


@dataclass
class Command:
    name: str
    description: str = field(default="")
    longdesc: str = field(default="")
    # when new Command is created we pass a list of dict, that is later converted to list of 'Command'
    subcommands: Union[List[dict], List['Command']] = field(default_factory=list)
    synopsis: Union[str, None] = None
    parent: Union[None, 'Command'] = None
    gocommand: GoCommand = None

    @property
    def folder(self):
        return self.gocommand.folder

    def __post_init__(self):
        self.gocommand = GoCommand(command=self)

        if not self.parent:
            path = f"generated/{self.folder}"
        elif self.parent and self.parent.parent:
            path = f"generated/{self.parent.folder}/{self.parent.parent.folder}/{self.folder}"
        elif self.parent:
            path = f"generated/{self.parent.folder}/{self.folder}"

        self.gocommand.path = path
        if not os.path.isdir(path):
            os.makedirs(path)

        with open(f"{path}/{self.folder}.go", "w", encoding="utf-8") as f:
            f.write(TEMPLATE_STRUCT.render(
                go=self.gocommand,
                cmd=self,
            ))

        self.subcommands = [Command(**_c, parent=self) for _c in self.subcommands]
        # handle main file
        MAIN_IMPORTS.append(self.gocommand)


if __name__ == '__main__':
    try: shutil.rmtree("./generated")
    except FileNotFoundError: pass

    os.makedirs("./generated")
    x = json.load(open("dump.json", "r", encoding="utf-8"))

    for cmd in x['subcommands']:
        COMMANDS.append(Command(**cmd))

    with open("./generated/main.go", "w", encoding="utf-8") as f:
        f.write(TEMPLATE_MAIN.render(
            imports=MAIN_IMPORTS,
        ))

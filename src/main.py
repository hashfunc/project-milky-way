import importlib.metadata

from typer import Typer

PROJECT_NAME = "project-milky-way"

app = Typer(
    no_args_is_help=True,
)


@app.command()
def main() -> None:
    pass


@app.command()
def version() -> None:
    __version__ = importlib.metadata.version(PROJECT_NAME)
    print(f"{PROJECT_NAME} {__version__}")


if __name__ == "__main__":
    app()

# mudora

<div align="center">
  <img src="assets/mudora.png" />
  <h3>The Book of Mudora</h3>

  <blockquote>The monoliths left by the Hylian people are inscribed with ancient script. If you find an inscription you cannot read, use this book and its meaning will become clear.</blockquote>
</div>

mudora is an ALttPR ROM inspection tool. It shows the item locations for a given ROM and can perform item searches if you're stuck. In addition to displaying item locations, mudora can attempt to solve the game for you, providing a shortest path candidate to defeat Ganon (defined as the path with the fewest required chests).

> [!WARNING]  
> Using this tool to cheat on races makes Link sad. ![Link fainting](/assets/link-faint.png)

# Usage

Run in CLI mode:

```sh
go run ./cmd/mudora --rom rom.sfc [item-query]

go run ./cmd/mudora --rom rom.sfc --solve

go run ./cmd/mudora --version
```

`-r`/`-s`/`-v` work as shorthand for `--rom`/`--solve`/`--version`.

Build a local web UI bundle:

```sh
go run ./cmd/mudora-web/build
```

Serve the UI with your favorite HTTP server, such as Python's built-in server:

```sh
# Assumes you've built a web bundle against a cloned repository, serving on port 8080
python -m http.server 8080 --directory ./cmd/mudora-web/web
```

# Contact

Ty Porter

tyler.b.porter@gmail.com

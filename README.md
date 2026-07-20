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

## Web UI

Build a local web UI bundle:

```sh
go run ./cmd/mudora-web/build
```

Serve the UI with your favorite HTTP server, such as Python's built-in server:

```sh
# Assumes you've built a web bundle against a cloned repository, serving on port 8080
python -m http.server 8080 --directory ./cmd/mudora-web/web
```

## CLI

```sh
go run ./cmd/mudora <flags>
```

Compile a binary from source and place it on `PATH` if you prefer:

```sh
go build ./cmd/mudora
```

Then use with `mudora <flags>`.

### Print out the item locations for a ROM

```sh
mudora -rom <ROM>
```

### Print the playthrough path for a ROM

```sh
mudora -rom <ROM> -solve
```

### Display the start screen code for the ROM

The start screen code are the items that are listed at the top of the start screen for a randomized game. They serve as a visual identifier confirming you're playing the right seed if you share ROMs, such as during a race.

```sh
mudora -rom <ROM> -hash
```

### Print the alttpr.com permalink for the ROM

All ROMs have an internal game ID that can be used to permalink the ROM download on [https://alttpr.com](https://alttpr.com). The game may or may not exist yet (especially true if you generated the ROM locally).

```sh
mudora -rom <ROM> -permalink
```

### Read continuous bytes at an address

You can also use Mudora to read the ROM directly for convenience. Values are printed in hexadecimal.

Read 16 bytes starting at `0x1234`.

```sh
mudora -rom <ROM> -read-bytes -start-byte 0x1234 -byte-count 16
```

It's a bit eaiser to use the shorthand for these flags:

```sh
mudora -r <ROM> -rb -sb 0x1234 -bc 16
```

### Available Flags

```
  -bc int
        shortand for -byte-count
  -byte-count int
        byte count for -read-bytes
  -hash
        print ROM hash items
  -permalink
        print the alttpr.com permalink hash embedded in the ROM
  -r string
        path to ALttPR ROM file (shorthand)
  -rb
        shortand for -read-bytes
  -reachable
        print all reachable items
  -read-bytes
        read raw bytes (requires -start-byte/-byte-count)
  -rom string
        path to ALttPR ROM file
  -s    shorthand for -solve
  -sb string
        shortand for -read-bytes
  -solve
        print the shortest path to Ganon's Tower
  -start-byte string
        start byte for -read-bytes
  -v    shorthand for -version
  -version
        print the mudora version
```

# Contact

Ty Porter

tyler.b.porter@gmail.com

# Trimlet DB

A read-only browser for **PostgreSQL** and **SQL Server**, in a small window driven only by the keyboard.

- Saved connections → database → tables with their row counts → a table: columns, keys and indexes,
  its `CREATE TABLE`, and all its rows, 50 per page.
- Rows can be copied as text or JSON, and exported to CSV, `INSERT`, or `DROP` + `CREATE` + `INSERT`.
- It only reads: there is no SQL prompt, and the scripts it writes are text for you, never run.
- The connection strings are kept in the system keychain (Keychain on Mac, Credential Manager on
  Windows); the list of connections in `trimletdb/connections.json` in the app settings folder.

## Download

The app for **macOS** (Mac with Apple chip: M1 and later) and **Windows** is on the
[Releases](https://github.com/trimlet/db/releases) page. It is a beta (0.x), and not signed yet: the first
time, macOS asks to confirm (right click on the app → Open), Windows shows SmartScreen (More info →
Run anyway).

## Build

Written in Go with [Fyne](https://fyne.io). To build it, with Go installed (and on a Mac the Xcode
command line tools):

```
go run fyne.io/tools/cmd/fyne@v1.7.3 package --release
```

or `go run .` to try it. The console window comes from [trimlet/gui](https://github.com/trimlet/gui),
which `go` fetches by itself.

## How it is made

The code is written with Claude Opus (Anthropic), under the direction and review of a developer with
almost 30 years of experience. Every feature starts from a real, repeated need, and most ideas are
turned down: the goal is to stay small.

## License

GPL-3.0, see [LICENSE](LICENSE).

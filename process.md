# The Solana MagicEden Floor tracker

---

## What is Floor Tracker?

A CLI tool which returns various info about a certain listing on MagicEden

## Usage

1. Run the executable via ``./floortracker` and input one of these subcommands:

   - `floor "<collection1>" "<collection2>" "<collection3>" ... <collectionX>`
   - `floor "<collection>"`

2. Or pass in the collection arguments directly in a shell command:
   - `./floortracker "<collection1>" "<collection2>" "<collection3>" ... <collectionX>`

## How it works

1. User inputs a command and argument/arguments.
2. Main module parses arguments and options.
3. `main` module calls `getStats()` function which returns a struct of stats data using `datafetcher` and `jsonworker` modules to fetch from ME API and format into JSON
4.

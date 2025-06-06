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
3. `main` module calls `solana` module, which makes a request to MagicEden's dev API and returns the fetched data in JSON form.
4. `main` module calls the ``formatter` module to format gotten data into a specified interface
5. `main` module displays data nicely in terminal

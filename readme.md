# The Solana MagicEden Floor tracker

---

## What is Floor Tracker?

A CLI tool which returns various info about a certain listing on MagicEden

## Usage

1. Run the executable via ``./floortracker` and input the collections you want to check
   - `./floortracker "<collection1>" "<collection2>" "<collection3>" ... <collectionX>`
2. Floortracker also supports JSON and CSV printing. Input these arguments after calling the executable:
   - For JSON: `./floortracker -j "<collection1>"...`
   - For CSV: `./floortracker -c "<collection1>"...`

## How it works

1. User inputs a command and argument/arguments.
2. Program concurrently fetches stats via the MagicEden API
3. The fetched stats are then exported in the preferred format

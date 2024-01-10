# [![Sia](https://sia.tech/assets/banners/sia-banner-expanded-hostd.png)](http://sia.tech)

A tool for migrating from the legacy siad host to hostd.

## Overview

`hostd` is an advanced Sia host solution created by the Sia Foundation, designed
to enhance the experience for storage providers within the Sia network. Tailored
for both individual and large-scale storage providers, `hostd` boasts a
user-friendly interface and a robust API, empowering providers to efficiently
manage their storage resources and revenue. `hostd` incorporates an embedded
web-UI, simplifying deployment and enabling remote management capabilities,
ensuring a smooth user experience across a diverse range of devices.

- A project roadmap is available on [GitHub](https://github.com/orgs/SiaFoundation/projects/3)

# Running

The utility is intended to be interactive. It will prompt for the location of
the legacy siad host database and convert it to the new format.

You will need to generate a new 12-word wallet seed and send your funds from
Sia-UI/siad to your new wallet address.

1. Download the latest release of `hostd`
   (https://sia.tech/host)
2. Extract the files and open terminal or command prompt in the directory
3. Run `hostd seed` to generate a new wallet seed. It will generate a new
   12-word seed and print the wallet address. **Write down all 12 words and
   double check it. You will need it to recover your funds.**
4. Send your funds from your old siad wallet to your new wallet.

## Migrating Sia-UI/Host Manager

1. Find your data directory

Operating System | Data Directory
---------------- | --------------
Windows          | `%APPDATA%\Sia-UI`
macOS            | `~/Library/Application Support/Sia-UI`
Linux            | `~/.config/Sia-UI`

2. Stop Sia-UI or Host Manager
3. Download the latest release of the migration utility (https://github.com/SiaFoundation/migrate-siad-host/releases)
4. Run the migration tool, answering any prompts
5. Wait for the migration to complete

## Migrating `siad`

1. Find your data directory
2. Stop siad
3. Download the latest release of the migration utility (https://github.com/SiaFoundation/migrate-siad-host/releases)
4. Run the migration tool, answering any prompts

## Running `hostd`

Setup `hostd`, making sure to start it in your existing data directory (https://docs.sia.tech/hosting/setup-guides)

# Building

`migrate` uses SQLite for its persistence. A gcc toolchain is required to build
the migration utility

```sh
go generate ./...
CGO_ENABLED=1 go build -o bin/ -tags='netgo timetzdata' -trimpath -a -ldflags '-s -w'  ./cmd/migrate
```


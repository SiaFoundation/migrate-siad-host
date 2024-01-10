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

```sh
./migrate
```


# Building

`hostd` uses SQLite for its persistence. A gcc toolchain is required to build
the migration utility

```sh
go generate ./...
CGO_ENABLED=1 go build -o bin/ -tags='netgo timetzdata' -trimpath -a -ldflags '-s -w'  ./cmd/migrate
```


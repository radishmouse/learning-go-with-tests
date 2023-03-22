
Learning Go by way of https://quii.gitbook.io/learn-go-with-tests/

Local dev on [NixOS](https://nixos.org/) using:

- `niv`
- `lorri`
- `direnv`


One-time setup (run in the same dir as `shell.nix`)

```sh
niv init && niv update nixpkgs -b nixpkgs-unstable && lorri init && direnv allow
```

After a couple minutes, you can go into any project folder and run

```sh
go test
```

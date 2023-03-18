let 
  sources = import ./nix/sources.nix;
  pkgs = import sources.nixpkgs {};

  packages = with pkgs; [
    go_1_20
    gopls
    gotools
  ];
in
pkgs.mkShell {
  buildInputs = packages;
}

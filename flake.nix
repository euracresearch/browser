{
  description = "LTER Browser Application";

  inputs = {
    nixpkgs.url = "nixpkgs";
    flake-utils.url = "github:numtide/flake-utils";
  };

  outputs = { self, nixpkgs, flake-utils }:
  flake-utils.lib.eachDefaultSystem (system:
  let
    pkgs = nixpkgs.legacyPackages.${system};

    packages = with pkgs; [
        go
        influxdb
    ];
  in
  {
  devShell = pkgs.mkShell {
        name = "lter-browser";
        buildInputs = packages;

        shellHook =
        ''
        '';
        };
    });
}

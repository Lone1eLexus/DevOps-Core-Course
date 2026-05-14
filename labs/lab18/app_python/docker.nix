{ pkgs ? import (fetchTarball "https://github.com/NixOS/nixpkgs/archive/nixos-24.05.tar.gz") {} }:

let
  app = import ./default.nix { inherit pkgs; };
in
pkgs.dockerTools.buildLayeredImage {
  name = "devops-info-service-nix";
  tag = "1.0.0";

  contents = [ app ];

  config = {
    Cmd = [ "${app}/bin/devops-info-service" ];
    ExposedPorts = {
      "8000/tcp" = {};
    };
    Env = [
      "PORT=8000"
      "HOST=0.0.0.0"
    ];
  };

  created = "1970-01-01T00:00:01Z";  # Fixed timestamp for reproducibility
}
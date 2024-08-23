{ pkgs ? import <nixpkgs> {} }:

pkgs.mkShell {
  buildInputs = with pkgs; [
    # General dependencies
    docker-compose
    docker
    git
    hoppscotch

    # Backend dependencies
    go
    libcap
    gcc
    air

    # Database dependencies
    postgresql

    # Frontend dependencies
    nodejs
  ];
}
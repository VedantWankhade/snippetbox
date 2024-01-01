let
    nixpkgs = fetchTarball "https://github.com/NixOS/nixpkgs/tarball/nixos-23.11";
    pkgs = import nixpkgs { config = {}; overlays = []; };
in

pkgs.mkShell {
    packages = with pkgs; [
        go
        gotools
        go-tools
        mysql-shell
    ];

    shellHook = ''
        echo "============= Snippetbox ==================="
        alias docker="sudo docker"
        echo "Entered Snippetbox dev environment."
        echo "For local database, please ensure that the docker is installed and create a mysql container."
        echo "------------------------"
        echo "Dependencies"
        go version
        echo "------------------------"
        echo "Git status"
        git status
        export PS1="(snippetbox nix-shell)$PS1"
    '';
}

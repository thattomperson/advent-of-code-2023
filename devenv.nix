{ pkgs, ... }:

{
  packages = [ pkgs.git ];
  languages.go.enable = true;
  pre-commit.hooks = {
    gotest.enable = true;
    gofmt.enable = true;
    nixfmt.enable = true;
  };
}

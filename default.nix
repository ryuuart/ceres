{
  sources ? import ./npins,
  system ? builtins.currentSystem,
  pkgs ? import sources.nixpkgs { inherit system; config = {}; overlays = []; },
}:
rec {
  shell = pkgs.mkShell {
    packages = with pkgs; [
      npins
      go
    ] ++ lib.optionals stdenv.isDarwin 
      (with darwin.apple_sdk.frameworks;
        [
          CoreFoundation
        ]
      );
  };
}

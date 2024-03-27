with import (fetchTarball "https://nixos.org/channels/nixos-unstable/nixexprs.tar.xz") { };
{
  devEnv = stdenv.mkDerivation {
    name = "dev";
    buildInputs = [ stdenv go_1_21 glibc ];
    shellHook = ''
      return
    '';
  };
}

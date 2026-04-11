{
  pkgs ? import <nixpkgs> { },
}:

pkgs.mkShell {
  nativeBuildInputs = with pkgs; [
    pkg-config
    go
  ];

  buildInputs = with pkgs; [
    portaudio
    glfw
    libGL
  ];

  shellHook = ''
    export CGO_ENABLED=1
  '';
}

{ pkgs ? import <nixpkgs> {} }:

pkgs.python3Packages.buildPythonPackage {
  pname = "devops-info-service";
  version = "1.0.0";
  src = ./.;
  format = "other"; 

  # Dependencies required by app.py
  # Note: Added prometheus-client and python-json-logger which are in app.py but missing from requirements.txt
  propagatedBuildInputs = with pkgs.python3Packages; [
    fastapi
    uvicorn
  ];

  nativeBuildInputs = [ pkgs.makeWrapper ];

  installPhase = ''
    mkdir -p $out/bin $out/lib
    cp app.py $out/lib/app.py
    
    # Create an executable wrapper that calls python with the correct PYTHONPATH
    makeWrapper ${pkgs.python3.interpreter} $out/bin/devops-info-service \
      --add-flags "$out/lib/app.py" \
      --set PYTHONPATH "$out/${pkgs.python3.sitePackages}" \
      --set HOST 0.0.0.0 \
      --set PORT 8000
  '';
}
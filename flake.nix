{
  description = "Goeather web service";

  inputs = {
    nixpkgs.url = "github:nixos/nixpkgs?ref=nixos-unstable";
    flake-utils.url = "github:numtide/flake-utils";
  };

  outputs =
    {
      self,
      nixpkgs,
      flake-utils,
    }:
    let
      nixosModule =
        {
          config,
          lib,
          pkgs,
          ...
        }:
        {
          options.services.goeather = {
            enable = lib.mkEnableOption "Goeather service";

            port = lib.mkOption {
              type = lib.types.port;
              default = 8000;
              description = "Port to listen on";
            };
          };

          config = lib.mkIf config.services.goeather.enable {
            systemd.services.goeather = {
              description = "Goeather Service";
              wantedBy = [ "multi-user.target" ];
              after = [ "network.target" ];
              serviceConfig = {
                ExecStart = "${self.packages.${pkgs.system}.default}/bin/goeather";
                Restart = "always";
                Type = "simple";
                DynamicUser = "yes";
              };
            };
          };
        };
    in
    (flake-utils.lib.eachDefaultSystem (
      system:
      let
        pkgs = nixpkgs.legacyPackages.${system};
      in
      {
        packages.default = pkgs.buildGoModule {
          pname = "goeather";
          version = "0.1.0";
          src = ./.;
          vendorHash = null;
        };

        app.default = {
          type = "app";
          program = "${self.packages.${system}.default}/bin/goeather";
        };
      }
    ))
    // {
      nixosModules.default = nixosModule;
    };

}

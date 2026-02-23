Package configs
---------------

Package configurations are defined in per-package YAML files and loaded from
the package directory as defined in the solution configuration.

User provided keys:
------------------

| Key               | Type             | Description                                                                                  |
| ----------------- | ---------------- | -------------------------------------------------------------------------------------------- |
| `provides`        | string/list      | package names that this one provides - used on dependency lookup                             |
| `type`            | string           | package type (eg. bundle, system, ...)                                                       |
| `sources`         | dict (see below) | specify how to retrieve source tree (eg. via git)                                            |
| `pkg-config`      | string/list      | for `system`-packages: pkg-config query                                                      |
| `build-depends`   | string/list      | build-time dependencies (resolved via either package names or provides                       |
| `depends`         | string/list      | runtime dependencies                                                                         |
| `source-dir`      | string           | source tree directory, preset by MPBT (per package), but can also be set manually if needed  |
| `buildsystem`     | string           | buildsystem backend ([see here](builders.md)) to be used for the package                     |
| `name`            | string           | package name (default: yaml file name w/o suffix and package config directory prefix)        |
| `install-prefix`  | string           | installation prefix (by default filled by MPBT)                                              |
| `parallel`        | integer          | number of parallel jobs (defaults to ${@SOLUTION::parallel})                                 |
| `enable-binpkg`   | bool             | enable binary packages (experimental)                                                        |

Automatic keys:
---------------

These keys are automatically filled in by MPBT and should not be overwritten.

| Key             | Description |
| --------------- | ------------------------------------------------------------------------------- |
| `@filename`     | file name of the package config yaml file                                       |
| `@basename`     | package name without subdir prefix (eg `libdrm` instead of `3rdparty/libdrm`)   |
| `@slug`         | short package name slug (used eg. for per-package directory names)              |
| `@PROJECT`      | link to the global project config object                                        |
| `@SOLUTION`     | link to the global solution config object                                       |
| `@statdir`      | directory for MPBT status files                                                 |
| `@destdir`      | DESTDIR prefix, if install should go to a staging area                          |

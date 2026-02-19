Package builders
================

MPBT offers several package builder backends, which are responsible for calling
the invidual package's build system. The builder selection is done by the
`builder` key inside the package config.

Builders are running in several stages: `prepare`, `configure`, `build`, `install`.
(there's also a separate `clean` stage, which isn't practically used yet)

autotools
---------

Runs standard autotools build systems.

|    Stage    | Description                                                     |
| ----------- | --------------------------------------------------------------- |
| `prepare`   | calls `./autogen.sh`                                            |
| `configure` | calls `./configure --prefix=...` (plus potentially extra args)  |
| `build`     | calls `make` (potentially with extra args, eg. parallel builds) |
| `install`   | calls `make install` (plus target directory, etc)               |
| `clean`     | calls `make clean`                                              |

meson
-----

Runs standard meson build - using build subdirectory.

|    Stage    | Description                                                     |
| ----------- | --------------------------------------------------------------- |
| `prepare`   | prepare build subdir                                            |
| `configure` | calls `meson setup`                                             |
| `build`     | calls `meson compile`                                           |
| `install`   | calls `meson install`                                           |
| `clean`     | removes the build subdirectory                                  |

The package keys `meson-args` and `meson-extra-args` (string lists) are appended
to the `meson setup` command line. Those can be defined within the package config
or overriden via solution config.

cmake
-----

Runs standard meson build (with make) - using build subdirectory.

|    Stage    | Description                                                     |
| ----------- | --------------------------------------------------------------- |
| `prepare`   | prepare build subdir                                            |
| `configure` | calls `cmake` for generating build files                        |
| `build`     | calls `cmake --build` in build subdirectory                     |
| `install`   | calls `cmake --install` in build subdirectory                   |
| `clean`     | removes the build subdirectory                                  |

The package keys `cmake-args` and `cmake-extra-args` (string lists) are appended
to the `cmake` command line. Those can be defined within the package config
or overriden via solution config.

none
----

Dummy builder that's doing nothing at all. Useful eg. when just some repos need to
be cloned, but no actual build steps for those.

exec
----

Executes command given lines.

Those are given as subkeys of the `commands` key within the package config - one for each
stage: `prepare`, `configure`, `build`, `install`, `clean`.

Note that if arguments are given, they need to be written as YAML list - otherwise arguments
won't be split properly.

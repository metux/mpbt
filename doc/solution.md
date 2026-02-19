Solution configs
================

Solution configurations are defined in YAML files. There has to be exactly one
active solution per build run (eg. given via command line).

The solution object is the one which of the packages actually should be built,
as well as giving extra build options (eg. choosing versions, flags, system
specific settings). From here, virtual package names can be resolved to real ones.

Having multiple solutions is a good way to support separate, eg. target specific,
configurations while not having to duplicate all the package configurations.
For example, some solution can be configured to use different sets of package
from the host system (instead of building on its own), use different pathes, etc.

User provided keys:
-------------------

| Key               | Type             | Description                                                                            |
| ----------------- | ---------------- | ---------------------------------------------------------------------------------------|
| `install-prefix`  | string           | global installation prefix (as it will be on the target) for the whole bundle/solution |
| `package-mapping` | dict (see below) | maps virtual package names to real ones for depenency resolution                       |
| `build`           | string/list      | list of (virtual) package names to be built                                            |
| `packages`        | string/list      | list of directories to search package configs from                                     |
| `package-config`  | string/list      | extra per-package settings (will be copied into individual packages                    |
| `parallel`        | integer          | number of parallel jobs (defaults to ${@PROJECT::parallel})                            |


Automatic keys:
---------------

These keys are automatically filled in by MPBT and should not be overwritten.

| Key             | Description                               |
| --------------- | ----------------------------------------- |
| `@PROJECT`      | link to the global project config object  |


Package resolution:
-------------------

The table `package-mapping` is used for resolving virtual package names to real ones.
This allows different per-solution mappings, eg. when some targets should use certain
system libraries while others have to have their own copies, or when there are different
implementations to choose from.

For names omitted here, MPBT tries to match up automatically:

* matching against real package names
* if a virtual name is provided only by one real package, using this one

If some name (eg. from `build` list or some dependency) can't be matched either from the
mapping or automatically, MPBT aborts very early.


Per-Package extra configuration:
--------------------------------

The `package-config` key allows setting/overriding extra settings on per-package basis:
for each package, there may be a subkey underneath the `package-config` key, which is
holding a key-value list of the fields to set within the individual package. All keys
listed here are copied into the invidual package's config object, before actual build starts.

That way, solutions may override generic package setting, eg. versions, build flags, etc.

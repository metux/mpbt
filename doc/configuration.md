Configuration
=============

The MPBT configuration mainly consists of several pieces:

* project: just in-memory object that's filled by the config loader
* packages: arbitrary number of per-package configs, loaded from dedicated directories
* solutions: the central place for fitting all together (eg. per-target configurations)

All these objects are derived from generic "SpecObject"s, providing access
to ["Magic dictionary"](https://github.com/metux/go-magicdict). Those are
hierarchical data structures, nested key-value lists, usually loaded from
YAML files, but with extra functionality like variable substitution. Those
objects are also linked into each other, so their data can be referenced
from within each other (eg. packages can refer to keys within solution or
package, etc).

Project
-------

This object represents the build run as a whole. It's not loaded from some
YAML file (yet), but only filled internally, holding all runtime information
(eg. current pathes, machine information, etc, etc). Other objects, eg. the
solution, are linked into it (and vice-versa).

Users rarely need to directly with this object.

Packages
--------

For each configured package there's one package object. These are usually
loaded from a directory (eg. as configured within solution) - the "name"
field is automatically filled from the individual file name (the directory
prefix and file name suffix removed).

See [here](package.md) for more information.

Solution
--------

The solution object is the one which of the packages actually should be built,
as well as giving extra build options (eg. choosing versions, flags, system
specific settings). From here, virtual package names can be resolved to real ones.

Having multiple solutions is a good way to support separate, eg. target specific,
configurations while not having to duplicate all the package configurations.
For example, some solution can be configured to use different sets of package
from the host system (instead of building on its own), use different pathes, etc.

See [here](solution.md) for more information.

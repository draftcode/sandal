Sandal
======

Sandal is a fault-aware model checker for message passing systems.

Install
-------

Sandal follows Go's standard package structure. The following command will install Sandal compiler to $GOPATH/bin:

    $ go get github.com/draftcode/sandal

In addition to Sandal compiler, you need to install NuSMV to model-check a converted model that the compiler outputs.

Usage
-----

Sandal compiler converts a model written in Sandal into a NuSMV module, and write it to STDOUT. For example, the following command will convert your Sandal model to a NuSMV module and execute NuSMV:

    $ sandal your_model.sandal | nusmv

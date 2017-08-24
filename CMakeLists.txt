cmake_minimum_required(VERSION 3.0)
project(quicken)

include(go_cmake/bld/cmake/GolangSimple.cmake)

set(GO_NAMESPACE github.com/matteojoliveau/quicken)

GO_GET(cobra github.com/spf13/cobra)
GO_GET(yaml gopkg.in/yaml.v2)
GO_COPY(cmd cmd)
GO_COPY(git modules/git)
GO_COPY(license modules/license)
GO_COPY(readme modules/readme)
GO_COPY(recipe recipe)
GO_COPY(utils utils)

set(MODULES cmd git license readme recipe utils)

ADD_GO_INSTALLABLE_PROGRAM(qk # executable name
                           quicken.go # `package main` source file
                           cobra
                           yaml
                           ${MODULES}) # everything else is a dependency

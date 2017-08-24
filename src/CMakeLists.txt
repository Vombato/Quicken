GO_GET(cobra github.com/spf13/cobra)
GO_GET(yaml gopkg.in/yaml.v2)
GO_COPY(cmd cmd)
GO_COPY(git modules/git)
GO_COPY(license modules/license)
GO_COPY(readme modules/readme)
GO_COPY(recipe recipe)
GO_COPY(utils utils)

ADD_GO_INSTALLABLE_PROGRAM(qk # executable name
                           quicken.go # `package main` source file
                           cobra
                           yaml) # everything else is a dependency

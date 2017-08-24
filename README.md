# Quicken
Stimulate your development! Quicken is a command line tool to generate software project based on clear, repeatable recipes.
***WARNING*** Under development, not ready for usage

Quicken is a software project scaffolder written in Go.
It allows to create projects based on a "recipe" that defines the language, needed tools (build tools like Maven, Gradle or SCons, VCSs like Git or Mercurial), license and other parameters.
It uses templates to setup files.

Example recipe:
```yaml
---
name: Test
description: This is a test project for Quicken!
owner:
    name: Name Surname
    email: name.sur@example.com
    organization: Example
    url: example.com
language: java
build_tool:
   name: maven
   version: 3.4
git:
    init: true
    remote:
        name: origin
        url: https://github.com/matteojoliveau/quicken
license: mit
readme: true
```

## Example usage
If in a folder containing a recipe.yml file,

`qk.exe.exe init`

To specify a particular recipe:

`qk.exe.exe init --recipe path/to/recipe.yml`





# Quicken
Stimulate your development! Quicken is a command line tool to generate software project based on clear, repeatable recipes.
***WARNING*** Under development, not ready for usage

Quicken is a software project scaffolder written in Go.
It allows to create projects based on a "recipe" that defines the language, needed tools (build tools like Maven, Gradle or SCons, VCSs like Git or Mercurial), license and other parameters.
It uses templates to setup files.

Example recipe:
```yaml
---
language: java
build_tool: 
   name: maven
   version: 3.5
vcs: 
   name: git
license: mit
readme: true
```

## Example usage
If in a folder containing a recipe.yml file,

`qk init`

To specify a particular recipe:

`qk init --recipe path/to/recipe.yml`





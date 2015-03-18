gf [![Build Status](http://drone.targetpivot.com/api/badge/git.target.com/gophersaurus/gf/status.svg?branch=master)](http://drone.targetpivot.com/git.target.com/gophersaurus/gf)
==

The Gophersaurus project tool.

> Requires The Following:
  * Git installed
  * $GOPATH set

## Install

`go get git.target.com/gophersaurus/gf`

## Create New Project

The `gf` command line tool does the following when creating a new project:
  1. determine implied project `import` path with `$GOPATH`.
  2. compare implied `import` path with the `git remote origin` value if provided.
  3. pull source code from `git.target.com/gophersaurus/gophersaurus`.
  4. update project name to the name provided.
  5. rewrite `import` paths to match your new project's path.
  6. set `git remote upstream` to `git.target.com/gophersaurus/gophersaurus`.
  7. set `git remote origin` to the origin value provided, if any.

Create a new project based on project name:
* `gf new {project-name}`
* `gf n {project-name}`

Create a new project and set `git remote origin`:
* `gf new -origin="https://github.com/org/project.git" {project-name}`
* `gf n -o="https://github.com/org/project.git" {project-name}`

gf tool [![Build Status](http://drone.targetpivot.com/api/badge/git.target.com/gophersaurus/gf/status.svg?branch=master)](http://drone.targetpivot.com/git.target.com/gophersaurus/gf)
==

The Gophersaurus project tool.  Use this to start a new Gophersaurus project.

#Install

`go get git.target.com/gophersaurus/gf`

#New Project

Create a new project based on project name:
* `gf new {project-name}`
* `gf n {project-name}`

Create a new project with `git remote origin`:
* `gf new -git="https://github.com/org/project.git" {project-name}`
* `gf n -g="https://github.com/org/project.git" {project-name}`

Create a new project based on `remote origin`:
* `gf new -git="https://github.com/org/project.git"`
* `gf n -g="https://github.com/org/project.git"`

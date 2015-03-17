gf tool [![Build Status](http://drone.targetpivot.com/api/badge/git.target.com/gophersaurus/gf/status.svg?branch=master)](http://drone.targetpivot.com/git.target.com/gophersaurus/gf)
==

The Gophersaurus project tool.  Use this to start a new Gophersaurus project.

#Install

`go get git.target.com/gophersaurus/gf`

#New Project

Create a new project based on project name:
* `gf new {project-name}`
* `gf n {project-name}`

Create a new project and set `git remote origin`:
* `gf new -origin="https://github.com/org/project.git" {project-name}`
* `gf n -o="https://github.com/org/project.git" {project-name}`

Create a new project based on `git remote origin`:
* `gf new -origin="https://github.com/org/project.git"`
* `gf n -o="https://github.com/org/project.git"`

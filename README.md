# gloomo 

[![Go Reference](https://pkg.go.dev/badge/github.com/elemir/gloomo.svg)](https://pkg.go.dev/github.com/elemir/gloomo)
[![Build & Test](https://github.com/elemir/gloomo/actions/workflows/test.yaml/badge.svg)](https://github.com/elemir/gloomo/actions/workflows/test.yaml)

Gloomo is a data oriented rendering engine for ebiten. Also gloomo provides some standard containers and repositories useful in a DAO game.

## Phylosophy

Gloomo follows ideas of domain driven design, clean architecture from Bob Martin and data-oriented architecture. It uses next conceptions:

* Component 
* Model 
* Repository provide 
* System

### Rendering

Mostly rendering follows gloomo phylosophy 

* node.Node component 
* Nodes is a special type of models that contains special component 

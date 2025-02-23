# gloomo 

[![Go Reference](https://pkg.go.dev/badge/github.com/elemir/gloomo.svg)](https://pkg.go.dev/github.com/elemir/gloomo)
[![Build & Test](https://github.com/elemir/gloomo/actions/workflows/test.yaml/badge.svg)](https://github.com/elemir/gloomo/actions/workflows/test.yaml)

Gloomo is a data oriented rendering engine for ebiten. Also gloomo provides some standard containers and repositories useful in a DAO game.

## Phylosophy

Gloomo follows some ideas of domain driven design, clean architecture from Bob Martin and data-oriented architecture. It based on the next ideas:

1. Orthogonality of data and behaviour. None of application logical units should abstract both of them.
2. All objects that represents behaviour should be constructed on the start of application.
3. Objects that represents data may have some additional methods if they are quasi-linear and don't have side effects.
4. Behaviours part of application should be devided on different layers. Data objects should be used for passing between layers.

### Layers

It uses next conceptions:

* Component 
* Model 
* Repository provide 
* System

### Rendering

Mostly rendering follows gloomo phylosophy 

* node.Node component 
* Nodes is a special type of models that contains special component 

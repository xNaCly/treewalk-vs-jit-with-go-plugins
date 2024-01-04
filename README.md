# Comparing Tree-walk Interpreters, Bytecode VMs and JIT compilation with embedding via Go-plugins

> Evaluating the trade-offs of using the Go-plugin API for JIT compilation while comparing the approach with common language evaluation patterns.

## Abstract

The goal of this paper is to evaluate whether or not the usage of the Go
plugin API is feasible for just-in-time compilation of a query language.
This evaluation is done based upon the criteria of the ease of usability,
performance and the robustness of the resulting implementation. For sake of
comparison a minimalist programming language is introduced. A tree-walk
interpreter, a bytecode compiler, the corresponding virtual machine and a
compiler for generating Go source code are implemented. The paper explores
the different possibilities for accessing the Go compiler, working with the
Go plugin API and highlights several benchmarks comparing the performance
of the three aforementioned language evaluation implementations.

## Project structure

The implementation, interpreter-backends (tree-wak, bytecode, go jit) and benchmarks are located in `src/`. The academic paper can be found in `paper/`.

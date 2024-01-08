# Comparing a Tree-walk Interpreter with JIT compilation and embedding via Go-plugins

> Evaluating the trade-offs of using the Go-plugin API for JIT compilation
> while comparing the approach with a tree-walk interpreter

## Abstract

The goal of this paper is to evaluate whether or not the usage of the Go
plugin API is feasible for just-in-time compilation of a query language
intended for a high performance in memory data store. This evaluation is
done based upon the criteria of the ease of usability, performance and the
robustness of the resulting implementation. For sake of comparison the
query language is introduced. The code a JIT would have generated for
several heavy queries is handwritten and benchmarked agains the same
expressions evaluated via the current tree walk interpreter. The paper
explores the different possibilities for accessing the Go compiler, working
with the Go plugin API and highlights several benchmarks comparing the
performance of the new JIT compiler and the previous language evaluation
implementation.

## Project structure

The implementation, interpreter-backends (tree-wak, bytecode, go jit) and
benchmarks are located in `src/`. The academic paper can be found in `paper/`.

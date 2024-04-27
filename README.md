# Comparing a Tree-walk Interpreter with JIT compilation and embedding via Go-plugins

> Evaluating the trade-offs of using the Go-plugin API for JIT compilation
> while comparing the approach with a tree-walk interpreter

## Abstract

The goal of this paper is to evaluate whether the usage of the Go plugin
API is feasible for just-in-time compilation of a query language intended
for a high performance in memory data storage. This evaluation is done
based upon the criteria of the ease of usability, performance and the
robustness of the resulting implementation. For the sake of comparison the
query language as well as its features are introduced. A just in time
compiler is implemented and benchmarked against the same expressions
evaluated with the currently employed tree walk interpreter. The paper
explores the different possibilities for accessing the Go compiler, working
with the Go plugin API and highlights several benchmarks comparing the
performance of the new JIT compiler and the previous language evaluation
implementation.

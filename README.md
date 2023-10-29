
# ml-data-builder

An efficient Golang utility that simplifies the process of generating Machine Learning Datasets from publicly available Social Network APIs. This tool is quite helpful with managing concurrent operations and preserving data, all while offering an intuitive API for straightforward implementation.

## Overview

Public Social Network APIs are frequently leveraged to build datasets used in training Machine Learning models. Models like [Tweet2Vec](https://arxiv.org/abs/1607.07514) that use this type of data to extract features or create embeddings are quite common. This tool is particularly useful for NLP-oriented models that can benefit from a large repository of structured (or even unstructured) text data.

Users often resort to finding specific publicly available general datasets or spend considerable time on setting an elaborate feature extraction pipeline which can detract time from actual feature engineering and model work. This tool aims to simplify such chores.

Despite being a newcomer to the Machine Learning space, valuable feedback is appreciated and taken seriously. The goal is to provide a genuinely useful tool that streamlines the process of ML data set creation.

## Roadmap

* ✔ Top-level Feature-based API

* ✔ Manage concurrency using Goroutines, channels, and other advanced techniques

* Upcoming features:

   * Caching operations to prevent redundant requests

   * Save functionality for different data formats

   * Multi-format API data support

   * Authentication support

   * Command-line functionality

* Additional feature suggestions are welcome!
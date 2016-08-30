/*
<!--
Copyright (c) 2016 Christoph Berger. Some rights reserved.
Use of this text is governed by a Creative Commons Attribution Non-Commercial
Share-Alike License that can be found in the LICENSE.txt file.

The source code contained in this file may import third-party source code
whose licenses are provided in the respective license files.
-->

<!--
NOTE: The comments in this file are NOT godoc compliant. This is not an oversight.

Comments and code in this file are used for describing and explaining a particular topic to the reader. While this file is a syntactically valid Go source file, its main purpose is to get converted into a blog article. The comments were created for learning and not for code documentation.
-->

+++
title = "Random thoughts"
description = "Random generators and their usage in Go"
author = "Christoph Berger"
email = "chris@appliedgo.net"
date = "2016-09-08"
publishdate = "2016-09-08"
domains = ["Algorithm and Data Structures"]
tags = ["random", "generator", ""]
categories = ["Tutorial"]
+++

How to generate random numbers, and the difference between math/rand and crypt/rand.

<!--more-->

## Generating randomness on deterministic machines

The ideal computer is completely deterministic. For every input, the output is foreseeable. Trying to generate random data on such a machine seems odd at first. But there are two ways to overcome the dichotomy between determinism and randomness.


### Exploiting sources of real randomness

Real-world computers are not quite the ideal machines that computer sciences would like them to be. There are a lot of sources in every computing device that produce more or less random data. Mouse movement, time between two keystrokes, the wall clock, counters like actual CPU usage, GPS receivers, movement sensors, and more.

There are even parts that can produce true random data. Electronic elements like transistors, diodes, or resistors can generate static noise. Turn up an amplifier that has no input device attached, and you hear - static noise. And static noise is truly random information. An analog/digital converter can turn the noise into never-ending sequences of random bits.

Still, these "natural" sources of random data suffer from asymmetries and systematic biases caused by various physical phenomena that are inherent to the chosen source. As a consequence, the generated random numbers are not uniformly distributed. Luckily, there are functions called "randomness extractors" that can fix this.


### Generating pseudo-random numbers

The second way is to simulate a source of random data. But how, if "randomness" is not part of the concept of a deterministic machine? The trick is to produce long sequences of bits and bytes that *appear* to be random. After a while, the sequence repeats, but for many consumers of random data this is perfectly fine.


## Go's `rand` packages


### math/rand


### crypt


## The code
*/

// ## Imports and globals
package main

/*
## Odds and Ends

### The code



### Links

[Randomness extractor](https://en.wikipedia.org/wiki/Randomness_extractor)

[/dev/random](https://en.wikipedia.org/wiki//dev/random)

[Yarrow algorithm](https://en.wikipedia.org/wiki/Yarrow_algorithm)


*/

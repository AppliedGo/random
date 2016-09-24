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
title = "A Random Blog Post"
description = "Random generators and their usage in Go"
author = "Christoph Berger"
email = "chris@appliedgo.net"
date = "2016-09-29"
draft = "true"
publishdate = "2016-09-29"
domains = ["Algorithm and Data Structures"]
tags = ["random", "mathematics", "cryptography"]
categories = ["Tutorial"]
+++

How to generate random numbers, and the difference between math/rand and crypto/rand.

<!--more-->

## Generating randomness on deterministic machines

The ideal computer is completely deterministic. For every input, the output is foreseeable. Trying to generate random data on such a machine seems odd at first. But there are two ways to overcome the dichotomy between determinism and randomness.


### Exploiting sources of real randomness

Real-world computers are not quite the ideal machines that computer sciences would like them to be. There are a lot of sources in every computing device that produce more or less random data. Mouse movement, time between two keystrokes, the wall clock, counters like actual CPU usage, GPS receivers, movement sensors, and more.

{{< figure src="/media/random/noise.png" class="imageLeft" alt="" >}}Electronic parts can also produce true random data. Transistors, diodes, or resistors can generate static noise. Turn up an amplifier that has no input device attached, and you hear - static noise. Tune an AM or FM radio between two stations, and you get - static noise.  And static noise is truly random information. An analog/digital converter can turn the noise into never-ending sequences of random bits.

Still, these "natural" sources of random data suffer from asymmetries and systematic biases caused by various physical phenomena that are inherent to the given source. Thus the generated random numbers are not uniformly distributed. Luckily, there are functions called ["randomness extractors"](https://en.wikipedia.org/wiki/Randomness_extractor) that can fix this, at the cost of a lower output rate.

### Generating pseudo-random numbers

The second way is to simulate a source of random data. But how, if "randomness" is not part of the concept of a deterministic machine? The trick is to produce long sequences of bits and bytes that *appear* to be random. After a while, the sequence repeats, but for many consumers of random data this is perfectly fine.

A very simple pseudo-random generator is a bit shift register with a feedback loop.

* At each clock cycle, all bits in the register are shifted to the right.
* The rightmost bit is added to the outgoing bit stream.
* At two (arbitrary) positions of the register, the bits are extracted and fed into an Exclusive-OR (XOR) gate. The result of the XOR operation is fed back to the first bit of the register.

(To recap: The XOR operation returns "true" if both input values are different, and "false" otherwise.)

If this sounds a bit too abstract, watch the animation below. The values "true" and "false" are represented by "1" and "0", respectively.

!HYPE[Bit Shift Register](bitshift.html)

At some point in time, however, the register contains a value that it contained earlier, and at this point, the cycle repeats.o

**In a generalized sense, this is how all pseudo-random number generators work. A deterministic algorithm produces a long series of seemingly random bits and bytes. Eventually the series will repeat, but depending on the algorithm, the output may still successfully pass various tests for randomness.**

## Go's `rand` packages

Go has two packages that generate random numbers: `math/rand` and `crypto/rand`. Now with the above in mind you surely already have an idea why there are two of them: One is a pseudo-random number generator, the other makes use of a real source of random data (provided by the operating system).

But yet - why do we need both? Can't we just use `crypto/rand` for everything and enjoy truly random numbers for all purposes?

You guessed it already: The answer is speed. True sources of randomness can produce only so many bits at a time. And the aforementioned randomness extractor reduces the throughput even more.

So unless you really need cryptographically secure random numbers, use `math/rand`. It will be sufficient for most of your needs.

Now let's have a quick look into both of these packages.

### math/rand



### crypto/rand

[/dev/random](https://en.wikipedia.org/wiki//dev/random)

## The code
*/

// ## Imports and globals
package main

/*
## Odds and Ends

### The code



### Links

[Randomness extractor](https://en.wikipedia.org/wiki/Randomness_extractor)

[Randomness tests](https://en.wikipedia.org/wiki/Randomness_tests)

[/dev/random](https://en.wikipedia.org/wiki//dev/random)

[Yarrow algorithm](https://en.wikipedia.org/wiki/Yarrow_algorithm)


*/

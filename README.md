Pathfinder
===

This is a tool written in Go to aid me in my Pathfinder campaign.

Features
---
Automatically generate your assets in any size! Woo!

Possible Future Features
---
When we get an actual database, it'll come out really purdity
I forsee a web interface for designing tiles. Creating generic events, etc...
Even output as google maps type thing.

Concurrency
---
The world generates with as many threads as your system can handle.

So, it seperates the world tiles into an even amount of squares, generates each space as its own area, then stiches those areas together.
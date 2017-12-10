Pathfinder
===

This is a tool written in Go to aid me in my Pathfinder campaign.



Concurrency
---
The world generates with as many threads as your system can handle.

So, it seperates the world tiles into an even amount of squares, generates each space as its own area, then stiches those areas together.
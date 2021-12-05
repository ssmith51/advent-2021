# Advent Day 05: Hyrdothermal Venture
You come across a field of hydrothermal vents on the ocean floor! These vents constantly produce large, opaque clouds, so it would be best to avoid them if possible.

## Puzzle 1
This one took me a few minutes to figure out. I decided to use a Multi Dimensonal Slice, which seemed to work well for me. 

The main issue I faced was in my interpertation of the problem (I'll blame that on being sleepy from Re:Invent). When I first started working on it I was missing the fact these were line segments and assumed you would count every number between coordinates. This was a fatal flaw and it wasn't till I finished what I thought was the correct solution did I realize it. 

Once realized it took me a while to reset, and the constant reminder that a Grid is read in X,Y but a Mutli Dimensonal Slice is indexed via Y,X led to much head banging (again, I'll blame it on Re:Invent travel lag :) ).

Overall, I feel there should be an easier way to do this than my solution. Since we're essentially dealing with Vectors I assume there is a library out there I could pull in to solve this in much less code. 

## Puzzle 2

## Summary
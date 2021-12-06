# Advent Day 06: Hyrdothermal Venture
The sea floor is getting steeper. Maybe the sleigh keys got carried this way?

A massive school of glowing lanternfish swims past. They must spawn quickly to reach such large numbers - maybe exponentially quickly? You should model their growth rate to be sure.

## Puzzle 1
This one took me a few minutes to figure out. I decided to use a Multi Dimensonal Slice, which seemed to work well for me. 

The main issue I faced was in my interpertation of the problem (I'll blame that on being sleepy from Re:Invent). When I first started working on it I was missing the fact these were line segments and assumed you would count every number between coordinates. This was a fatal flaw and it wasn't till I finished what I thought was the correct solution did I realize it. 

Once realized it took me a while to reset, and the constant reminder that a Grid is read in X,Y but a Mutli Dimensonal Slice is indexed via Y,X led to much head banging (again, I'll blame it on Re:Invent travel lag :) ).

Overall, I feel there should be an easier way to do this than my solution. Since we're essentially dealing with Vectors I assume there is a library out there I could pull in to solve this in much less code. 

## Puzzle 2
Trig! It's been a while since I've looked at anything related to Trigonometry! Fortunately, this second portion was fairly easy. Since I had mistakenly coded for diagonals as part of the bingo puzzle I already had most of the logic ready to go. Again, the hardest part of the problem was the Y,X Multi Dimensonal slice and making sure the correct points were marked. 

## Summary
This took me much longer than I anticpiated, however once I understood the expected outcome it was fairly easy to complete. Continuing with my logic, it feels like there should be an easier way to do this than Multi Dimensonal arrays and a bunch of loops...
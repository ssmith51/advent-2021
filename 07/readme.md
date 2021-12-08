# Advent Day 07: The Treachery of Whales
A giant whale has decided your submarine is its next meal, and it's much faster than you are. There's nowhere to run!

Suddenly, a swarm of crabs (each in its own tiny submarine - it's too deep for them otherwise) zooms in to rescue you! They seem to be preparing to blast a hole in the ocean floor; sensors indicate a massive underground cave system just beyond where they're aiming!

The crab submarines all need to be aligned before they'll have enough power to blast a large enough hole for your submarine to get through. However, it doesn't look like they'll be aligned before the whale catches you! Maybe you can help?

## Puzzle 1
This was a puzzler... thinking on the last problem I realized I would need to think more fundamental math than I have worked recently (cough... trig... cough...). Quickly found that the Median was what I was looking for. Sadly, I looked over the fact that fuel was caculated off the Median and was not the Median itself. After realizing that, solution was simple. 

## Puzzle 2
And, you have now broken me Advent of Code. It took [Scott Morris](https://github.com/scott-morris) walking me through some basic logic to realize what to do. I even looked at his JS to see what logic he was using. After my drive home I realized the following: 

1. We can't easily calculate the fastest soltuion at once
2. We can determine a reasonable range to look at to reduce our fuel calcualtions
3. We need to start with our answer from puzzle one (Median in my case) and use the Standard Deviation off that to get a small range to calcuation (i.e. don't brute force the entire list). 

Once that clicked, I took 1 Standard Deviation above and below, calculated the fuel costs for each potential, and took the lowest fuel cost of all of those results. 

Oh, I also had my factorial logic wrong... I'd like to blame that on distractions but it was really just a coding mistake... :) 

## Summary
I'm a little worried about how long I can go till I get stumped on the math,not the coding, with AoC. Hopefully, between those competing with me and a bunch of rubber ducking, we can all make it to X-Mas :) 

Additionally, this is the fist puzzle I pulled in an external library to reduce the amount of code I need to write: [Stats](https://pkg.go.dev/github.com/montanaflynn/stats). I'm not ashamed by that! Kudos to the developer for publishing that libaray! 
# Advent Day 06: Hyrdothermal Venture
The sea floor is getting steeper. Maybe the sleigh keys got carried this way?

A massive school of glowing lanternfish swims past. They must spawn quickly to reach such large numbers - maybe exponentially quickly? You should model their growth rate to be sure.

## Puzzle 1
This seemed too easy, we have a bunch of fish, each fish counts down to 0 at which point we rotate that fish back to position 6 and create a new fish that starts counting down from 8 -> 0. When those new fish give birth at position 0 they go to position 6, and a new fish get created at 8, and so on. 

However, the word __exponentially__ stood out (was highlighted too). I guessed some type of hooked in the second puzzle would change, such as fush reproduce on different days and grow faster etc... 

I thought about it some and realized all we care about are the number of fish per day in their 'cycle'. I mapped it out by the days in the cycle, some fish started at the 2nd day in the cycle, 3rd day, etc.. Then I started movving them closer to day 0 in the cyle, at which point they reproduce, set the count of the fish on day 8 to the number on day 0 and increase the number of fish on day 6 of the 'cycle' by the number of fish in day 0 of the 'cycle. 

The biggest I ran into was adding the new fish at day 6 & 8 after moving fish down the day cycle. I had to pull the number at the start then move them down. 

## Puzzle 2
Boom! Change the total number od days and this worked like a champ!

## Summary
The biggest part was figuring out how to not brute force this. I feel the last few I've been brute forcing them so I spent some time thinking about it. I thougt there would be some type of log factor I could use per fish in the initial list but there wasn't. It also helped that I was reading on Friday a lot of people saying future problems may run a long time and try to exhaust resources. The __exponential__ keyword hinted at that. 
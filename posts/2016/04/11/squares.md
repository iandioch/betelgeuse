---
title: 'Project Euler #135 & #136 Postmortem'
categories: ["updates", "reports"]
tags: ["maths", "programming", "project euler", "java"]
---

I've been doing some Project Euler problems in the past few days, and regained my proud position in the top 20 in Ireland. I'm currently in joint 18th place, with a noble 98 problems solved. So close to that magic 100.

My most recent conquests were #135 and #136, done this morning.

I wanted to try some mathsy ones, so I looked at these. [Here](https://projecteuler.net/problem=135) is #135:

> Given the positive integers, x, y, and z, are consecutive terms of an arithmetic progression, the least value of the positive integer, n, for which the equation, x2 − y2 − z2 = n, has exactly two solutions is n = 27:
>
> 342 − 272 − 202 = 122 − 92 − 62 = 27
>
> It turns out that n = 1155 is the least value which has exactly ten solutions.
>
> How many values of n less than one million have exactly ten distinct solutions?

It looks almost brute forceable! But I decided to do some digging first. The most important line I almost overlooked was the fact that `x`, `y`, and `z` are in an arithmetic sequence, so they have a common difference. Knowing this, I could represent them in terms of this difference `d`:

`z + d = y`

`y + d = x`

`z + 2d = x`

The original equation is `x^2 - y^2 - z^2 = n`, which can be rewritten with the above information as `(z + 2d)^2 - (z + d)^2 - (z)^2 = n`, or `-z^2 + 2dz + 3d^2 = n`. We only have two variables to worry about here, `z` and `d`, so we can start iterating and finding answers!

\#135 asks for the number of `n`s below one million with exactly 10 solutions, so it's simple to create an array of integers of length one million, initialised to zero, and start iterating over `z` and `d` to find matching solutions, and increment that array element.

At the end, all that needs to be done is count over that array and find the number of indexes that have the value 10.

However, it is not quick.

Many of the `z` and `d` pairs that I was iterating over were not suitable. The expression `-z^2 + 2dz + 3d^2` was often evaluating to be less than zero, or over our maximum limit of 1,000,000. This is wasted effort. I tried to find the the `z` which, given a certain `d`, would cause the expression to give negative numbers from then on, so we could safely break out of the loop then and skip the lost causes. I tried factorising the equation, failed miserably at doing the simple maths, and so my beloved [@CianLR](http://github.com/CianLR) off the top of his head told me that `3d^2 + 2dz - z^2 = 0` factorised to `(3d - z)(d + z) = 0`. Seeing as `(d+z) = 0` results in `z` and `d` being of opposite signs, and therefore never occurring in my loops, I knew that the tipping point was when `d = z/3`. 

A few off-by-one errors later, and I had a solution that works great for both #135, and the almost identical #136 (which just has a maximum limit of 50 million instead of 1 million), running in 7.286s and 16.090s respectively.

My Java source code is [here for #135](https://github.com/iandioch/solutions/blob/master/project_euler/135/Solution.java) and [here for #136](https://github.com/iandioch/solutions/blob/master/project_euler/136/Solution.java).

May the spirits of the maths gods be with you,

Noah

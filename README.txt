CSE486 Capstone Team Assignment Program
=======================================

Created by Shantanu Bala (shantanu@asu.edu) as an honors contract assignment for CSE486 under
Dr. Debra Calliss. The program may be used to radomly assign teams to projects
for the CSE485/CSE486 capstone class at Arizona State University. The compiled
executable will take a CSV file as input with the project preferences of
teams. According to the most and least preferred projects of teams, the
program will attempt to create an optimal assignment of teams to projects.


Procedure Followed by Program
=============================

The program executes the following procedure:

1. Select teams in a random order to receive their first pick
2. If the team's first pick is not available, attempt second or third pick
3. If none of the teams preferred projects are available, select randomly
4. Ensure that no team receives a project that is in their least preferred list
5. Output a CSV file with team assignments (with overall statistics)


CSV Data Formatting
===================

The structure of the input data is most easily understood by opening the sample `groups.csv` file. In it, you will find a list of teams. The first three columns following the team name are the "top 3 picks" in order, and the last three columns are the "bottom 3 picks" in order.

After a single blank space, a full list of projects is also included in the first column after the teams.

The total number of teams and projects is flexible -- any number can be used. For within-team assignment, the input data is also flexible. You can include only the "top 2" and "bottom 2" picks -- the program will simply split the selections by half of the total number. For example:

lambda	p1	p4	p5	p6	p0	p2

This indicates that p1, p4, and p5 are the top 3 picks for team lambda, and p6, p0, and p2 are the bottom three picks.

If only two selections are made, the row will look like:

lambda	p1	p4	p5	p6

In this case, p1 and p4 are interpreted as the top 2 picks, and p5 and p6 are the bottom 2 picks for team lambda.

This allows for complete flexibility over the number of teams, the number of selections per team, and the total number of projects.

Compiling the Source
=====================

The code is written in Go (https://golang.org/), a programming language made
at Google. To compile, simply install Go and run the `go build` command inside
of the directory containing the source code file `assign.go`. See the web
page for more details: https://golang.org/cmd/go/


Using Precompiled Binary
========================

A precompiled EXE file for Windows is included. To run, simply execute the
following command:

	./assign-groups.exe ./groups.csv

In the aforementioned commnd, the EXE file is the included compiled binary. The CSV file is the included list of sample groups for assignment.

The output of the program is stored in `output.csv` in the same folder, but can also be accessed within the terminal window itself.

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

# multifasta

A quick (*bad*) rewrite of a shell script to combine files to a multi
fasta format to Go.

## Purpose

An original super basic shell based tool written by someone else `01_example_output.fasta`
to join together multiple fasta style files into a single fasta file with > headers based on
their original file names

This was then rewritten to be Go based because

* An excuse to force myself to write more Go
* Be compiled to an exe that can be more easily run on windows

## Usage

Basic usage is to just run it & pass any number as the input. Output will then go to a new file
with the default name of `multifasta_output_YYYYMMDD-HHMMSS.fasta` in the same directory as the
first input file (input files are sorted cause windows was giving it a weird order). There is an
`-output` option to name your own output file or use `-` for STDOUT

The use of the default output file name allows for one to just drag a selection of files onto the
.exe in windows. This further allows the .exe to be added to the "Send To" collection in windows to
then just right click & select & select "Send to multifasta" to then generate your combined file

# License
Licensed under the MIT license, see LICENSE for full details

#!/bin/bash

set -e
echo $'\n'"#############################"$'\n'"### MultiFasta File Maker ###"$'\n'"#############################"$'\n'
folder="Inputfiles"
multifasta="Multi Fasta Output"$'\n'

echo "What is the name of the folder with your sequence files?"$'\n'"Press Enter for default ('Input Files')"
read folder				## read custom folder name for input files
folder=${folder:-"Input files"}		## use custom name or (default) "Input files" if no custom name used
echo "What would you like to call the new MultiFasta file?"$'\n'"Press Enter for default ('Multifasta_output.txt')"
read multifasta				## read custom name for output file name
multifasta=${multifasta:-"Multifasta_output"}	 ## use custom name as variable or (default) "Multifasta_output"

cd "$folder"	## cd to folder with seq files (not sure if it would be more efficient to just include the pathname to read the files rather than changing directory)
seq_file_names=(*.seq)			## assign array of file names for all .seq files

for ((i=0;i<"${#seq_file_names[@]}";i++)); 	##begin loop
do
  seq_value=$(<${seq_file_names[$i]})		## extract raw sequence from input file "i"

  sequence_label=${seq_file_names[$i]#*_}		## extract sequence name from input file "i"
  sequence_label=${sequence_label%.seq}		## generate sequence name 

  fasta=">$sequence_label sequence exported from ${seq_file_names[$i]}\r\n$seq_value\n" ## restructure input file data into multi-fasta format, eol = \r\n to improve readability in various external progs

  printf "$fasta" >> "$multifasta.txt"		## write next line to multifasta output file 
done


echo ""
echo ""
echo "Your merged fasta file '$multifasta.txt'"
echo ""
echo "is in the ${folder} folder."

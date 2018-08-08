#!/bin/bash

seq 1 1000000 | while read -r line
do
	echo $line 1>&2
done

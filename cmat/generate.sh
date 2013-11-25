#!/bin/bash

dir=../mat
gofiles=.gofiles

ls $dir | grep '\.go$' >$gofiles
for f in `cat $gofiles`
do
	echo $dir/$f -\> ./$f
	cp $dir/$f ./
	go fmt $f >/dev/null
	sed 's/float64/complex128/g' $f         > tmp && mv tmp $f
	sed 's/^package mat$/package cmat/g' $f > tmp && mv tmp $f
	go fmt $f >/dev/null
	chmod a-w $f
done

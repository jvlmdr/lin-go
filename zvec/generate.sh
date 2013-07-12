#!/bin/bash

dir=../vec
gofile=.gofiles

ls $dir | grep '\.go$' | grep -v '_real[_.]' >$gofile
for f in `cat $gofile`
do
	cp $dir/$f .
	go fmt $f >/dev/null
	sed -i '' 's/float64/complex128/g' $f
	sed -i '' 's/^package vec$/package zvec/g' $f
	go fmt $f >/dev/null
	chmod a-w $f
done

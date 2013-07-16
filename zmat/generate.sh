#!/bin/bash

dir=../mat
gofile=.gofiles

vec='github.com\/jackvalmadre\/lin-go\/vec'
zvec='github.com\/jackvalmadre\/lin-go\/zvec'

ls $dir | grep '\.go$' | grep -v '_real[_.]' >$gofile
for f in `cat $gofile`
do
	echo $dir/$f -\> ./$f
	cp $dir/$f ./
	go fmt $f >/dev/null
	sed -i 's/float64/complex128/g' $f
	sed -i 's/^package mat$/package zmat/g' $f
	sed -i "s/\"$vec\"/\"$zvec\"/g" $f
	sed -i "s/vec\./zvec\./g" $f
	go fmt $f >/dev/null
	chmod a-w $f
done

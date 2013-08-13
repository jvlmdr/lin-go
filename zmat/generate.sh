#!/bin/bash

dir=../mat
gofiles=.gofiles

vec='github.com\/jackvalmadre\/lin-go\/vec'
zvec='github.com\/jackvalmadre\/lin-go\/zvec'

ls $dir | grep '\.go$' | grep -v '_real[_.]' >$gofiles
for f in `cat $gofiles`
do
	echo $dir/$f -\> ./$f
	cp $dir/$f ./
	go fmt $f >/dev/null
	sed 's/float64/complex128/g' $f         > tmp && mv tmp $f
	sed 's/^package mat$/package zmat/g' $f > tmp && mv tmp $f
	sed "s/\"$vec\"/\"$zvec\"/g" $f         > tmp && mv tmp $f
	sed "s/vec\./zvec\./g" $f               > tmp && mv tmp $f
	go fmt $f >/dev/null
	chmod a-w $f
done

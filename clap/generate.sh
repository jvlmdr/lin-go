#!/bin/bash

dir=../lapack
gofiles=.gofiles

ls $dir | grep '\.go$' >$gofiles
for f in `cat $gofiles`
do
	echo $dir/$f -\> ./$f
	cp $dir/$f ./$f
	go fmt $f >/dev/null
	sed 's/float64/complex128/g' $f           > tmp && mv tmp $f
	sed 's/^package lapack$/package clap/g' $f > tmp && mv tmp $f

	sed -E "s/dsy(ev)/zhe\1/g" $f             > tmp && mv tmp $f
	sed -E "s/d(ge|po|sy)([a-f]*)/z\1\2/g" $f > tmp && mv tmp $f
	sed -E "s/D(GE|PO|SY)([A-Z]*)/Z\1\2/g" $f > tmp && mv tmp $f
	sed "s/doublereal/doublecomplex/g" $f > tmp && mv tmp $f
	go fmt $f >/dev/null
done

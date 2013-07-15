#!/bin/bash

realfiles=.realfiles

vec='github.com\/jackvalmadre\/lin-go\/vec'
zvec='github.com\/jackvalmadre\/lin-go\/zvec'
mat='github.com\/jackvalmadre\/lin-go\/mat'
zmat='github.com\/jackvalmadre\/lin-go\/zmat'

ls | grep '\.go$' | grep '_real[_.]' >$realfiles
for d in `cat $realfiles`
do
	z=`echo $d | sed 's/_real\([_.]\)/_cmplx\1/g'`
	echo $d -\> $z
	cp $d $z
	go fmt $z >/dev/null
	sed -i '' 's/float64/complex128/g' $z
	sed -i '' "s/\"$vec\"/\"$zvec\"/g" $z
	sed -i '' "s/\"$mat\"/\"$zmat\"/g" $z
	sed -i '' "s/vec\./zvec\./g" $z
	sed -i '' "s/mat\./zmat\./g" $z
	sed -i '' "s/Solve\([(A-Z]\)/SolveComplex\1/g" $z
	sed -i '' "s/RealLU/ComplexLU/g" $z
	sed -i '' "s/dgesv/zgesv/g" $z
	sed -i '' "s/DGESV/ZGESV/g" $z
	sed -i '' "s/doublereal/doublecomplex/g" $z
	go fmt $z >/dev/null
	chmod a-w $z
done

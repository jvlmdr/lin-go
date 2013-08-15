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
	sed 's/float64/complex128/g' $z > tmp && mv tmp $z
	sed "s/\"$vec\"/\"$zvec\"/g" $z > tmp && mv tmp $z
	sed "s/\"$mat\"/\"$zmat\"/g" $z > tmp && mv tmp $z
	sed "s/vec\./zvec\./g" $z > tmp && mv tmp $z
	sed "s/mat\./zmat\./g" $z > tmp && mv tmp $z
	sed -E "s/Solve(FullRank|Square|Cond|Symm|PosDef)(Mat)?(NoCopy)?\(/Solve\1\2Cmplx\3(/g" $z \
		> tmp && mv tmp $z
	sed -E "s/(LU|LDL)(NoCopy)?\(/\1Cmplx\2(/g" $z > tmp && mv tmp $z
	sed -E "s/(LU|LDL)Fact/\1FactCmplx/g" $z > tmp && mv tmp $z
	sed -E "s/d(ge|po|sy)([a-z]*)/z\1\2/g" $z > tmp && mv tmp $z
	sed -E "s/D(GE|PO|SY)([A-Z]*)/Z\1\2/g" $z > tmp && mv tmp $z
	sed "s/doublereal/doublecomplex/g" $z > tmp && mv tmp $z
	go fmt $z >/dev/null
	chmod a-w $z
done

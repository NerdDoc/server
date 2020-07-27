#!/usr/bin/env bash

/usr/bin/lscpu |grep 'Arch'| grep 'arm' &> /dev/null
if [ $? == 0 ]; then
	export LD_LIBRARY_PATH="$(pwd)/res/lib/"
	ln -sf $(pwd)/res/libs/arm $(pwd)/res/lib
	bash -i
else 
	export LD_LIBRARY_PATH="$(pwd)/res/lib/"
	ln -sf $(pwd)/res/libs/x86 $(pwd)/res/lib
	bash -i
fi



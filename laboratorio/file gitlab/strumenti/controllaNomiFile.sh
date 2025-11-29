#!/bin/bash

#https://stackoverflow.com/questions/58637568/for-loop-in-linux-treats-pattern-as-filename-when-no-files-exist
shopt -s nullglob # expands the glob to empty string when there are no matching files in the directory.

if
    test "$#" -ne 2
then
    echo "Usage: $0 <path consegne (già unzippate)> <file elenco nomi>"
    exit 1
fi

DIRCONSEGNE=$1
NOMI=$2

if
    ! test -d $DIRCONSEGNE
then
    echo "$DIRCONSEGNE non accessibile"
    exit 2
fi

if
    ! test -s $NOMI
then
    echo "$NOMI non esistente o vuoto"
    exit 3
fi

################################################################################
# cerca ogni go
# verifica se presente in NOMI
find $DIRCONSEGNE -name '*.go' | grep -v _test | sort | while
    read sorgente
do
    BASENAME=$(basename $sorgente)

    if
        grep -q $BASENAME $NOMI
    then
        echo "$sorgente: nome OK!"
    else
        echo "$sorgente: NON presente nella lista dei nomi!!!"
    fi
done
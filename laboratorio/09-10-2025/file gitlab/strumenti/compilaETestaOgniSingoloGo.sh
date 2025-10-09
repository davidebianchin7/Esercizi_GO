#!/bin/bash

#https://stackoverflow.com/questions/58637568/for-loop-in-linux-treats-pattern-as-filename-when-no-files-exist
shopt -s nullglob # expands the glob to empty string when there are no matching files in the directory.


# verifica che vengano passati DUE parametri
if
    test "$#" -ne 2
then
    echo "Usage: $0 <path consegne (già unzippate)> <path bbtest corrispondente>"
    exit 1
fi

DIRCONSEGNE=$PWD/$1
DIRBBTEST=$PWD/$2




# verifica che i DUE parametri siano directory e esistano
if
    ! test -d $DIRCONSEGNE
then
    echo "$DIRCONSEGNE non accessibile"
    exit 2
fi

if
    ! test -d $DIRBBTEST
then
    echo "$DIRBBTEST non accessibile"
    exit 3
fi




################################################################################
# esegue i BB test se ci sono
bbtest() {
    #echo TESTS...
    for testinput in $DIRBBTEST/$EXENAME.*.input; do
        #echo $testinput

        expected=${testinput//input/expected}
        #echo $expected
        #ls -l

        # fa il `diff` tra l'output dell'esecuzione del programma (con redirezione input) e l'out expected
        diff -y -W 180 <(timeout 5 ./$EXENAME <$testinput | head -n 200) $expected
    done
}

################################################################################
# cerca ogni go
# lo compila singolarmente
# a partire da dir corrente
# e applica i test blackbox se esistono
find $DIRCONSEGNE -name '*.go' | grep -v _test | sort | while
    read sorgente
do
    DIRNAME=$(dirname $sorgente)
    BASENAME=$(basename $sorgente)
    EXENAME=$(basename $sorgente .go)
    #echo $DIRNAME - $BASENAME

    # entra nella dir
    pushd $DIRNAME >/dev/null

    echo '---'
    echo "# $sorgente"

    # https://stackoverflow.com/questions/74522713/what-are-possible-go-build-ldflags-options


    # compila
    go build $BASENAME &> $BASENAME.compile

    if
        test -s $BASENAME.compile # se l'out della compilazione non è vuoto ci sono ERRORI COMPILAZIONE
    then
        echo "## errori di compilazione"
        echo '```'
        cat $BASENAME.compile
        echo '```'
        mv $BASENAME.compile $BASENAME.compile.ERR
    else
        rm $BASENAME.compile
        #echo "# $DIRNAME"
        echo '```'
        bbtest
        echo '```'
    fi

    # rimuove eseguibile (non serve più)
    rm $EXENAME

    # esce dalla dir
    popd >/dev/null
done #| tee compila-result.md

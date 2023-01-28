#!/bin/bash

### BEGIN url_extract
function url_extract(){
    if [[ -p /dev/stdin ]]; then
        text_to_search="$(cat -)"
    else
        text_to_search="$*"
    fi
    regex_pattern="(https?|s?ftp|file|[a-z]+)://[a-zA-Z0-9_.](:[0-9]{2,5})?([a-zA-Z0-9_.,/#!?&;=%:~*-]+)?"
    echo "$text_to_search" | grep --color=never -Eo "$regex_pattern"
    # Here's a different approach, if limited to HTML, which extracts URLs from inside <A> tag's href attribute's value using Perl's so-called “lookaround” regex syntax:
    # perl_regex_pattern='(?<=href=")[^"]*(?=")'
    # echo "$text_to_search" | grep -Po $perl_regex_pattern
}
### E N D url_extract


### BEGIN url_deref
function url_deref(){
    if [[ -p /dev/stdin ]]; then
        export url_to_trace="$(cat -)"
    else
        export url_to_trace="$*"
    fi
    curl -w "\n%{url_effective}\n" -I -L -s -S $url_to_trace -o /dev/null
    # curl -d "url=$url_to_trace" https://deref.link/deref | jq
    unset url_to_trace
}
### E N D url_deref


### BEGIN function url_unescape
function url_unescape() {
    if [[ -p /dev/stdin ]]; then
        export url_to_unescape="$(cat -)"
    else
        export url_to_unescape="$*"
    fi
    python3 $(dirname $BASH_SOURCE)/url_unescape.py
    unset url_to_unescape
}
### E N D function url_unescape

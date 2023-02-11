#!/bin/bash

### BEGIN url_extract
function url_extract(){
    if [[ -p /dev/stdin ]]; then
        text_to_search_in="$(cat -)"
    else
        text_to_search_in="$*"
    fi
    regex_pattern="(dict|file|s?t?ftps?|gophers?|https?|imaps?|ldaps?|mqtt|pop3s?|rtmp|rtsp|scp|smbs?|smtps?|telnet)://[a-zA-Z0-9_.]+(:[0-9]{2,5})?([a-zA-Z0-9_.,/#!?&;=%:~*-]+)?"
    echo "$text_to_search_in" | grep --color=never -Eo "$regex_pattern"
    # Here's a different approach, if limited to HTML, which extracts URLs from inside <A> tag's href attribute's value using Perl's so-called “lookaround” regex syntax:
    # perl_regex_pattern='(?<=href=")[^"]*(?=")'
    # echo "$text_to_search_in" | grep -Po $perl_regex_pattern
}
### E N D url_extract


### BEGIN url_deref
function url_deref(){
    if [[ -p /dev/stdin ]]; then
        url_to_trace="$(cat -)"
    else
        url_to_trace="$*"
    fi
    target_url=$(curl -w "%{url_effective}" -L -s -S $url_to_trace -o /dev/null)
    echo -n "$target_url"
}
### E N D url_deref


### BEGIN function url_unescape
function url_unescape() {
    if [[ -p /dev/stdin ]]; then
        url_to_unescape="$(cat -)"
    else
        url_to_unescape="$*"
    fi
    python3 $(dirname $BASH_SOURCE)/url_unescape.py "$url_to_unescape"
}
### E N D function url_unescape

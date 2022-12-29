#!/bin/bash

### BEGIN url_clean
function url_clean(){
    if [[ -p /dev/stdin ]]; then
        url_to_clean="$(cat -)"
    else
        url_to_clean="$*"
    fi
    unwanted_params=$(cat ./unwanted_params.txt)
    for param in $unwanted_params ; do {
        # if param doesn't contain ? symbol
        if [[ "$param" != *"?"* ]]; then
            # if param ends with * then replace * with .*
            param=$(echo "$param" | sed -E "s/[*]$/.*/")
            url_to_clean=$(echo "$url_to_clean" | sed -E "s/[&]?${param}=[^&]*//g")
        # if param contains ? symbol
        else
            url_without_protocol=$(echo "$url_to_clean" | sed -E 's|https?://||')
            url_domain=$(echo "$url_without_protocol" | sed -E 's|/.*$||')
            params_domain_name=$(echo "$param" | sed -E 's/[?].+$//')
            # if url_to_clean's domain name contains param's domain name
            if [[ "$url_domain" == *"$params_domain_name"* ]]; then
                param_without_domain_name=$(echo "$param" | sed -E 's/.+[?]//')
                # if param ends with * then replace * with .*
                param=$(echo "$param_without_domain_name" | sed -E "s/[*]$/.*/")
                url_to_clean=$(echo "$url_to_clean" | sed -E "s/[&]?${param}=[^&]*//g")
            fi
        fi
    }; done
    # remove final ? symbol if no params left
    url_to_clean=$(echo "$url_to_clean" | sed -E "s/[?]$//")
    echo "$url_to_clean"
    unset url_to_clean
}
### E N D url_clean


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

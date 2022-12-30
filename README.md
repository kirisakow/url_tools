# My URL tools written in BASH and Go

And six lines of Python, too.

### BASH functions `url_deref` and `url_unescape`

First, don't forget to run

```sh
source ./bash_functions.sh
```

* BASH function `url_deref` is based on `curl`. It follows URL redirects and prints only the target URL. The URL can be given either as an argument or as STDIN stream:

```sh
# process a URL as an argument:
url_deref "https://stackoverflow.com/a/70819429/4883320"
# or process a URL as STDIN stream:
echo "https://stackoverflow.com/a/70819429/4883320" | url_deref
```
Result:
```
https://stackoverflow.com/questions/70817657/cant-get-firefox-extension-logs-to-show-up/70819429#70819429
```

* BASH function `url_unescape` uses Python's `urllib.parse.unquote` in order to unescape non-ASCII characters in a URL and make it prettier. Here again, the URL can be given either as an argument or as STDIN stream, and the two BASH functions `url_deref` and `url_unescape` are meant to complete each other and can be combined:

```sh
long_url_with_redirect=https://france24.nlfrancemm.com/m/surl/200243/517183/yD0Vqr_mEaDTwJcBJSIuyA==/link_13/HztCd5MALBSiwyWcdZpQvGZuP+L2dlD0fqSjv4DZVsqW+MUvK7a2X8uUILOWdBCiVjMwqEsKsY+9dh7nVfSCzxyxWHUs7tbSQxU3Ok5bOrTyAvRPCKsURxr+LisJ58BR28mFkT2aLLItU7iBkLrHfB5MoWOY3+x0YHcH5Z66LNg-L0J2ND8pSiAw4qzu0Dz19Meq-zbPfN7-MLR6V9LeeQGpxifPQCKMU5nmaVyQUXRZDgDLx+sLPRlzIr--Oc3bzV0X+jgm6SfsBYhxruKPQz70kvNSgAGeNQPgEtBR0AC-m92X8EDJI2th4UFqBvwNeU-rRJx1wgsydqUjrVsLi6-0og9XJILZ3hSboC3S85wB3AW2D6PP7SDuZkDhaTGLG03mmkCipwsPwW2-8UhTLniSzKA054euZqG9vo+Ve3gJrO9QYwQ64EjKTplSScUZVZMok0OhhCg9C3dW1M-tQ1Hd19YpdgWP8U9Tl0xyPmJmOZUAamPUyZJR569tdI+hW-g7tMx9T90eAAstFzj86hQISpD7cKeV3PvMJj+MV8K2668OTZULlrocfGSXTyMbDc0ZaSroLe0nrpbHSjmRWgUisF-z2Rq2+7XzUGmrtcS3sYgpMag2QemK68TzVlqu2CaK2B97jIyZNOyuHpbBKPNYRM58mu+D7-9KTnysI-YcH93Fmh33mRv1fyVlxCpmm0PoZXmZd7x7klL6-JStwhei33DpD-qRUAlmo93xOlzO9xJQxjUpZaG1qM2xn9e+WAfwVIA3ouw8slY0W5PjCRmqOjtB4bSIWANjsLrKkAAwzHm-BCcfeWFjzA+PlQXJ3jV4WNaTkek91lEF0aPbWoxUplU0xV+610tu3sKnjM4=

# process a URL as STDIN stream:
url_deref "$long_url_with_redirect" | url_unescape
# or process a URL as an argument:
url_unescape $(url_deref "$long_url_with_redirect")
# result:
https://www.france24.com/fr/amériques/20221227-le-blizzard-du-siècle-fait-au-moins-50-morts-aux-états-unis-le-bilan-risque-de-s-alourdir?xtor=EPR-300&_ope=eyJndWlkIjoiN2ZiZTFiYWI1YWRiMTI1ZGJmMzRkMDdhNWQzNGQ2ZWIifQ==
```

### Go function `url_clean`

First, don't forget to run

```sh
cd /path/to/url_clean.go
go build
```

The `url_clean` function checks a URL against a 120+ long list of garbage query parameters (`unwanted_params.txt`) which it removes from the URL. Originally written in BASH, this function has been refactored in Go to be faster. This function can process a URL either as an argument or as STDIN stream, and can be combined with the two aforementioned BASH functions:

```sh
# process a URL as STDIN stream:
url_deref "$long_url_with_redirect" | url_unescape | ./url_clean
# or process a URL as an argument:
./url_clean $(url_deref "$long_url_with_redirect" | url_unescape)
# result:
https://www.france24.com/fr/amériques/20221227-le-blizzard-du-siècle-fait-au-moins-50-morts-aux-états-unis-le-bilan-risque-de-s-alourdir
```

Unlike the two BASH functions which only process one URL at a time, the `url_clean` function can process multiple URLs either as multiple arguments or as a multiline STDIN stream.

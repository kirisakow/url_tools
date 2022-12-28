# My URL tools written in BASH and Go

And six lines of Python, too.

### BASH functions `url_deref` and `url_unescape`

First, don't forget to

```sh
source ./bash_functions.sh
```

* BASH function `url_deref` is based on `curl`. It follows URL redirects and prints only the target URL. The URL can be given either as an argument or as STDIN:

```sh
url_deref "https://france24.nlfrancemm.com/m/surl/200243/517183/yD0Vqr_mEaDTwJcBJSIuyA==/link_13/HztCd5MALBSiwyWcdZpQvGZuP+L2dlD0fqSjv4DZVsqW+MUvK7a2X8uUILOWdBCiVjMwqEsKsY+9dh7nVfSCzxyxWHUs7tbSQxU3Ok5bOrTyAvRPCKsURxr+LisJ58BR28mFkT2aLLItU7iBkLrHfB5MoWOY3+x0YHcH5Z66LNg-L0J2ND8pSiAw4qzu0Dz19Meq-zbPfN7-MLR6V9LeeQGpxifPQCKMU5nmaVyQUXRZDgDLx+sLPRlzIr--Oc3bzV0X+jgm6SfsBYhxruKPQz70kvNSgAGeNQPgEtBR0AC-m92X8EDJI2th4UFqBvwNeU-rRJx1wgsydqUjrVsLi6-0og9XJILZ3hSboC3S85wB3AW2D6PP7SDuZkDhaTGLG03mmkCipwsPwW2-8UhTLniSzKA054euZqG9vo+Ve3gJrO9QYwQ64EjKTplSScUZVZMok0OhhCg9C3dW1M-tQ1Hd19YpdgWP8U9Tl0xyPmJmOZUAamPUyZJR569tdI+hW-g7tMx9T90eAAstFzj86hQISpD7cKeV3PvMJj+MV8K2668OTZULlrocfGSXTyMbDc0ZaSroLe0nrpbHSjmRWgUisF-z2Rq2+7XzUGmrtcS3sYgpMag2QemK68TzVlqu2CaK2B97jIyZNOyuHpbBKPNYRM58mu+D7-9KTnysI-YcH93Fmh33mRv1fyVlxCpmm0PoZXmZd7x7klL6-JStwhei33DpD-qRUAlmo93xOlzO9xJQxjUpZaG1qM2xn9e+WAfwVIA3ouw8slY0W5PjCRmqOjtB4bSIWANjsLrKkAAwzHm-BCcfeWFjzA+PlQXJ3jV4WNaTkek91lEF0aPbWoxUplU0xV+610tu3sKnjM4="
```
```
https://www.france24.com/fr/am%c3%a9riques/20221227-le-blizzard-du-si%c3%a8cle-fait-au-moins-50-morts-aux-%c3%a9tats-unis-le-bilan-risque-de-s-alourdir?xtor=EPR-300&_ope=eyJndWlkIjoiN2ZiZTFiYWI1YWRiMTI1ZGJmMzRkMDdhNWQzNGQ2ZWIifQ%3D%3D
```
```sh
echo "https://france24.nlfrancemm.com/m/surl/200243/517183/yD0Vqr_mEaDTwJcBJSIuyA==/link_13/HztCd5MALBSiwyWcdZpQvGZuP+L2dlD0fqSjv4DZVsqW+MUvK7a2X8uUILOWdBCiVjMwqEsKsY+9dh7nVfSCzxyxWHUs7tbSQxU3Ok5bOrTyAvRPCKsURxr+LisJ58BR28mFkT2aLLItU7iBkLrHfB5MoWOY3+x0YHcH5Z66LNg-L0J2ND8pSiAw4qzu0Dz19Meq-zbPfN7-MLR6V9LeeQGpxifPQCKMU5nmaVyQUXRZDgDLx+sLPRlzIr--Oc3bzV0X+jgm6SfsBYhxruKPQz70kvNSgAGeNQPgEtBR0AC-m92X8EDJI2th4UFqBvwNeU-rRJx1wgsydqUjrVsLi6-0og9XJILZ3hSboC3S85wB3AW2D6PP7SDuZkDhaTGLG03mmkCipwsPwW2-8UhTLniSzKA054euZqG9vo+Ve3gJrO9QYwQ64EjKTplSScUZVZMok0OhhCg9C3dW1M-tQ1Hd19YpdgWP8U9Tl0xyPmJmOZUAamPUyZJR569tdI+hW-g7tMx9T90eAAstFzj86hQISpD7cKeV3PvMJj+MV8K2668OTZULlrocfGSXTyMbDc0ZaSroLe0nrpbHSjmRWgUisF-z2Rq2+7XzUGmrtcS3sYgpMag2QemK68TzVlqu2CaK2B97jIyZNOyuHpbBKPNYRM58mu+D7-9KTnysI-YcH93Fmh33mRv1fyVlxCpmm0PoZXmZd7x7klL6-JStwhei33DpD-qRUAlmo93xOlzO9xJQxjUpZaG1qM2xn9e+WAfwVIA3ouw8slY0W5PjCRmqOjtB4bSIWANjsLrKkAAwzHm-BCcfeWFjzA+PlQXJ3jV4WNaTkek91lEF0aPbWoxUplU0xV+610tu3sKnjM4=" | url_deref
```
```
https://www.france24.com/fr/am%c3%a9riques/20221227-le-blizzard-du-si%c3%a8cle-fait-au-moins-50-morts-aux-%c3%a9tats-unis-le-bilan-risque-de-s-alourdir?xtor=EPR-300&_ope=eyJndWlkIjoiN2ZiZTFiYWI1YWRiMTI1ZGJmMzRkMDdhNWQzNGQ2ZWIifQ%3D%3D
```

* BASH function `url_unescape` uses Python's `urllib.parse.unquote` in order to unescape non-ASCII characters in a URL and make it prettier. Here again, the URL can be given either as an argument or as STDIN:

```sh
url_unescape "https://www.france24.com/fr/am%c3%a9riques/20221227-le-blizzard-du-si%c3%a8cle-fait-au-moins-50-morts-aux-%c3%a9tats-unis-le-bilan-risque-de-s-alourdir?xtor=EPR-300&_ope=eyJndWlkIjoiN2ZiZTFiYWI1YWRiMTI1ZGJmMzRkMDdhNWQzNGQ2ZWIifQ%3D%3D"
```
```
https://www.france24.com/fr/amériques/20221227-le-blizzard-du-siècle-fait-au-moins-50-morts-aux-états-unis-le-bilan-risque-de-s-alourdir?xtor=EPR-300&_ope=eyJndWlkIjoiN2ZiZTFiYWI1YWRiMTI1ZGJmMzRkMDdhNWQzNGQ2ZWIifQ==
```
```sh
echo "https://www.france24.com/fr/am%c3%a9riques/20221227-le-blizzard-du-si%c3%a8cle-fait-au-moins-50-morts-aux-%c3%a9tats-unis-le-bilan-risque-de-s-alourdir?xtor=EPR-300&_ope=eyJndWlkIjoiN2ZiZTFiYWI1YWRiMTI1ZGJmMzRkMDdhNWQzNGQ2ZWIifQ%3D%3D" | url_unescape
```
```
https://www.france24.com/fr/amériques/20221227-le-blizzard-du-siècle-fait-au-moins-50-morts-aux-états-unis-le-bilan-risque-de-s-alourdir?xtor=EPR-300&_ope=eyJndWlkIjoiN2ZiZTFiYWI1YWRiMTI1ZGJmMzRkMDdhNWQzNGQ2ZWIifQ==
```

The two BASH functions `url_deref` and `url_unescape` are meant to complete each other and can be combined:

```sh
url_deref "https://france24.nlfrancemm.com/m/surl/200243/517183/yD0Vqr_mEaDTwJcBJSIuyA==/link_13/HztCd5MALBSiwyWcdZpQvGZuP+L2dlD0fqSjv4DZVsqW+MUvK7a2X8uUILOWdBCiVjMwqEsKsY+9dh7nVfSCzxyxWHUs7tbSQxU3Ok5bOrTyAvRPCKsURxr+LisJ58BR28mFkT2aLLItU7iBkLrHfB5MoWOY3+x0YHcH5Z66LNg-L0J2ND8pSiAw4qzu0Dz19Meq-zbPfN7-MLR6V9LeeQGpxifPQCKMU5nmaVyQUXRZDgDLx+sLPRlzIr--Oc3bzV0X+jgm6SfsBYhxruKPQz70kvNSgAGeNQPgEtBR0AC-m92X8EDJI2th4UFqBvwNeU-rRJx1wgsydqUjrVsLi6-0og9XJILZ3hSboC3S85wB3AW2D6PP7SDuZkDhaTGLG03mmkCipwsPwW2-8UhTLniSzKA054euZqG9vo+Ve3gJrO9QYwQ64EjKTplSScUZVZMok0OhhCg9C3dW1M-tQ1Hd19YpdgWP8U9Tl0xyPmJmOZUAamPUyZJR569tdI+hW-g7tMx9T90eAAstFzj86hQISpD7cKeV3PvMJj+MV8K2668OTZULlrocfGSXTyMbDc0ZaSroLe0nrpbHSjmRWgUisF-z2Rq2+7XzUGmrtcS3sYgpMag2QemK68TzVlqu2CaK2B97jIyZNOyuHpbBKPNYRM58mu+D7-9KTnysI-YcH93Fmh33mRv1fyVlxCpmm0PoZXmZd7x7klL6-JStwhei33DpD-qRUAlmo93xOlzO9xJQxjUpZaG1qM2xn9e+WAfwVIA3ouw8slY0W5PjCRmqOjtB4bSIWANjsLrKkAAwzHm-BCcfeWFjzA+PlQXJ3jV4WNaTkek91lEF0aPbWoxUplU0xV+610tu3sKnjM4=" | url_unescape
```
```
https://www.france24.com/fr/amériques/20221227-le-blizzard-du-siècle-fait-au-moins-50-morts-aux-états-unis-le-bilan-risque-de-s-alourdir?xtor=EPR-300&_ope=eyJndWlkIjoiN2ZiZTFiYWI1YWRiMTI1ZGJmMzRkMDdhNWQzNGQ2ZWIifQ==
```

### Go function `url_clean`

First, don't forget to

```sh
go build
```

Originally written in BASH, the `url_clean` function has been refactored in Go to be faster. It checks a URL against a 120+ long dictionary of `unwanted_params.txt` which it removes from the URL. Here again, the URL can be given either as an argument or as STDIN, and can be combined with the two other functions:

```sh
url_deref "https://france24.nlfrancemm.com/m/surl/200243/517183/yD0Vqr_mEaDTwJcBJSIuyA==/link_13/HztCd5MALBSiwyWcdZpQvGZuP+L2dlD0fqSjv4DZVsqW+MUvK7a2X8uUILOWdBCiVjMwqEsKsY+9dh7nVfSCzxyxWHUs7tbSQxU3Ok5bOrTyAvRPCKsURxr+LisJ58BR28mFkT2aLLItU7iBkLrHfB5MoWOY3+x0YHcH5Z66LNg-L0J2ND8pSiAw4qzu0Dz19Meq-zbPfN7-MLR6V9LeeQGpxifPQCKMU5nmaVyQUXRZDgDLx+sLPRlzIr--Oc3bzV0X+jgm6SfsBYhxruKPQz70kvNSgAGeNQPgEtBR0AC-m92X8EDJI2th4UFqBvwNeU-rRJx1wgsydqUjrVsLi6-0og9XJILZ3hSboC3S85wB3AW2D6PP7SDuZkDhaTGLG03mmkCipwsPwW2-8UhTLniSzKA054euZqG9vo+Ve3gJrO9QYwQ64EjKTplSScUZVZMok0OhhCg9C3dW1M-tQ1Hd19YpdgWP8U9Tl0xyPmJmOZUAamPUyZJR569tdI+hW-g7tMx9T90eAAstFzj86hQISpD7cKeV3PvMJj+MV8K2668OTZULlrocfGSXTyMbDc0ZaSroLe0nrpbHSjmRWgUisF-z2Rq2+7XzUGmrtcS3sYgpMag2QemK68TzVlqu2CaK2B97jIyZNOyuHpbBKPNYRM58mu+D7-9KTnysI-YcH93Fmh33mRv1fyVlxCpmm0PoZXmZd7x7klL6-JStwhei33DpD-qRUAlmo93xOlzO9xJQxjUpZaG1qM2xn9e+WAfwVIA3ouw8slY0W5PjCRmqOjtB4bSIWANjsLrKkAAwzHm-BCcfeWFjzA+PlQXJ3jV4WNaTkek91lEF0aPbWoxUplU0xV+610tu3sKnjM4=" | url_unescape | ./url_clean_go
```
```
https://www.france24.com/fr/amériques/20221227-le-blizzard-du-siècle-fait-au-moins-50-morts-aux-états-unis-le-bilan-risque-de-s-alourdir
```
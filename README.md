README
====

The overall package currently has a couple depedencies, 
the general flow is 

* create rmd documents prepended with `__` (eg `__myEx.Rmd`). 
* output resulting dashboards to the demos folder
* for now manually, but will automate
    * the additional to the data/items.toml files
    * move the final image for each representative dashboard to static/images/full
* re-run hugo to generate 
* rename public folder to /docs and copy in demos folder


## misc dependency info

make on windows being lazy without adding hugo to path via

```
.\hugo_0.17_<rest-of-binary-name>.exe
```

make thumbnails of all files in 'full' images folder via


```
go run make_thumbnails.go
```
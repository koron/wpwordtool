# Wikipedia Word Tool

Wikipedia の記事見出しを java-ugmatcha-suite で取り扱える形式に変換するツール。

詳細は
<https://github.com/koron/java-ugmatcha-suite/blob/master/doc/algorithm.md>
および
<https://github.com/koron/java-ugmatcha-suite/blob/master/wikidict/README.md>
を参照してください。

## Getting started

### Install or Update


```console
$ GO111MODULE=on go get -u -i github.com/koron/wpwordtool
```

### Convert sub command

```console
$ wpwordtool convert
```

Load `jawiki-latest-all-titles-in-ns0.gz` (default of `-ja`) and
`enwiki-latest-all-titles-in-ns0.gz` (default of `-en`), then
save converted dictionary files `tmp/wikiwords.stt` and `tmp/wikiwords.stw`.
Basename of output files is `tmp/wikiwords` which is default of `-out`.

### Abstract sub command

Extract all abstracts from wikipedia's article XML files.

`jawiki-YYYYMMDD-abstractX.xml.gz`.

How to use:

```console
$ wpwordtool abstract jawiki-20190701-abstract*.xml.gz > abstract.txt
2019/07/24 11:58:09 extracting from jawiki-20190701-abstract1.xml.gz
2019/07/24 11:58:11 extracting from jawiki-20190701-abstract2.xml.gz
2019/07/24 11:58:13 extracting from jawiki-20190701-abstract3.xml.gz
2019/07/24 11:58:14 extracting from jawiki-20190701-abstract4.xml.gz
2019/07/24 11:58:15 extracting from jawiki-20190701-abstract5.xml.gz
2019/07/24 11:58:16 extracting from jawiki-20190701-abstract6.xml.gz

$ wc -l -c abstract.txt
  1157686 141950858 abstract.txt
```

## Dictionaries

Download these files.

*   <https://dumps.wikimedia.org/jawiki/> - YYYYMMDD/jawiki-YYYYMMDD-all-title-in-ns0.gz
*   <https://dumps.wikimedia.org/enwiki/> - YYYYMMDD/enwiki-YYYYMMDD-all-title-in-ns0.gz

# ColorConverter.go

[![Build Status](https://travis-ci.org/SimonWaldherr/ColorConverterGo.svg?branch=master)](https://travis-ci.org/SimonWaldherr/ColorConverterGo)
[![GoDoc](https://godoc.org/github.com/SimonWaldherr/ColorConverterGo?status.svg)](https://godoc.org/github.com/SimonWaldherr/ColorConverterGo)
[![License MIT](http://img.shields.io/badge/license-MIT-red.svg)](http://opensource.org/licenses/MIT)
[![Coverage Status](https://img.shields.io/coveralls/SimonWaldherr/ColorConverterGo.svg)](https://coveralls.io/r/SimonWaldherr/ColorConverterGo)

Convert between RGB, HSL, HSV, CMYK and HEX color defining with these GO functions under MIT-License  
translated from https://github.com/SimonWaldherr/ColorConverter.js

## about

License:   MIT  
Version: 0.0.6  
Date:  10.2014  

## download

* ```go get github.com/SimonWaldherr/ColorConverterGo```
* git: ```git@github.com:SimonWaldherr/ColorConverter.go.git```
* [zip](https://github.com/SimonWaldherr/ColorConverter.go/archive/master.zip) from GitHub

## matrix

  . | RGB | HSL | HSV | YUV | Hex | CMYK
----|-----|-----|-----|-----|-----|-----
RGB |     |  X  |  .  |  X  |  X  |  X
HSL |  X  |     |  .  |  .  |  .  |  X
HSV |  X  |  .  |     |  .  |  .  |  X
YUV |  .  |  .  |  .  |     |  .  |  .
Hex |  X  |  .  |  .  |  .  |     |  X
CMYK|  X  |  .  |  .  |  .  |  .  |   

## contact

Feel free to contact me via [eMail](mailto:contact@simonwaldherr.de) or on [Twitter](http://twitter.com/simonwaldherr). This software will be continually developed. Suggestions and tips are always welcome.

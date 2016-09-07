package main

import (
    "testing"
    "github.com/stretchr/testify/assert"
)

func TestParse(t *testing.T) {
    assert := assert.New(t)

    body := `
<link rel="stylesheet" href="//www.dpfile.com/mod/app-shop-full-map/1.0.1/css/index.css" type="text/css"/>
    <div class="mod aside-mod">
        <div class="map">
            <div id="map" class="container"></div>
            <a id="J_map-show" class="map-zoom"><i class="icon"></i></a>
        </div>
    </div>
    <script>
        (function(config){var gid=function(id){return document.getElementById(id)};var map=gid("map");map.innerHTML='<img src="http://apis.map.qq.com/ws/staticmap/v2/?key=I3OBZ-MBSRQ-WBJ5P-G5VZS-QGAIF-Y7B27&size=240*90&center={lat},{lng}&zoom=15&markers=icon:http://i2.dpfile.com/lib/1.0/map/img/marker.png|{lat},{lng}" />'.replace(/\{(\w+)\}/g,function(all,key){return config[key]})})
        ({lng:116.460608,lat:39.895087});
    </script>
<div class="J_midas-4"></div>
`

    lng, lat, found := Parse(body)
    assert.True(found)
    assert.Equal("116.460608", lng)
    assert.Equal("39.895087", lat)

    body2 := `whatever`
    _, _, found = Parse(body2)
    assert.False(found)
}

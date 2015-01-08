ready = ->
  console.log 'a'
  $('.picture').each ->
    $(@).height($('.picture img')[0].width)
 
throttled_ready = _.throttle(ready, 100)
$(document).ready(throttled_ready)
$(document).on('page:load', throttled_ready)
$(window).on('resize', throttled_ready)
ready = ->
  $('.picture').each ->
    $(@).height($('.picture img')[0].width)

  $('#picture-section').waypoint
    handler: ->
      $('.navbar').toggleClass('scrolled')
    offset: '30%'

throttled_ready = _.throttle(ready, 100)
$(document).ready(throttled_ready)
$(document).on('page:load', throttled_ready)
$(window).on('resize', throttled_ready)

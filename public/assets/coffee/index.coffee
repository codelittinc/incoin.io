$(document).ready ->
  $("#main").append "<div class='angled-border'></div>"
  resizeBorder()

  $(window).resize ->
    resizeBorder()

# Since borders don't allow width in percentage we have to calculate it
resizeBorder = ->
  $('.angled-border').css('border-left-width', "#{$(window).width()}px")
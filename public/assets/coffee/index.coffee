$(document).ready ->
  $("#main").append "<div class='angled-border bottom'></div>"
  $("#ready-for").append "<div class='angled-border top'></div>"


  resizeBorder()

  $(window).resize ->
    resizeBorder()

# Since borders don't allow width in percentage we have to calculate it
resizeBorder = ->
  $('.angled-border.bottom').css('border-left-width', "#{$(window).width()}px")
  $('.angled-border.top').css('border-right-width', "#{$(window).width()}px")
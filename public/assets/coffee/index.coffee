$(document).ready ->
  $("#main").append "<div class='angled-border bottom'></div>"
  $("#ready-for").append "<div class='angled-border top'></div>"


  resizeBorder()

  $(window).resize ->
    resizeBorder()

  # Show and hide employee and employer section on the feature page
  # Show employer container
  $('#employer-tab').on "click", ->
    $('#employee-tab').removeClass("active")
    $('.employee-description').removeClass("active")
    $('#employer-tab').addClass("active")
    $('.employers-description').addClass("active")

  # Show employee container
  $('#employee-tab').on "click", ->
    $('#employer-tab').removeClass("active")
    $('.employers-description').removeClass("active")
    $('#employee-tab').addClass("active")
    $('.employee-description').addClass("active")

# Since borders don't allow width in percentage we have to calculate it
resizeBorder = ->
  $('.angled-border.bottom').css('border-left-width', "#{$(window).width()}px")
  $('.angled-border.top').css('border-right-width', "#{$(window).width()}px")


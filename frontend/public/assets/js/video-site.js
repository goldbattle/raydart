$(function() {
  // Handle loading of players
  if(typeof videojs != 'undefined') {
    videojs.options.flash.swf = "/assets/swf/video-js.swf";
  }
});


$(document).ready(function(){
  // Don't display right click menu
  $('.video-js').bind('contextmenu',function() { return false; });
});
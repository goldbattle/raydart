$(document).ready(function() {
  // Define viewer count button
  // Add listeners here if needed
  videojs.viewer_count = videojs.Button.extend({
    init: function(player, options){
      videojs.Button.call(this, player, options);
    }
  });

  // Our plugin
  videojs.plugin('view-count', function() {
    // Define our properties
    var props = {
      className: 'vjs-viewer-count',
      innerHTML: '<span class="glyphicon glyphicon-user" style="margin-right:5px;"></span><span id="viewer-count">0</span>', 
      role: 'button'
    };
    // Create the option data
    var options = { 'el' : videojs.Component.prototype.createEl(null,props) };
    // Create object
    var pl_viewer_count = new videojs.viewer_count(this, options);
    // Add it to the control bar
    this.controlBar.el().appendChild(pl_viewer_count.el());
  });

  // Function that calls our api, and updates the viewer count
  function update_viewer_count() {
    // Start our viewer count update
    $.getJSON("/api/v1/viewercount/").success(function(json_data){
      // Update our viewer count on the player
      $('#viewer-count').html(json_data.viewer_count);
    })
  }
  // Refresh the viewer count every 60s
  setInterval(update_viewer_count, 60000);
  // Call it on load
  update_viewer_count();

  // Define our bitrate selection button
  // Adding a click listener here
  videojs.bitrate_select = videojs.Button.extend({
    init: function(player, options, sources){
      videojs.Button.call(this, player, options);
      // Enable our click method
      this.on('click', this.onClick);
      // Load in our resolutions
      var content = '<ul style="list-style:none;padding:0px;margin:0px;">';
      for (var i in sources) {
        content += '<li>'+sources[i]['data-res']+'</li>';
      }
      content += '</ul>';
      // Load content into popup
      $('#bitrate-popup').html(content);
      // Add our listener
      $('#bitrate-popup li').on("click", function() {
        // Load in our resolutions
        var player = $('#stream-video').get(0).player;
        var new_url = '';
        // Loop through each source, and select the right one
        for (var i in player.options().sources) {
          if(player.options().sources[i]['data-res'] == $(this).text())
            new_url = player.options().sources[i].src
        }
        // Switch the player to the new stream
        if(new_url != '') {
          // Change to our new source
          player.src({ type: "rtmp/mp4", src: new_url });
          // Pause and replay
          player.pause();
          player.play();
          sleep(500);
          player.pause();
          player.play();
        }
        // Toggle our popup
        $('#bitrate-popup').fadeToggle('fast');
        $('.vjs-control-bar').toggleClass('bitrate-active');
        // Update current button
        $('#bitrate-current').html($(this).text())
        // Debug
        //console.log("Switching to: "+$(this).text())
        //console.log(player.currentSrc())
      });
    }
  });

  // This function handles setting the location of the popup
  function update_bitrate_popup() {
    // Get our location
    var offset = $('.vjs-bitrate-select').offset();
    var height = $('#bitrate-popup').height();
    // Calculate our positions
    var top = (offset.top - height) + "px";
    var right = offset.left + "px";
    // Move to the new position
    $('#bitrate-popup').css( {
      'position': 'absolute',
      'left': right,
      'top': top
    });
  }

  // OnResize - Window resize, fix our bitrate window
  $(window).resize(function(){
    update_bitrate_popup();
  });

  // OnClick - Handle displaying the selection
  videojs.bitrate_select.prototype.onClick = function() {
    // Update our position
    update_bitrate_popup()
    // Toggle our popup
    $('#bitrate-popup').fadeToggle('fast');
    $('.vjs-control-bar').toggleClass('bitrate-active');
    // Recalc just incase
    update_bitrate_popup();
  };

  // Our bitrate plugin
  videojs.plugin('bitrate-select', function() {
    // Load in our resolutions
    var sources = $('#stream-video').get(0).player.options().sources;
    // Define our properties, select the first stream
    var props = {
      className: 'vjs-bitrate-select',
      innerHTML: '<span class="glyphicon glyphicon-cog" style="margin-right:5px;"></span><span id="bitrate-current">'+sources[0]['data-res']+'</span>', 
      role: 'button'
    };
    // Create the option data
    var options = { 'el' : videojs.Component.prototype.createEl(null,props) };
    // Create object
    var pl_bitrate_select = new videojs.bitrate_select(this, options, sources);
    // Add it to the control bar
    this.controlBar.el().appendChild(pl_bitrate_select.el());
  });

  function sleep(milliseconds) {
    var start = new Date().getTime();
    for (var i = 0; i < 1e7; i++) {
      if ((new Date().getTime() - start) > milliseconds){
        break;
      }
    }
  }
});
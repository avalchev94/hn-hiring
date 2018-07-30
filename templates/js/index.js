$(function(){
  function onSocketMessage(message) {
    var jsonMessage = eval("("+message.data+")")
    switch (jsonMessage.type) {
      case 0: // new job post found
        posts.append(
          $("<li>").append(
            $("<span>").append(jsonMessage.message.text)
          )
        );
        break;
      case 1: // Nothing was found
        posts.append(
          $("<span>").append(jsonMessage.message)
        );
        break;
      case 2: // Search finished
        posts.append(
          $("<span>").append(jsonMessage.message)
        );
        break;
      case 3: // Incorrect expression
        alert(jsonMessage.message.error + " at " + jsonMessage.message.column);
        break;
    }
  }

  function onSocketClose() {
    alert("socket connection closed? refresh your page.")
  }

  function sendSocketData() {
    var expression = $("input[name='expression']").val()
    var items = [];
    $(".selectpicker option:selected").each( function(){
      items.push(parseInt($(this).val(), 10));
    });
    var message = JSON.stringify({
        "items": items,
        "expression": expression
    })
    socket.send(message)
  }

  var socket = null;
  var posts = $("#posts");

  if (!window["WebSocket"]) {
    alert("Error: WebBrowser does not support websockets");
  } else {
    // initialize select data
    $(".selectpicker option:first").attr("selected", "selected")

    // initialize the socket and its callbacks
    socket = new WebSocket("ws://localhost:8080/hire");
    socket.onclose = onSocketClose;
    socket.onmessage = onSocketMessage;
    socket.onopen = sendSocketData;

    // when search button is clicked, send socket data.
    $("#search").click(sendSocketData);
  }
}); 
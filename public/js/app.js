$(function() {

  messageHandler = {
    goodbye: function() {
      console.log("goodbye:", "Connection has been closed!");
    },
    welcome: function(data) {
      console.log("data:", data);
      $("h1").text(data.data);
      uuid = data.uuid;
      console.log("uuid:", uuid);
      $("h3").text(uuid);
    },
    displayMessage: function(message) {
      display = $("#chat-display");
      display.append("<p><strong>" + message.username + ":</strong> " + message.message + "</p>");
      display.scrollTop(display.prop("scrollHeight"));
    },
    postMessage: function(message) {
      $.ajax({
        method: "POST",
        contentType: "applicatin/json",
        dataType: "json",
        url: "/message",
        data: JSON.stringify(message),
        success: function(data, status, jqxhr) {
          $("#message-text").val("");
        },
        failure: function(errMsg) {
          alert(errMsg);
        }
      });
    }
  }

  $("#message-text").keypress(function(event) {
    if(event.which == 13) {
      event.preventDefault();

      messageTextBox = $(event.currentTarget);
      message = { username: "jcarley", message: messageTextBox.val() }
      messageHandler.postMessage(message);
    }
  });

  goButton = $("#go-button").click(function(e) {
    messageTextBox = $("#message-text");
    message = { username: "jcarley", message: messageTextBox.val() }
    messageHandler.postMessage(message);
  });

  connect = function() {
    ws = new WebSocket("wss://" + window.location.host + "/ws");
    ws.onopen = function(e) {
      console.log("onopen:", arguments);
    };
    ws.onclose = function(e) {
      messageHandler.goodbye();
    };
    ws.onmessage = function(e) {
      d = JSON.parse(e.data);
      message = d["NewValue"];
      messageHandler.displayMessage(message);
    };
  };

  connect();
});


// {
  // "NewValue":
    // {
    // "created":"2015-05-31T21:15:27.081000089Z",
    // "id":"007bfc3e-12cf-4ed2-8a0f-71d8bc02f0cf",
    // "message":"this is a message",
    // "username":"jcarley"
  // },
  // "OldValue":null
// }


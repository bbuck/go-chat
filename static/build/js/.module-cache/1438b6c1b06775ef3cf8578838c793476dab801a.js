define(["reactjs", "components/message_box"], function(React, MessageBox) {
  return React.createClass({
    render: function() {
      return (
        React.createElement("h1", null, "Hello, World!")
      );
    }
  });
});
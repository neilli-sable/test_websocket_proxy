var app = require("express")();
var http = require("http").createServer(app);
var io = require("socket.io")(http);

app.get("/", (req, res) => {
  res.sendFile(__dirname + "/index.html");
});

io.on("connection", (socket) => {
  socket.on("chat message", (msg) => {
    console.log(socket.request.headers);
    console.log("message: " + msg);
  });
});

http.listen(12345, () => {
  console.log("listening on *:12345");
});

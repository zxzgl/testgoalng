define(
	"main",
	[
		"MessageList"
	],
	function(MessageList) {
		var ws = new WebSocket("ws://192.168.2.110:8080/entry");
		var list = new MessageList(ws);
		ko.applyBindings(list);
	}
);

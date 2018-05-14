(function(){
	"use strict";
	const Message = function(){
		this.body = "";
		this.username = "";
		this.type = 0;
	};


	const app = new Vue({
		el  : "#app",
		data: {
			messageMode  : 0,
			messages     : [],
			newMessage   : new Message(),
			emojiList    : [..."😀😁😂🤣😃😄😅😆😉😊😋😎😍😘😗😙😚🙂🤗🤩🤔🤨😐😑😶🙄😏😣😥😮🤐😯😪😫😴😌😛😜😝"],
			selectedEmoji: null,
		},
		created(){
			this.getMessages();
		//console.log(this.moji);
		},
		methods: {
			getMessages(){
				fetch("/api/messages").then(response => response.json()).then(data => {
					this.messages = data.result;
					/*for(let i=data.result.length;i--;){
				let type=~~(Math.random()*2);
				data.result[i].type=0;
				if(type===1){
					data.result[i].type=1;
					data.result[i].body="絵";
				}
			}*/
				});
			},
			sendMessage(){
				const message = this.newMessage;
				const emojiCharRegEx =
/^[\u{1f300}-\u{1f5ff}\u{1f900}-\u{1f9ff}\u{1f600}-\u{1f64f}\u{1f680}-\u{1f6ff}\u{2600}-\u{26ff}\u{2700}-\u{27bf}\u{1f1e6}-\u{1f1ff}\u{1f191}-\u{1f251}\u{1f004}\u{1f0cf}\u{1f170}-\u{1f171}\u{1f17e}-\u{1f17f}\u{1f18e}\u{3030}\u{2b50}\u{2b55}\u{2934}-\u{2935}\u{2b05}-\u{2b07}\u{2b1b}-\u{2b1c}\u{3297}\u{3299}\u{303d}\u{00a9}\u{00ae}\u{2122}\u{23f3}\u{24c2}\u{23e9}-\u{23ef}\u{25b6}\u{23f8}-\u{23fa}]$/u;
				console.log(this.newMessage.type);
				if(this.newMessage.type === 1 && !this.newMessage.body.match(emojiCharRegEx)){
					alert("Emojiが選択されていません");
					return;
				}
				return;
				fetch("/api/messages", {
					method: "POST",
					body  : JSON.stringify(message),
				})
					.then(response => response.json())
					.then(response => {
						if (response.error){
							alert(response.error.message);
							return;
						}
						this.messages.push(response.result);
						this.clearMessage();
					})
					.catch(error => {
						console.log(error);
					});
			},
			removeMessage(id){
				return fetch(`/api/messages/${id}`, {
					method: "DELETE",
				})
					.then(response => response.json())
					.then(response => {
						if (response.error){
							alert(response.error.message);
							return;
						}
						this.messages = this.messages.filter(m => {
							return m.id !== id;
						});
					});
			},
			updateMessage(updatedMessage){
				return fetch(`/api/messages/${updatedMessage.id}`, {
					method: "PUT",
					body  : JSON.stringify(updatedMessage),
				})
					.then(response => response.json())
					.then(response => {
						if (response.error){
							alert(response.error.message);
							return;
						}
						const index = this.messages.findIndex(m => {
							return m.id === updatedMessage.id;
						});
						Vue.set(this.messages, index, updatedMessage);
					});
			},
			clearMessage(){
				this.newMessage = new Message();
			},
		},
	});
})();

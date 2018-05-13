(function() {
  'use strict';
  const Message = function() {
    this.body = ''
	 this.username = ''
	 this.type = 0;
  };

  const app = new Vue({
    el: '#app',
    data: {
      messages: [],
      newMessage: new Message()
    },
    created() {
      this.getMessages();
    },
    methods: {
      getMessages() {
        fetch('/api/messages').then(response => response.json()).then(data => {
			this.messages = data.result;
			for(let i=data.result.length;i--;){
				let type=~~(Math.random()*2);
				data.result[i].type=0;
				if(type===1){
					data.result[i].type=1;
					data.result[i].body="絵";
				}
			}
        });
      },
      sendMessage() {
        const message = this.newMessage;
        fetch('/api/messages', {
          method: 'POST',
          body: JSON.stringify(message)
        })
          .then(response => response.json())
          .then(response => {
            if (response.error) {
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
      removeMessage(id) {
        return fetch(`/api/messages/${id}`, {
          method: 'DELETE'
        })
        .then(response => response.json())
        .then(response => {
          if (response.error) {
            alert(response.error.message);
            return;
          }
          this.messages = this.messages.filter(m => {
            return m.id !== id
          })
        })
      },
      updateMessage(updatedMessage) {
        return fetch(`/api/messages/${updatedMessage.id}`, {
          method: 'PUT',
          body: JSON.stringify(updatedMessage),
        })
        .then(response => response.json())
        .then(response => {
            if (response.error) {
              alert(response.error.message);
              return;
            }
            const index = this.messages.findIndex(m => {
              return m.id === updatedMessage.id
            })
            Vue.set(this.messages, index, updatedMessage)
        })
      },
      clearMessage() {
        this.newMessage = new Message();
      }
    }
  });
})();

Vue.component("message", {
// Tutorial 1-1. ユーザー名を表示しよう
	props: ["id", "body", "type", "username", "removeMessage", "updateMessage"],
	data(){
		return {
			editing   : false,
			editedBody: null,
		};
	},
	// Tutorial 1-1. ユーザー名を表示しよう
	template: `
	<div class="message">
		<div v-if="editing">
			<div class="row">
				<textarea v-model="editedBody" class="u-full-width"></textarea>
				<button v-on:click="doneEdit">Save</button>
				<button v-on:click="cancelEdit">Cancel</button>
			</div>
		</div>
		<div class="message-body" v-else>
			<span><span v-bind:class="{isStamp: type}">{{ body }}</span> - {{ username }}</span>
			<span class="action-button u-pull-right" v-on:click="remove">&#10007;</span>
			<span v-if="!type" class="action-button u-pull-right" v-on:click="edit">&#9998;</span>
		</div>
	</div>
`,
	methods: {
		remove(){
			this.removeMessage(this.id);
		},
		edit(){
			this.editing = true;
			this.editedBody = this.body;
		},
		cancelEdit(){
			this.editing = false;
			this.editedBody = null;
		},
		doneEdit(){
			const {id, editedBody:body, username, type} = this;
			this.updateMessage({id, body, username, type})
				.then(response => {
					this.cancelEdit();
				});
		},
	},
});

Vue.component("emoji-sel", {
	props: {
		emoji: {
			type    : String,
			required: true,
		},
		emojiRadioId: {
			type    : String,
			required: true,
		},
		value: {
			type    : String,
			required: true,
		},
	},
	template: `
	<div class="emoji-inp">
		<input type="radio" name="emojiRadio" :id="emojiRadioId" :value="emoji" v-model="updVal">
		<label :for="emojiRadioId">{{emoji}}</label>
	</div>
	`,
	computed: {
		updVal: {
			get(){
				return this.value;
			},
			set(val){
				this.$emit("input", val);
			},
		},
	},
});

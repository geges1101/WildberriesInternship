<template xmlns:v-on="http://www.w3.org/1999/xhtml">
  <div>
    <form v-on:submit.prevent="createMessage">
      <div class="input-group">
        <input v-model.trim="msgBody" type="text" class="form-control" placeholder="What's happening?">
        <div class="input-group-append">
          <button class="btn btn-primary" type="submit">Message</button>
        </div>
      </div>
    </form>

    <div class="mt-4">
      <Message v-for="message in message" :key="message.id" :message="message" />
    </div>
  </div>
</template>

<script>
import { mapState } from 'vuex';
import Message from '@/components/Message';

export default {
  data() {
    return {
      msgBody: '',
    };
  },
  computed: mapState({
    messages: (state) => state.messages,
  }),
  methods: {
    createMessage() {
      if (this.msgBody.length != 0) {
        this.$store.dispatch('createMessage', { body: this.msgBody });
        this.msgBody = '';
      }
    },
  },
  components: {
    Message,
  },
};
</script>

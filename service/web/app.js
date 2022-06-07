const app = Vue.createApp({
  data() {
    return {
      api: 'https://api.{{ app }}',
      hostname: '',
      output: '',
      ip: '',
      selected: 'ping',
      contact: false,
      subject: '',
      message: '',
      email: '',
      uuid: ''
    }
  },
  methods: {
    main() {
      fetch(this.api + '/' + this.selected + '/v1/' + this.hostname)
        .then(response => response.json())
        .then(data =>
          this.output = data.message)
        .catch(error => {
          console.error(error);
          this.output = '501 Not Implemented';
        })
    },
    active(index) {
      if (index != this.selected) {
        this.output = '';
        this.selected = index;
      }
    },
    send() {
      fetch(this.api + '/send/submit', {
        method: 'post',
          headers: {
            'Content-Type' : 'application/json'
          },
          body: JSON.stringify({
            "subject": this.subject,
            "message": this.message,
            "email": this.email,
            "channel": "slack"
          })
        })
        .then(response => response.json())
        .catch(error => {
          console.error(error);
        })
      this.subject = '';
      this.message = '';
      this.email = '';
      this.contact = false;
    }
  },
  mounted() {
    fetch(this.api + '/whoami/ip')
      .then(response => response.json())
      .then(data =>
        this.ip = data.message)
      .catch(error => console.error(error));
    fetch(this.api + '/uuid/generate')
      .then(response => response.json())
      .then(data =>
        this.uuid = data.uuid)
      .catch(error => console.error(error))
  }
})

app.mount('#app')

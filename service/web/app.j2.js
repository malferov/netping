const app = Vue.createApp({
  data() {
    return {
      api: 'https://api.{{ app }}',
      hostname: '',
      placeholder: 'Еnter IP address or hostname',
      output: '',
      ip: '',
      selected: 'ping',
      contact: false,
      button: "Send message",
      subject: '',
      message: '',
      email: '',
      uuid: ''
    }
  },
  methods: {
    call() {
      this.output = '';
      let path = this.hostname;
      if (this.selected == 'portcheck') {
        path = path.replace(':', '/');
      }
      fetch(this.api + '/' + this.selected + '/v1/' + path)
        .then(response => response.json())
        .then(data =>
          this.output = data.message)
        .catch(error => {
          console.error(error);
          this.output = error;
        })
    },
    active(index) {
      if (index != this.selected) {
        this.output = '';
        this.hostname = '';
        this.selected = index;
        if (index == 'portcheck') {
          this.placeholder = 'hostname:port'
        } else {
          this.placeholder = 'Еnter IP address or hostname'
        }
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
      this.button = "Message sent!";
      setTimeout(() => {
          this.contact = false;
          this.subject = '';
          this.message = '';
          this.email = '';
          this.button = "Send message";
      }, 2000);
    },
    async copy(s) {
      await navigator.clipboard.writeText(s);
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

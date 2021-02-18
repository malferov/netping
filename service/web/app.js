const app = new Vue({
  el: '#app',
  data: {
    api: 'https://api.{{ app }}',
    hostname: '',
    output: '',
    ip: '',
    selected: 'ping',
    contact: false,
    subject: '',
    message: '',
    email: ''
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
      fetch(this.api + '/send', {
        method: 'post',
          headers: {
            'Content-Type' : 'application/json'
          },
          body: JSON.stringify({
            "subject": this.subject,
            "message": this.message,
            "email": this.email
          })
        })
        .then(response => response.json())
        .catch(error => {
          console.error(error);
        })
      this.contact = false;
    }
  },
  mounted() {
    fetch(this.api + '/whoami/ip')
      .then(response => response.json())
      .then(data =>
        this.ip = data.message)
      .catch(error => console.error(error))
  }
})

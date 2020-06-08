const app = new Vue({
  el: '#app',
  data: {
    api: 'https://api.{{ app }}',
    hostname: '',
    output: '',
    ip: '',
  },
  methods: {
    main() {
      fetch(this.api + '/ping/v1/' + this.hostname)
        .then(response => response.json())
        .then(data =>
          this.output = data.message.replace(/(?:\r\n|\r|\n)/g, '<br>'))
        .catch(error => console.error(error))
    },
    init() {
      fetch(this.api + '/whoami/ip')
        .then(response => response.json())
        .then(data =>
          this.ip = data.message)
        .catch(error => console.error(error))
    }
  }
})

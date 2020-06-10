const app = new Vue({
  el: '#app',
  data: {
    api: 'https://api.{{ app }}',
    hostname: '',
    output: '',
    ip: '',
    selected: 'ping',
  },
  methods: {
    main() {
      fetch(this.api + '/' + this.selected + '/v1/' + this.hostname)
        .then(response => response.json())
        .then(data =>
          this.output = data.message.replace(/(?:\r\n|\r|\n)/g, '<br>'))
        .catch(error => {
          console.error(error);
          this.output = '404 Not Found';
        })
    },
    active(index) {
      if (index != this.selected) {
        this.output = '';
        this.selected = index;
      }
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

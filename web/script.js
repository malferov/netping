const app = new Vue({
  el: '#app',
  data: {
    api: 'https://api.{{ app }}',
    hostname: '',
    output: '',
  },
  methods: {
    main() {
      fetch(this.api + '/ping/v1/' + this.hostname)
        .then(response => response.json())
        .then(data =>
          this.output = data.message.replace(/(?:\r\n|\r|\n)/g, '<br>'))
        .catch(error => console.error(error))
    }
  }
})

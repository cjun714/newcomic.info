tmpl = `
  <ul id="array-rendering">
    <li v-for="info in comics">{{ info.name }}</li>
  </ul>
  `;

Vue.component('comic-list', {
  // new Vue({
  template: tmpl,
  data: function () {
    return {
      comics: [],
    }
  },
  mounted: function () {
    var self = this;
    fetch('/api/comics')
      .then(response => response.json())
      .then(data => self.comics = data)
      .catch(error => console.log(error));
  },
  methods: {
    go: (targetUrl => window.location.href = targetUrl)
  }
})

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

Vue.component('paginate', VuejsPaginate)
Vue.component('paging', {
  template: `
    <paginate :pageCount="2000" :page-range="10" :margin-pages="1" :containerClass="'pagination'"
      :clickHandler="clickCallback"></paginate>
  `,
  methods: {
    clickCallback: function (page) {
      console.log(page)
    }
  }
})

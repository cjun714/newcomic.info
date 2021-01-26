fetch('/api/comics/1')
  .then(response => response.json())
  .then(data => {
    v.$refs.list.comics = data.data
    v.$refs.index.pageCount = data.count / 20
  })
  .catch(error => console.log(error));

tmpl = `
  <ul id="array-rendering">
    <li v-for="info in comics">{{ info.name }}</li>
  </ul>
  `;
Vue.component('comic-list', {
  template: tmpl,
  data: function () {
    return {
      comics: [],
    }
  },
  methods: {
    go: (targetUrl => window.location.href = targetUrl)
  }
})

Vue.component('paginate', VuejsPaginate)
Vue.component('index', {
  template: `
    <paginate :pageCount="pageCount" :page-range="10" :margin-pages="1" :containerClass="'pagination'"
      :clickHandler="clickCallback"></paginate>
  `,
  data: function () {
    return {
      pageCount: 0
    }
  },
  methods: {
    clickCallback: function (page) {
      url = '/api/comics/'
      fetch(url + page)
        .then(response => response.json())
        .then(data => {
          v.$refs.list.comics = data.data
          v.$refs.index.pageCount = data.count / 20
        })
        .catch(error => console.log(error));
    }
  }
})

function clickCallback(page) {}

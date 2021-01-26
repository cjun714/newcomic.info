page = 1

idx = getParameters()["page"]
if (idx != null) {
  page = idx
}

fetch('/api/comics/' + page)
  .then(response => response.json())
  .then(data => {
    v.$refs.list.comics = data.data
    v.$refs.index.pageCount = data.count / 20
    v.$refs.index.page = page
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
    <paginate v-model="page" :pageCount="pageCount" :page-range="10" :margin-pages="1" :containerClass="'pagination'"
      :clickHandler="clickCallback"></paginate>
  `,
  data: function () {
    return {
      pageCount: 0,
      page: 1
    }
  },
  created: function () {},
  methods: {
    clickCallback: function (page) {
      // url = '/api/comics/'
      window.location.href = '/page/index.html' + '?page=' + page
      // fetch(url + page)
      //   .then(response => response.json())
      //   .then(data => {
      //     v.$refs.list.comics = data.data
      //     v.$refs.index.pageCount = data.count / 20
      //   })
      //   .catch(error => console.log(error));
    }
  }
})

function getParameters() {
  var vars = {};
  var parts = window.location.href.replace(/[?&]+([^=&]+)=([^&]*)/gi, function (m, key, value) {
    vars[key] = value;
  });
  return vars;
}

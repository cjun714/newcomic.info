currentPage = 1

idx = getParameters()["page"]
if (idx != null) {
  currentPage = Number(idx)
}

fetch('/api/comics/' + String(currentPage))
  .then(response => response.json())
  .then(data => {
    v.$refs.list.comics = data.data
    v.$refs.index.pageCount = Math.round(data.count / 20)
  })
  .catch(error => console.log(error));

Vue.component('comic-list', {
  template: `
  <ul id="array-rendering">
    <li v-for="info in comics">{{ info.name }}</li>
  </ul>
  `,
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
    <paginate v-model="page" :pageCount="pageCount" :page-range="10" :margin-pages="1" :containerClass="'pagination'" :clickHandler="clickCallback"></paginate>
  `,
  data: function () {
    return {
      pageCount: 0,
      page: currentPage
    }
  },
  created: function () {
    // v.$refs.index.page = 3
  },
  methods: {
    clickCallback: function (p) {
      // url = '/api/comics/'
      window.location.href = '/page/index.html' + '?page=' + p
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
